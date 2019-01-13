package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestTransaction_CreatePayment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	amountToCapture := 5.0

	mockauthReq := database.NewMockAuthorizationRequest(controller)
	mockCard := database.NewMockCard(controller)
	mockTx := database.NewMockTransaction(controller)

	mockDB := database.NewMockDataBase(controller)
	mockDB.EXPECT().AuthorizationRequest().Return(mockauthReq).AnyTimes()
	mockDB.EXPECT().Card().Return(mockCard).AnyTimes()
	mockDB.EXPECT().Transaction().Return(mockTx).AnyTimes()

	mockMiddleware := NewMockMiddleware(controller)
	mockMiddleware.EXPECT().DataBase().Return(mockDB).AnyTimes()

	t.Run("should fail due to db error", func(t *testing.T) {

		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(amountToCapture)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		txMiddleware := &TransactionMiddleware{
			middleware: mockMiddleware,
		}

		res, err := txMiddleware.CreatePayment(authReq.ID, amountToCapture)

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should run", func(t *testing.T) {

		card, _ := model.NewCard("ownerID")
		card.IncrementAccountBalance(20.0)
		card.BlockAmount(amountToCapture)

		authReq, _ := model.NewAuthorizationRequest("merchantID", card.ID, 10.0)
		authReq.Approve()

		mockauthReq.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockauthReq.EXPECT().Update(gomock.Any()).Return(nil)

		mockCard.EXPECT().Read(card.ID).Return(card, nil)
		mockCard.EXPECT().Update(gomock.Any()).Return(nil)

		mockTx.EXPECT().Write(gomock.Any()).Return(nil)

		txMiddleware := &TransactionMiddleware{
			middleware: mockMiddleware,
		}

		res, err := txMiddleware.CreatePayment(authReq.ID, amountToCapture)

		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, res.Amount, amountToCapture)
		require.Equal(t, res.Sender, card.Owner)
		require.Equal(t, res.Receiver, authReq.Merchant)
	})
}
