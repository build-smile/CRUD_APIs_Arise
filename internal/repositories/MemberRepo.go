package repositories

import (
	"github.com/build-smile/CRUD_APIs_Arise/infrastructure"
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"gorm.io/gorm"
)

type MemberRepo struct {
	db *gorm.DB
}

func NewMemberRepo() port.MemberRepo {
	return &MemberRepo{
		db: infrastructure.DB,
	}
}

func (r MemberRepo) GetMembers() ([]port.Member, error) {
	var res []port.Member
	err := r.db.Table("member").Where("is_active = ?", true).
		Order("id ASC").
		Find(&res)

	return res, err.Error
}

func (r MemberRepo) GetMember(id int) (*port.Member, error) {
	var res *port.Member
	err := r.db.Table("member").First(&res, id).Error
	return res, err
}

func (r MemberRepo) CreateMember(req port.Member) error {
	err := r.db.Table("member").Create(&req)
	return err.Error
}

func (r MemberRepo) UpdateMember(id int, req port.Member) error {
	err := r.db.Table("member").Where("id = ?", id).Updates(&req).Error
	if err != nil {
		return err
	}
	return nil
}
func (r MemberRepo) DeleteMember(id int) error {
	err := r.db.Table("member").Where("id = ?", id).Update("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}
