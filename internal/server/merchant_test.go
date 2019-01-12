package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestMerchant_Create(t *testing.T) {

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockMerchant := database.NewMockMerchant(controller)
		mockMerchant.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Merchant().Return(mockMerchant)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/merchant", nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCard := database.NewMockMerchant(controller)
		mockCard.EXPECT().Write(gomock.Any()).Return(errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Merchant().Return(mockCard)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/merchant", nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
