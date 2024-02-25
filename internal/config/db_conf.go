package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DbName         string `yaml:"db_name"`
	CollectionName string `yaml:"collection_name"`
}

func GetConnection() (*Config, error) {
	confData, err := os.ReadFile("./internal/config/conf.yaml")
	if err != nil {
		return nil, err
	}

	var conf Config

	err = yaml.Unmarshal(confData, &conf)
	if err != nil {
		return nil, err
	}

	result := Config{
		conf.DbName,
		conf.CollectionName,
	}

	return &result, nil
}
