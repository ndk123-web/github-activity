package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var GhToken string = ""

type GhTokenConfig struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func SetGhToken(token string) {

	path, err := ConfigPath()
	if err != nil {
		fmt.Println("Error getting config path:", err)
		return
	}

	cfg := &GhTokenConfig{
		Token:     token,
		CreatedAt: time.Now(),
	}

	err = WriteConfigFile(*cfg, path)
	if err != nil {
		fmt.Println("Error writing config file:", err)
	}

	GhToken = token
	fmt.Println("- GitHub Token has been set successfully.")
}

func LoadGhToken() (string, error) {
	path, err := ConfigPath()
	if err != nil {
		return "", err
	}

	var cfg GhTokenConfig
	data, err := os.ReadFile(path)

	if err = json.Unmarshal(data, &cfg); err != nil {
		return "", err
	}

	return cfg.Token, nil
}
