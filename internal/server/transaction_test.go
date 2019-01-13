package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestTransaction_Create(t *testing.T) {

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

	captureBody := &CaptureBody{
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
