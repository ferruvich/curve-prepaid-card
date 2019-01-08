package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewUserHandler(t *testing.T) {
	user, err := NewUserHandler(testdata.GetMockContext(t))

	require.NotNil(t, user)
	require.NoError(t, err)
	require.NotNil(t, user.middleware)
}
