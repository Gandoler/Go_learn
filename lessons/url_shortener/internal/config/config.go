package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" envDefault:"development"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" envDefault:"/data" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDR" envDefault:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" envDefault:"4s"`
	idleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" envDefault:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH %s does not exist", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg

}
