package controller

import (
	"example/dto"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (ctrl *Controller) Login(c *fiber.Ctx) error{
	validate := validator.New()
	
	var req loginReq

	if err := c.BodyParser(&req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "body can not parsed",
		})
	}

	if err := validate.Struct(req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "body can not validate",
		})
	}

	user := dto.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	jwtKey, err := ctrl.service.Login(user)

	if err != nil{
		if err.Error() == "invalid credentials"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "wrong password",
			})
		}

		if err.Error() == "user not found" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name: "Authorization",
		Value: fmt.Sprintf("Bearer %v", jwtKey),
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login succesfully",
	})
}