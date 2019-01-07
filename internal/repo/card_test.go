package repo

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewCardRepo(t *testing.T) {
	repo, err := NewCardRepo(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, repo)
}