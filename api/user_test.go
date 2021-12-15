package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/maciejlaufer/todoornottodo/db/mock"
	db "github.com/maciejlaufer/todoornottodo/db/sqlc"
	"github.com/maciejlaufer/todoornottodo/util"
	"github.com/stretchr/testify/require"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		GetUserById(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).
		Return(user, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%s", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchUser(t, recorder.Body, formatUserResponse(user))
}

func randomUser() db.User {
	return db.User{
		ID:        uuid.New(),
		Email:     util.RandomEmail(),
		Password:  util.RandomString(10),
		FirstName: db.NewNullString(util.RandomString(9)),
		LastName:  db.NewNullString(util.RandomString(8)),
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user userResponse) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser userResponse
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
