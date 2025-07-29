package services_test

import (
	"errors"
	services "github.com/build-smile/CRUD_APIs_Arise/internal/core/services/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/repositories"
	"testing"
	"time"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateMemberSvc_Execute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	t.Run("success", func(t *testing.T) {
		mockRepo := repositories.NewMemberRepoMock()

		svc := services.NewupdateMemberSvc(mockRepo)

		req := domain.UpdateMemberReq{
			Id:   1,
			Name: "Updated Name",
		}

		expectedMember := port.Member{
			Name:      req.Name,
			UpdatedAt: time.Now(),
		}

		mockRepo.On("UpdateMember", req.Id, mock.MatchedBy(func(m port.Member) bool {
			return m.Name == expectedMember.Name
		})).Return(nil)

		err := svc.Execute(c, req)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := repositories.NewMemberRepoMock()
		svc := services.NewupdateMemberSvc(mockRepo)

		req := domain.UpdateMemberReq{
			Id:   1,
			Name: "Updated Name",
		}

		expectedError := errors.New("database error")
		mockRepo.On("UpdateMember", req.Id, mock.MatchedBy(func(m port.Member) bool {
			return m.Name == req.Name
		})).Return(expectedError)

		err := svc.Execute(c, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
}
