package action

import (
	"Todo/models"
	"errors"
	"github.com/Aurivena/answer"
)

var (
	errorPriorityIsInvalid = errors.New("priority is invalid")
)

func (a *Action) Create(input *models.TodoInput, session string) (*models.TodoOutput, answer.ErrorCode) {
	out, err := a.domains.Create(input, session)
	if err != nil {
		if errors.As(err, &errorPriorityIsInvalid) {
			return nil, answer.BadRequest
		}
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

func (a *Action) Delete(id int, session string) answer.ErrorCode {
	if err := a.domains.Delete(id, session); err != nil {
		return answer.InternalServerError
	}
	return answer.NoContent
}

func (a *Action) Update(input *models.TodoInput, id int, session string) answer.ErrorCode {
	if err := a.domains.Update(input, id, session); err != nil {
		if errors.As(err, &errorPriorityIsInvalid) {
			return answer.BadRequest
		}
		return answer.InternalServerError
	}
	return answer.NoContent
}

func (a *Action) UpdateDone(input *models.DoneChange, id int, session string) answer.ErrorCode {
	if err := a.domains.UpdateDone(input, id, session); err != nil {
		return answer.InternalServerError
	}
	return answer.NoContent
}
