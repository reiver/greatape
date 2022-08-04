package types

type FollowerDTO struct {
	Handle string `json:"handle" validate:"required"`
}

type FollowerResponse struct {
	Handler string `json:"handler"`
}
