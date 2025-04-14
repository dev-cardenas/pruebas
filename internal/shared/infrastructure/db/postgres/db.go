package postgres

import (
	"context"
	"fmt"
	"pruebas/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *configs.DatabaseConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
