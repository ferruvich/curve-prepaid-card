package middleware

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
	"github.com/ferruvich/curve-challenge/testdata"
)

const (
	userID = "someID"
)

func TestNewUserMiddleware(t *testing.T) {
	userMiddleware, err := NewUserMiddleware(testdata.GetMockContext(t))

	require.NoError(t, err)
	require.NotNil(t, userMiddleware)
}

func TestUserMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := repo.NewMockUser(controller)
	mockRepo.EXPECT().Write(
		context.Background(),
		gomock.Any(),
	).Return(nil)

	userMiddleware := &UserMiddleware{
		repo: mockRepo,
	}

	user, err := userMiddleware.Create(context.Background())

	require.NoError(t, err)
	require.NotNil(t, user)
}

func TestUserMiddleware_Read(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := repo.NewMockUser(controller)
	mockRepo.EXPECT().Read(
		context.Background(),
		userID,
	).Return(&model.User{ID: userID}, nil)

	userMiddleware := &UserMiddleware{
		repo: mockRepo,
	}

	user, err := userMiddleware.Read(context.Background(), userID)

	require.NoError(t, err)
	require.NotNil(t, user)
}
