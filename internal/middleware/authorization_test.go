package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestAuthorizationRequest_Create(t *testing.T) {

	merchantID := "merchant_ID"
	cardID := "card_ID"
	amount := 10.0

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Write(gomock.Any()).Return(
			errors.New("error"),
		)

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(cardID).Return(&model.Card{
			ID: cardID, AccountBalance: amount, AvailableBalance: amount,
		}, nil)
		mockCardDB.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(3)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		authReq, err := mockAuthRequestMiddleware.Create(
			merchantID, cardID, amount,
		)

		require.Nil(t, authReq)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Write(gomock.Any()).Return(
			nil,
		)

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(cardID).Return(&model.Card{
			ID: cardID, AccountBalance: amount, AvailableBalance: amount,
		}, nil)
		mockCardDB.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(3)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		authReq, err := mockAuthRequestMiddleware.Create(
			merchantID, cardID, amount,
		)

		require.NotNil(t, authReq)
		require.True(t, authReq.Approved)
		require.NoError(t, err)
	})
}

func TestAuthorizationRequest_Revert(t *testing.T) {

	authReq := &model.AuthorizationRequest{
		ID: "authID", Merchant: "merchantID", Card: "cardID",
		Amount: 10.0,
	}

	t.Run("should fail due to db error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockAuthReqDB.EXPECT().Update(gomock.Any()).Return(errors.New("error"))

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(authReq.Card).Return(&model.Card{
			ID: authReq.Card, AccountBalance: authReq.Amount,
			AvailableBalance: 0.0,
		}, nil)
		mockCardDB.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB).Times(2)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(4)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		err := mockAuthRequestMiddleware.Revert(
			authReq.ID, 5.0,
		)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockAuthReqDB := database.NewMockAuthorizationRequest(controller)
		mockAuthReqDB.EXPECT().Read(authReq.ID).Return(authReq, nil)
		mockAuthReqDB.EXPECT().Update(gomock.Any()).Return(nil)

		mockCardDB := database.NewMockCard(controller)
		mockCardDB.EXPECT().Read(authReq.Card).Return(&model.Card{
			ID: authReq.Card, AccountBalance: authReq.Amount,
			AvailableBalance: 0.0,
		}, nil)
		mockCardDB.EXPECT().Update(gomock.Any()).Return(nil)

		mockDB := database.NewMockDataBase(controller)
		mockDB.EXPECT().AuthorizationRequest().Return(mockAuthReqDB).Times(2)
		mockDB.EXPECT().Card().Return(mockCardDB).Times(2)

		mockMiddleware := NewMockMiddleware(controller)
		mockMiddleware.EXPECT().DataBase().Return(mockDB).Times(4)

		mockAuthRequestMiddleware := &AuthorizationRequestMiddleware{
			middleware: mockMiddleware,
		}

		err := mockAuthRequestMiddleware.Revert(
			authReq.ID, 5.0,
		)

		require.NoError(t, err)
	})
}
