package usecase

import (
	"clean-code/__mock__/repomock"
	"clean-code/__mock__/usecasemock"
	"clean-code/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductUseCaseTestSuite struct {
	suite.Suite
	prm *repomock.ProductRepoMock
	uuc *usecasemock.UomUseCaseMock
	puc ProductUseCase
}

func (suite *ProductUseCaseTestSuite) SetupTest() {
	suite.prm = new(repomock.ProductRepoMock)
	suite.uuc = new(usecasemock.UomUseCaseMock)
	suite.puc = NewProductUseCase(suite.prm, suite.uuc)
}

func TestProductUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ProductUseCaseTestSuite))
}

func (suite *ProductUseCaseTestSuite) TestGetByName_Failed() {

	suite.prm.On("FindByName", "product a").Return([]model.Product{}, errors.New("error"))
	gotProducts, gotErr := suite.puc.GetByName("product a")
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
	assert.NotEqual(suite.T(), []model.Product{}, gotProducts)
}

func (suite *ProductUseCaseTestSuite) TestGetByName_Success() {
	mockData := []model.Product{
		{
			ID:    "1",
			Name:  "product a",
			Price: 100,
			Uom: model.Uom{
				ID:   "1",
				Name: "kg",
			},
		},
	}
	suite.prm.On("FindByName", mockData[0].Name).Return(mockData, nil)
	gotProducts, gotErr := suite.puc.GetByName(mockData[0].Name)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotProducts)
}

func (suite *ProductUseCaseTestSuite) TestDelete_InvalidProductId() {
	suite.prm.On("FindById", "1").Return(model.Product{}, errors.New("error"))
	suite.prm.On("DeleteById", "1").Return(nil)
	gotErr := suite.puc.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestDelete_Failed() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}
	suite.prm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.prm.On("DeleteById", mockData.ID).Return(errors.New("failed to delete product"))
	gotErr := suite.puc.Delete(mockData.ID)
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestDelete_Success() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}
	suite.prm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.prm.On("DeleteById", mockData.ID).Return(nil)
	gotErr := suite.puc.Delete(mockData.ID)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestGetAll_Failed() {
	suite.prm.On("FindByAll").Return(nil, errors.New("data product is empty"))
	got, err := suite.puc.GetAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestGetAll_Success() {
	mockProducts := []model.Product{
		{
			ID:    "1",
			Name:  "product a",
			Price: 1000,
			Uom: model.Uom{
				ID:   "1",
				Name: "pcs",
			},
		},
	}

	suite.prm.On("FindByAll").Return(mockProducts, nil)
	got, err := suite.puc.GetAll()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockProducts, got)
	assert.Equal(suite.T(), 1, len(got))
}

func (suite *ProductUseCaseTestSuite) TestUpdate_InvalidProductId() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	suite.uuc.On("GetById", mockData.Uom.ID).Return(model.Uom{}, nil)
	mockData.ID = "2"
	suite.prm.On("FindById", mockData.ID).Return(model.Product{}, errors.New("failed to find product"))
	gotErr := suite.puc.Update(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestUpdate_InvalidUom() {
	suite.uuc.On("GetById", "yy").Return(model.Uom{}, errors.New("id uom is not found"))
	gotErr := suite.puc.Update(model.Product{ID: "1", Name: "product a", Price: 100, Uom: model.Uom{ID: "yy"}})
	assert.Error(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestUpdate_EmptyFail() {
	// ID is required
	suite.prm.On("UpdateById", model.Product{}).Return(errors.New("error"))
	gotErr := suite.puc.Update(model.Product{})
	assert.Error(suite.T(), gotErr)
	// Name is required
	suite.prm.On("UpdateById", model.Product{ID: "1"}).Return(errors.New("error"))
	gotErr = suite.puc.Update(model.Product{ID: "1"})
	assert.Error(suite.T(), gotErr)
	// price must be greater than zero
	suite.prm.On("UpdateById", model.Product{ID: "1", Name: "product a", Price: 0}).Return(errors.New("error"))
	gotErr = suite.puc.Update(model.Product{ID: "1", Name: "product a", Price: 0})
	assert.Error(suite.T(), gotErr)
}

func (suite *ProductUseCaseTestSuite) TestUpdate_Failed() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	suite.uuc.On("GetById", mockData.Uom.ID).Return(mockData.Uom, nil)
	suite.prm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.prm.On("UpdateById", mockData).Return(errors.New("failed to update product"))
	got := suite.puc.Update(mockData)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestUpdate_Success() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	suite.uuc.On("GetById", mockData.Uom.ID).Return(mockData.Uom, nil)
	suite.prm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.prm.On("UpdateById", mockData).Return(nil)
	got := suite.puc.Update(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestCreateNew_Success() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	suite.uuc.On("GetById", mockData.Uom.ID).Return(mockData.Uom, nil)
	suite.prm.On("Save", mockData).Return(nil)
	got := suite.puc.CreateNew(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestCreateNew_Failed() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	suite.uuc.On("GetById", mockData.Uom.ID).Return(mockData.Uom, nil)
	suite.prm.On("Save", mockData).Return(errors.New("failed to create new product"))
	got := suite.puc.CreateNew(mockData)
	assert.Error(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestCreateNew_InvalidUom() {
	suite.uuc.On("GetById", "YY").Return(model.Uom{}, errors.New("uom is not found"))
	got := suite.puc.CreateNew(model.Product{ID: "1", Name: "product 1", Price: 100, Uom: model.Uom{ID: "YY"}})
	assert.Error(suite.T(), got)
}

func (suite *ProductUseCaseTestSuite) TestCreateNew_EmptyFail() {
	// ID empty
	suite.prm.On("Save", model.Product{}).Return(errors.New("error"))
	gotErr := suite.puc.CreateNew(model.Product{})
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
	// name empty
	suite.prm.On("Save", model.Product{ID: "1"}).Return(errors.New("error"))
	gotErr = suite.puc.CreateNew(model.Product{ID: "1"})
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
	// price must greater than zero
	suite.prm.On("Save", model.Product{ID: "1", Name: "product a", Price: 0})
	gotErr = suite.puc.CreateNew(model.Product{ID: "1", Name: "product a", Price: 0})
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
}
