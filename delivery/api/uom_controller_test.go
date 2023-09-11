package api

import (
	"clean-code/__mock__/usecasemock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UomControllerTestSuite struct {
	suite.Suite
	uucm   *usecasemock.UomUseCaseMock
	router *gin.Engine
}

func (suite *UomControllerTestSuite) SetupTest() {
	suite.uucm = new(usecasemock.UomUseCaseMock)
	suite.router = gin.Default()
}

func TestUomControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UomControllerTestSuite))
}

// func (suite *UomControllerTestSuite) TestCreateUom_Success() {
// 	mockData := model.Uom{
// 		ID:   "1",
// 		Name: "kg",
// 	}
// 	suite.uucm.On("CreateNew", mockData).Return(nil)
// 	mockRg := suite.router.Group("/api/v1")
// 	NewUomController(suite.uucm, mockRg).Route()

// 	recorder := httptest.NewRecorder()

// 	payloadMarshal, err := json.Marshal(mockData)
// 	assert.Nil(suite.T(), err)

// 	request, err := http.NewRequest(http.MethodPost, "/uoms", bytes.NewBuffer(payloadMarshal))
// 	assert.Nil(suite.T(), err)

// 	suite.router.ServeHTTP(recorder, request)
// 	// simulasikan pengambilan response hasil dari serve diatas
// 	response := recorder.Body.Bytes()

// 	var uomResponse model.Uom
// 	err = json.Unmarshal(response, &uomResponse)
// 	// assert.Nil(suite.T(), err)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
// 	// assert.Equal(suite.T(), mockData, uomResponse)
// }
