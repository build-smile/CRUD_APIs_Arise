package services_test

import (
	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	services "github.com/build-smile/CRUD_APIs_Arise/internal/core/services/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestMembers_Execute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	t.Run("success", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMembersSvc(mockRepo)

		// Define the request
		req := domain.GetMemberReq{
			Id: 1,
		}

		// Define the expected member
		expectedMember := []port.Member{
			{
				Id:   req.Id,
				Name: "Test Member",
			},
		}

		// Set up the mock expectation
		mockRepo.On("GetMembers").Return(expectedMember, nil)

		// Execute the service method
		_, err := svc.Execute(c)

		// Assert no error and correct response
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
	t.Run("not found", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMembersSvc(mockRepo)

		// Define the request

		// Set up the mock expectation for not found error
		mockRepo.On("GetMembers").Return(nil, gorm.ErrRecordNotFound)

		// Execute the service method
		_, err := svc.Execute(c)

		// Assert error and nil response
		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMembersSvc(mockRepo)

		// Define the request

		// Set up the mock expectation for not found error
		mockRepo.On("GetMembers").Return(nil, gorm.ErrDuplicatedKey)

		// Execute the service method
		_, err := svc.Execute(c)

		// Assert error and nil response
		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}
