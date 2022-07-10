package command

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/aligoren/netenv/config"
)

type Command struct {
	CommandText string
	Config      *config.Config
	Con         net.Conn
}

func (c *Command) parseCommand() []string {
	return strings.Split(strings.ToLower(c.CommandText), ":")
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

		return errors.New("ip is not allowed")
	}

	return nil
}

func (c *Command) checkAuth() ([]string, error) {

	// The first and second fields always used for auth purpose
	commands := c.parseCommand()

	if c.Config.Global.Auth.Enabled {

		if strings.ToLower(commands[0]) != "auth" {
			return nil, errors.New("auth is required")
		}

		username := commands[1]
		password := commands[2]

		if username != c.Config.Global.Auth.Username || password != c.Config.Global.Auth.Password {
			return nil, errors.New("username or password is wrong")
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

// HandleCommand it takes command as a string and takes config.
// it will always check the authentication
func (c *Command) HandleCommand() string {

	commands, err := c.checkAuth()
	if err != nil {
		return fmt.Sprintf("%s\n", err)
	}

	cmd := commands[len(commands)-1]

	if cmd == "echo" {
		return "Hello :)\n"
	}

	return ""
}
