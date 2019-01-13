package server

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

func TestServer_Routers(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		routers := server.Routers()

		require.NotNil(t, routers)
	})
}

func TestServer_DataBase(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{
			db: &database.Service{},
		}

		db := server.DataBase()

		require.NotNil(t, db)
		require.Equal(t, db, server.db)
	})
}

func TestServer_NewAuthorizationRequestHandler(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		authReq := server.NewAuthorizationRequestHandler()

		require.NotNil(t, authReq)
		require.Equal(t, authReq.(*AuthorizationRequestHandler).server, server)
	})
}

func TestServer_NewCardHandler(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		card := server.NewCardHandler()

		require.NotNil(t, card)
		require.Equal(t, card.(*CardHandler).server, server)
	})
}

func TestServer_NewMerchantHandler(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		merchant := server.NewMerchantHandler()

		require.NotNil(t, merchant)
		require.Equal(t, merchant.(*MerchantHandler).server, server)
	})
}

func TestServer_NewTransactionHandler(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		tx := server.NewTransactionHandler()

		require.NotNil(t, tx)
		require.Equal(t, tx.(*TransactionHandler).server, server)
	})
}

func TestServer_NewUserHandler(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		server := &Service{}

		user := server.NewUserHandler()

		require.NotNil(t, user)
		require.Equal(t, user.(*UserHandler).server, server)
	})
}
