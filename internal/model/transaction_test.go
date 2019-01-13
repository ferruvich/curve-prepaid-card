package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		tx, err := newTransaction(
			"sender", "receiver", 10.0,
		)

		require.NoError(t, err)
		require.NotNil(t, tx)
	})
}

func TestNewPaymentTransaction(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		tx, err := NewPaymentTransaction(
			"sender", "receiver", 10.0,
		)

		require.NoError(t, err)
		require.NotNil(t, tx)
		require.Equal(t, tx.Type, payment)
	})
}

func TestNewRefundTransaction(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		tx, err := NewRefundTransaction(
			"sender", "receiver", 10.0,
		)

		require.NoError(t, err)
		require.NotNil(t, tx)
		require.Equal(t, tx.Type, refund)
	})
}
