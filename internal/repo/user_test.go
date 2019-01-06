package repo

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewUserRepo(t *testing.T) {
	repo, err := NewUserRepo(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, repo)
}
