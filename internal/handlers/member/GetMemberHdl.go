package handlers

import (
	"net/http"
	"strconv"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/gin-gonic/gin"
)

type GetMemberHdl struct {
	service domain.GetMemberSvc
}

func NewGetMemberHdl(service domain.GetMemberSvc) *GetMemberHdl {
	return &GetMemberHdl{
		service: service,
	}
}

func (h *GetMemberHdl) Handle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	req := domain.GetMemberReq{
		Id: id,
	}
	res, err := h.service.Execute(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
