package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewMerchantHandler(t *testing.T) {
	merchant, err := NewMerchantHandler(testdata.GetMockContext(t))

	require.NotNil(t, merchant)
	require.NoError(t, err)
	require.IsType(t, &Merchant{}, merchant)
	require.NotNil(t, merchant.(*Merchant).middleware)
}
