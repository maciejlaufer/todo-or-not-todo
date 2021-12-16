package api

import (
	"bytes"
	"database/sql"
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
	userId := user.ID.String()

	testCases := []struct {
		name          string
		userID        string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: userId,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, formatUserResponse(user))
			},
		},
		{
			name:   "NotFound",
			userID: userId,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			userID: userId,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InvalidID",
			userID: "123e4567-e89b-12d3-a456-4266141740",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/%s", tc.userID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})

	}
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
