package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"scrobble-status/internal/config"
)

type Status struct {
	StatusText       string `json:"status_text"`
	StatusEmoji      string `json:"status_emoji"`
	StatusExpiration int64  `json:"status_expiration"`
}

type Profile struct {
	Status Status `json:"profile"`
}

func SendStatus(emoji string, status string, durationSeconds int) error {
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}
	body, err := json.Marshal(Profile{
		Status: Status{
			StatusText:       status,
			StatusEmoji:      emoji,
			StatusExpiration: time.Now().Add(time.Second * time.Duration(durationSeconds)).Unix(),
		},
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://slack.com/api/users.profile.set", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+conf.AccessToken)
	_, err = http.DefaultClient.Do(req)
	return err
}
