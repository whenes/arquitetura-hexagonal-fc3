package cli_test

import (
	"fmt"
	"testing"
	"github.com/codeedu/go-hexagonal/adapters/cli"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mock_application "github.com/codeedu/go-hexagonal/mocks/application"
	"github.com/golang/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)
	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	assert.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled",
	result, err = cli.Run(productServiceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled",
	result, err = cli.Run(productServiceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s\n Name: %s\n Price: %f\n Status: %s",
	productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
}