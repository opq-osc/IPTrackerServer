package main

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var defaultImage []byte

func init() {
	defaultImage, _ = base64.StdEncoding.DecodeString(
		"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",
	)
}

type ResponseData struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func ensureKey(c *fiber.Ctx) error {
	key := c.Params("key")
	if key == "" {
		return c.JSON(&ResponseData{Code: 1, Msg: "??????key"})
	}
	return c.Next()
}

func info(c *fiber.Ctx) error {
	return c.JSON(&ResponseData{Code: 0, Msg: "ok", Result: GetItems(c.Params("key"))})
}

func check(c *fiber.Ctx) error {
	return c.JSON(&ResponseData{Code: 0, Msg: "ok", Result: CheckKey(c.Params("key"))})
}

func insert(c *fiber.Ctx) error {
	InsertItem(c.Params("key"), c.IP(), string(c.Request().Header.UserAgent()))

	if location := c.Query("r"); location != "" {
		return c.Redirect(location)
	}

	c.Set("Content-Type", "image/png")
	return c.Send(defaultImage)
}

func createServer(verbose bool) *fiber.App {
	app := fiber.New(fiber.Config{GETOnly: true, DisableStartupMessage: true})

	if verbose {
		app.Use(logger.New())
	}

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "image/png")
		return c.Send(defaultImage)
	})

	app.Get("/:key.info", info, ensureKey)

	app.Get("/:key.check", check, ensureKey)

	app.Get("/:key", insert, ensureKey)

	return app
}
