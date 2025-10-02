package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yml:"env" env-default:"local"`
	StoragePath string        `yml:"storage_path" env-required:"true"`
	TokenTTl    time.Duration `yml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yml:"grpc"`
}

type GRPCConfig struct {
	Port    int `yml:"port" `
	Timeout int `yml:"timeout" `
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found" + path)
	}
	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config" + err.Error())
	}
	return &config

}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "config.yml", "config file path")
	flag.Parse()
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
