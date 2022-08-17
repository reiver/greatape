package activitypub

import "time"

const Public = "https://www.w3.org/ns/activitystreams#Public"

type Activity struct {
	Context   interface{} `json:"@context"`
	ID        string      `json:"id,omitempty"`
	Type      string      `json:"type,omitempty"`
	Actor     string      `json:"actor,omitempty"`
	Object    interface{} `json:"object,omitempty"`
	From      string      `json:"from,omitempty"`
	To        interface{} `json:"to,omitempty"`
	InReplyTo string      `json:"inReplyTo,omitempty"`
	Content   string      `json:"content,omitempty"`
	Published time.Time   `json:"published,omitempty"`
}
