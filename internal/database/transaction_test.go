package database

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestTransaction_Write(t *testing.T) {

	tx, _ := model.NewPaymentTransaction("senderID", "receiverID", 10.0)
	db := &sql.DB{}

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)
		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), tx.ID, tx.Sender, tx.Receiver, tx.Amount,
			tx.Date, tx.Type,
		).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		txBD := &TransactionDataBase{
			service: mockDB,
		}

		err := txBD.Write(tx)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)
		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), tx.ID, tx.Sender, tx.Receiver, tx.Amount,
			tx.Date, tx.Type,
		).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(nil, nil)
		mockDB.EXPECT().GetConnection().Return(db)

		txBD := &TransactionDataBase{
			service: mockDB,
		}

		err := txBD.Write(tx)

		require.NoError(t, err)
	})
}
