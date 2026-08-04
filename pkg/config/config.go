package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Logger   LoggerConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	ConnectionString string
	MigrationsPath   string
}

type LoggerConfig struct {
	Level string
}

type ServerConfig struct {
	Addr string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("HTTP_ADDR", ":8080")
	viper.SetDefault("MIGRATIONS_PATH", "migrations")

	_ = viper.ReadInConfig()

	connectionString := viper.GetString("CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("CONNECTION_STRING is required")
	}

	cfg := &Config{
		Database: DatabaseConfig{
			ConnectionString: connectionString,
			MigrationsPath:   viper.GetString("MIGRATIONS_PATH"),
		},
		Logger: LoggerConfig{
			Level: viper.GetString("LOG_LEVEL"),
		},
		Server: ServerConfig{
			Addr: viper.GetString("HTTP_ADDR"),
		},
	}

	return cfg, nil
}
