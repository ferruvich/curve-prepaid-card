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

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Write(&sql.DB{}, card)

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
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, nil,
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Write(&sql.DB{}, card)

		require.NoError(t, err)
	})
}

func TestCard_Read(t *testing.T) {

	card := &model.Card{ID: "id"}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		resCard, err := cardDB.Read(&sql.DB{}, "id")

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
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, nil,
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		resCard, err := cardDB.Read(&sql.DB{}, "id")

		require.NotNil(t, resCard)
		require.NoError(t, err)
	})
}

func TestCard_Update(t *testing.T) {

	card, _ := model.NewCard("userID")

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), card.ID, card.Owner, card.AccountBalance,
			(card.AccountBalance - card.AvailableBalance),
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Update(&sql.DB{}, card)

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
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, nil,
		)

		cardDB := &CardDataBase{
			service: mockDB,
		}

		err := cardDB.Update(&sql.DB{}, card)

		require.NoError(t, err)
	})
}
