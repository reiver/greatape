package types

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

type SignupDTO struct {
	LoginDTO
	Username string `json:"username" validate:"username"`
}

type ProfileDTO struct {
	DisplayName    string `json:"display_name"`
	Bio            string `json:"bio"`
	Github         string `json:"github"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	PrivateProfile bool   `json:"private_profile"`
}

type UserResponse struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	DisplayName    string `json:"display_name"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	Bio            string `json:"bio"`
	Github         string `json:"github"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	ApiKey         string `json:"api_key"`
	PublicKey      string `json:"publicKey"`
	Actor          string `json:"actor,omitempty"`
	Webfinger      string `json:"webfinger,omitempty"`
	PrivateProfile bool   `json:"private_profile"`
}

type KeyResponse struct {
	PrivateKey string `json:"privateKey"`
}

type AccessResponse struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
