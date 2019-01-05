package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserHandler(t *testing.T) {
	user := NewUserHandler()

	require.NotNil(t, user)
	require.IsType(t, &User{}, user)
	require.NotNil(t, user.(*User).middleware)
}
