package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/internal/repo"
	"github.com/ferruvich/curve-challenge/testdata"
)

const (
	ownerID = "someID"
)

func TestNewCardMiddleware(t *testing.T) {
	cardMiddleware, err := NewCardMiddleware(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, cardMiddleware)
}

func TestCardMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCardRepo := repo.NewMockCard(controller)
	mockCardRepo.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	cardMiddleware := &CardMiddleware{
		repo: mockCardRepo,
	}
	merchant, err := cardMiddleware.Create(context.Background(), ownerID)

	require.NoError(t, err)
	require.NotNil(t, merchant)
}
