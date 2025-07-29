package handlers

import (
	"net/http"
	"strconv"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/gin-gonic/gin"
)

type UpdateMemberHdl struct {
	service domain.UpdateMemberSvc
}

func NewUpdateMemberHdl(service domain.UpdateMemberSvc) *UpdateMemberHdl {
	return &UpdateMemberHdl{
		service: service,
	}
}

func (h *UpdateMemberHdl) Handle(c *gin.Context) {
	req := domain.UpdateMemberReq{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	req.Id = id
	err = h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}
