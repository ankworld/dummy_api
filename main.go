package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Pivot struct {
	X int64   `json:"x"`
	Y float64 `json:"y"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/info/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"result":   "success",
			"mac":      "24:6F:28:4D:2F:35",
			"firmware": "0.4",
			"ssid":     "apesp",
			"sensor":   "['BNO055','sensorEx']",
			"opts":     "['opt1','opt2']",
		})
	})

	app.Post("/command/", func(c *fiber.Ctx) error {
		now := time.Now()
		nanos := now.UnixNano()
		min := 10.0
		max := 30.0
		rnumber := min + rand.Float64()*(max-min)
		body := c.Body()
		return c.JSON(fiber.Map{
			"result": "success",
			"data": []Pivot{
				{
					X: nanos,
					Y: rnumber,
				},
			},
			"body": body,
		})
	})

	app.Listen(":5000")
}
