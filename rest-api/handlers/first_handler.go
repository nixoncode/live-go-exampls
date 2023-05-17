// Package handlers
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 13:27
package handlers

import "github.com/gofiber/fiber/v2"

func TestHandler(c *fiber.Ctx) error {
	return c.SendString("testing")
}
