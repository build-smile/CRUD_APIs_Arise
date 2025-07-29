package domain

import "github.com/gin-gonic/gin"

type GetMembersSvc interface {
	Execute(c *gin.Context) ([]GetMembersRes, error)
}

type GetMembersRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
