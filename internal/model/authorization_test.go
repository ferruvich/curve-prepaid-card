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

func TestAuthorizationRequest_Revert(t *testing.T) {
	tests := map[string]struct {
		card            *AuthorizationRequest
		amountToReverse float64
		expectingError  bool
	}{
		"should fail due to insufficient reversable amount": {
			card: &AuthorizationRequest{
				Amount: 10.0, Reversed: 3.0, Approved: true,
			}, amountToReverse: 10.0, expectingError: true,
		},
		"should increment successfully": {
			card: &AuthorizationRequest{
				Amount: 10.0, Reversed: 3.0, Approved: true,
			}, amountToReverse: 3.0, expectingError: false,
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

func TestAuthorizationRequest_Capture(t *testing.T) {
	tests := map[string]struct {
		card            *AuthorizationRequest
		amountToCapture float64
		expectingError  bool
	}{
		"should fail due to insufficient capturable amount": {
			card: &AuthorizationRequest{
				Amount: 10.0, Reversed: 3.0, Captured: 0.0, Approved: true,
			},
			amountToCapture: 10.0, expectingError: true,
		},
		"should capture successfully": {
			card: &AuthorizationRequest{
				Amount: 10.0, Reversed: 3.0, Captured: 0.0, Approved: true,
			},
			amountToCapture: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.Capture(test.amountToCapture)

			if test.expectingError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAuthorizationRequest_Refund(t *testing.T) {
	tests := map[string]struct {
		card           *AuthorizationRequest
		amountToRefund float64
		expectingError bool
	}{
		"should fail due to insufficient refundable amount": {
			card: &AuthorizationRequest{
				Amount: 10.0, Captured: 5.0, Approved: true,
			}, amountToRefund: 6.0, expectingError: true,
		},
		"should refund successfully": {
			card: &AuthorizationRequest{
				Amount: 10.0, Captured: 5.0, Approved: true,
			}, amountToRefund: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.Refund(test.amountToRefund)

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
