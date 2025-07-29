package services

import (
	"net/http"
	"time"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/build-smile/CRUD_APIs_Arise/utils"
	"github.com/gin-gonic/gin"
)

type updateMemberSvc struct {
	repo port.MemberRepo
}

func NewupdateMemberSvc(repo port.MemberRepo) domain.UpdateMemberSvc {
	return &updateMemberSvc{
		repo: repo,
	}
}

func (s *updateMemberSvc) Execute(_ *gin.Context, req domain.UpdateMemberReq) error {
	memberReq := port.Member{
		Name:      req.Name,
		UpdatedAt: time.Now(),
	}
	err := s.repo.UpdateMember(req.Id, memberReq)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
