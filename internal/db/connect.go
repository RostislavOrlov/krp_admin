package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbUrl)

	if err != nil {
		//return nil, errors.Wrap(err, "unable create pgxpool")
	}

	err = retry.Do(func() error {
		var err error
		err = pool.Ping(context.Background())
		return err
	},
		retry.Attempts(10),
		retry.OnRetry(func(n uint, err error) {

		}))

	if err != nil {
		//return err, errors.Wrap(err, "unable ping database")
	}

	return pool, nil
}
