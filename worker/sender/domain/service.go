package domain

import (
	"errors"
	"go-queue/config"
	"go-queue/worker/sender"
	"log"
)

type Service interface {
	AddQueue(patient sender.Patient) (sender.Patient, error)
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
	patient, err := s.repository.GetAllListOfNumber()
	err = emitter.Push(patient, "printqueue")
}

func (s service) AddQueue(patient sender.Patient) (sender.Patient, error) {
	save, err := s.repository.Save(patient)
	if err != nil {
		return save, errors.New("service : error occured when added number qo queue")
	}
	log.Println("successfully add %v to queue", save)
	return save, nil
}

func NewService(repository Repository) Service {
	return &service{repository}
}
