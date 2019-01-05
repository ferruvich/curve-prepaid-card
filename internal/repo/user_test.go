package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserRepo(t *testing.T) {
	repo, err := NewUserRepo()

	require.NoError(t, err)
	require.NotNil(t, repo)
}
