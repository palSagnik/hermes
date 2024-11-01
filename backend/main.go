package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/palSagnik/hermes/config"
	"github.com/palSagnik/hermes/database"
	"github.com/palSagnik/hermes/middleware"
	"github.com/palSagnik/hermes/router"
)
func init() {
	go middleware.CleanupVisitors()
}

func main() {

	// loop till database is initialised
	for {
		if err := database.ConnectDB(); err != nil {
			log.Warn(err)
			log.Info("waiting for 30 seconds before trying again")
			time.Sleep(time.Second * 30)
			continue
		}
		break
	}

	// creating tables
	err := database.MigrateUp()
	if err != nil {
		log.Fatal(err)
	}

	// Initialising *fiber.App
	app := fiber.New()
	app.Use(recover.New())
	router.SetUpRoutes(app)

	log.Fatal(app.Listen(config.APP_PORT))
}
