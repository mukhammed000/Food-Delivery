package postgres

import (
	"auth/config"
	stg "auth/storage"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db   *sql.DB
	Rdb  *redis.Client
	auth stg.AuthService
}

func NewPostgresStorage(cfg config.Config) (stg.InitRoot, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword,
		cfg.POSTGRES_HOST, cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.REDIS_HOST, cfg.REDIS_PORT),
	})

	return &Storage{Db: db, Rdb: redisClient}, nil
}

func (s *Storage) Auth() stg.AuthService {
	if s.auth == nil {
		s.auth = &AuthStorage{db: s.Db, rdb: s.Rdb}
	}
	return s.auth
}
