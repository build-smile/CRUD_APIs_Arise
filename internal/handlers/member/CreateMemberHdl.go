package handlers

import (
	"net/http"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/gin-gonic/gin"
)

type CreateMemberHdl struct {
	service domain.CreateMemberSvc
}

func NewCreateMemberHdl(service domain.CreateMemberSvc) *CreateMemberHdl {
	return &CreateMemberHdl{
		service: service,
	}
}

func (h *CreateMemberHdl) Handle(c *gin.Context) {
	req := domain.CreateMemberReq{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	err = h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}
