package main

import (
	"log"
	"strings"

	"github.com/farzai/app/handlers"
	"github.com/farzai/app/pkg/validation"
	"github.com/farzai/app/services/google_recaptcha"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	Addr = ":3000"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: NewAppErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(NewRecaptchaMiddleware([]string{
		// "/api/v1/example-to-protect",
	}))

	handlers.ApplySystemRoutes(app)

	if err := app.Listen(Addr); err != nil {
		log.Fatal(err)
	}
}

// NewAppErrorHandler creates a new error handler for the app
func NewAppErrorHandler() func(*fiber.Ctx, error) error {
	return func(c *fiber.Ctx, err error) error {
		// Check instance of ValidationError
		if _, ok := err.(validation.ValidationError); ok {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
				"errors": err.(validation.ValidationError).Errors().All(),
			})
		}

		if _, ok := err.(*fiber.Error); ok {
			return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
}


const (
	// HeaderKeyName is the name of the header key to get the recaptcha token
	HeaderKeyName = "x-recaptcha"
)

// NewRecaptchaMiddleware creates a new recaptcha middleware
func NewRecaptchaMiddleware(verifyPaths []string) func(*fiber.Ctx) error {
	// Normalize the paths
	for i, path := range verifyPaths {
		verifyPaths[i] = "/" + strings.Trim(strings.TrimSpace(path), "/")
	}

	// Return the middleware
	return func(c *fiber.Ctx) error {
		// Check headers
		headers := c.GetReqHeaders()

		// Loop through the verify paths
		for _, path := range verifyPaths {
			// Check if the current path is not the same as the path to verify
			if ! strings.HasPrefix(path, "/") {
				continue
			}

			if _, ok := headers[HeaderKeyName]; !ok {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "Recaptcha is required",
				})
			}

			// Get the recaptcha token
			token := strings.TrimSpace(headers[HeaderKeyName][0])
			if token == "" {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "Recaptcha is required",
				})
			}

			googleRecaptcha := google_recaptcha.NewWithDefaultClient(google_recaptcha.RecaptchaConfig{
				SiteKey: "site-key",
				SecretKey: "secret-key",
			})

			if err := googleRecaptcha.VerifyV3(token); err != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "Recaptcha is invalid",
				})
			}

			return c.Next()
		}

		return c.Next()
	}
}
