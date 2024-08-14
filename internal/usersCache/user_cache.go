package usercache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
	userdb "weatherbot/internal/userDB"
	"weatherbot/logger"

	"github.com/go-redis/redis"
)

var (
	ErrKeyNotExist    = errors.New("key does not exist")
	ErrParsingResults = errors.New("result parsing error")
)

type CacheManager struct {
	lg  *logger.SLogger
	mu  *sync.RWMutex
	cli *redis.Client
}

type RedisCfg struct {
	Host     string
	Port     string
	Username string
	Pass     string
	ID       int
}

func New(cfg RedisCfg) *CacheManager {
	log := logger.NewSLogger()
	url := fmt.Sprintf("redis://%s:%s@%s:%s/%d", cfg.Username, cfg.Pass, cfg.Host, cfg.Port, cfg.ID)
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal(context.Background(), err)
	}
	cli := redis.NewClient(opts)
	return &CacheManager{
		lg:  log,
		mu:  &sync.RWMutex{},
		cli: cli,
	}
}

// User data organization in redis:
// key: "user:{id}" value: list{city, status}

func (cm *CacheManager) GetUser(id int64) (*userdb.User, error) {
	cm.mu.RLock()
	result, err := cm.cli.LRange(fmt.Sprintf("user:%d", id), 0, -1).Result()
	cm.mu.RUnlock()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrKeyNotExist
		}
		return nil, errors.New("cache manager error: " + err.Error())
	}
	st, err := strconv.Atoi(result[1])
	if err != nil {
		return nil, ErrParsingResults
	}
	return &userdb.User{
		City:   result[0],
		Status: int32(st),
	}, nil
}

func (cm *CacheManager) SetUser(u *userdb.User) error {
	cm.mu.Lock()
	err := cm.cli.Set(fmt.Sprintf("user:%d", u.Id), []string{u.City, strconv.FormatInt(int64(u.Status), 10)}, time.Hour).Err()
	cm.mu.Unlock()
	if err != nil {
		return errors.New("cache manager error: " + err.Error())
	}
	return nil
}
