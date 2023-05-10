package domain

import (
	"errors"
	"go-queue/publisher/config"
	"log"
	"strconv"
)

type Service interface {
	AddQueue(number int) (int, error)
	SendListOfQueue()
}

type service struct {
	repository Repository
}

func (s service) SendListOfQueue() {
	emitter, err := config.CreateEventEmitter()
	if err != nil {
		return
	}
	list, err := s.repository.GetAllListOfNumber()
	err = emitter.Push(list, "printqueue")
}

func (s service) AddQueue(number int) (int, error) {
	save, err := s.repository.Save(number)
	if err != nil {
		return 0, errors.New("service : error occured when added number qo queue")
	}
	log.Println("successfully add " + strconv.Itoa(save) + " to queue")
	return save, nil
}

func NewService(repository Repository) Service {
	return &service{repository}
}
