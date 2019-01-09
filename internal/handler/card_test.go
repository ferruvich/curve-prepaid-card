package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/testdata"
)

func TestNewCardHandler(t *testing.T) {
	card, err := NewCardHandler(testdata.GetMockContext(t))

	require.NotNil(t, card)
	require.NoError(t, err)
	require.NotNil(t, card.middleware)
}
