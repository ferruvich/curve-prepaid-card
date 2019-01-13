package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

func TestAuthorizationRequest_Write(t *testing.T) {

	authReq, _ := model.NewAuthorizationRequest("merchantID", "cardID", 10.0)
	authReq.Approve()
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), authReq.ID, authReq.Merchant, authReq.Card,
			authReq.Approved, authReq.Amount, authReq.Reversed, authReq.Captured, authReq.Refunded,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing authorization"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		authRequest := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		err := authRequest.Write(authReq)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), authReq.ID, authReq.Merchant, authReq.Card,
			authReq.Approved, authReq.Amount, authReq.Reversed, authReq.Captured, authReq.Refunded,
		).Return(
			&pipelineStmt{},
		)
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, nil,
		)
		mockDB.EXPECT().GetConnection().Return(db)

		authRequest := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		err := authRequest.Write(authReq)

		require.NoError(t, err)
	})
}

func TestAuthorizationRequest_Read(t *testing.T) {

	authReq := &model.AuthorizationRequest{ID: "id", Approved: true}
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)
		mockDB.EXPECT().newPipelineStmt(gomock.Any(), authReq.ID).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(
			nil, errors.New("error writing auth request"),
		)
		mockDB.EXPECT().GetConnection().Return(db)

		authReqDB := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		resAuthReq, err := authReqDB.Read("id")

		require.Nil(t, resAuthReq)
		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)

		mockDB.EXPECT().newPipelineStmt(gomock.Any(), authReq.ID).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(nil, nil)
		mockDB.EXPECT().GetConnection().Return(db)

		authReqDB := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		resAuthReq, err := authReqDB.Read("id")

		require.NotNil(t, resAuthReq)
		require.NoError(t, err)
	})
}

func TestAuthorizationRequest_Update(t *testing.T) {

	authReq, _ := model.NewAuthorizationRequest("merchantID", "cardID", 10.0)
	authReq.Approve()
	db := &sql.DB{}

	t.Run("should return error due to error on db", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)
		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), authReq.ID, authReq.Merchant, authReq.Card,
			authReq.Approved, authReq.Amount, authReq.Reversed, authReq.Captured, authReq.Refunded,
		).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(nil, errors.New("error writing card"))
		mockDB.EXPECT().GetConnection().Return(db)

		authReqDB := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		err := authReqDB.Update(authReq)

		require.Error(t, err)
	})

	t.Run("should run", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDB := NewMockDataBase(controller)
		mockDB.EXPECT().newPipelineStmt(
			gomock.Any(), authReq.ID, authReq.Merchant, authReq.Card,
			authReq.Approved, authReq.Amount, authReq.Reversed, authReq.Captured, authReq.Refunded,
		).Return(&pipelineStmt{})
		mockDB.EXPECT().withTransaction(db, gomock.Any()).Return(nil, nil)
		mockDB.EXPECT().GetConnection().Return(db)

		authReqDB := &AuthorizationRequestDataBase{
			service: mockDB,
		}

		err := authReqDB.Update(authReq)

		require.NoError(t, err)
	})
}
