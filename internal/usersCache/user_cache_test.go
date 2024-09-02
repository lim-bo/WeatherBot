package usercache_test

import (
	"errors"
	"reflect"
	"testing"
	"weatherbot/internal/bot"
	userdb "weatherbot/internal/userDB"
	usercache "weatherbot/internal/usersCache"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
)

var ErrNotNil = errors.New("not nil error")

func TestRedisCache(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	mockedCli := redismock.NewNiceMock(redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	}))
	// Setting usercache repository with mocked redis client
	uc := usercache.NewWithClient(mockedCli)

	u := userdb.User{Id: 0, City: "test", Status: bot.DefaultState}
	// Testing SetUser method:
	// clean run:
	mockedCli.On("RPUSH").Return(redis.NewStatusResult("", nil))
	err = uc.SetUser(&u)
	if err != nil {
		t.Error("error expected: nil, got: ", err)
	}

	// Testing GetUser method:
	// clean run:
	mockedCli.On("LRANGE").Return(redis.NewStringSliceResult([]string{"test", "0"}, nil))
	gotU, err := uc.GetUser(u.Id)
	if err != nil {
		t.Error("expected nil error, got: ", err)
	}
	if !reflect.DeepEqual(u, *gotU) {
		t.Error("results don't match, got: ", gotU)
	}
}
