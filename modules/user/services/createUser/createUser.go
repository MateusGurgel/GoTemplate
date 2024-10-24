package createUser

import (
	"goTemplate/modules/user"
	"goTemplate/modules/user/infra"
)

func CreatUser(createUserDTO *DTO) (*user.User, error) {
	createdUser, err := infra.Create(user.User{Password: createUserDTO.Password, Email: createUserDTO.Email})
	return createdUser, err
}
