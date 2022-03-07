package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BotToken string `json:"bot_token"`
}

func LoadConfig(path string) (*Config, error) {

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := json.Unmarshal(jsonBytes, &c); err != nil {
		return nil, err
	}

	return &c, nil

}
