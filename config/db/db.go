package db

import (
	"context"
	"fmt"
	"log"
	"task_manager/constants"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var CTX = context.Background()

func InitializeDatabase() {
	var err error
	config, _ := pgxpool.ParseConfig(constants.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 32
	config.MinConns = 16
	DB, err = pgxpool.NewWithConfig(CTX, config)

	fmt.Println(config.MaxConns)

	if err != nil {
		log.Fatal(err)
	}

	DB.Exec(CTX, `
		CREATE TABLE IF NOT EXISTS users (
        	id BIGSERIAL PRIMARY KEY,
        	name VARCHAR(30) NOT NULL,
        	email vARCHAR(50) NOT NULL UNIQUE,
        	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	_, err = DB.Exec(CTX, `

		CREATE TABLE IF NOT EXISTS todos (
        	id BIGSERIAL PRIMARY KEY,
        	title varchar(100) NOT NULL,
        	user_id BIGINT NOT NULL,
			description VARCHAR(250),
        	status todo_status DEFAULT 'pending',
        	due_date TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	// defer DB.Close()
	log.Printf("Database connection established.")
}
