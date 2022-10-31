package entity

type entity struct {
	IdField      int64  `json:"id"`
	PayloadField string `json:"payload"`
}

func (entity *entity) Id() int64 {
	return entity.IdField
}

func (entity *entity) Payload() string {
	if entity.PayloadField == "" {
		entity.PayloadField = "{}"
	}

	return entity.PayloadField
}

func (entity *entity) SetPayload(payload string) {
	entity.PayloadField = payload
}
