package config

import (
	"WarungMakan/manager"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func NewConfig() *Config {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "warung_makan"
	dbUser := "postgres"
	dbPassword := "5434jago"

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager
	return config
}
