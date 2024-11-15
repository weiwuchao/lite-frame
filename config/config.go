package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	MySql MySQLInfo `yaml:"mysql"`
}

type MySQLInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"dbname"`
}

func NewConfig() (*Config, error) {
	fileByte, err := os.ReadFile("F:\\practice\\src\\lite-frame\\config\\config.yaml")
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(fileByte, cfg)
	return cfg, err
}

func NewConfigTest() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	err = yaml.NewDecoder(file).Decode(cfg)
	return cfg, err
}
