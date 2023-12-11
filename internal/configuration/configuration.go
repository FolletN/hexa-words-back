package configuration

import (
	"fmt"
	"os"
	"strconv"

	"hexacrosswords/internal/api"
	"hexacrosswords/internal/db"
)

type Configuration struct {
	Server api.Configuration
	DB     db.Configuration
}

func ReadConfiguration() (Configuration, error) {
	serverConfiguration, err := ReadServerConfiguration()
	if err != nil {
		return Configuration{}, err
	}

	dbConfiguration := ReadDatabaseConfiration()

	return Configuration{
		Server: serverConfiguration,
		DB:     dbConfiguration,
	}, nil
}

const (
	serverPortEnv     = "SERVER_PORT"
	defaultServerPort = 8080
)

func ReadServerConfiguration() (api.Configuration, error) {
	serverConfig := api.Configuration{
		Port: defaultServerPort,
	}
	portString := os.Getenv(serverPortEnv)
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			return api.Configuration{}, fmt.Errorf("failed to read env var %s : %w", serverPortEnv, err)
		}
		serverConfig.Port = port
	}
	return serverConfig, nil
}

const (
	databasePortEnv     = "DATABASE_PORT"
	databaseAddressEnv  = "DATABASE_ADDRESS"
	databasePasswordEnv = "DATABASE_PASSWORD"
	databaseUserEnv     = "DATABASE_USER"
	databaseDatabaseEnv = "DATABASE_DATABASE"
)

func ReadDatabaseConfiration() db.Configuration {
	return db.Configuration{
		User:     os.Getenv(databaseUserEnv),
		Password: os.Getenv(databasePasswordEnv),
		Address:  os.Getenv(databaseAddressEnv),
		Port:     os.Getenv(databasePortEnv),
		Database: os.Getenv(databaseDatabaseEnv),
	}
}
