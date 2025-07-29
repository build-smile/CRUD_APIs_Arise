package services

import (
	domain "github.com/build-smile/CRUD_APIs_Arise/internal/core/domain/member"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/gin-gonic/gin"
)

type getMembersSvc struct {
	repo port.MemberRepo
}

func NewGetMembersSvc(repo port.MemberRepo) domain.GetMembersSvc {
	return &getMembersSvc{
		repo: repo,
	}
}

func (s getMembersSvc) Execute(c *gin.Context) ([]domain.GetMembersRes, error) {
	orderRes, err := s.repo.GetMembers()
	res := s.buildRes(orderRes)
	return res, err
}

func (s getMembersSvc) buildRes(m []port.Member) []domain.GetMembersRes {
	memberList := make([]domain.GetMembersRes, 0)
	for _, it := range m {
		member := domain.GetMembersRes{Id: it.Id, Name: it.Name}
		memberList = append(memberList, member)
	}
	return memberList
}
