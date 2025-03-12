package database

import "os"

type DbConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Name:     os.Getenv("MYSQL_DB"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
	}
}
