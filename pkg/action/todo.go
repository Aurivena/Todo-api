package action

import (
	"Todo/models"
	"github.com/Aurivena/answer"
)

func (a *Action) Create(input *models.TodoInput, session string) (*models.TodoOutput, answer.ErrorCode) {
	out, err := a.domains.Create(input, session)
	if err != nil {
		return nil, answer.InternalServerError
	}
	return out, answer.OK
}

func (a *Action) Get(session string) ([]models.TodoOutput, answer.ErrorCode) {
	out, err := a.domains.Get(session)
	if err != nil {
		return nil, answer.InternalServerError
	}
	return out, answer.OK
}

func (a *Action) Delete(id int) answer.ErrorCode {
	if err := a.domains.Delete(id); err != nil {
		return answer.InternalServerError
	}
	return answer.NoContent
}

func (a *Action) Update(input *models.TodoInput, id int) answer.ErrorCode {
	if err := a.domains.Update(input, id); err != nil {
		return answer.InternalServerError
	}
	return answer.NoContent
}

func (a *Action) UpdateDone(input *models.DoneChange, id int) answer.ErrorCode {
	if err := a.domains.UpdateDone(input, id); err != nil {
		return answer.InternalServerError
	}
	return answer.NoContent
}
