package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func MustLoad() *Config {
	configPath := fetchPath()
	if configPath == "" {
		panic("No config provided")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

  var cfg Config
  if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
    panic(fmt.Sprintf("Errored to read config: %e", err))
  }

  return &cfg
}

func fetchPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
