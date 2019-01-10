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

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), merchant.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, errors.New("error writing merchant"),
		)

		merchantDB := &MerchantDataBase{
			service: mockDB,
		}

		err := merchantDB.Write(&sql.DB{}, merchant)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), merchant.ID,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(&sql.DB{}, gomock.Any()).Return(
			nil, nil,
		)

		merchantDB := &MerchantDataBase{
			service: mockDB,
		}

		err := merchantDB.Write(&sql.DB{}, merchant)

		require.NoError(t, err)
	})
}
