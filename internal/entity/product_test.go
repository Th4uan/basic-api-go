package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Pc Gamer", 2000.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, p.Name, "Pc Gamer")
	assert.Equal(t, p.Price, 2000.00)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 2000.00)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Pc Gamer", 0.0)
	assert.Nil(t, p)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Pc Gamer", 2000.00)
	assert.Nil(t, err)
	assert.Nil(t, p.Validate())
	assert.NotNil(t, p)
}
