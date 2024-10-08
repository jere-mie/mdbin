package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"math/rand"

	"bytes"

	"github.com/joho/godotenv"
	"github.com/yuin/goldmark"
)

//go:embed templates
var TemplateAssets embed.FS

//go:embed static/*
var StaticAssets embed.FS

//go:embed version.txt
var Version string

type Snippet struct {
	gorm.Model
	ShortID string `gorm:"unique;not null"`
	Content string `gorm:"not null"`
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// connect to and init db
	db, err := gorm.Open(sqlite.Open("mdbin.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db: %v\n", err)
	}
	if err := db.AutoMigrate(&Snippet{}); err != nil {
		log.Fatalf("Failed to migrate db: %v\n", err)
	}

	// embed /templates directory into binary
	engine := django.NewPathForwardingFileSystem(http.FS(TemplateAssets), "/templates", ".html")
	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page for 404 or 500
			if code == 404 || code == 500 {
				err = ctx.Render(fmt.Sprintf("%d", code), fiber.Map{})
			}

			if err != nil {
				// In case ctx.Render fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			// Return from handler
			return nil
		},
	})
	app.Use(logger.New())

	// embed /static assets into binary
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(StaticAssets),
		PathPrefix: "static",
		Browse:     false,
	}))

	PORT := GetEnv("PORT", "3000")

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendFile("./static/favicon.ico")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/policy", func(c *fiber.Ctx) error {
		return c.Render("policy", fiber.Map{})
	})

	app.Get("/abuse", func(c *fiber.Ctx) error {
		return c.Render("abuse", fiber.Map{})
	})

	app.Get("/s/:shortID", func(c *fiber.Ctx) error {
		shortID := c.Params("shortID")
		var snippet Snippet
		if err := db.First(&snippet, "short_id = ?", shortID).Error; err != nil {
			return c.Render("404", fiber.Map{})
		}

		// render Content using goldmark
		var buf bytes.Buffer
		if err := goldmark.New().Convert([]byte(snippet.Content), &buf); err != nil {
			return c.Render("500", fiber.Map{})
		}

		return c.Render("snippet", fiber.Map{
			"Content": buf.String(),
		})
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		content := c.FormValue("content")
		if content == "" {
			return c.Status(400).SendString("Content cannot be empty")
		}
		shortID := GenerateShortID()
		if err := db.Create(&Snippet{
			ShortID: shortID,
			Content: content,
		}).Error; err != nil {
			return c.Render("500", fiber.Map{})
		}
		return c.Redirect("/s/" + shortID)
	})

	log.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", PORT)))
}

// Helper function to get environment variables with a fallback value
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// Generate a random base32 string of length 8 for creating "short links" to files
// 32^8 > 1 trillion combinations, which should be plenty for our use case
func GenerateShortID() string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	length := 8
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
