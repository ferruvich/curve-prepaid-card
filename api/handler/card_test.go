package handler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/testdata"
)

func TestNewCardHandler(t *testing.T) {
	card, err := NewCardHandler(testdata.GetMockContext(t))

	require.NotNil(t, card)
	require.NoError(t, err)
	require.IsType(t, &Card{}, card)
	require.NotNil(t, card.(*Card).middleware)
}
