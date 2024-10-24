package infra

import (
	"goTemplate/infra/database"
	"goTemplate/modules/user"
)

func Create(user user.User) (*user.User, error) {
	db := database.Database.GetDatabase()

	result := db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
