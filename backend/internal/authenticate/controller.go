package authenticate


import (
	"github.com/gofiber/fiber/v2"
	
)

type AuthenticateController struct {
	storage *AuthenticateStorage
}

func NewAuthenticateController(storage *AuthenticateStorage) *AuthenticateController {
	return &AuthenticateController{
		storage: storage,
	}
}


func (a *AuthenticateController) register (c *fiber.Ctx) error {

	c.Accepts("application/json")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "this is ok",
	})
	
}
