package config

import (
	"flag"
	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port     int    `yaml:"port"`
	Env      string `yaml:"env"`
	Database Postgres
}

type Postgres struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func MustInit() *Config {
	path := fetchPath()

	_, err := os.Stat(path)
	if err != nil {
		panic("failed to find cfg file")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read config file" + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to unmarshal" + err.Error())
	}
	return &cfg
}

func (p *Postgres) GetDSN() string {
	return "postgres://postgres:postgres@localhost:5000/auth?sslmode=disable"
}

func fetchPath() string {
	var path string
	_ = godotenv.Load()

	flag.StringVar(&path, "c", "", "path to cfg file")
	flag.Parse()

	if path == "" {
		path = "config/config.yaml"
	}

	return path
}
