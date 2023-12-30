package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	service struct {
		app_mode    string `yaml:"app_mode"`
		server_port int    `yaml:"server_port"`
	}
	mysql struct {
		host   string `yaml:"host"`
		port   int    `yaml:"port"`
		user   string `yaml:"user"`
		passwd string `yaml:"passwd"`
		dbname string `yaml:"dbname"`
	}
}

func Load() {

}
