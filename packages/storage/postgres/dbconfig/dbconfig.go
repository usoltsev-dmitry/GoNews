package dbconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
}

// ConnString формирует и возвращает строку соединения на основе конфигурации.
func (c *Config) ConnString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DbName,
		c.Database.SSLMode,
	)
}

// LoadConfig загружает конфигурацию из файла.
func LoadConfig(filename string) (*Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
