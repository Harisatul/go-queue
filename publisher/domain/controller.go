package domain

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (r *controller) AddQueue(ctx *fiber.Ctx) error {
	params := ctx.Query("number")
	atoi, err2 := strconv.Atoi(params)
	if err2 != nil {
		log.Println(err2)
		return err2
	}
	queue, err2 := r.service.AddQueue(atoi)
	log.Println(queue)
	return err2
}
