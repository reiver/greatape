package activitypub

import (
	"encoding/json"
	"fmt"
	"time"
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

func NewNote(from, to, content string) *Note {
	return &Note{
		Context:      ActivityStreams,
		To:           []string{to},
		Content:      content,
		Type:         TypeNote,
		AttributedTo: from,
	}
}

func NewPublicNote(from, content string) *Note {
	return NewNote(from, Public, content)
}

func (note *Note) Wrap(username, publicUrl, uniqueIdentifier string) *Activity {
	return &Activity{
		Context:   ActivityStreams,
		Type:      TypeCreate,
		ID:        fmt.Sprintf("%s/u/%s/posts/%s", publicUrl, username, uniqueIdentifier),
		To:        note.To,
		Actor:     fmt.Sprintf("%s/u/%s", publicUrl, username),
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
