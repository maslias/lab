package config

import "github.com/maslias/envstruct"

type Config struct {
	App       AppConfig
	Common    CommonConfig
	Web       WebConfig
	User      UserConfig
	Health    HealthConfig
	Forwarder ForwarderConfig
}

type CommonConfig struct {
	Consul struct {
		Addr string
	}
}

type AppConfig struct {
	Enviroment string
	Version    string
}

type WebConfig struct {
	Http struct {
		Server struct {
            Port string `default:"localhost:3000"`
			Name string `default:"web-http"`
		}
	}
}

type ForwarderConfig struct {
	Grpc struct {
		Server struct {
			Addr string `default:"2000"`
			Name string `default:"forwarder-grpc"`
		}
	}
	Http struct {
		Server struct {
            Port string `default:"localhost:3001"`
			Name string `default:"forwarder-http"`
		}
	}
}

type HealthConfig struct {
	Http struct {
		Server struct {
            Port string `default:"localhost:3002"`
			Name string `default:"health-http"`
		}
	}
}

type UserConfig struct {
	Grpc struct {
		Server struct {
            Addr string `default:"localhost:2001"`
			Name string `default:"user-grpc"`
		}
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := envstruct.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
