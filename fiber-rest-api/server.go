package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	const port = ":3000"

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]interface{}{
			"status":       "success",
			"errorMessage": nil,
			"data": map[string]string{
				"message": "Hello, World!",
			},
		})
	})

	err := app.Listen(port)
	if err != nil {
		return
	}
}