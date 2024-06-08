package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ClientID         string `json:"client_id"`
	ClientSecret     string `json:"client_secret"`
	SlackRedirectURI string `json:"slack_redirect_uri"`
	WSInfoURI        string `json:"ws_info_uri"`
	AccessToken      string `json:"access_token"`
}

func GetConfig() (conf Config, err error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &conf)
	return
}

func SetAccessToken(token string) error {
	var conf Config
	conf, err := GetConfig()
	if err != nil {
		return err
	}
	conf.AccessToken = token
	confJSON, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", confJSON, 0644)
	return err
}
