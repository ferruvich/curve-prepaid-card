package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestUser_Create(t *testing.T) {

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockUserDB := database.NewMockUser(controller)
		mockUserDB.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUserDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockUserMiddleware := &UserMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockUserMiddleware.Create()

		require.Nil(t, card)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockUserDB := database.NewMockUser(controller)
		mockUserDB.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUserDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockUserMiddleware := &UserMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockUserMiddleware.Create()

		require.NotNil(t, card)
		require.NoError(t, err)
	})
}

func TestUser_Read(t *testing.T) {

	userID := "userID"
	user := &model.User{}

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockUserDB := database.NewMockUser(controller)
		mockUserDB.EXPECT().Read(userID).Return(nil, errors.New("error"))

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUserDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockUserMiddleware := &UserMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockUserMiddleware.Read(userID)

		require.Nil(t, card)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockUserDB := database.NewMockUser(controller)
		mockUserDB.EXPECT().Read(userID).Return(user, nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().User().Return(mockUserDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockUserMiddleware := &UserMiddleware{
			middleware: mockMiddleware,
		}

		resCard, err := mockUserMiddleware.Read(userID)

		require.NotNil(t, resCard)
		require.NoError(t, err)
	})
}
