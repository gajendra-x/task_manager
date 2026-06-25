package todo

import (
	"task_manager/config/db"
)

func CreateTodoService(todo Todo) (Todo, error) {
	var data Todo

	query := `INSERT INTO todos (title, user_id) VALUES ($1, $2)
			  RETURNING id, title, user_id, is_completed, created_at`
	err := db.DB.QueryRow(
		db.CTX,
		query,
		todo.Title,
		todo.UserId,
	).Scan(
		&data.ID,
		&data.Title,
		&data.UserId,
		&data.IsCompleted,
		&data.CreatedAt,
	)

	if err != nil {
		return Todo{}, err
	}

	return data, nil
}

func GetUsersTodoService(userID int64, limit, page int) ([]Todo, error) {
	offset := (page - 1) * limit

	query := `SELECT id, title, user_id, is_completed, created_at
              FROM todos
              WHERE user_id = $1
              ORDER BY created_at DESC
              LIMIT $2 OFFSET $3`
	var data []Todo

	rows, err := db.DB.Query(db.CTX, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.UserId,
			&todo.IsCompleted,
			&todo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, todo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return data, nil
}
