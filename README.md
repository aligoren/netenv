## Netenv

This is a project to get environment variables from the network.

## Description

When you need to make your environment variables public for your network, you have to share them with everyone. You can share your .env files through the network using Netenv.

You only need to listen to TCP with some arguments

## Simple Output
```json
{name:ali}${MYSQL_USERNAME:admin}${TEST:false}
```

## Example netenv File

```yaml
global:
  addr: ":8080"
  auth:
    enabled: true # set this true if you want to use authentication
    username: admin
    password: admin
    iplist:
      - 0.0.0.0
      - 127.0.0.1

envfiles:
  project1:
    default: dev
    environments:
      dev:
        path: /home/ali/pyt/test.env
        excludes:
          - MYSQL_USERNAME
  project2:
    default: dev
    environments:
      dev:
        path: /home/ali/pyt/project2_dev.env
      stage:
        path: /home/ali/pyt/project2_stage.env
      prod:
        path: /home/ali/pyt/project2_prod.env
```