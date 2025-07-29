package services_test

import (
	"errors"
	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	services "github.com/build-smile/CRUD_APIs_Arise/internal/core/services/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteMemberSvc_Execute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	t.Run("success", func(t *testing.T) {
		mockRepo := repositories.NewMemberRepoMock()

		svc := services.NewCreateMemberSvc(mockRepo)

		req := domain.CreateMemberReq{
			Name: "Updated Name",
		}

		mockRepo.On("CreateMember", mock.Anything).Return(nil)

		err := svc.Execute(c, req)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewCreateMemberSvc(mockRepo)

		req := domain.CreateMemberReq{
			Name: "Test",
		}

		expectedError := errors.New("database error")
		mockRepo.On("CreateMember", mock.Anything).Return(expectedError)

		err := svc.Execute(c, req)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
