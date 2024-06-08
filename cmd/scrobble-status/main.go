package main

import (
	"fmt"
	"net/http"

	"scrobble-status/internal/config"
	"scrobble-status/internal/slack"
	"scrobble-status/internal/webscrobbler"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	if conf.AccessToken == "" {
		err = slack.OpenOAuthPage()
		if err != nil {
			panic(err)
		}
	}

	wsAuthUrl, err := webscrobbler.GetWSAuthUrl()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting server. To connect to web scrobbler, please click this link: %s\n", wsAuthUrl)

	http.HandleFunc("/auth", slack.HandleAuth)
	http.HandleFunc("/webscrobbler", webscrobbler.HandleWebScrobbler)
	err = http.ListenAndServe(":8564", nil)
	if err != nil {
		panic(err)
	}
}
