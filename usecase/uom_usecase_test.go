package usecase

import (
	"clean-code/__mock__/repomock"
	"clean-code/model"
	"clean-code/model/dto"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UomUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UomRepoMock
	uuc UomUseCase
}

func (suite *UomUseCaseTestSuite) SetupTest() {
	suite.urm = new(repomock.UomRepoMock)
	suite.uuc = NewUomUseCase(suite.urm)
}

func TestUomUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UomUseCaseTestSuite))
}

func (suite *UomUseCaseTestSuite) TestDelete_InvalidUomId() {
	suite.urm.On("FindById", "1").Return(model.Uom{}, errors.New("error"))
	suite.urm.On("DeleteById", "1").Return(nil)
	gotErr := suite.uuc.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *UomUseCaseTestSuite) TestDelete_Failed() {
	mockData := model.Uom{
		ID:   "1",
		Name: "pcs",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.urm.On("DeleteById", mockData.ID).Return(errors.New("failed to delete uom"))
	got := suite.uuc.Delete(mockData.ID)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *UomUseCaseTestSuite) TestDelete_Success() {
	mockData := model.Uom{
		ID:   "1",
		Name: "pcs",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.urm.On("DeleteById", mockData.ID).Return(nil)
	got := suite.uuc.Delete(mockData.ID)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
	assert.Equal(suite.T(), nil, got)
}

func (suite *UomUseCaseTestSuite) TestPaging_Failed() {
	suite.urm.On("Paging", dto.PageRequest{
		Page: 1,
		Size: 0,
	}).Return(nil, dto.PageRequest{}, errors.New("error"))
	gotUoms, gotPaging, gotErr := suite.uuc.Paging(dto.PageRequest{})
	assert.Error(suite.T(), gotErr)
	assert.Nil(suite.T(), gotUoms)
	assert.Equal(suite.T(), dto.Paging{}, gotPaging)
}

func (suite *UomUseCaseTestSuite) TestPaging_Success() {
	mockUoms := []model.Uom{
		{
			ID:   "1",
			Name: "pcs",
		},
	}
	mockPaging := dto.Paging{
		Page:       1,
		Size:       1,
		TotalRows:  1,
		TotalPages: 1,
	}
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 1,
	}
	suite.urm.On("Paging", mockPageRequest).Return(mockUoms, mockPaging, nil)
	gotUoms, gotPaging, gotErr := suite.uuc.Paging(mockPageRequest)
	assert.Nil(suite.T(), gotErr)
	assert.Equal(suite.T(), mockUoms, gotUoms)
	assert.Equal(suite.T(), mockPaging, gotPaging)
}

func (suite *UomUseCaseTestSuite) TestGetAll_Failed() {
	suite.urm.On("FindByAll").Return([]model.Uom{}, errors.New("data is empty"))
	gotUoms, gotErr := suite.uuc.GetAll()
	assert.Error(suite.T(), gotErr)
	assert.Equal(suite.T(), []model.Uom{}, gotUoms)
}

func (suite *UomUseCaseTestSuite) TestGetAll_Success() {
	mockData := []model.Uom{
		{
			ID:   "1",
			Name: "pcs",
		},
	}
	suite.urm.On("FindByAll").Return(mockData, nil)
	gotUoms, gotErr := suite.uuc.GetAll()
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotUoms)
}

func (suite *UomUseCaseTestSuite) TestUpdate_InvalidUomId() {
	suite.urm.On("FindById", "1").Return(model.Uom{}, errors.New("id is not found"))
	suite.urm.On("UpdateById", model.Uom{ID: "1", Name: "pcs"}).Return(nil)
	gotErr := suite.uuc.Update(model.Uom{ID: "1", Name: "pcs"})
	assert.Error(suite.T(), gotErr)
}

func (suite *UomUseCaseTestSuite) TestUpdate_EmptyFail() {
	// ID is required
	suite.urm.On("UpdateById", model.Uom{}).Return(errors.New("id is required"))
	gotErr := suite.uuc.Update(model.Uom{})
	assert.Error(suite.T(), gotErr)
	// Name is required
	suite.urm.On("UpdateById", model.Uom{ID: "1"}).Return(errors.New("name is required"))
	gotErr = suite.uuc.Update(model.Uom{ID: "1"})
	assert.Error(suite.T(), gotErr)
	suite.urm.On("UpdateById", model.Uom{ID: "1", Name: "pcs"}).Return(nil)
	gotErr = suite.uuc.Update(model.Uom{ID: "1"})
	assert.Error(suite.T(), gotErr)
}

func (suite *UomUseCaseTestSuite) TestUpdate_Failed() {
	mockData := model.Uom{
		ID:   "1",
		Name: "pcs",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.urm.On("UpdateById", mockData).Return(errors.New("failed to update uom"))
	gotErr := suite.uuc.Update(mockData)
	assert.Error(suite.T(), gotErr)
	assert.NotNil(suite.T(), gotErr)
}

func (suite *UomUseCaseTestSuite) TestUpdate_Success() {
	mockData := model.Uom{
		ID:   "1",
		Name: "pcs",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, nil)
	suite.urm.On("UpdateById", mockData).Return(nil)
	got := suite.uuc.Update(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *UomUseCaseTestSuite) TestGetById_Failed() {
	mockData := model.Uom{
		ID:   "1",
		Name: "akbar",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, errors.New("failed to find uom"))
	gotUom, gotErr := suite.uuc.GetById(mockData.ID)
	assert.Error(suite.T(), gotErr)
	assert.NotEqual(suite.T(), mockData, gotUom)
}

func (suite *UomUseCaseTestSuite) TestGetById_Success() {
	mockData := model.Uom{
		ID:   "1",
		Name: "akbar",
	}
	suite.urm.On("FindById", mockData.ID).Return(mockData, nil)
	gotUom, gotErr := suite.uuc.GetById(mockData.ID)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotUom)
}

func (suite *UomUseCaseTestSuite) TestCreateNew_EmptyFail() {
	suite.urm.On("Save", model.Uom{}).Return(errors.New("id is required"))
	gotErr := suite.uuc.CreateNew(model.Uom{})
	assert.Error(suite.T(), gotErr)
	suite.urm.On("Save", model.Uom{ID: "1"}).Return(errors.New("name is required"))
	gotErr = suite.uuc.CreateNew(model.Uom{ID: "1"})
	assert.Error(suite.T(), gotErr)
}

func (suite *UomUseCaseTestSuite) TestCreateNew_Failed() {
	mockData := model.Uom{
		ID:   "1",
		Name: "akbar",
	}
	suite.urm.On("Save", mockData).Return(errors.New("failed to create new uom"))
	got := suite.uuc.CreateNew(mockData)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *UomUseCaseTestSuite) TestCreateNew_Success() {
	mockData := model.Uom{
		ID:   "1",
		Name: "akbar",
	}
	suite.urm.On("Save", mockData).Return(nil)
	got := suite.uuc.CreateNew(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}
