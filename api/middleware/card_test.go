package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/api/model"
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

	mockUserRepo := repo.NewMockUser(controller)
	mockUserRepo.EXPECT().Read(
		context.Background(),
		ownerID,
	).Return(&model.User{ID: ownerID}, nil)

	mockCardRepo := repo.NewMockCard(controller)
	mockCardRepo.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	userMiddleware := &UserMiddleware{
		repo: mockUserRepo,
	}

	cardMiddleware := &CardMiddleware{
		repo: mockCardRepo,
	}

	merchant, err := cardMiddleware.Create(context.Background(), ownerID, userMiddleware)

	require.NoError(t, err)
	require.NotNil(t, merchant)
}
