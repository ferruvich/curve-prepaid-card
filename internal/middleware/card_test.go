package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestCard_Create(t *testing.T) {

	ownerID := "ownerID"

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockCardMiddleware.Create(ownerID)

		require.Nil(t, card)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Write(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockCardMiddleware.Create(ownerID)

		require.NotNil(t, card)
		require.NoError(t, err)
	})
}

func TestCard_GetCard(t *testing.T) {

	cardID := "cardID"
	card := &model.Card{}

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(cardID).Return(
			nil, errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		card, err := mockCardMiddleware.GetCard(cardID)

		require.Nil(t, card)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(cardID).Return(card, nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		resCard, err := mockCardMiddleware.GetCard(cardID)

		require.NotNil(t, resCard)
		require.NoError(t, err)
	})
}

func TestCard_Deposit(t *testing.T) {

	amountToDeposit := 10.0

	t.Run("should fail due to db error", func(t *testing.T) {

		card := &model.Card{
			ID: "ID", Owner: "ownerID", AccountBalance: 0.0,
			AvailableBalance: 0.0,
		}
		charged := &model.Card{
			ID: "ID", Owner: "ownerID", AccountBalance: amountToDeposit,
			AvailableBalance: amountToDeposit,
		}

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(card.ID).Return(card, nil)
		mockCardDB.EXPECT().Update(charged).Return(
			errors.New("error"),
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(2)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		err := mockCardMiddleware.Deposit(
			card.ID, amountToDeposit,
		)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {

		card := &model.Card{
			ID: "ID", Owner: "ownerID", AccountBalance: 0.0,
			AvailableBalance: 0.0,
		}
		charged := &model.Card{
			ID: "ID", Owner: "ownerID", AccountBalance: amountToDeposit,
			AvailableBalance: amountToDeposit,
		}

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(card.ID).Return(card, nil)
		mockCardDB.EXPECT().Update(charged).Return(
			nil,
		)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(2)

		mockCardMiddleware := &CardMiddleware{
			middleware: mockMiddleware,
		}

		err := mockCardMiddleware.Deposit(
			card.ID, amountToDeposit,
		)

		require.NoError(t, err)
	})
}
