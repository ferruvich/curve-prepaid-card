package database

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestMerchant_Write(t *testing.T) {

	merchant, _ := model.NewMerchant()
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), merchant.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing merchant"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		merchantDB := &MerchantDataBase{
			service: mockDB,
		}

		err := merchantDB.Write(merchant)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), merchant.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		merchantDB := &MerchantDataBase{
			service: mockDB,
		}

		err := merchantDB.Write(merchant)

		require.NoError(t, err)
	})
}
