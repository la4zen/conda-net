package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Http struct {
	IP   string `json:"ip,omitempty"`
	Port string `json:"port,omitempty"`
}

type Database struct {
	IP       string `json:"ip,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DB       string `json:"db,omitempty"`
}

type Config struct {
	Http     `json:"http,omitempty"`
	Database `json:"database,omitempty"`
}

func NewConfig(configpath string) Config {
	config := &Config{
		Http{
			IP:   "127.0.0.1",
			Port: "8080",
		},
		Database{
			IP:       "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DB:       "postgres",
		},
	}
	file, err := os.Open(configpath)
	if err != nil {
		log.Println("Using default config.")
		return *config
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Failed to read config. Using default config.")
		return *config
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Println("Failed to unmarshal configuration. Using default config.")
		return *config
	}
	return *config
}
