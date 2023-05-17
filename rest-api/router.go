// Package main
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 13:24
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixoncode/live-go-exampls/handlers"
)

func generateApp() *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	libraryGroup := app.Group("/library")
	libraryGroup.Get("/", handlers.TestHandler)

	return app
}
