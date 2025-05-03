package persistence

import (
	"Todo/models"
	"github.com/jmoiron/sqlx"
)

type TodoPersistence struct {
	db *sqlx.DB
}

func NewTodoPersistence(db *sqlx.DB) *TodoPersistence {
	return &TodoPersistence{db: db}
}

func (p *TodoPersistence) Create(input *models.TodoInput, session string) (*models.TodoOutput, error) {
	var out models.TodoOutput
	err := p.db.Get(&out, `INSERT INTO "Todo" (session, title, description, priority) 
		VALUES ($1, $2, $3, $4) RETURNING id, title, description, priority, date, done`,
		session, input.Title, input.Description, input.Priority)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (p *TodoPersistence) Get(session string) ([]models.TodoOutput, error) {
	var out []models.TodoOutput

	err := p.db.Select(&out, `SELECT id,title,description,priority,date,done FROM "Todo" WHERE session = $1`, session)
	if err != nil {
		return nil, err
	}

	return out, err
}

func (p *TodoPersistence) Delete(id int) error {
	_, err := p.db.Exec(`DELETE FROM "Todo" WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *TodoPersistence) Update(input *models.TodoInput, id int) error {
	_, err := p.db.Exec(`UPDATE "Todo" SET title=$1, description=$2, priority=$3 WHERE id=$4`, input.Title, input.Description, input.Priority, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *TodoPersistence) UpdateDone(input *models.DoneChange, id int) error {
	_, err := p.db.Exec(`UPDATE "Todo" SET done = $1 WHERE id = $2`, input.Done, id)
	if err != nil {
		return err
	}
	return nil
}
