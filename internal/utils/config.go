package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GRPC struct {
	GRPCPort string `yaml:"grpc_port"`
}

type DATABASE struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	DBSslMode  string `yaml:"db_ssl_mode"`
	DBDriver   string `yaml:"db_driver"`
}
type KAFKA struct {
	KafkaBroker string `yaml:"broker_kafka"`
	KafkaTopic  string `yaml:"topic"`
}

type APP struct {
	AppChunkSize int    `yaml:"chunk_size"`
	Logfile      string `yaml:"logfile"`
}
type Config struct {
	Grpc     GRPC     `yaml:"grpc"`
	Database DATABASE `yaml:"database"`
	Kafka    KAFKA    `yaml:"kafka"`
	App      APP      `yaml:"app"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
