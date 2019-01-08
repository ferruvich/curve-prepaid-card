package repo

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewAuthorizationRequestRepo(t *testing.T) {
	repo, err := NewAuthorizationRequestRepo(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, repo)
}
