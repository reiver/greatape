package entity

type entity struct {
	IdField        int64   `json:"id"`
	SortOrderField float32 `json:"sort_order"`
	PayloadField   string  `json:"payload"`
}

func (entity *entity) Id() int64 {
	return entity.IdField
}

func (entity *entity) SortOrder() float32 {
	return entity.SortOrderField
}

func (entity *entity) SetSortOrder(sortOrder float32) {
	entity.SortOrderField = sortOrder
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
