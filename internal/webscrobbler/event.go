package webscrobbler

import "fmt"

type EventType string

const (
	WSNowPlaying     EventType = "nowplaying"
	WSResumedPlaying EventType = "resumedplaying"
	WSPaused         EventType = "paused"
	WSScrobble       EventType = "scrobble"
	WSLoved          EventType = "loved"
)

type ParsedSong struct {
	Track       string `json:"track"`
	Artist      string `json:"artist"`
	Duration    int    `json:"duration"`
	CurrentTime int    `json:"currentTime"`
}

type ProcessedSong struct {
	Track    string `json:"track"`
	Artist   string `json:"artist"`
	Duration int    `json:"duration"`
}

type Metadata struct {
	UserPlayCount int `json:"userPlayCount"`
}

type Song struct {
	Parsed    ParsedSong    `json:"parsed"`
	Processed ProcessedSong `json:"processed"`
	Metadata  Metadata      `json:"metadata"`
}

type WebScrobblerEvent struct {
	EventName       EventType `json:"eventName"`
	TimestampMillis int64     `json:"time"`
	Data            struct {
		Song Song `json:"song"`
	} `json:"data"`
}

func (s Song) getArtist() string {
	processedArtist := s.Processed.Artist
	if processedArtist != "" {
		return processedArtist
	}
	return s.Parsed.Artist
}

func (s Song) getTrack() string {
	processedTrack := s.Processed.Track
	if processedTrack != "" {
		return processedTrack
	}
	return s.Parsed.Track
}

func (s Song) getDuration() int {
	processedDuration := s.Processed.Duration
	if processedDuration != 0 {
		return processedDuration
	}
	return s.Parsed.Duration
}

func (s Song) getCurrentTime() int {
	return s.Parsed.CurrentTime
}

func (s Song) getPlayCountString() string {
	return fmt.Sprintf("%d回再生", s.Metadata.UserPlayCount)
}

func (s Song) getSlackStatus() string {
	return fmt.Sprintf("%s「%s」（%s）", s.getArtist(), s.getTrack(), s.getPlayCountString())
}
