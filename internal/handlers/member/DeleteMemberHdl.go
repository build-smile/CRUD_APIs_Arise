package handlers

import (
	"net/http"
	"strconv"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/gin-gonic/gin"
)

type DeleteMemberHdl struct {
	service domain.DeleteMemberSvc
}

func NewDeleteMemberHdl(service domain.DeleteMemberSvc) *DeleteMemberHdl {
	return &DeleteMemberHdl{
		service: service,
	}
}

func (h *DeleteMemberHdl) Handle(c *gin.Context) {
	req := domain.DeleteMemberReq{}
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
