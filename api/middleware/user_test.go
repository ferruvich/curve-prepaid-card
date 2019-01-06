package middleware

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/curve-challenge/internal/repo"
)

func TestNewUserMiddleware(t *testing.T) {
	userMiddleware, err := NewUserMiddleware()

	require.NoError(t, err)
	require.NotNil(t, userMiddleware)
}

func TestUserMiddleware_Create(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := repo.NewMockUser(controller)
	mockRepo.EXPECT().Write(gomock.Any()).Return(nil)

	userMiddleware := &UserMiddleware{
		repo: mockRepo,
	}

	user, err := userMiddleware.Create()

	require.NoError(t, err)
	require.NotNil(t, user)
}
