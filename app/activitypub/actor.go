package activitypub

import (
	"encoding/json"
	"time"
)

type Actor struct {
	Context           []interface{} `json:"@context"`
	Followers         string        `json:"followers"`
	Following         string        `json:"following"`
	Id                string        `json:"id"`
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
	Icon              Media         `json:"icon,omitempty"`
	Image             Media         `json:"image,omitempty"`
}

type Media struct {
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
	var actor Actor
	err := json.Unmarshal(data, &actor)
	return actor, err
}

func (actor *Actor) Marshal() ([]byte, error) {
	return json.Marshal(actor)
}
