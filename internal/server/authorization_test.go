package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestAuthorizationRequest_Capture(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockauthReq := database.NewMockAuthorizationRequest(controller)
	mockCard := database.NewMockCard(controller)
	mockTx := database.NewMockTransaction(controller)

	mockDB := database.NewMockDataBase(controller)
	mockDB.EXPECT().AuthorizationRequest().Return(mockauthReq).AnyTimes()
	mockDB.EXPECT().Card().Return(mockCard).AnyTimes()
	mockDB.EXPECT().Transaction().Return(mockTx).AnyTimes()

	server := &Service{db: mockDB}

	captureBody := &AmountBody{
		Amount: 5.0,
	}
	captureBodyBytes, _ := json.Marshal(captureBody)

	router := server.Routers()

	t.Run("should run", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(captureBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(nil)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "capture",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(captureBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to wrong body", func(t *testing.T) {
		route := strings.Join([]string{
			"/authorization", "testID", "capture",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusBadRequest, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(captureBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "capture",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(captureBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}

func TestAuthorizationRequest_Refund(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockauthReq := database.NewMockAuthorizationRequest(controller)
	mockCard := database.NewMockCard(controller)
	mockTx := database.NewMockTransaction(controller)

	mockDB := database.NewMockDataBase(controller)
	mockDB.EXPECT().AuthorizationRequest().Return(mockauthReq).AnyTimes()
	mockDB.EXPECT().Card().Return(mockCard).AnyTimes()
	mockDB.EXPECT().Transaction().Return(mockTx).AnyTimes()

	server := &Service{db: mockDB}

	refundBody := &AmountBody{
		Amount: 5.0,
	}
	refundBodyBytes, _ := json.Marshal(refundBody)

	router := server.Routers()

	t.Run("should run", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(refundBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()
		authReq.Capture(refundBody.Amount)

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(nil)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "refund",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(refundBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to wrong body", func(t *testing.T) {
		route := strings.Join([]string{
			"/authorization", "testID", "refund",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusBadRequest, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(refundBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()
		authReq.Capture(refundBody.Amount)

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "refund",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(refundBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}

func TestAuthorizationRequest_Revert(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockauthReq := database.NewMockAuthorizationRequest(controller)
	mockCard := database.NewMockCard(controller)

	mockDB := database.NewMockDataBase(controller)
	mockDB.EXPECT().AuthorizationRequest().Return(mockauthReq).AnyTimes()
	mockDB.EXPECT().Card().Return(mockCard).AnyTimes()

	server := &Service{db: mockDB}

	captureBody := &AmountBody{
		Amount: 5.0,
	}
	captureBodyBytes, _ := json.Marshal(captureBody)

	router := server.Routers()

	t.Run("should run", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(captureBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "revert",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(captureBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusAccepted, testRecorder.Code)
	})

	t.Run("should fail due to wrong body", func(t *testing.T) {
		route := strings.Join([]string{
			"/authorization", "testID", "revert",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusBadRequest, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(captureBody.Amount)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(errors.New("error"))

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		route := strings.Join([]string{
			"/authorization", authReq.ID, "revert",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(captureBodyBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
