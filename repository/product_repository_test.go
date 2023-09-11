package repository

import (
	"clean-code/model"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    ProductRepository
}

func (suite *ProductRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewProductRepository(suite.mockDB)
}

func TestProductRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepoTestSuite))
}

func (suite *ProductRepoTestSuite) TestSave_Success() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}
	expectedSql := "INSERT INTO product"
	suite.mockSQL.ExpectExec(expectedSql).WithArgs(mockData.ID, mockData.Name, mockData.Price, mockData.Uom.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	got := suite.repo.Save(mockData)
	assert.Nil(suite.T(), got)
	assert.NoError(suite.T(), got)
}

func (suite *ProductRepoTestSuite) TestSave_Fail() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}
	expectedSql := "INSERT INTO product"
	suite.mockSQL.ExpectExec(expectedSql).WithArgs(mockData.ID, mockData.Name, mockData.Price, mockData.Uom.ID).WillReturnError(errors.New("error"))

	got := suite.repo.Save(mockData)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *ProductRepoTestSuite) TestFindAll_Fail() {
	expectedQuery := "SELECT p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id"
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnError(errors.New("error"))
	gotProduct, gotErr := suite.repo.FindByAll()
	assert.Error(suite.T(), gotErr)
	assert.Nil(suite.T(), gotProduct)
}

func (suite *ProductRepoTestSuite) TestFindById_Success() {
	mockData := model.Product{
		ID:    "1",
		Name:  "product a",
		Price: 1000,
		Uom: model.Uom{
			ID:   "1",
			Name: "pcs",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "price", "uom_id", "uom_name"})
	rows.AddRow(mockData.ID, mockData.Name, mockData.Price, mockData.Uom.ID, mockData.Uom.Name)

	expectedQuery := "SELECT p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id WHERE p.id=$1"
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(mockData.ID).WillReturnRows(rows)
	gotProduct, gotErr := suite.repo.FindById(mockData.ID)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotProduct)
}

func (suite *ProductRepoTestSuite) TestFindById_Fail() {
	expectedQuery := "SELECT  p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id WHERE p.id=$1"
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs("12").WillReturnError(errors.New("error"))
	gotProduct, gotErr := suite.repo.FindById("12")
	assert.Error(suite.T(), gotErr)
	assert.Equal(suite.T(), model.Product{}, gotProduct)
}
