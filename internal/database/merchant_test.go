package database

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/testdata"
)

func TestNewMerchantdatabase(t *testing.T) {
	database, err := NewMerchantdatabase(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, database)
}
