package domain

import "github.com/gin-gonic/gin"

type DeleteMemberSvc interface {
	Execute(c *gin.Context, req DeleteMemberReq) error
}
type DeleteMemberReq struct {
	Id int `json:"id"`
}

type DeleteMemberRes struct {
}
