package repo

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewMerchantRepo(t *testing.T) {
	repo, err := NewMerchantRepo(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, repo)
}
