package model

import "time"

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}