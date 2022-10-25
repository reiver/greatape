package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
} // @name LoginRequest

type LoginResponse struct {
	User User `json:"user"`
	Auth Auth `json:"auth"`
} // @name LoginResponse

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
} // @name SignupRequest

type SignupResponse struct {
	Code string `json:"code"`
} // @name SignupResponse

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
} // @name VerifyRequest

type VerifyResponse struct {
	User User `json:"user"`
	Auth Auth `json:"auth"`
} // @name VerifyResponse

type User struct {
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
	PublicKey      string `json:"public_key"`
	Actor          string `json:"actor"`
	Webfinger      string `json:"webfinger"`
	PrivateProfile bool   `json:"private_profile"`
} // @name User

type Auth struct {
	Token string `json:"token"`
} // @name Auth
