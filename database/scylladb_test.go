package database_test

import (
	"testing"

	"github.com/SergioVenicio/grpc_gtw/config"
	"github.com/SergioVenicio/grpc_gtw/database"
	"github.com/SergioVenicio/grpc_gtw/models"
	"github.com/SergioVenicio/grpc_gtw/repositories"
	"github.com/brianvoe/gofakeit"
	"github.com/scylladb/gocqlx/v3/gocqlxtest"
	"github.com/stretchr/testify/assert"
)

func createUser() models.User {
	id := gofakeit.Int32()
	name := gofakeit.Name()
	email := gofakeit.Email()
	return models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func TestScyllaDB(t *testing.T) {
	cfg := config.NewConfig()
	cluster := gocqlxtest.CreateCluster()
	db := database.NewScyllaDB[models.User](cluster, cfg)

	assert := assert.New(t)

	t.Run("CreateUser", func(t *testing.T) {
		t.Parallel()
		err := db.Save(repositories.UserMetadata, createUser())
		assert.Nil(err)
	})

	t.Run("GetUser", func(t *testing.T) {
		t.Parallel()
		user := createUser()
		err := db.Save(repositories.UserMetadata, user)
		assert.Nil(err)

		dbUser, err := db.Get(repositories.UserMetadata, models.User{ID: user.ID})
		assert.Nil(err)

		assert.Equal(dbUser, &user)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		t.Parallel()
		user := createUser()
		err := db.Save(repositories.UserMetadata, user)
		assert.Nil(err)

		dbUser, err := db.Get(repositories.UserMetadata, models.User{ID: user.ID})
		assert.Nil(err)

		assert.Equal(dbUser, &user)

		newEmail := gofakeit.Email()
		err = db.Update(
			repositories.UserMetadata,
			models.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: newEmail,
			},
		)
		assert.Nil(err)

		dbUser, err = db.Get(repositories.UserMetadata, models.User{ID: user.ID})
		assert.Nil(err)
		assert.Equal(dbUser.Email, newEmail)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		t.Parallel()

		user := createUser()
		err := db.Save(repositories.UserMetadata, user)
		assert.Nil(err)

		dbUser, err := db.Get(repositories.UserMetadata, models.User{ID: user.ID})
		assert.Nil(err)

		assert.Equal(dbUser, &user)

		err = db.Delete(repositories.UserMetadata, user.ID)
		assert.Nil(err)

		dbUser, err = db.Get(repositories.UserMetadata, models.User{ID: user.ID})
		assert.NotNil(err)
		assert.Nil(dbUser)
	})
}
