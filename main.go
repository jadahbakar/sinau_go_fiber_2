package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/sinau_go_fiber_2/database"
)

func main() {
	// Connect to database
	_, err := database.Connect()

	if err != nil {
		fmt.Printf("Error Database: %v", err)
	}

	configFiber := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	}

	app := fiber.New(configFiber) // call the New() method - used to instantiate a new Fiber App
	// app.Use(middleware.Logger())

	app.Listen(":3000") // listen/Serve the new Fiber app on port 3000

}
