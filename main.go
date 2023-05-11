package main

import (
	"github.com/gofiber/fiber/v2"
	"go-queue/config"
	domain2 "go-queue/worker/sender/domain"
	"log"
	"time"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to initialize the database: %v", err)
	}
	defer db.Close()
	repository := domain2.NewRepository(db)
	service := domain2.NewService(repository)
	controller := domain2.NewController(service)

	if err != nil {
		log.Println(err)
	}
	//log.Println(queue)
	app := fiber.New()

	api := app.Group("/api")

	api.Post("addqueue", controller.AddQueue)

	ticker := time.Tick(5 * time.Minute)

	// Start a goroutine to execute the method periodically
	go func() {
		for range ticker {
			service.SendListOfQueue()
		}
	}()

	err = app.Listen(":8080")
	if err != nil {
		log.Println(err)
	}

	// Keep the main goroutine running
	select {}

}
