package main

import (
	"log"

	"pranshul9987/Slack-Snip/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/app", routes.RootRoute)
	app.Post("/activity-route", routes.ActivityRoute)

	log.Fatal(app.Listen(":3000"))
}
