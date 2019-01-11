package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

func TestNewMiddleware(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		db := &database.Service{}

		middleware := NewMiddleware(db)

		require.NotNil(t, middleware)
		require.NotNil(t, middleware.(*Service).db)
		require.Equal(t, middleware.(*Service).db, db)
	})
}

func TestMiddleware_DataBase(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		db := &database.Service{}

		middleware := &Service{
			db: db,
		}

		resDB := middleware.DataBase()

		require.NotNil(t, resDB)
		require.Equal(t, middleware.db, resDB)
	})
}

func TestMiddleware_AuthorizationRequest(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		middleware := &Service{}

		authReq := middleware.AuthorizationRequest()

		require.NotNil(t, authReq)
		require.Equal(t, middleware, authReq.(*AuthorizationRequestMiddleware).middleware)
	})
}

func TestMiddleware_Card(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		middleware := &Service{}

		card := middleware.Card()

		require.NotNil(t, card)
		require.Equal(t, middleware, card.(*CardMiddleware).middleware)
	})
}

func TestMiddleware_Merchant(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		middleware := &Service{}

		merchant := middleware.Merchant()

		require.NotNil(t, merchant)
		require.Equal(t, middleware, merchant.(*MerchantMiddleware).middleware)
	})
}

func TestMiddleware_User(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		middleware := &Service{}

		user := middleware.User()

		require.NotNil(t, user)
		require.Equal(t, middleware, user.(*UserMiddleware).middleware)
	})
}
