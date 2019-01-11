package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

func TestMerchant_Create(t *testing.T) {

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockMerchantDB := database.NewMockMerchant(controller)
		mockMerchantDB.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Merchant().Return(mockMerchantDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockMerchantMiddleware := &MerchantMiddleware{
			middleware: mockMiddleware,
		}

		merchant, err := mockMerchantMiddleware.Create()

		require.Nil(t, merchant)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockMerchantDB := database.NewMockMerchant(controller)
		mockMerchantDB.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Merchant().Return(mockMerchantDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockMerchantMiddleware := &MerchantMiddleware{
			middleware: mockMiddleware,
		}

		merchant, err := mockMerchantMiddleware.Create()

		require.NotNil(t, merchant)
		require.NoError(t, err)
	})
}
