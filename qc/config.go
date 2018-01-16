package qc

import (
	"encoding/json"
	"log"
)

type Config struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	Protocol          string `json:"protocol"`
	URI               string `json:"uri"`
	AppId	string	`json:"app_id"`
}

type PoxiedConfig struct{
	Type string `json:"type"`
	Config json.RawMessage `json:"config"`
}

type Proxy struct{
	LogFile string `json:"log_file"`
	QYConfig Config `json:"qy_config"`
	PConfig PoxiedConfig `json:"poxied_config"`
	logger *log.Logger
	conn  Connector
}

