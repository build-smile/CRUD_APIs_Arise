package port

import "time"

type MemberRepo interface {
	GetMembers() ([]Member, error)
	CreateMember(req Member) error
	GetMember(id int) (*Member, error)
	UpdateMember(id int, req Member) error
	DeleteMember(id int) error
}

type Member struct {
	Id        int       `gorm:"id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	IsActive  bool      `gorm:"is_active"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

// TableName overrides the table name used by User to `profiles`
func (Member) TableName() string {
	return "member"
}
