package domain

import "github.com/gin-gonic/gin"

type GetMemberSvc interface {
	Execute(c *gin.Context, req GetMemberReq) (*GetMemberRes, error)
}
type GetMemberReq struct {
	Id int
}

type GetMemberRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
