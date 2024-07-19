package userdb

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserManager struct {
	mu   sync.RWMutex
	pool *pgxpool.Pool
}

func (um *UserManager) GetUserPreferences(id int32) (string, error) {
	um.mu.RLock()
	row := um.pool.QueryRow(context.Background(), "SELECT user_id, city FROM preferences where user_id = $1", id)
	um.mu.RUnlock()
	var out string
	err := row.Scan(&out)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (um *UserManager) SetUserPreference(id int32, cityName string) error {
	um.mu.Lock()
	tg, err := um.pool.Exec(context.Background(), "UPDATE preferences SET city = $1 WHERE user_id = $2", cityName, id)
	um.mu.RUnlock()
	if tg.RowsAffected() == 0 || err != nil {
		return err
	}
	return nil
}
