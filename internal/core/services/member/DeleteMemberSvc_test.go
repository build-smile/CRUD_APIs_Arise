package services_test

import (
	"errors"
	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	services "github.com/build-smile/CRUD_APIs_Arise/internal/core/services/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestDeleteMember_Execute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	t.Run("success", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewDeleteMemberSvc(mockRepo)

		// Define the request
		req := domain.DeleteMemberReq{
			Id: 1,
		}

		// Set up the mock expectation
		mockRepo.On("DeleteMember", req.Id).Return(nil)

		// Execute the service method
		err := svc.Execute(c, req)

		// Assert no error and correct response
		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
	t.Run("not found", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewDeleteMemberSvc(mockRepo)

		// Define the request
		req := domain.DeleteMemberReq{
			Id: 1,
		}
		// Set up the mock expectation for not found error
		mockRepo.On("DeleteMember", req.Id).Return(errors.New("record not found"))

		// Execute the service method
		err := svc.Execute(c, req)

		// Assert error and nil response
		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewDeleteMemberSvc(mockRepo)

		// Define the request
		req := domain.DeleteMemberReq{
			Id: 1,
		}
		// Set up the mock expectation for not found error
		mockRepo.On("DeleteMember", req.Id).Return(gorm.ErrDuplicatedKey)

		// Execute the service method
		err := svc.Execute(c, req)

		// Assert error and nil response
		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})
}
