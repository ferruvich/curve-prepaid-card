package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/internal/repo"
	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewMerchantMiddleware(t *testing.T) {
	merchantMiddleware, err := NewMerchantMiddleware(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, merchantMiddleware)
}

func TestMerchantMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := repo.NewMockMerchant(controller)
	mockRepo.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	merchantMiddleware := &MerchantMiddleware{
		repo: mockRepo,
	}

	user, err := merchantMiddleware.Create(context.Background())

	require.NoError(t, err)
	require.NotNil(t, user)
}
