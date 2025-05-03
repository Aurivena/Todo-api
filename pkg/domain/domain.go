package domain

import (
	"Todo/models"
	"Todo/pkg/persistence"
)

type Todo interface {
	Create(input *models.TodoInput, session string) (*models.TodoOutput, error)
	Delete(id int) error
	Update(input *models.TodoInput, id int) error
	UpdateDone(input *models.DoneChange, id int) error
	Get(session string) ([]models.TodoOutput, error)
}

type Domain struct {
	Todo
}

func NewDomain(persistence *persistence.Persistence, cfg *models.ConfigService) *Domain {
	return &Domain{
		Todo: NewTodoDomain(persistence),
	}
}
