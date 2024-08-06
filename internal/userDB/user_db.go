package userdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dbPool interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
}

type UserManager struct {
	Mu   *sync.RWMutex
	Pool dbPool
}

type User struct {
	Id     int64
	City   string
	Status int32
}

type DBConfig struct {
	User   string
	Pass   string
	DBName string
	Host   string
	Port   string
}

func NewUserDB(cfg DBConfig) *UserManager {
	pool, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName),
	)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	return &UserManager{
		Pool: pool,
		Mu:   &sync.RWMutex{},
	}
}

func (um *UserManager) GetUserPreferences(id int64) (string, error) {
	um.Mu.RLock()
	row := um.Pool.QueryRow(context.Background(), "SELECT user_id, city FROM preferences where user_id = $1;", id)
	um.Mu.RUnlock()
	var out string
	var respID int32
	err := row.Scan(&respID, &out)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (um *UserManager) SetUserPreference(id int64, cityName string) error {
	um.Mu.RLock()
	tg, err := um.Pool.Exec(context.Background(), "UPDATE preferences SET city = $1 WHERE user_id = $2;", cityName, id)
	um.Mu.RUnlock()
	if err == nil && tg.RowsAffected() == 0 {
		err = um.CreateUserPreferences(id, cityName)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func (um *UserManager) CreateUserPreferences(id int64, cityName string) error {
	um.Mu.Lock()
	defer um.Mu.Unlock()
	_, err := um.Pool.Exec(context.Background(), "INSERT INTO preferences(user_id, city) VALUES ($1, $2);", id, cityName)
	return err
}

// Returns all info about user.
func (um *UserManager) GetUser(id int64) (*User, error) {
	um.Mu.RLock()
	row := um.Pool.QueryRow(context.Background(), "SELECT city, status FROM preferences WHERE user_id = $1;", id)
	um.Mu.RUnlock()
	var out User
	err := row.Scan(&out.City, &out.Status)
	if err != nil {
		return nil, err
	}
	out.Id = id
	return &out, nil
}

func (um *UserManager) CheckUserExist(id int64) (bool, error) {
	um.Mu.RLock()
	row := um.Pool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM preferences WHERE user_id = $1);", id)
	um.Mu.RUnlock()
	var result bool
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (um *UserManager) SetUser(u User) error {
	um.Mu.Lock()
	defer um.Mu.Unlock()
	tg, err := um.Pool.Exec(context.Background(), "UPDATE preferences SET city = $1, status = $2 WHERE user_id = $3;",
		u.City,
		u.Status,
		u.Id,
	)
	if err != nil {
		return err
	}
	if tg.RowsAffected() == 0 {
		return errors.New("SetUser error: user doesn't exist")
	}
	return nil
}

func (um *UserManager) CreateUser(id int64) error {
	um.Mu.Lock()
	defer um.Mu.Unlock()
	_, err := um.Pool.Exec(context.Background(), "INSERT INTO preferences(user_id) VALUES ($1);", id)
	return err
}
