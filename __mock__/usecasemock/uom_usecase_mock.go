package usecasemock

import (
	"clean-code/model"
	"clean-code/model/dto"

	"github.com/stretchr/testify/mock"
)

type UomUseCaseMock struct {
	mock.Mock
}

func (u *UomUseCaseMock) Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error) {
	args := u.Called(payload)
	if args.Get(2) != nil {
		return nil, dto.Paging{}, args.Error(2)
	}
	return args.Get(0).([]model.Uom), args.Get(1).(dto.Paging), nil
}

func (u *UomUseCaseMock) GetByName(name string) ([]model.Uom, error) {
	args := u.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Uom), nil
}

func (u *UomUseCaseMock) CreateNew(payload model.Uom) error {
	return u.Called(payload).Error(0)
}

func (u *UomUseCaseMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

func (u *UomUseCaseMock) GetAll() ([]model.Uom, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Uom), nil
}

func (u *UomUseCaseMock) GetById(id string) (model.Uom, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.Uom{}, args.Error(1)
	}
	return args.Get(0).(model.Uom), nil
}

func (u *UomUseCaseMock) Update(payload model.Uom) error {
	return u.Called(payload).Error(0)
}
