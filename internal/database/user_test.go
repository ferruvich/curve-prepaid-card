package database

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/testdata"
)

func TestNewUserdatabase(t *testing.T) {
	database, err := NewUserdatabase(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, database)
}
