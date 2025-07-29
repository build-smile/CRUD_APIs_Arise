package domain

import "github.com/gin-gonic/gin"

type UpdateMemberSvc interface {
	Execute(c *gin.Context, req UpdateMemberReq) error
}
type UpdateMemberReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateMemberRes struct {
}
