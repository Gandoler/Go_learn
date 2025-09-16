package config

import (
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" envDefault:"development"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" envDefault:"/data" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
type HTTPServer struct {
	Address      string        `yaml:"address" env:"HTTP_ADDR" envDefault:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" envDefault:"4s"`
	idle_timeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" envDefault:"60s"`
}

func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"

	}
}
