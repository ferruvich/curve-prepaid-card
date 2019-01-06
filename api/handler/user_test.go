package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserHandler(t *testing.T) {
	user, err := NewUserHandler()

	require.NotNil(t, user)
	require.NoError(t, err)
	require.IsType(t, &User{}, user)
	require.NotNil(t, user.(*User).middleware)
}
