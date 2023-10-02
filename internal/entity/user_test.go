package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Roberto", "r@r.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Roberto", user.Name)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Roberto", "r@r.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
