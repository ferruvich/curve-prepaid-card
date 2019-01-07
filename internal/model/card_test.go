package model

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
)

func TestNewCard(t *testing.T) {
	userID := uuid.New().String()

	card, err := NewCard(userID)

	require.NoError(t, err)
	require.NotNil(t, card)
	require.Equal(t, card.Owner, userID)
	require.Equal(t, card.AccountBalance, 0.0)
	require.Equal(t, card.AvailableBalance, 0.0)
}

func TestCard_IncrementAccountBalance(t *testing.T) {
	userID := uuid.New().String()

	card, err := NewCard(userID)

	require.NoError(t, err)
	require.NotNil(t, card)

	amountToIncrement := 10.0

	card.IncrementAccountBalance(amountToIncrement)

	require.Equal(t, card.AccountBalance, amountToIncrement)
	require.Equal(t, card.AvailableBalance, amountToIncrement)

	card.IncrementAccountBalance(amountToIncrement)

	require.Equal(t, card.AccountBalance, amountToIncrement*2)
	require.Equal(t, card.AvailableBalance, amountToIncrement*2)
}

func TestCard_ReverseAmountBlocked(t *testing.T) {

	tests := map[string]struct {
		card            *Card
		amountToReverse float64
		expectingError  bool
	}{
		"should fail due to high available balance": {
			card:            &Card{AccountBalance: 10.0, AvailableBalance: 3.0},
			amountToReverse: 10.0, expectingError: true,
		},
		"should increment successfully": {
			card:            &Card{AccountBalance: 10.0, AvailableBalance: 3.0},
			amountToReverse: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.ReverseAmountBlocked(test.amountToReverse)

			if test.expectingError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestCard_PayAmount(t *testing.T) {

	tests := map[string]struct {
		card           *Card
		amountToPay    float64
		expectingError bool
	}{
		"should fail due to low account balance": {
			card:        &Card{AccountBalance: 10.0, AvailableBalance: 3.0},
			amountToPay: 10.0, expectingError: true,
		},
		"should pay successfully": {
			card:        &Card{AccountBalance: 10.0, AvailableBalance: 3.0},
			amountToPay: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.PayAmount(test.amountToPay)

			if test.expectingError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCard_BlockAmount(t *testing.T) {

	tests := map[string]struct {
		card           *Card
		amountToBlock  float64
		expectingError bool
	}{
		"should fail due to low available balance": {
			card:          &Card{AvailableBalance: 3.0},
			amountToBlock: 10.0, expectingError: true,
		},
		"should pay successfully": {
			card:          &Card{AvailableBalance: 3.0},
			amountToBlock: 3.0, expectingError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.card.BlockAmount(test.amountToBlock)

			if test.expectingError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
