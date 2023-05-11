package domain

import (
	"github.com/jmoiron/sqlx"
	"go-queue/worker/sender"
)

type Repository interface {
	Save(patient sender.Patient) (sender.Patient, error)
	GetAllListOfNumber() (sender.Patient, error)
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) GetAllListOfNumber() (sender.Patient, error) {
	var patient sender.Patient
	err := r.db.Get(&patient, "SELECT * FROM patients WHERE is_publish = false LIMIT 1")
	if err != nil {
		return patient, err
	}

	// Update the retrieved patient's is_publish to true
	_, err = r.db.Exec("UPDATE patients SET is_publish = true WHERE identifier_number = $1", patient.IdentifierNumber)
	if err != nil {
		return patient, err
	}

	return patient, nil
}

func (r *repository) Save(patient sender.Patient) (sender.Patient, error) {
	_, err := r.db.NamedExec("INSERT INTO patients (name, queue_number, identifier_number, is_publish) VALUES (:name, :queue_number, :identifier_number, :is_publish)", patient)
	if err != nil {
		return sender.Patient{}, err
	}
	return patient, nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
