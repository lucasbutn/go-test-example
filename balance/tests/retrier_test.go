package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"test-example/balance"
	"testing"
)


func TestSuiteRetrier(t *testing.T) {

	t.Run("Execute Ok no retries", testExecuteOkNoRetries)
	t.Run("Execute Ok two retries", testExecuteOkTwoRetries)
	t.Run("Execute error all retries", testExecuteErrorAllRetries)

}

func testExecuteOkNoRetries(t *testing.T) {
	retrier := balance.NewDelayedRetrier(3, 0)

	var execTimes int
	err := retrier.Run(func() error {
		execTimes++
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, execTimes)
}

func testExecuteOkTwoRetries(t *testing.T) {
	retrier := balance.NewDelayedRetrier(3, 0)

	var execTimes int
	err := retrier.Run(func() error {
		execTimes++
		if execTimes <= 2 {
			return errors.New("some error happened")
		}
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, 3, execTimes)
}

func testExecuteErrorAllRetries(t *testing.T) {
	retrier := balance.NewDelayedRetrier(3, 0)

	var execTimes int
	err := retrier.Run(func() error {
		execTimes++
		return errors.New("some error happened")
	})

	assert.Error(t, err)
	assert.Equal(t, 4, execTimes)
}


