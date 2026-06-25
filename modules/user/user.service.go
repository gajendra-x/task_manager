package user

import config "task_manager/config/db"

func CreateUSerService(payload CreateUserPayload) (User, error) {
	var userData User

	query := `INSERT INTO users (name, email) VALUES ($1, $2)
			  RETURNING id, name, email, created_at
			`
	err := config.DB.QueryRow(
		config.CTX,
		query,
		payload.Name,
		payload.Email).Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return userData, nil
}

func GetUSerService(email string) (User, error) {
	var userData User

	query := `SELECT * FROM users WHERE email = $1`
	err := config.DB.QueryRow(
		config.CTX,
		query,
		email).Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return userData, nil
}
