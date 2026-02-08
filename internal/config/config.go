package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName     string `json:"app_name"`
	Version     string `json:"version"`
	Currency    string `json:"currency"`
	DemoAccount struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	} `json:"demoAccount"`
}

var GlobalConfig Config

func Init() error {
	data, err := os.ReadFile("config.json ")
	if err != nil {
		return fmt.Errorf("ошибка чтения json %w", err)
	}

	if err := json.Unmarshal(data, &GlobalConfig); err != nil {
		return fmt.Errorf("ошибка парсинга config.json %w", err)
	}
	return nil
}

func GetConfig() Config {
	return GlobalConfig
}
