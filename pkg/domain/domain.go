package domain

import (
	"Todo/models"
	"Todo/pkg/persistence"
)

type Todo interface {
	Create(input *models.TodoInput, session string) (*models.TodoOutput, error)
	Delete(id int, session string) error
	Update(input *models.TodoInput, id int, session string) error
	UpdateDone(input *models.DoneChange, id int, session string) error
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
