package activitypub

import (
	"encoding/json"
	"time"
)

type Actor struct {
	Context           []interface{} `json:"@context"`
	Followers         string        `json:"followers"`
	Following         string        `json:"following"`
	ID                string        `json:"id"`
	Type              string        `json:"type"`
	PreferredUsername string        `json:"preferredUsername"`
	Inbox             string        `json:"inbox"`
	Outbox            string        `json:"outbox"`
	Playlists         string        `json:"playlists"`
	Name              string        `json:"name"`
	PublicKey         PublicKey     `json:"publicKey"`
	Url               string        `json:"url"`
	Summary           string        `json:"summary"`
	Published         time.Time     `json:"published"`
	Icon              Icon          `json:"icon,omitempty"`
	Image             Icon          `json:"image,omitempty"`
}

type Icon struct {
	Height    int64  `json:"height,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Width     int64  `json:"width,omitempty"`
}

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

func UnmarshalActor(data []byte) (Actor, error) {
	var r Actor
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Actor) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
