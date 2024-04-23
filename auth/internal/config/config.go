package config

import (
	"flag"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port     int      `yaml:"port"`
	Env      string   `yaml:"env"`
	Database Postgres `yaml:"db"`
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
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", p.User, p.Password, p.Host, p.Port, p.Name)
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
