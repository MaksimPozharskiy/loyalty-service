package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(connectParams string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectParams)
	if err != nil {
		return nil, err
	}

	err = initDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB(db *sql.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
		user_id serial PRIMARY KEY,
        username VARCHAR ( 50 ) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
		order_numbers INTEGER[],
		accural_value DECIMAL(10, 2) NOT NULL DEFAULT 0.00
		created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := db.Exec(schema)
	return err
}
