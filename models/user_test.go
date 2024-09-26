package models_test

import (
	"testing"

	"github.com/SergioVenicio/grpc_gtw/models"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	t.Parallel()

	id := gofakeit.Int32()
	name := gofakeit.Name()
	email := gofakeit.Email()
	user := models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	assert := assert.New(t)
	assert.Equal(user.ID, id)
	assert.Equal(user.Name, name)
	assert.Equal(user.Email, email)
}
