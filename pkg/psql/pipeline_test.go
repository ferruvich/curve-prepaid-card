package psql

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// fakeSQLResult implements sql.Result in order to make us able to
// mock Transaction functions such as Exec
type fakeSQLresult struct{}

// Since we do not care what these functions are returning
// we make them as simple as possible
func (f *fakeSQLresult) LastInsertId() (int64, error) {
	return 1, nil
}

func (f *fakeSQLresult) RowsAffected() (int64, error) {
	return 1, nil
}

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
			pipelineStmt := NewPipelineStmt(nil, test.query, test.args...)

			require.NotNil(t, pipelineStmt)
		})
	}
}

func TestPipelineStmt_Exec(t *testing.T) {
	pipelineStmt := &PipelineStmt{
		query: "SQL INSERT Query", args: []interface{}{},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTransaction := NewMockTransaction(controller)

	mockTransaction.EXPECT().Exec("SQL INSERT Query").Return(
		&fakeSQLresult{}, nil,
	)

	sqlResult, err := pipelineStmt.Exec(mockTransaction)

	require.NoError(t, err)
	require.NotNil(t, sqlResult)
}

func TestRunPipeline(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTransaction := NewMockTransaction(controller)

	mockTransaction.EXPECT().Exec("Q1").Return(
		&fakeSQLresult{}, nil,
	)
	mockTransaction.EXPECT().Exec("Q2").Return(
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
