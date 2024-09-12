package postgres

import (
	"database/sql"
	"delivery/config"
	stg "delivery/storage"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db   *sql.DB
	auth stg.DeliveryService
}

func NewPostgresStorage(cfg config.Config) (stg.InitRoot, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword,
		cfg.PostgresHost, cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}

func (s *Storage) Delivery() stg.DeliveryService {
	if s.auth == nil {
		s.auth = &DeliveryStorage{db: s.Db}
	}
	return s.auth
}
