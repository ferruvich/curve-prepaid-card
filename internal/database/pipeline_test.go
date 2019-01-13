package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestDataBase_RunPipeline(t *testing.T) {

	fakeError := errors.New("error")

	t.Run("should fail due to error in makeQuery", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := NewMockDBTransaction(controller)
		mockTransaction.EXPECT().Query(gomock.Any()).Return(
			nil, fakeError,
		)

		service := &Service{}
		pipelineStatement := &pipelineStmt{}

		res, err := service.runPipeline(
			mockTransaction, pipelineStatement,
		)

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := NewMockDBTransaction(controller)
		mockTransaction.EXPECT().Query(gomock.Any()).Return(
			&sql.Rows{}, nil,
		)

		service := &Service{}
		pipelineStatement := &pipelineStmt{}

		res, err := service.runPipeline(
			mockTransaction, pipelineStatement,
		)

		require.NoError(t, err)
		require.NotNil(t, res)
	})
}

func TestDataBase_NewPipelineStms(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		pipeline := service.newPipelineStmt("query")

		require.NotNil(t, pipeline)
	})
}

func TestPipeline_MakeQuery(t *testing.T) {

	fakeError := errors.New("error")

	t.Run("should fail due to error in makeQuery", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := NewMockDBTransaction(controller)
		mockTransaction.EXPECT().Query(gomock.Any()).Return(
			nil, fakeError,
		)

		pipelineStatement := &pipelineStmt{}

		res, err := pipelineStatement.makeQuery(
			mockTransaction,
		)

		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockTransaction := NewMockDBTransaction(controller)
		mockTransaction.EXPECT().Query(gomock.Any()).Return(
			&sql.Rows{}, nil,
		)

		pipelineStatement := &pipelineStmt{}

		res, err := pipelineStatement.makeQuery(
			mockTransaction,
		)

		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
