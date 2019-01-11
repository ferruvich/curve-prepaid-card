package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

func TestAuthorizationRequest_Create(t *testing.T) {
	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		authReq, err := mockAuthRequestMiddleware.Create(
			"merchant_ID", "card_ID", 10.0,
		)

		require.Nil(t, authReq)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Write(gomock.Any()).Return(
			nil,
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		authReq, err := mockAuthRequestMiddleware.Create(
			"merchant_ID", "card_ID", 10.0,
		)

		require.NotNil(t, authReq)
		require.NoError(t, err)
	})
}
