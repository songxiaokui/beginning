package model

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Notes       string    `json:"notes"`
	RemindAt    time.Time `json:"remind_at"`
	CreatedAt   time.Time `json:"created_at"`
	IsCompleted bool      `json:"completed"`
}
