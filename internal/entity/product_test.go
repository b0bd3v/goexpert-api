package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 1000)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)

	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 1000.0, product.Price)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("", 1000)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	product, err := NewProduct("Product 1", 0)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	product, err := NewProduct("Product 1", -1)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductIsValid(t *testing.T) {
	product, err := NewProduct("Product 1", 1000)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
