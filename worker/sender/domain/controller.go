package domain

import (
	"github.com/gofiber/fiber/v2"
	"go-queue/worker/sender"
	"log"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (r *controller) AddQueue(ctx *fiber.Ctx) error {
	var patient sender.Patient

	// Parse the request body and map it to the Patient struct
	err := ctx.BodyParser(&patient)
	if err != nil {
		log.Println(err)
		return err
	}

	queue, err := r.service.AddQueue(patient)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(queue)
	return nil
}
