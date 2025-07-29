package handlers

import (
	"net/http"

	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/gin-gonic/gin"
)

type GetMembersHdl struct {
	service domain.GetMembersSvc
}

func NewGetMembersHdl(service domain.GetMembersSvc) *GetMembersHdl {
	return &GetMembersHdl{
		service: service,
	}
}

func (h *GetMembersHdl) Handle(c *gin.Context) {
	res, err := h.service.Execute(c)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
