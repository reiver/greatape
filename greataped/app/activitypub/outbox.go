package activitypub

import "encoding/json"

type Outbox struct {
	Context      string      `json:"@context"`
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	TotalItems   int         `json:"totalItems"`
	OrderedItems interface{} `json:"orderedItems,omitempty"`
}

func UnmarshalOutbox(data []byte) (Outbox, error) {
	var o Outbox
	err := json.Unmarshal(data, &o)
	return o, err
}

func (o *Outbox) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
