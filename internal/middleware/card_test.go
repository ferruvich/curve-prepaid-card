package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/internal/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
	"github.com/ferruvich/curve-challenge/testdata"
)

const (
	ownerID         = "someUserID"
	cardID          = "someCardID"
	amountToDeposit = 10.0
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
	card, err := cardMiddleware.Create(context.Background(), ownerID)

	require.NoError(t, err)
	require.NotNil(t, card)
}

func TestCardMiddleware_GetCard(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCardRepo := repo.NewMockCard(controller)
	mockCardRepo.EXPECT().Read(
		context.Background(), cardID,
	).Return(&model.Card{
		ID: cardID,
	}, nil)

	cardMiddleware := &CardMiddleware{
		repo: mockCardRepo,
	}
	card, err := cardMiddleware.GetCard(context.Background(), cardID)

	require.NoError(t, err)
	require.NotNil(t, card)
	require.Equal(t, card.ID, cardID)
}

func TestCardMiddleware_Deposit(t *testing.T) {

	mockCard := &model.Card{
		ID: cardID, AvailableBalance: 0.0, AccountBalance: 0.0,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCardRepo := repo.NewMockCard(controller)
	mockCardRepo.EXPECT().Read(
		context.Background(), cardID,
	).Return(mockCard, nil)
	mockCardRepo.EXPECT().Update(
		context.Background(), mockCard,
	).Return(nil)

	cardMiddleware := &CardMiddleware{
		repo: mockCardRepo,
	}
	err := cardMiddleware.Deposit(context.Background(), cardID, amountToDeposit)

	require.NoError(t, err)
	require.Equal(t, mockCard.AccountBalance, amountToDeposit)
	require.Equal(t, mockCard.AvailableBalance, amountToDeposit)
}
