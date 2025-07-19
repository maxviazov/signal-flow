package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	// Add fields and methods as needed for your Postgres repository
	pool *pgxpool.Pool
}

func New() {

}
