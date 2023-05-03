package activitypub

import "encoding/json"

type Followers struct {
	Context      string      `json:"@context"`
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	TotalItems   int         `json:"totalItems"`
	OrderedItems interface{} `json:"orderedItems,omitempty"`
}

func UnmarshalFollowers(data []byte) (Followers, error) {
	var followers Followers
	err := json.Unmarshal(data, &followers)
	return followers, err
}

func (followers *Followers) Marshal() ([]byte, error) {
	return json.Marshal(followers)
}
