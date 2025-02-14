package config

import "plc_project/internal/database"

type Config struct {
	MariaDB database.DBConfig
	Redis   struct {
		Host     string
		Port     int
		Password string
	}
}

func LoadConfig() *Config {
	return &Config{
		MariaDB: database.DBConfig{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "Danilo@34333528",
			Database: "plc_config",
		},
		Redis: struct {
			Host     string
			Port     int
			Password string
		}{
			Host:     "localhost",
			Port:     6379,
			Password: "",
		},
	}
}
