package webscrobbler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"scrobble-status/internal/slack"
)

var lastEventTime int64 = 0

func HandleWebScrobbler(w http.ResponseWriter, r *http.Request) {
	var event WebScrobblerEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusOK)
		return
	}
	if event.TimestampMillis < lastEventTime {
		return
	}
	lastEventTime = event.TimestampMillis

	switch event.EventName {
	case WSNowPlaying:
		err = slack.SendStatus(":musical_note:", event.Data.Song.getSlackStatus(), event.Data.Song.getDuration())
	case WSResumedPlaying:
		err = slack.SendStatus(":musical_note:", event.Data.Song.getSlackStatus(), event.Data.Song.getDuration()-event.Data.Song.getCurrentTime())
	case WSPaused:
		err = slack.SendStatus(":double_vertical_bar:", event.Data.Song.getSlackStatus(), 600)
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}
