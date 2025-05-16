package models

import "time"

type TodoInput struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Priority      string    `json:"priority"`
	DateCompleted time.Time `json:"dateCompleted"`
	Done          bool      `json:"done" db:"done"`
}

type TodoOutput struct {
	ID            int       `json:"id" db:"id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	Priority      string    `json:"priority" db:"priority"`
	DateCreated   time.Time `json:"date" db:"date_created"`
	DateCompleted time.Time `json:"dateCompleted" db:"date_completed"`
	Done          bool      `json:"done" db:"done"`
}
