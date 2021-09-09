package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() (*Config, error) {
	viper.AddConfigPath(Path)
	viper.SetConfigType("toml")
	viper.SetConfigName("env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config = &Config{new(Environment), new(Application)}

	env := viper.GetString("environment")
	if env == Dev {
		config.e.isDev = true
	} else if env == Pre {
		config.e.isDev = true
	} else if env == Prod {
		config.e.isDev = true
	} else {
		return nil, errors.New("environment is invalid")
	}

	viper.SetConfigName(fmt.Sprintf("application-%s", env))
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(config.a)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) Env() string {
	if c.e.isDev {
		return Dev
	} else if c.e.isPre {
		return Pre
	} else if c.e.isProd {
		return Prod
	}
	return ""
}

func (c *Config) IsDev() bool {
	return c.e.isDev
}

func (c *Config) IsPre() bool {
	return c.e.isPre
}

func (c *Config) IsProd() bool {
	return c.e.isProd
}

func (c *Config) Get() *Application {
	return c.a
}

func Env() string {
	if config.e.isDev {
		return Dev
	} else if config.e.isPre {
		return Pre
	} else if config.e.isProd {
		return Prod
	}
	return ""
}

func IsDev() bool {
	return config.e.isDev
}

func IsPre() bool {
	return config.e.isPre
}

func IsProd() bool {
	return config.e.isProd
}

func Get() *Application {
	return config.a
}
