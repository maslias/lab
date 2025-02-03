## What is Envstruct?

Envstruct is a lightweight package for loading OS enviroment variables into nested structs for golang.
By default Envstruct use the name structure of the nested struct to look for the enviroment variable.

## Install

```
go get github.com/maslias/envstruct

```

## Usage Example

The following example shows how to use the Envstruct package.

```
import "github.com/maslias/envstruct"

type Config struct {
	Web WebConfig
    App AppConfig
}

type WebConfig struct {
	Http struct {
		Server struct {
			Port string
			Name string `default:myServer"`
		}
	}
}

type AppConfig struct {
	Version    string
Enviroment string `env:"special_key"`
}

cfg := Config{}
err := envstruct.Parse(&cfg)

```

Example of enviroment variables.

```
WEB_HTTP_SERVER_PORT=5000
APP_VERSION=0.0.2
SPECIAL_KEY=development

```

## Tips

- The Parse() function gets a pointer of a struct.
- It supports nested structs and structs pointers.
- The name structure of the nested struct is the default eviroment variable to look for.
- You can use struct tags for specific enviroment variable names and/or default values when the variables are not found.

## Optional Struct Tags

```
`env:"special_key"`
`default:"this is my default value"`

```

## More options

```
Delimeter             string = "_"
Preterm               string = "CUSTOMERXX"
Capitalize            bool   = true
TagKeyForValue        string = "env"
TagKeyForDefaultValue string = "default"

```

## What Envstruct does not do

Envstruct does not load your .env file in the enviroment variable. Use your prefert way to do that, like [GoDoEnv](https://github.com/joho/godotenv) or a Makefile with the following bash script:

```
define setup_env
	$(eval ENV_FILE := ./$(1).env)
	@echo " - load env file:  $(ENV_FILE)"
	$(eval include ./$(1).env)
	$(eval export sed 's/=.*//' ./$(1).env)
endef

env:
	$(call setup_env,dev)

run: env;
	@go run .

```

## Inspired by

Envstruct is heavily inspired by:

- (https://github.com/golobby/env)
- (https://github.com/spf13/viper)

## License

Maslias Envstruct is released under the [MIT License](https://github.com/maslias/envstruct/blob/master/LICENSE).

