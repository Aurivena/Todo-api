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
	err := p.db.Get(&out, `INSERT INTO "Todo" (session, title, description, priority,date_completed) 
		VALUES ($1, $2, $3, $4,$5) RETURNING id, title, description, priority, date_created,date_completed, done`,
		session, input.Title, input.Description, input.Priority, input.DateCompleted)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (p *TodoPersistence) Get(session string) ([]models.TodoOutput, error) {
	var out []models.TodoOutput

	err := p.db.Select(&out, `SELECT id,title,description,priority,date_created, date_completed,done FROM "Todo" WHERE session = $1`, session)
	if err != nil {
		return nil, err
	}

	return out, err
}

func (p *TodoPersistence) Delete(id int, session string) error {
	_, err := p.db.Exec(`DELETE FROM "Todo" WHERE id = $1 AND session = $2`, id, session)
	if err != nil {
		return err
	}
	return nil
}

func (p *TodoPersistence) Update(input *models.TodoInput, id int, session string) error {
	_, err := p.db.Exec(`UPDATE "Todo" SET title=$1, description=$2, priority=$3, date_completed = $4, done = $5 WHERE id=$6 AND session = $7`, input.Title, input.Description, input.Priority, input.DateCompleted, input.Done, id, session)
	if err != nil {
		return err
	}

	return nil
}
