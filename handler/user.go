package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/go-web-fiber/service"
)

type userHandler struct {
	userRest service.UserService
}

func NewUserHandler(userRest service.UserService) userHandler {
	return userHandler{userRest: userRest}
}

func (h userHandler) CreateTable(c *fiber.Ctx) error {

	err := h.userRest.ServiceCreateDB()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON("{status: true}")
}

func (h userHandler) UserSignUp(c *fiber.Ctx) error {

	username := c.Params("username")
	password := c.Params("password")
	data := service.UserRequre{
		Username: username,
		Password: password,
	}

	if username == "" || password == "" {
		return fiber.ErrUnprocessableEntity
	}

	res, err := h.userRest.ServiceSignUp(data)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h userHandler) UserSignIn(c *fiber.Ctx) error {

	user := service.UserRequre{
		Username:  c.Params("username"),
		Password: c.Params("password"),
	}
	if user.Username == "" || user.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	res, err := h.userRest.ServiceGetId(user)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h userHandler) GetAllData(c *fiber.Ctx) error {

	res, err := h.userRest.ServiceGetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
