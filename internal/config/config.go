package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Env      string         `yaml:"env" env-default:"local"`
		Version  string         `yaml:"version"`
		GRPC     GRPCConfig     `yaml:"grpc"`
		Postgres PostgresConfig `yaml:"postgres"`
		Redis    RedisConfig    `yaml:"redis"`
	}

	GRPCConfig struct {
		Port    string `yaml:"port" env-default:"50051"`
		Timeout string `yaml:"timeout"`
	}

	PostgresConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	}

	RedisConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Db       int    `yaml:"db"`
		Password string `yaml:"password"`
	}
)

func MustLoad() (error, *Config) {
	op := "config.MustLoad"

	configPath := fetchConfigPath()
	if configPath == "" {
		return fmt.Errorf("%s: %s", op, "config path is empty"), nil
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) (error, *Config) {
	op := "config.MustLoadPath"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", op, err), nil
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return fmt.Errorf("%s: %s", op, err), nil
	}

	return nil, &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
