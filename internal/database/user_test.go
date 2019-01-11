package database

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestUser_Write(t *testing.T) {

	user, _ := model.NewUser()
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), user.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		userDB := &UserDataBase{
			service: mockDB,
		}

		err := userDB.Write(user)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), user.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		userDB := &UserDataBase{
			service: mockDB,
		}

		err := userDB.Write(user)

		require.NoError(t, err)
	})
}

func TestMerechant_Read(t *testing.T) {

	user := &model.User{ID: "id"}
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), user.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing card"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		userDB := &UserDataBase{
			service: mockDB,
		}

		resCard, err := userDB.Read("id")

		require.Nil(t, resCard)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), user.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		userDB := &UserDataBase{
			service: mockDB,
		}

		resCard, err := userDB.Read("id")

		require.NotNil(t, resCard)
		require.NoError(t, err)
	})
}
