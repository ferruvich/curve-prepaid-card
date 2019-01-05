package psql

import (
	"testing"

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

func TestPipelineStmt_Exec(t *testing.T) {
	// TODO
}

func TestRunPipeline(t *testing.T) {
	// TODO
}
