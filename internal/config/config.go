package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/" + configFileName, nil
}

func ReadConfig() (Config, error) {
	// Get home dir
	cfgPath, err := getConfigDir()
	if err != nil {
		return Config{}, fmt.Errorf("user home directory not found")
	}

	// Read and unmarshal config file
	file, err := os.Open(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("config path not found")
	}
	defer file.Close()

	var config Config
	fileBytes, _ := io.ReadAll(file)

	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error while reading config: %v", err)
	}

	return config, nil
}

func (c Config) SetUser(currentUser string) error {
	c.CurrentUserName = currentUser

	// Write to file
	cfgPath, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("user home directory not found")
	}

	file, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("config path not found")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
