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

func TestGetMember_Execute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	t.Run("success", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMemberSvc(mockRepo)

		// Define the request
		req := domain.GetMemberReq{
			Id: 1,
		}

		// Define the expected member
		expectedMember := port.Member{
			Id:   req.Id,
			Name: "Test Member",
		}

		// Set up the mock expectation
		mockRepo.On("GetMember", req.Id).Return(&expectedMember, nil)

		// Execute the service method
		res, err := svc.Execute(c, req)

		// Assert no error and correct response
		assert.NoError(t, err)
		assert.Equal(t, expectedMember.Id, res.Id)
		assert.Equal(t, expectedMember.Name, res.Name)

		mockRepo.AssertExpectations(t)
	})
	t.Run("not found", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMemberSvc(mockRepo)

		// Define the request
		req := domain.GetMemberReq{
			Id: 1,
		}

		// Set up the mock expectation for not found error
		mockRepo.On("GetMember", req.Id).Return(nil, gorm.ErrRecordNotFound)

		// Execute the service method
		res, err := svc.Execute(c, req)

		// Assert error and nil response
		assert.Error(t, err)
		assert.Nil(t, res)

		mockRepo.AssertExpectations(t)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		// Mock the repository and service
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewGetMemberSvc(mockRepo)

		// Define the request
		req := domain.GetMemberReq{
			Id: 1,
		}

		// Set up the mock expectation for not found error
		mockRepo.On("GetMember", req.Id).Return(nil, gorm.ErrDuplicatedKey)

		// Execute the service method
		res, err := svc.Execute(c, req)

		// Assert error and nil response
		assert.Error(t, err)
		assert.Nil(t, res)

		mockRepo.AssertExpectations(t)
	})

}
