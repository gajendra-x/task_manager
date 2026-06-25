package todo

import "time"

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	UserId      int64     `json:"user_id"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoPayload struct {
	Title  string `json:"title"`
	UserId int64  `json:"user_id"`
}
