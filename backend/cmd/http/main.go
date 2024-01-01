package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kasuki2/go-starter.git/config"
	"github.com/kasuki2/go-starter.git/internal/authenticate"
	"github.com/kasuki2/go-starter.git/pkg/shutdown"
)

func main() {
	var exitCode int

	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()

	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	cleanup, err := run(env)

	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()

}

func run(env config.EnvVars) (func(), error) {
	
	app, cleanup, err := buildServer()
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen(env.Port)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer() (*fiber.App, func(), error) {

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})
	

	api := app.Group("/api")

	authStore := authenticate.NewAuthenticateStorage()
	authController := authenticate.NewAuthenticateController(authStore)
	authenticate.AddAuthenticateRoutes(api, authController)
	return app, func() {
		//storage.CloseMongo(db)
	}, nil


}