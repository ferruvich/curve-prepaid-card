package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/testdata"
)

func TestNewAuthorizationRequestHandler(t *testing.T) {
	authReq, err := NewAuthoziationRequestHandler(testdata.GetMockContext(t))

	require.NotNil(t, authReq)
	require.NoError(t, err)
	require.NotNil(t, authReq.middleware)
}
