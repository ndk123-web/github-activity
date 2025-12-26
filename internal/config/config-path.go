package config

import (
	"encoding/json"
	// "fmt"
	"os"
	"path/filepath"
)

func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	// fmt.Println("Home Dir: ", home)
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, ".gh-activity", "config.json")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return "", err
		}
		file, err := os.Create(path)
		if err != nil {
			return "", err
		}
		defer file.Close()
	}
	return path, nil
}

func WriteConfigFile(cfg GhTokenConfig, path string) error {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, []byte(data), 0600)
}
