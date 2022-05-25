package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		S3    `yaml:"s3"`
		Kafka `yaml:"kafka"`
		//RMQ  `yaml:"rabbitmq"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// S3 -.
	S3 struct {
		URL             string `env-required:"true" yaml:"url" env:"S3_URL"`
		AccessKeyID     string `env-required:"true" yaml:"accessKeyID"               env:"S3_ID"`
		SecretAccessKey string `env-required:"true" yaml:"secretAccessKey"               env:"S3_KEY"`
		UseSSL          bool   `env-required:"true" yaml:"useSSL"               env:"USE_SSL"`
	}

	//Kafka -.
	Kafka struct {
		url string `env-required:"true" yaml:"url" env:"KAFKA_URL"`
	}

	// RMQ -.
	//RMQ struct {
	//	ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
	//	ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
	//	URL            string `env-required:"true"                            env:"RMQ_URL"`
	//}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
