package createUserController

import (
	"github.com/gofiber/fiber/v2"
	"goTemplate/modules/user/services/createUser"
	"regexp"
)

func Handle(ctx *fiber.Ctx) error {

	body := new(input)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if body.Email == "" || body.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email or Password is empty",
		})
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(body.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email address",
		})
	}

	if len(body.Password) < 8 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password must be at least 8 characters",
		})
	}

	createdUser, err := createUser.CreatUser(&createUser.DTO{Email: body.Email, Password: body.Password})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": createdUser,
	})
}
