package webscrobbler

import (
	"net/url"

	"scrobble-status/internal/config"
)

func GetWSAuthUrl() (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	path := "https://web-scrobbler.com/webhook"
	values := url.Values{
		"applicationName": {"Web Scrobbler Slack Status"},
		"userApiUrl":      {conf.WSInfoURI},
	}
	return path + "?" + values.Encode(), nil
}
