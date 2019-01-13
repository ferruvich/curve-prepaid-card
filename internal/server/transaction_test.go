package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestTransaction_GetList(t *testing.T) {

	txs := []*model.Transaction{}
	userID := "userID"
	cardID := "cardID"

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := database.NewMockTransaction(controller)
		mockTransaction.EXPECT().GetListByCard(cardID).Return(txs, nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Transaction().Return(mockTransaction)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID, "transaction",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodGet, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusOK, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := database.NewMockTransaction(controller)
		mockTransaction.EXPECT().GetListByCard(cardID).Return(nil, errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Transaction().Return(mockTransaction)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID, "transaction",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodGet, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
