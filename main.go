package main

import (
	"fmt"
	"os"
	"strings"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pakawatkung/go-web-fiber/handler"
	"github.com/pakawatkung/go-web-fiber/repository"
	"github.com/pakawatkung/go-web-fiber/service"
	"github.com/spf13/viper"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	viper.SetConfigName("vipers")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	os.Remove(viper.GetString("db.file"))
	db, err := sqlx.Open(viper.GetString("db.driver"), viper.GetString("db.file"))
	if err != nil {
		println("Error Cannot Open the Database...")
		panic(err)
	}
	defer db.Close()

	fmt.Println(viper.GetString("key.jwt"))

	userRepo := repository.NewUserdb(db)
	userSrv := service.NewUserService(userRepo)
	userApi := handler.NewUserHandler(userSrv)

	err = userSrv.ServiceCreateDB()
	if err != nil {
		fmt.Println("46 main line error")
		panic(err)
	}

	app := fiber.New()
	app.Use("/user", jwtware.New(jwtware.Config{ 
		SigningMethod: "HS256",
		SigningKey: []byte(viper.GetString("key.jwt")),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))
	app.Get("/user", userApi.GetAllData)
	app.Post("/signup/:username/:password", userApi.UserSignUp)
	app.Post("/signin/:username/:password", userApi.UserSignIn)

	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}