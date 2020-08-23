package balance

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBalanceOk(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)
	mockService.EXPECT().GetBalance(gomock.Eq("12345")).Times(1).Return(&Balance{"12345", 100.0}, nil)

	controller,err := NewController(mockService)
	require.Empty(t, err)

	// Act
	writer := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/balances/12345", nil)
	require.Nil(t, err)

	controller.GetBalance(writer, request)

	// Assert
	assert.Equal(t, http.StatusOK, writer.Result().StatusCode)

	bytes := writer.Body.Bytes()
	balance := Balance{}
	err = json.Unmarshal(bytes, &balance)

	assert.Nil(t, err)
	assert.Equal(t, 100.0, balance.Total)
	assert.Equal(t, "12345", balance.UserId)
}

func TestGetBalanceBadRequest(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)

	controller,err := NewController(mockService)
	require.Empty(t, err)

	// Act
	writer := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/balances/", nil)
	require.Nil(t, err)

	controller.GetBalance(writer, request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, writer.Result().StatusCode)
}


func TestGetBalanceNotFound(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockService(mockCtrl)
	mockService.EXPECT().GetBalance(gomock.Any()).Times(1).Return(nil, errors.New("an unexpected internal error"))

	controller,err := NewController(mockService)
	require.Empty(t, err)

	// Act
	writer := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/balances/232342", nil)
	require.Nil(t, err)

	controller.GetBalance(writer, request)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, writer.Result().StatusCode)
}
