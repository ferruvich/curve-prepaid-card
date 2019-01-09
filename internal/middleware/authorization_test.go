package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/testdata"
)

const (
	merchantID      = "somemerchantID"
	amountToRequest = 10.0
)

func TestNewAuthorizationRequestMiddleware(t *testing.T) {
	authReqMiddleware, err := NewAuthorizationRequestMiddleware(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, authReqMiddleware)
}

func TestAuthorizationRequestMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCarddatabase := database.NewMockAuthorizationRequest(controller)
	mockCarddatabase.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	cardMiddleware := &AuthorizationRequestMiddleware{
		database: mockCarddatabase,
	}
	card, err := cardMiddleware.Create(context.Background(), merchantID, cardID, amountToRequest)

	require.NoError(t, err)
	require.NotNil(t, card)
}
