package activitypub

import "encoding/json"

type Object struct {
	Context interface{} `json:"@context"`
	Type    string      `json:"type"`
}

func UnmarshalObject(data []byte) (Object, error) {
	var object Object
	err := json.Unmarshal(data, &object)
	return object, err
}

func (object *Object) Marshal() ([]byte, error) {
	return json.Marshal(object)
}
