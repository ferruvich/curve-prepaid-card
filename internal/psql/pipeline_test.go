package psql

import (
	sql "database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewPipelineStmt(t *testing.T) {

	tests := map[string]struct {
		query string
		args  []interface{}
	}{
		"empty args": {
			query: "SQL Query",
		},
		"with args": {
			query: "SQL Query with args",
			args:  []interface{}{"arg1", "arg2"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pipelineStmt := NewPipelineStmt(test.query, test.args...)

			require.NotNil(t, pipelineStmt)
		})
	}
}

func TestPipelineStmt_Query(t *testing.T) {

	fakeRows := &sql.Rows{}

	pipelineStmt := &PipelineStmt{
		query: "SQL INSERT Query", args: []interface{}{},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTransaction := NewMockTransaction(controller)

	mockTransaction.EXPECT().Query("SQL INSERT Query").Return(
		fakeRows, nil,
	)

	sqlResult, err := pipelineStmt.Query(mockTransaction)

	require.NoError(t, err)
	require.NotNil(t, sqlResult)
}

func TestRunPipeline(t *testing.T) {

	fakeRows := &sql.Rows{}

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTransaction := NewMockTransaction(controller)

	mockTransaction.EXPECT().Query("Q1").Return(
		fakeRows, nil,
	)
	mockTransaction.EXPECT().Query("Q2").Return(
		nil, errors.Errorf("Error"),
	)

	tests := map[string]struct {
		pipelineStmt   *PipelineStmt
		expectingError bool
	}{
		"it should fail executing query": {
			pipelineStmt: &PipelineStmt{
				query: "Q2",
			}, expectingError: true,
		},
		"it should not fail": {
			pipelineStmt: &PipelineStmt{
				query: "Q1",
			}, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := RunPipeline(mockTransaction, test.pipelineStmt)

			if test.expectingError {
				require.Nil(t, res)
				require.Error(t, err)
			} else {
				require.NotNil(t, res)
				require.NoError(t, err)
			}
		})
	}
}
