package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
}

type APIConfig struct {
	Port string `envconfig:"API_PORT" default:":8000"`
}

type PostgresConfig struct {
	DatabaseName string `envconfig:"DB_NAME" default:"transfer-me"`
	User         string `envconfig:"DB_USER" default:"postgres"`
	Password     string `envconfig:"DB_PASS" default:"postgres"`
	Host         string `envconfig:"DB_HOST" default:"127.0.0.1"`
	Port         string `envconfig:"DB_PORT" default:"5432"`
	SSLMode      string `envconfig:"DB_SSLMODE" default:"disable"`
}

func (p PostgresConfig) DSN() string {
	connectString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		p.User, p.Password, p.Host, p.Port, p.DatabaseName)

	if p.SSLMode != "" {
		connectString = fmt.Sprintf("%s sslmode=%s",
			connectString, p.SSLMode)
	}

	return connectString
}

func (p PostgresConfig) URL() string {
	if p.SSLMode == "" {
		p.SSLMode = "disable"
	}

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.DatabaseName, p.SSLMode)

	return connectString
}

func LoadConfig() (*Config, error) {
	var config Config
	noPrefix := ""
	err := envconfig.Process(noPrefix, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
