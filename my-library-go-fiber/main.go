package main

import (
	"LearnGoLanguage/my-library-go-fiber/pkg/configs"
	"LearnGoLanguage/my-library-go-fiber/pkg/middleware"
	"LearnGoLanguage/my-library-go-fiber/pkg/routes"
	utils2 "LearnGoLanguage/my-library-go-fiber/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils2.StartServerWithGracefulShutdown(app)
}
