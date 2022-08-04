package activitypub

type Activity struct {
	Context   string      `json:"@context"`
	ID        string      `json:"id,omitempty"`
	Type      string      `json:"type,omitempty"`
	Actor     string      `json:"actor,omitempty"`
	Object    interface{} `json:"object,omitempty"`
	From      string      `json:"from,omitempty"`
	To        interface{} `json:"to,omitempty"`
	InReplyTo string      `json:"inReplyTo,omitempty"`
	Content   string      `json:"content,omitempty"`
}
