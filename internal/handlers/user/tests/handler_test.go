package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"your-project-name/internal/auth"
	"your-project-name/internal/db"
	"your-project-name/internal/handlers/user"
	"your-project-name/internal/middleware/tests"
	"your-project-name/internal/server"
	"your-project-name/internal/store/mock"
)

func TestGetProfile(t *testing.T) {
	testCases := []struct {
		name      string
		setupAuth func(t *testing.T,
			request *http.Request,
			tokenMaker auth.Maker,
			userID uuid.UUID,
			duration time.Duration)
		body          gin.H
		buildStubs    func(store *mock.MockStore, uid uuid.UUID)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker auth.Maker, userID uuid.UUID, duration time.Duration) {
				tests.AddAuthorization(t, request, tokenMaker, userID, duration)
			},
			buildStubs: func(store *mock.MockStore, uid uuid.UUID) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(uid)).
					Times(1).
					Return(db.User{
						ID:       uid,
						Username: "testuser",
						Email:    "test@example.com",
						FullName: "Test User",
					}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				if recorder.Code != http.StatusOK {
					t.Logf("Response body: %s", recorder.Body.String())
				}
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body)
			},
		},
		// ... other test cases
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := mock.NewMockStore(ctrl)
			testServer := server.NewTestServer(t, mockStore)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/users/me"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
			require.NoError(t, err)

			uid, _ := uuid.NewUUID()
			tc.setupAuth(t, request, testServer.TokenMaker, uid, time.Minute)
			tc.buildStubs(mockStore, uid)

			testServer.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotResponse user.ProfileResponse
	err = json.Unmarshal(data, &gotResponse)
	require.NoError(t, err)

	require.NotEmpty(t, gotResponse.ID)
	require.NotEmpty(t, gotResponse.Username)
	require.NotEmpty(t, gotResponse.Email)
}
