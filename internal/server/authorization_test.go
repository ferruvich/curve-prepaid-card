package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

func TestAuthorizationRequest_Create(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReq := database.NewMockAuthorizationRequest(controller)
		mockAuthReq.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReq)

		server := &Service{
			db: mockDB,
		}

		authReq := &AuthorizationRequestBody{
			MerchantID: "merchantID", CardID: "cardID", Amount: 10.0,
		}

		authReqBytes, _ := json.Marshal(authReq)

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/authorization", bytes.NewReader(authReqBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReq := database.NewMockAuthorizationRequest(controller)
		mockAuthReq.EXPECT().Write(gomock.Any()).Return(errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReq)

		server := &Service{
			db: mockDB,
		}

		authReq := &AuthorizationRequestBody{
			MerchantID: "merchantID", CardID: "cardID", Amount: 10.0,
		}

		authReqBytes, _ := json.Marshal(authReq)

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/authorization", bytes.NewReader(authReqBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
