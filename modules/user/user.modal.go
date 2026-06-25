package user

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
