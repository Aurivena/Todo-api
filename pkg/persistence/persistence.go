package persistence

import (
	"Todo/models"
	"github.com/jmoiron/sqlx"
)

type Todo interface {
	Create(input *models.TodoInput, session string) (*models.TodoOutput, error)
	Delete(id int) error
	Update(input *models.TodoInput, id int) error
	UpdateDone(input *models.DoneChange, id int) error
	Get(session string) ([]models.TodoOutput, error)
}

type Persistence struct {
	Todo
}

type Sources struct {
	BusinessDB *sqlx.DB
}

func NewPersistence(sources *Sources) *Persistence {
	return &Persistence{
		Todo: NewTodoPersistence(sources.BusinessDB),
	}
}
