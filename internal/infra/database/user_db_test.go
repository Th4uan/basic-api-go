package database

import (
	"testing"

	"github.com/Th4uan/basic-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Thauan", "tht@g.com", "1235")
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Thauan", "tht@g.com", "1235")
	userDb := NewUser(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.ID, userFound.ID)
	assert.NotNil(t, userFound.Password)
}
