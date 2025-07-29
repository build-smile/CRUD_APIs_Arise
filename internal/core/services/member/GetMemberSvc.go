package services

import (
	"errors"
	"net/http"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/build-smile/CRUD_APIs_Arise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getMemberSvc struct {
	repo port.MemberRepo
}

func NewGetMemberSvc(repo port.MemberRepo) domain.GetMemberSvc {
	return &getMemberSvc{
		repo: repo,
	}
}

func (s getMemberSvc) Execute(c *gin.Context, req domain.GetMemberReq) (*domain.GetMemberRes, error) {
	member, err := s.repo.GetMember(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, err.Error())
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	return &domain.GetMemberRes{
		Id:   member.Id,
		Name: member.Name,
	}, err
}
