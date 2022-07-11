package command

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/aligoren/netenv/config"
	"github.com/aligoren/netenv/internal/helpers"
	"github.com/joho/godotenv"
)

type Command struct {
	CommandText string
	Config      *config.Config
	Con         net.Conn
}

func (c *Command) parseCommand() []string {
	return strings.Split(c.CommandText, ":")
}

func (c *Command) checkIpList() error {
	ipList := c.Config.Global.Auth.IPList

	if len(ipList) > 0 {
		for _, ip := range ipList {
			addr := strings.Split(c.Con.RemoteAddr().String(), ":")
			if ip == addr[0] {
				return nil
			}
		}

		return errors.New(IP_IS_NOT_ALLOWED)
	}

	return nil
}

func (c *Command) checkAuth() ([]string, error) {

	// The first and second fields always used for auth purpose
	commands := c.parseCommand()

	if c.Config.Global.Auth.Enabled {

		if strings.ToLower(commands[0]) != "auth" {
			return nil, errors.New(AUTH_IS_REQUIRED)
		}

		username := commands[1]
		password := commands[2]

		if username != c.Config.Global.Auth.Username || password != c.Config.Global.Auth.Password {
			return nil, errors.New(USERNAME_OR_PASSWORD_IS_WRONG)
		}

		err := c.checkIpList()
		if err != nil {
			return nil, err
		}

		// we make them empty
		commands[0] = ""
		commands[1] = ""
	}

	return commands, nil
}

func parseCustomVariables(selectedVariables string, data map[string]string, excludes []string) (map[string]string, error) {

	customVariables := make(map[string]string)
	variables := make([]string, len(data))

	if selectedVariables == "" || selectedVariables == "*" {
		for variable := range data {
			variables = append(variables, variable)
		}
	} else {
		variables = strings.Split(selectedVariables, ",")
	}

	hasExcludes := len(excludes) > 0

	for _, variable := range variables {
		value := data[variable]
		if value != "" {
			isInExcludedList := hasExcludes && helpers.IsSliceContainsString(variable, excludes)
			if !isInExcludedList {
				customVariables[variable] = data[variable]
			}
		}
	}

	if len(customVariables) == 0 {
		return nil, errors.New(THERE_IS_NO_KEY_VALUE_PAIR_FOUND)
	}

	return customVariables, nil
}

func (c *Command) parseEnv(commands []string) (map[string]string, error) {
	var file string
	var environment string
	var selectedVariables string

	for _, commmand := range commands {
		splittedCommand := strings.Split(commmand, "$")
		if len(splittedCommand) > 0 {
			switch strings.ToLower(splittedCommand[0]) {
			case "file":
				file = splittedCommand[1]
			case "env":
				environment = splittedCommand[1]
			case "var":
				selectedVariables = splittedCommand[1]
			}
		}
	}

	cfg := c.Config.EnvFiles[file]

	if environment == "" {
		environment = cfg.Default
	}

	env := cfg.Environments[environment]

	dotFile, err := os.OpenFile(env.Path, os.O_RDONLY, 0744)
	if err != nil {
		return nil, err
	}

	data, err := godotenv.Parse(dotFile)
	if err != nil {
		return nil, err
	}

	customVariables, err := parseCustomVariables(selectedVariables, data, env.Excludes)

	if err != nil {
		return nil, err
	}

	return customVariables, nil
}

// HandleCommand it takes command as a string and takes config.
// it will always check the authentication
func (c *Command) HandleCommand() string {

	commands, err := c.checkAuth()
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	cmds := commands

	if c.Config.Global.Auth.Enabled {
		cmds = commands[3:]
	}

	if cmds[0] == "echo" {
		return HELLO
	}

	variables, err := c.parseEnv(cmds)
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	var output strings.Builder

	for key, variable := range variables {
		output.WriteString(fmt.Sprintf("{%s:%s}$", key, variable))
	}

	value := strings.TrimRight(output.String(), "$")

	return fmt.Sprintf("%s\n", value)
}
