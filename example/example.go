package example

import (
	"fmt"
	"strconv"	
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
)

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func Exam()  {
	
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	//Midleware
	app.Use("/error", func(c *fiber.Ctx) error {
		c.Locals("name", "bond")
		fmt.Println("before")
		err := c.Next()
		fmt.Println("After")
		return err
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))

	//GET
	app.Get("/employee", func(c *fiber.Ctx) error {
		return c.SendString("get employee")
	})

	//POST
	app.Post("/employee", func(c *fiber.Ctx) error {
		return c.SendString("post employee")
	})

	//Parameter, ? is optional parameter
	app.Get("/employee/:id/:name?", func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := c.Params("name")
		return c.SendString("id: " + id + " " + name)
	})

	//Parameter Int
	app.Get("/employees/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString("/Param Int : " + strconv.Itoa(id))
	})

	//Query
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		return c.SendString("query " + name)
	})

	//Query Parser
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	//Wildcards
	app.Get("/w/*", func(c *fiber.Ctx) error {
		war := c.Params("*")
		return c.SendString(war)
	})

	//Error
	app.Get("/error", func(c *fiber.Ctx) error {
		name := c.Locals("name")
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("test error path %v", name))
	})

	//Group, และมันทำงานเป็น midleware ได้	
	v1 := app.Group("v1", func(c *fiber.Ctx) error {
		c.Set("version","v1")
		return c.Next()
	})
	v1.Get("/app", func(c *fiber.Ctx) error {
		return c.SendString("v1 app")
	})

	v2 := app.Group("v2", func(c *fiber.Ctx) error {
		c.Set("version","v2")
		return c.Next()
	})
	v2.Get("/app", func(c *fiber.Ctx) error {
		return c.SendString("v2 app")
	})

	//Mount
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("login")
	})

	app.Mount("/user", userApp)

	//Server
	app.Server().MaxConnsPerIP = 10
	app.Get("/server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 30)
		return c.SendString("server")
	})

	//Environment
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseUrl": c.BaseURL(),
			"Hostname": c.Hostname(),
			"IP": c.IP(),
			"IPS": c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path": c.Path(),
			"Protocal": c.Protocol(),
			"Subdomains": c.Subdomains(),
		})
	})

	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Println("Is json ",c.Is("json"))
		fmt.Println(string(c.Body()))
		data := map[string]interface{}{}
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}
		return c.SendString("body path")
	})

	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}