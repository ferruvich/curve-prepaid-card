package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestCard_Create(t *testing.T) {

	userID := "userID"

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Write(gomock.Any()).Return(errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}

func TestCard_GetCard(t *testing.T) {

	userID := "userID"
	cardID := "cardID"
	card := &model.Card{
		ID: cardID, Owner: userID,
	}

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(gomock.Any()).Return(card, nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID,
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodGet, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		response := &model.Card{}
		json.Unmarshal([]byte(testRecorder.Body.String()), response)

		require.Equal(t, http.StatusOK, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(gomock.Any()).Return(nil, errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID,
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodGet, route, nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}

func TestCard_Deposit(t *testing.T) {

	userID := "userID"
	cardID := "cardID"

	t.Run("should run", func(t *testing.T) {

		card := &model.Card{
			ID: cardID, Owner: userID,
		}
		deposit := &DepositRequest{
			Amount: 10.0,
		}
		depositBytes, _ := json.Marshal(deposit)
		charged := &model.Card{
			ID: cardID, Owner: userID, AccountBalance: deposit.Amount,
			AvailableBalance: deposit.Amount,
		}

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(gomock.Any()).Return(card, nil)
		mockCard.EXPECT().Update(charged).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard).Times(2)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID, "deposit",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(depositBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusNoContent, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {

		// card := &model.Card{
		// 	ID: cardID, Owner: userID,
		// }
		deposit := &DepositRequest{
			Amount: 10.0,
		}
		depositBytes, _ := json.Marshal(deposit)

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockCard(controller)
		mockCard.EXPECT().Read(gomock.Any()).Return(nil, errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		route := strings.Join([]string{
			"/user", userID, "card", cardID, "deposit",
		}, "/")

		testRequest := httptest.NewRequest(
			http.MethodPost, route, bytes.NewReader(depositBytes),
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
