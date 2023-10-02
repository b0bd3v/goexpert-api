package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go_api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)

	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.NoError(t, err)

	assert.NotEmpty(t, product.ID)

}

func TestProduct_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)

	assert.NoError(t, err)

	db.Create(product)

	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())

	assert.NoError(t, err)

	var productFound entity.Product

	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Error(t, err)
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestProduct_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)

	assert.NoError(t, err)

	db.Create(product)

	productDB := NewProduct(db)

	productFound, err := productDB.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}

func TestProduct_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	_ = db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.00)

	assert.NoError(t, err)

	db.Create(product)

	productDB := NewProduct(db)

	productFound, err := productDB.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	productFound.Name = "Product 2"
	productFound.Price = 20.00

	err = productDB.Update(productFound)

	assert.NoError(t, err)

	productFound, err = productDB.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 2", productFound.Name)
	assert.Equal(t, 20.00, productFound.Price)
}
