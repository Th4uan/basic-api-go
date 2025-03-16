package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Th4uan/basic-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 10.0)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotNil(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Produto %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	produtos, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, produtos, 10)
	assert.Equal(t, "Produto 1", produtos[0].Name)
	assert.Equal(t, "Produto 6", produtos[5].Name)

	produtos, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, produtos, 10)
	assert.Equal(t, "Produto 11", produtos[0].Name)
	assert.Equal(t, "Produto 16", produtos[5].Name)
}
