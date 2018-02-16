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

type wioConfig struct {
	Jira jiraConfig
}

func readConfig(filename string) (cfg wioConfig) {
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatal(err)
	}

	return
}
