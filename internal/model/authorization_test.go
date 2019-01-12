package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewAuthorizationRequest(t *testing.T) {
	merchantID := uuid.New().String()
	cardID := uuid.New().String()
	amount := 10.0

	authRequest, err := NewAuthorizationRequest(merchantID, cardID, amount)

	require.NoError(t, err)
	require.NotNil(t, authRequest)
	require.Equal(t, authRequest.Merchant, merchantID)
	require.Equal(t, authRequest.Card, cardID)
	require.False(t, authRequest.Approved)
	require.Equal(t, authRequest.Amount, amount)
	require.Zero(t, authRequest.Reversed)
}

func TestAuthorizationRequest_Reverse(t *testing.T) {
	tests := map[string]struct {
		card            *AuthorizationRequest
		amountToReverse float64
		expectingError  bool
	}{
		"should fail due to insufficient reversable amount": {
			card:            &AuthorizationRequest{Amount: 10.0, Reversed: 3.0},
			amountToReverse: 10.0, expectingError: true,
		},
		"should increment successfully": {
			card:            &AuthorizationRequest{Amount: 10.0, Reversed: 3.0},
			amountToReverse: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.Revert(test.amountToReverse)

			if test.expectingError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAuthorizationRequest_Approve(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		authReq := &AuthorizationRequest{}

		authReq.Approve()

		require.True(t, authReq.Approved)
	})
}
