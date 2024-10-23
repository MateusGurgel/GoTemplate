package createUser

import (
	"goTemplate/modules/user"
	"goTemplate/modules/user/infra/userRepository"
)

func CreatUser(createUserDTO *DTO) (*user.User, error) {
	createdUser, err := userRepository.Create(user.User{Password: createUserDTO.Password, Email: createUserDTO.Email})

	return createdUser, err
}
