package main

import (
	"github.com/Vansh3140/credit-card-validator/routes" // Importing custom routes for handling requests
	"github.com/gofiber/fiber/v2"                       // Importing Fiber framework for creating web servers
	"log"                                               // Importing log package for logging messages
	"os"                                                // Importing os package to handle operating system functions
	"os/signal"                                         // Importing os/signal to handle system signals
	"syscall"                                           // Importing syscall for signal handling constants
)

func main() {
	// Create a new Fiber app instance
	app := fiber.New()

	// Define a POST route to check mail using the custom route handler
	app.Post("/api/v1/check", routes.CheckMail)

	// Create a channel to listen for OS signals for graceful shutdown
	stop := make(chan os.Signal, 1)

	// Notify the stop channel for specific signals (e.g., SIGINT, SIGTERM, SIGTSTP)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	// Start the server in a separate goroutine to listen for incoming HTTP requests
	go func() {
		// Log an error and exit if the server fails to start
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Block the main goroutine until a signal is received
	<-stop
	log.Println("Received shutdown signal, shutting down...")

	// Gracefully shut down the server when a signal is received
	// This allows for cleanup of resources before the application exits
	if err := app.Shutdown(); err != nil {
		// Log any errors that occur during the shutdown process
		log.Printf("Error shutting down: %v\n", err)
	} else {
		// Log a successful server shutdown message
		log.Println("Server shut down gracefully")
	}
}
