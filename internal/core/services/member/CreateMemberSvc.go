package services

import (
	"net/http"
	"time"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/build-smile/CRUD_APIs_Arise/utils"
	"github.com/gin-gonic/gin"
)

type createMemberSvc struct {
	repo port.MemberRepo
}

func NewCreateMemberSvc(repo port.MemberRepo) domain.CreateMemberSvc {
	return &createMemberSvc{
		repo: repo,
	}
}

func (s createMemberSvc) Execute(c *gin.Context, req domain.CreateMemberReq) error {
	memberReq := port.Member{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	err := s.repo.CreateMember(memberReq)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
