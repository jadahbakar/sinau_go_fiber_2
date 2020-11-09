package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jadahbakar/sinau_go_fiber_2/config"
)

type pgRepository struct {
	dbHandler *pgxpool.Pool
}

// NewPgCustomerRepo function
func Connect() (*pgRepository, error) {

	var err error

	p := config.Config("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here - hack ??
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing Port env: str to int")
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &pgRepository{
		dbHandler: dbpool,
	}, nil
}
