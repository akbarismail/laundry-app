package repomock

import (
	"clean-code/model"
	"clean-code/model/dto"

	"github.com/stretchr/testify/mock"
)

type UomRepoMock struct {
	mock.Mock
}

func (u *UomRepoMock) Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error) {
	args := u.Called(payload)
	if args.Get(2) != nil {
		return nil, dto.Paging{}, args.Error(2)
	}
	return args.Get(0).([]model.Uom), args.Get(1).(dto.Paging), nil
}

func (u *UomRepoMock) FindByName(name string) ([]model.Uom, error) {
	args := u.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Uom), nil
}

func (u *UomRepoMock) UpdateById(uom model.Uom) error {
	return u.Called(uom).Error(0)
}

func (u *UomRepoMock) DeleteById(id string) error {
	return u.Called(id).Error(0)
}

func (u *UomRepoMock) FindByAll() ([]model.Uom, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Uom), nil
}

func (u *UomRepoMock) FindById(id string) (model.Uom, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Uom{}, args.Error(1)
	}
	return args.Get(0).(model.Uom), nil
}

func (u *UomRepoMock) Save(uom model.Uom) error {
	return u.Called(uom).Error(0)
}
