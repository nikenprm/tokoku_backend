package config

import (
	"fmt"
)

type DBParam struct {
	Type string
	Host string
	Port string
	User string
	Pass string
	Name string
}

func (params *DBParam) MigrationString() (string, error) {
	if params.Type == "postgres" {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", params.User, params.Pass, params.Host, params.Port, params.Name), nil
	} else {
		return "", fmt.Errorf("Unknown database type: %s", params.Type)
	}
}

func (params *DBParam) ConnectionString() (string, error) {
	if params.Type == "postgres" {
		return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", params.Name, params.User, params.Pass, params.Host, params.Port), nil
	} else {
		return "", fmt.Errorf("Unknown database type: %s", params.Type)
	}
}
