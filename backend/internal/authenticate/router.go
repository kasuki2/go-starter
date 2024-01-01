package authenticate

import (
	"github.com/gofiber/fiber/v2"
	
)

func AddAuthenticateRoutes(api fiber.Router, controller *AuthenticateController) {

	
	api.Get("/register", controller.register)

}