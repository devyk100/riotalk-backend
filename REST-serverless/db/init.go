package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sync"
)

var DBQueries *Queries
var Pool *pgxpool.Pool
var once sync.Once

func InitDb(context context.Context) error {
	once.Do(func() {

	})
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context, config)
	if err != nil {
		return err
	}

	queries := New(pool)
	DBQueries = queries
	Pool = pool
	return nil
}
