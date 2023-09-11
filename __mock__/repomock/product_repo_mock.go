package repomock

import (
	"clean-code/model"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

func (p *ProductRepoMock) FindByName(name string) ([]model.Product, error) {
	args := p.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), nil
}

func (p *ProductRepoMock) DeleteById(id string) error {
	return p.Called(id).Error(0)
}

func (p *ProductRepoMock) FindByAll() ([]model.Product, error) {
	args := p.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), nil
}

func (p *ProductRepoMock) FindById(id string) (model.Product, error) {
	args := p.Called(id)
	if args.Get(1) != nil {
		return model.Product{}, args.Error(1)
	}
	return args.Get(0).(model.Product), nil
}

func (p *ProductRepoMock) Save(product model.Product) error {
	return p.Called(product).Error(0)
}

func (p *ProductRepoMock) UpdateById(product model.Product) error {
	// args := p.Called(product)
	// if args.Get(0) != nil {
	// 	return args.Error(0)
	// }
	// return nil

	return p.Called(product).Error(0)
}
