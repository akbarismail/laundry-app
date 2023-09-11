package repository

import (
	"clean-code/model"
	"clean-code/model/dto"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UomRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    UomRepository
}

func (suite *UomRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewUomRepository(suite.mockDB)
}

func TestUomRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UomRepositoryTestSuite))
}

func (suite *UomRepositoryTestSuite) TestPaging_Failed() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.Uom{
		{
			ID:   "1",
			Name: "pcs",
		},
	}
	expectedQuery := `SELECT id, name FROM uom LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnError(errors.New("error"))
	gotUoms, gotPaging, gotErr := suite.repo.Paging(dto.PageRequest{})
	assert.Error(suite.T(), gotErr)
	assert.Nil(suite.T(), gotUoms)
	assert.Equal(suite.T(), 0, gotPaging.TotalRows)

	// TODO Error Select Count
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, uom := range mockData {
		rows.AddRow(uom.ID, uom.Name)
	}
	expectedQuery = `SELECT id, name FROM uom LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(((mockPageRequest.Page - 1) * mockPageRequest.Size), mockPageRequest.Size).WillReturnRows(rows)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM uom`)).WillReturnError(errors.New("error"))
	gotUoms, gotPaging, gotErr = suite.repo.Paging(mockPageRequest)
	assert.Error(suite.T(), gotErr)
	assert.Nil(suite.T(), gotUoms)
	assert.Equal(suite.T(), 0, gotPaging.TotalRows)
}

func (suite *UomRepositoryTestSuite) TestPaging_Success() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.Uom{
		{
			ID:   "1",
			Name: "pcs",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, uom := range mockData {
		rows.AddRow(uom.ID, uom.Name)
	}
	expectedQuery := `SELECT id, name FROM uom LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(((mockPageRequest.Page - 1) * mockPageRequest.Size), mockPageRequest.Size).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM uom`)).WillReturnRows(rowCount)

	gotUoms, gotPaging, gotErr := suite.repo.Paging(mockPageRequest)
	assert.Nil(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotUoms)
	assert.Equal(suite.T(), 1, gotPaging.TotalRows)
}
