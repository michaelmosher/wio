package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type jiraConfig struct {
	Hostname string
	Username string
	Password string
}

type databaseConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type wioConfig struct {
	Jira jiraConfig
	DB   databaseConfig
}

func readConfig(filename string) (cfg wioConfig) {
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatal(err)
	}

	return
}
