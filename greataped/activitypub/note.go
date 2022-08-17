package activitypub

import (
	"config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Context      string   `json:"@context" validate:"activitystream"`
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type"`
	To           []string `json:"to"`
	AttributedTo string   `json:"attributedTo"`
	InReplyTo    string   `json:"inReplyTo,omitempty"`
	Content      string   `json:"content"`
}

func (note *Note) Wrap(username string) *Activity {
	return &Activity{
		Context:   ActivityStreams,
		Type:      TypeCreate,
		ID:        fmt.Sprintf("%s://%s/u/%s/posts/%s", config.PROTOCOL, config.DOMAIN, username, uuid.New().String()),
		To:        note.To,
		Actor:     fmt.Sprintf("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
		Published: time.Now(),
		Object:    note,
	}
}

func UnmarshalNote(data []byte) (Note, error) {
	var note Note
	err := json.Unmarshal(data, &note)
	return note, err
}

func (note *Note) Marshal() ([]byte, error) {
	return json.Marshal(note)
}
