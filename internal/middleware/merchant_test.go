package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/testdata"
)

func TestNewMerchantMiddleware(t *testing.T) {
	merchantMiddleware, err := NewMerchantMiddleware(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, merchantMiddleware)
}

func TestMerchantMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockdatabase := database.NewMockMerchant(controller)
	mockdatabase.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	merchantMiddleware := &MerchantMiddleware{
		database: mockdatabase,
	}

	merchant, err := merchantMiddleware.Create(context.Background())

	require.NoError(t, err)
	require.NotNil(t, merchant)
}
