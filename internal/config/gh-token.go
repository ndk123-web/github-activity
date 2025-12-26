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

	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		return "", err
	}

	// fmt.Println("Created At: ", cfg.CreatedAt)

	if cfg.CreatedAt.IsZero() {
		fmt.Println("- Warning: Token creation time missing. Please reset token.")
		return cfg.Token, nil
	} else {
		// fmt.Println("- GitHub Token was set on:", cfg.CreatedAt.Format("2006-01-02 15:04:05"))

		if time.Since(cfg.CreatedAt) > 90*24*time.Hour {
			fmt.Println("- Warning: GitHub Token is older than 90 days. Consider updating it for better security.")
		}
	}

	GhToken = cfg.Token

	return cfg.Token, nil
}
