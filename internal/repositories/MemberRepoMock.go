package repositories

import (
	"github.com/build-smile/CRUD_APIs_Arise/internal/core/port"
	"github.com/stretchr/testify/mock"
)

// MemberRepoMock is a mock implementation of port.MemberRepo interface
type MemberRepoMock struct {
	mock.Mock
}

func NewMemberRepoMock() *MemberRepoMock {
	return &MemberRepoMock{}
}

// GetMember mocks the GetMember method
func (m *MemberRepoMock) GetMember(id int) (*port.Member, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*port.Member), args.Error(1)
}

// GetMembers mocks the GetMembers method
func (m *MemberRepoMock) GetMembers() ([]port.Member, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]port.Member), args.Error(1)
}

// CreateMember mocks the CreateMember method
func (m *MemberRepoMock) CreateMember(member port.Member) error {
	args := m.Called(member)
	return args.Error(0)
}

// UpdateMember mocks the UpdateMember method
func (m *MemberRepoMock) UpdateMember(id int, member port.Member) error {
	args := m.Called(id, member)
	return args.Error(0)
}

// DeleteMember mocks the DeleteMember method
func (m *MemberRepoMock) DeleteMember(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
