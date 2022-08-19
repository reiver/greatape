package activitypub

import "encoding/json"

type OrderedCollection struct {
	Context      string      `json:"@context"`
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type,omitempty"`
	TotalItems   int         `json:"totalItems"`
	OrderedItems interface{} `json:"orderedItems,omitempty"`
}

func NewOrderedCollection(id string, items interface{}, length int) *OrderedCollection {
	return &OrderedCollection{
		Context:      "https://www.w3.org/ns/activitystreams",
		ID:           id,
		Type:         "OrderedCollection",
		TotalItems:   length,
		OrderedItems: items,
	}
}

func UnmarshalOrderedCollection(data []byte) (OrderedCollection, error) {
	var o OrderedCollection
	err := json.Unmarshal(data, &o)
	return o, err
}

func (o *OrderedCollection) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
