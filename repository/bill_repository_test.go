package repository

import (
	"clean-code/model"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    BillRepository
}

func (suite *BillRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewBillRepository(suite.mockDB)
}

func TestBillRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BillRepositoryTestSuite))
}

func (suite *BillRepositoryTestSuite) TestSave_Success() {
	mockBill := model.Bill{
		ID:         "1",
		BillDate:   time.Now(),
		EntryDate:  time.Now(),
		FinishDate: time.Now(),
		Employee: model.Employee{
			ID: "1",
		},
		Customer: model.Customer{
			ID: "1",
		},
	}
	mockBillDetail := model.BillDetail{
		ID:   "1",
		Bill: mockBill,
		Product: model.Product{
			ID:    "1",
			Name:  "product 1",
			Price: 10000,
			Uom: model.Uom{
				ID:   "1",
				Name: "pcs",
			},
		},
		ProductPrice: 10000,
		Qty:          2,
	}
	suite.mockSQL.ExpectBegin()
	suite.mockSQL.ExpectExec("INSERT INTO bill").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSQL.ExpectExec("INSERT INTO bill_detail").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSQL.ExpectCommit()
	gotErr := suite.repo.Save(mockBill, mockBillDetail)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
}
