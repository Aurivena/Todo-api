package persistence

import (
	"Todo/models"
	"github.com/jmoiron/sqlx"
)

type Todo interface {
	Create(input *models.TodoInput, session string) (*models.TodoOutput, error)
	Delete(id int, session string) error
	Update(input *models.TodoInput, id int, session string) error
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
