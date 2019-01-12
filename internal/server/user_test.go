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

func TestUser_Create(t *testing.T) {

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockUser := database.NewMockUser(controller)
		mockUser.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUser)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/user", nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusCreated, testRecorder.Code)
	})

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockUser := database.NewMockUser(controller)
		mockUser.EXPECT().Write(gomock.Any()).Return(errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUser)

		server := &Service{
			db: mockDB,
		}

		router := server.Routers()

		testRequest := httptest.NewRequest(
			http.MethodPost, "/user", nil,
		)

		testRecorder := httptest.NewRecorder()

		router.ServeHTTP(testRecorder, testRequest)

		require.Equal(t, http.StatusInternalServerError, testRecorder.Code)
	})
}
