package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	mockdb "learn-until-die/db/mock"
	db "learn-until-die/db/sqlc"
	"learn-until-die/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountId     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				requireBody(t, recorder.Body, account)
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:      "Not found",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "Internal Server Error",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:      "Bad Request",
			accountId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).
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
			url := fmt.Sprintf("/account/%d", tc.accountId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

// func TestCreateAccountAPI(t *testing.T) {
// 	account := randomAccount()

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"currency": account.Currency,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.CreateAccountParams{
// 					Owner: account.Owner,
// 					Currency: account.Currency,
// 					Balance: 0,
// 				}
// 				store.EXPECT().CreateAccount(arg.Any(), gomock.Eq(arg)).
// 				Times(1).
// 				Return(account, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		}
// 	}
// }

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBody(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getAccount db.Account
	err = json.Unmarshal(data, &getAccount)
	require.NoError(t, err)
	require.Equal(t, account, getAccount)
}
