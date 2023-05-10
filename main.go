package main

import (
	"github.com/gofiber/fiber/v2"
	"go-queue/publisher/domain"
	"log"
	"time"
)

func main() {
	list := []int{4, 1, 21, 1, 1}
	repository := domain.NewRepository(&list)
	service := domain.NewService(repository)
	controller := domain.NewController(service)
	queue, err := service.AddQueue(12)
	if err != nil {
		log.Println(err)
	}
	log.Println(queue)
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
