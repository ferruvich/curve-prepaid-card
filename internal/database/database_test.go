package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataBase_AuthorizationRequest(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		authReq := service.AuthorizationRequest()

		require.NotNil(t, authReq)
		require.Equal(t, service, authReq.(*AuthorizationRequestDataBase).service)
	})
}

func TestDataBase_Card(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		card := service.Card()

		require.NotNil(t, card)
		require.Equal(t, service, card.(*CardDataBase).service)
	})
}

func TestDataBase_Merchant(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		merchant := service.Merchant()

		require.NotNil(t, merchant)
		require.Equal(t, service, merchant.(*MerchantDataBase).service)
	})
}

func TestDataBase_Transaction(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		tx := service.Transaction()

		require.NotNil(t, tx)
		require.Equal(t, service, tx.(*TransactionDataBase).service)
	})
}

func TestDataBase_User(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		service := &Service{}

		user := service.User()

		require.NotNil(t, user)
		require.Equal(t, service, user.(*UserDataBase).service)
	})
}
