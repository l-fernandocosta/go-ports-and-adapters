package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	mock "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock.NewMockProductInterface(ctrl)
	persistence := mock.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).Times(1)

	service := application.ProductService{Persistence: persistence}
	res, err := service.Get("123")

	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, product, res)

}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock.NewMockProductInterface(ctrl)
	persistence := mock.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

	service := application.ProductService{Persistence: persistence}
	res, err := service.Create("Teste", 12.0)

	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, product, res)

}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	res, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, res)

	res, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, res)

}
