package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser()

	require.NoError(t, err)
	require.NotNil(t, user)
}
