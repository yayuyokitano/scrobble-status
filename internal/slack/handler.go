package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"scrobble-status/internal/config"
)

type Token struct {
	AuthedUser struct {
		AccessToken string `json:"access_token"`
	} `json:"authed_user"`
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	conf, err := config.GetConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	code := r.URL.Query().Get("code")
	values := url.Values{
		"code":          {code},
		"client_id":     {conf.ClientID},
		"client_secret": {conf.ClientSecret},
	}
	req, err := http.NewRequest("POST", "https://slack.com/api/oauth.v2.access", strings.NewReader(values.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res.StatusCode != 200 {
		http.Error(w, res.Status, res.StatusCode)
		return
	}
	var token Token
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if token.AuthedUser.AccessToken == "" {
		http.Error(w, "access token is empty", http.StatusUnauthorized)
		return
	}
	err = config.SetAccessToken(token.AuthedUser.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		fmt.Println(err)
	}
}
