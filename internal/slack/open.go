package slack

import (
	"net/url"

	"scrobble-status/internal/browser"
	"scrobble-status/internal/config"
)

func OpenOAuthPage() error {
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}
	path := "https://slack.com/oauth/v2/authorize"
	values := url.Values{
		"user_scope":   {"users.profile:write"},
		"redirect_uri": {conf.SlackRedirectURI},
		"client_id":    {conf.ClientID},
	}
	err = browser.OpenURL(path + "?" + values.Encode())
	return err
}
