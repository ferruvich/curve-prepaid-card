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
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestAuthorizationRequest_Create(t *testing.T) {

	merchantID := "merchant_ID"
	cardID := "card_ID"
	amount := 10.0

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReq := database.NewMockAuthorizationRequest(controller)
		mockAuthReq.EXPECT().Write(gomock.Any()).Return(nil)

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(cardID).Return(&model.Card{
			ID: cardID, AccountBalance: amount, AvailableBalance: amount,
		}, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReq)
		mockDB.EXPECT().Card().Return(mockCard).Times(2)

		server := &Service{db: mockDB}

		authReq := &AuthorizationRequestBody{
			MerchantID: merchantID, CardID: cardID, Amount: amount,
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

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(cardID).Return(&model.Card{
			ID: cardID, AccountBalance: amount, AvailableBalance: amount,
		}, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReq)
		mockDB.EXPECT().Card().Return(mockCard).Times(2)

		server := &Service{db: mockDB}

		authReq := &AuthorizationRequestBody{
			MerchantID: merchantID, CardID: cardID, Amount: amount,
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
