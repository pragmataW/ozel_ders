package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userBody struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Insert(data interface{}) error {
	fmt.Printf("%v inserted to database", data)
	return nil
}



var validate = validator.New()

func main() {
	app := fiber.New()

	app.Post("/save-user", func(c *fiber.Ctx) error {
		var body userBody

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "body can not parsed",
			})
		}

		if err := validate.Struct(body); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "body field are not true",
			})
		}

		if err := Insert(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "body can not inserted to database",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "user created",
		})
	})

	app.Listen(":8080")
}
