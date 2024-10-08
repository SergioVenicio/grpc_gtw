package repositories

import (
	"github.com/SergioVenicio/grpc_gtw/config"
	"github.com/SergioVenicio/grpc_gtw/database"
	"github.com/SergioVenicio/grpc_gtw/models"

	"github.com/scylladb/gocqlx/v3/table"
)

var UserMetadata = table.Metadata{
	Name:    "user",
	Columns: []string{"id", "name", "email"},
	PartKey: []string{"id"},
}

type Users struct {
	db *database.ScyllaDB[models.User]
}

func (u *Users) Save(newUser models.User) error {
	return u.db.Save(UserMetadata, newUser)
}

func (u *Users) Update(newUser models.User) error {
	return u.db.Update(UserMetadata, newUser)
}

func (u *Users) Get(id int32) (*models.User, error) {
	dbUser, err := u.db.Get(UserMetadata, models.User{ID: id})
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (u *Users) Delete(id int32) error {
	err := u.db.Delete(UserMetadata, id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(cfg *config.Config) *Users {
	return &Users{database.NewScyllaDB[models.User](nil, cfg)}
}
