package types

type FollowerDTO struct {
	Target string `json:"target" validate:"required"`
	Handle string `json:"handle" validate:"required"`
}

type FollowerResponse struct {
	Target string `json:"target"`
	Handle string `json:"handle"`
}
