// Package rest_api
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 12:00
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
