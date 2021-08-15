package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber Rest API",
	})

	const port = ":3000"

	app.Static("/static", "./files")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]interface{}{
			"status":       "success",
			"errorMessage": nil,
			"data": map[string]string{
				"message": "Hello, World!",
			},
		})
	})

	app.Get("/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.JSON(map[string]interface{}{
			"status":       "success",
			"errorMessage": nil,
			"data": map[string]string{
				"message": fmt.Sprintf("Hello, %s!", name),
			},
		})
	})

	err := app.Listen(port)
	if err != nil {
		return
	}
}
