package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestCard_Write(t *testing.T) {

	card, _ := model.NewCard("userID")
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Write(card)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Write(card)

		require.NoError(t, err)
	})
}

func TestCard_Read(t *testing.T) {

	card := &model.Card{ID: "id"}
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		resCard, err := cardDB.Read("id")

		require.Nil(t, resCard)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		resCard, err := cardDB.Read("id")

		require.NotNil(t, resCard)
		require.NoError(t, err)
	})
}

func TestCard_Update(t *testing.T) {

	card, _ := model.NewCard("userID")
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Update(card)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Update(card)

		require.NoError(t, err)
	})
}
