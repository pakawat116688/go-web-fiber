package service

type UserRequre struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserJwtToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
	JwtToken string `json:"jwttoken"`
}

type UserService interface {
	ServiceCreateDB() error
	ServiceSignUp(UserRequre) (*UserRequre, error)
	ServiceGetAll() ([]UserRequre, error)
	ServiceGetId(UserRequre) (*UserJwtToken, error)
}