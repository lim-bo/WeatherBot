package userdb_test

import (
	"sync"
	"testing"
	userdb "weatherbot/internal/userDB"

	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
)

func TestDBGet(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockpool := pgxpoolmock.NewMockPgxIface(ctrl)

	var requestID int32 = 1
	requestCity := "TEST"

	mockpool.EXPECT().QueryRow(gomock.Any(), pgxpoolmock.QueryContains("SELECT user_id, city FROM"), []interface{}{requestID}).Return(
		pgxpoolmock.NewRow(requestID, requestCity),
	)
	um := &userdb.UserManager{
		Pool: mockpool,
		Mu:   &sync.RWMutex{},
	}
	respCity, err := um.GetUserPreferences(int64(requestID))
	if respCity != requestCity || err != nil {
		t.Fatal("results don't match")
	}

	mockpool.EXPECT().QueryRow(gomock.Any(), pgxpoolmock.QueryContains("SELECT user_id, city FROM"), []interface{}{int32(0)}).Return(
		pgxpoolmock.NewRow(int32(0), "").WithError(pgx.ErrNoRows),
	)
	_, err = um.GetUserPreferences(0)
	if err == nil {
		t.Fatal("error must not be nil")
	}
}

func TestDBSet(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockpool := pgxpoolmock.NewMockPgxIface(ctrl)

	um := &userdb.UserManager{
		Pool: mockpool,
		Mu:   &sync.RWMutex{},
	}
	requestCity := "TEST"
	requestID := int32(1)
	testCommandTag := []byte("UPDATE 1")
	mockpool.EXPECT().Exec(gomock.Any(), pgxpoolmock.QueryContains("UPDATE preferences SET"),
		[]interface{}{
			requestCity,
			requestID,
		},
	).Return(
		testCommandTag,
		nil,
	)
	err := um.SetUserPreference(int64(requestID), requestCity)
	if err != nil {
		t.Fatal("update error: ", err)
	}
	mockpool.EXPECT().Exec(gomock.Any(), pgxpoolmock.QueryContains("UPDATE preferences SET"),
		[]interface{}{
			requestCity,
			requestID,
		},
	).Return(
		testCommandTag,
		pgx.ErrNoRows,
	)
	err = um.SetUserPreference(int64(requestID), requestCity)
	if err == nil {
		t.Fatal("err must not be nil")
	}

	// Testing scenario when there is no this user in db
	testCommandTag = []byte("INSERT 1")
	mockpool.EXPECT().Exec(gomock.Any(), pgxpoolmock.QueryContains("INSERT INTO preferences"),
		[]interface{}{
			requestID,
			requestCity,
		},
	).Return(
		testCommandTag,
		nil,
	)
	err = um.CreateUserPreferences(int64(requestID), requestCity)
	if err != nil {
		t.Fatal("insertion error: ", err)
	}
	mockpool.EXPECT().Exec(gomock.Any(), pgxpoolmock.QueryContains("INSERT INTO preferences"),
		[]interface{}{
			requestID,
			requestCity,
		},
	).Return(
		testCommandTag,
		pgx.ErrNoRows,
	)
	err = um.CreateUserPreferences(int64(requestID), requestCity)
	if err == nil {
		t.Fatal("err must not be nil")
	}
}
