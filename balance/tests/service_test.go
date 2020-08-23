package tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"test-example/balance"
	"test-example/balance/mocks"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BalanceTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	mockRetrier *mocks.MockRetrier
	mockClient  *mocks.MockClient
	service     balance.Service
}


// Make sure all common variables is set before each test
func (suite *BalanceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRetrier = mocks.NewMockRetrier(suite.mockCtrl)

	suite.mockClient = mocks.NewMockClient(suite.mockCtrl)

	var err error
	suite.service, err = balance.NewService(suite.mockClient, suite.mockRetrier)

	assert.Nil(suite.T(), err)
}

// Make sure all teardown are called befoe eache tests.
// mockCrl.Finish will trigger all validations
func (suite *BalanceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *BalanceTestSuite) TestGetBalance() {

	suite.mockClient.EXPECT().GetAllMovements(gomock.Any()).Times(1).Return([]*balance.Movement{
		{"1", 1597191115, "DEPOSIT", 100.0},
		{"1", 1597193100, "PURCHASE", -25.0},
		{"1", 1597196872, "PURCHASE", -15.0}}, nil)

	suite.mockRetrier.EXPECT().Run(gomock.Any()).Times(1).
		Do(func(arg func() error) {
			arg()
	}).Return(nil)

	balanc, err := suite.service.GetBalance("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &balance.Balance{"1", 60.0}, balanc)

}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *BalanceTestSuite) TestGetBalanceError() {

	retrierErr :=  errors.New("Any Error")

	suite.mockClient.EXPECT().GetAllMovements(gomock.Any()).Times(1).Return(nil, retrierErr)

	suite.mockRetrier.EXPECT().Run(gomock.Any()).Times(1).
		Do(func(arg func() error) {
			arg()
		}).Return(retrierErr)

	_, err := suite.service.GetBalance("1")

	assert.Error(suite.T(), retrierErr)
	assert.Error(suite.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestBalanceTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceTestSuite))
}
