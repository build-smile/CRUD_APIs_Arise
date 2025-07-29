package services

import (
	"net/http"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/build-smile/CRUD_APIs_Arise/utils"
	"github.com/gin-gonic/gin"
)

type DeleteMemberSvc struct {
	repo port.MemberRepo
}

func NewDeleteMemberSvc(repo port.MemberRepo) domain.DeleteMemberSvc {
	return &DeleteMemberSvc{
		repo: repo,
	}
}

func (s DeleteMemberSvc) Execute(c *gin.Context, req domain.DeleteMemberReq) error {
	err := s.repo.DeleteMember(req.Id)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
