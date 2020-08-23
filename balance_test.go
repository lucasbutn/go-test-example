package main

import (
	"encoding/json"
	. "github.com/bunniesandbeatings/goerkin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"test-example/balance"
	"test-example/balance/mocks"
	"testing"
)

/*Scenario: Retrieve balanace for a user with movements
	GIVEN a user with movements: 100$, -15$, -25$
	WHEN its balance is requested
	THEN the balance should be 60$
 */

var client *mocks.MockClient
var controller balance.Controller

func TestBalance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	client = mocks.NewMockClient(mockCtrl)

	retrier := balance.NewDelayedRetrier(2, 500)

	srv, err := balance.NewService(client, retrier)
	require.Nil(t, err)

	controller, err = balance.NewController(srv)
	require.Nil(t, err)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Balance Suite")
}

var _= Describe("A user with movements", func() {
		// add movements
	client.EXPECT().GetAllMovements(gomock.Eq("12345")).Times(1).Return([]*balance.Movement{
		{"12345", 1597191115, "DEPOSIT", 100.0},
		{"12345", 1597193100, "PURCHASE", -25.0},
		{"12345", 1597196872, "PURCHASE", -15.0}}, nil)

	Context("its balance is requested", func() {
		writer := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodGet, "/balances/12345", nil)
		Expect(err).Should(BeNil())

		controller.GetBalance(writer, request)

		It("the balance should be 60$", func() {

			Expect(writer.Result().StatusCode).Should(Equal(http.StatusOK))

			bytes := writer.Body.Bytes()
			balance := balance.Balance{}

			err := json.Unmarshal(bytes, &balance)

			Expect(err).Should(BeNil())

			Expect(balance.Total).Should(Equal(60.0))
		})
	})
})



// In Goerkin's way
var _= Describe("Feature: In to get the balance of the users, as an api consumer, I want to get the full balance of different users", func() {

	var (
		writer *httptest.ResponseRecorder
	)

	steps := Define(func(define Definitions) {
		define.Given(`^A user with movements like (-?\d+), (-?\d+) & (-?\d+)$`, func(mov1, mov2, mov3 string) {
			// add movements

			m1, _ := strconv.ParseFloat(mov1,64)
			m2, _ := strconv.ParseFloat(mov2,64)
			m3, _ := strconv.ParseFloat(mov3,64)

			client.EXPECT().GetAllMovements(gomock.Eq("12345")).Times(1).Return([]*balance.Movement{
				{"12345", 1597191115, "DEPOSIT", m1},
				{"12345", 1597193100, "PURCHASE", m2},
				{"12345", 1597196872, "PURCHASE", m3}}, nil)
		})

		define.When(`^Its balance is requested$`, func() {
			writer = httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, "/balances/12345", nil)
			Expect(err).Should(BeNil())

			controller.GetBalance(writer, request)

		})

		define.Then(`^The balance should be (\d+)$`, func(totalS string) {
			Expect(writer.Result().StatusCode).Should(Equal(http.StatusOK))

			bytes := writer.Body.Bytes()
			balance := balance.Balance{}

			err := json.Unmarshal(bytes, &balance)

			Expect(err).Should(BeNil())

			total, _ := strconv.ParseFloat(totalS,64)

			Expect(balance.Total).Should(Equal(total))
		})
	})

	Scenario("Retrieve balanace for a user with movements", func() {
		steps.Given("A user with movements like 100, -15 & -20")
		steps.When("Its balance is requested")
		steps.Then("The balance should be 65")
	})

})

