package domain

import "github.com/gin-gonic/gin"

type CreateMemberSvc interface {
	Execute(c *gin.Context, req CreateMemberReq) error
}
type CreateMemberReq struct {
	Name      string `json:"name"`
	GroupCode int    `json:"groupCode"`
}

type CreateMemberRes struct {
}
