package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig(filename string) (*DBConfig, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	dbConfig := &DBConfig{
		Host:     cfg.Section("database").Key("host").String(),
		Port:     cfg.Section("database").Key("port").MustInt(),
		User:     cfg.Section("database").Key("user").String(),
		Password: cfg.Section("database").Key("password").String(),
		DBName:   cfg.Section("database").Key("name").String(),
		SSLMode:  cfg.Section("database").Key("sslmode").String(),
	}

	return dbConfig, nil
}
