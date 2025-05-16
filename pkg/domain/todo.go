package domain

import (
	"Todo/models"
	"Todo/pkg/persistence"
)

type TodoDomain struct {
	pers *persistence.Persistence
}

func (t *TodoDomain) Create(input *models.TodoInput, session string) (*models.TodoOutput, error) {
	return t.pers.Create(input, session)
}

func (t *TodoDomain) Delete(id int, session string) error {
	return t.pers.Delete(id, session)
}

func (t *TodoDomain) Update(input *models.TodoInput, id int, session string) error {
	return t.pers.Update(input, id, session)
}

func (t *TodoDomain) Get(session string) ([]models.TodoOutput, error) {
	return t.pers.Get(session)
}

func NewTodoDomain(pers *persistence.Persistence) *TodoDomain {
	return &TodoDomain{pers: pers}
}
