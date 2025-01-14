package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	*OrdersRepository
}

func New(dsn string) (*Repository, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		fmt.Println(dsn)
		log.Fatal("Cannot connect to Postgres")
	}

	return &Repository{
		// DB:                db,
		OrdersRepository: &OrdersRepository{DB: db},
	}, nil
}
