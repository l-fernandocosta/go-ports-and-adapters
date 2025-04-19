package application_test

import (
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductEnableAndDisableWithSuccess(t *testing.T) {
	// Test case 1: Enable product
	product := application.NewProduct("1", "Product A", 100.0)

	product.Enable()
	assert.Equal(t, product.GetStatus(), application.ENABLED)

	// Test case 2: Disable product should throw an error if the price is not 0
	err := product.Disable()

	assert.Equal(t, err.Error(), "product price must be 0 to disable")

	// Test case 3: Disable product
	product.Price = 0.0
	err = product.Disable()
	assert.Nil(t, err)
	assert.Equal(t, product.GetStatus(), application.DISABLED)
}

func TestProductEnableAndDisableWithError(t *testing.T) {
	// Test case 1: Enable product
	product := application.NewProductWithStatus("1", "Product A", 0.0, application.DISABLED)

	err := product.Enable()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "product price must be greater than 0 to enable")
}

func TestProductIsValidMethod(t *testing.T) {
	product := application.Product{}

	product.Price = 10.0
	product.Name = "Hello"
	product.ID = uuid.NewV4().String()
	product.Status = application.DISABLED

	valid, err := product.IsValid()
	assert.Nil(t, err)
	assert.True(t, valid)

	// invalid product
	product.Price = -10
	product.Name = ""
	product.ID = ""
	product.Status = "invalid"

	valid, err = product.IsValid()

	log.Print(err)
	assert.NotNil(t, err)
	assert.False(t, valid)
}
