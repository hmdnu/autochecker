package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	User []struct {
		Token string
		Uid   string
		Name  string
	}
	CheckTime string
}

func GetConfig() (Config, error) {
	var config Config

	file, err := os.Open("config.json")
	if err != nil {
		return config, fmt.Errorf("error opening config.json : %w", err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
