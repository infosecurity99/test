package postgres

import (
	"context"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dealerRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewDealerRepo(db *pgxpool.Pool, log logger.ILogger) storage.IDealerStorage {
	return &dealerRepo{
		db:  db,
		log: log,
	}
}

func (d *dealerRepo) AddSum(ctx context.Context, totalSum int) error {
	//ozini sum: ga qoshish kerak total sum -> update
	query := `update dealer set sum = sum + $1 where id = '1cfd84e6-72cb-4135-a802-85d10e4183ea'`
	if rowsAffected, err := d.db.Exec(ctx, query, &totalSum); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			d.log.Error("error is while rows affected", logger.Error(err))
			return err
		}
		d.log.Error("error is while updating dealer sum", logger.Error(err))

		return err
	}
	return nil
}
