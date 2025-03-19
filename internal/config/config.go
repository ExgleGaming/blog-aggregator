package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

// SetUser This set the username and write it to the config
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return Write(*c)
}

// Read This is to read the config file and return it in usable format
func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer data.Close()

	decoder := json.NewDecoder(data)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// getConfigFilePath Retrieves the config file path and returns it. If not at location it will create it for you.
func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFileName), nil
}

// Write This is the function called to write to the config file
func Write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := os.Create(path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(data)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
