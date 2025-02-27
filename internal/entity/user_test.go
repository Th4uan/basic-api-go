package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Thauan", "tht@g.com", "1235")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "Thauan")
	assert.Equal(t, user.Email, "tht@g.com")
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Thauan", "tht@g.com", "1235")
	assert.Nil(t, err)
	assert.NotEqual(t, user.Password, "1235")
	assert.True(t, user.ValidatePassword("1235"))
	assert.False(t, user.ValidatePassword("1234"))
}
