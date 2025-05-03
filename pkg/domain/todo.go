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

func (t *TodoDomain) Delete(id int) error {
	return t.pers.Delete(id)
}

func (t *TodoDomain) Update(input *models.TodoInput, id int) error {
	return t.pers.Update(input, id)
}

func (t *TodoDomain) UpdateDone(input *models.DoneChange, id int) error {
	return t.pers.UpdateDone(input, id)
}

func (t *TodoDomain) Get(session string) ([]models.TodoOutput, error) {
	return t.pers.Get(session)
}

func NewTodoDomain(pers *persistence.Persistence) *TodoDomain {
	return &TodoDomain{pers: pers}
}
