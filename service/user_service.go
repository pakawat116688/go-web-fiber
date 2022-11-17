package service

import (
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pakawatkung/go-web-fiber/repository"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService  {
	return userService{userRepo: repo}
}

func (s userService) ServiceCreateDB() error {
	err := s.userRepo.CreateDB()
	if err != nil {
		return err
	}
	return nil
}

func (s userService) ServiceSignUp(user UserRequre) (*UserRequre, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}


	data := repository.UserCreate{
		Username: user.Username,
		Password: string(password),
	}

	result, err := s.userRepo.SignUp(data)
	if err != nil {
		return nil, err
	}

	dataReponse := UserRequre{
		Username: result.Username,
		Password: result.Password,
	}

	return &dataReponse, nil
}

func (s userService) ServiceGetAll() ([]UserRequre, error) {

	result, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	data := []UserRequre{}
	for _, user := range result {
		userReq := UserRequre{
			Username: user.Username,
			Password: user.Password,
		}
		data = append(data, userReq)
	}
	return data, nil
}

func (s userService) ServiceGetId(user UserRequre) (*UserJwtToken, error)  {

	res, err := s.userRepo.GetId(user.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	cliams := jwt.StandardClaims{
		Issuer: strconv.Itoa(res.Id),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(viper.GetString("key.jwt")))
	if err != nil {
		return nil, err
	}

	users := UserJwtToken{
		Username: res.Username,
		Password: res.Password,
		JwtToken: token,
	}

	return &users, nil
}