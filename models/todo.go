package models

import "time"

type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type TodoOutput struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Priority    string    `json:"priority" db:"priority"`
	Date        time.Time `json:"date" db:"date"`
	Done        bool      `json:"done" db:"done"`
}

type DoneChange struct {
	Done bool `json:"done"`
}
