package user_input_adapter

import (
	"errors"
	user_use_case "example/api/use-cases/user-use-case"
)

type UserInputAdapter struct {
	userUseCase user_use_case.UserUseCase
}
type UserInputAdapterOptions struct {
	UserUseCase user_use_case.UserUseCase
}

func NewUserInputAdapter(options UserInputAdapterOptions) (*UserInputAdapter, error) {

	if options.UserUseCase == nil {
		return nil, errors.New("")
	}
	return &UserInputAdapter{userUseCase: options.UserUseCase}, nil
}

func (a UserInputAdapter) CreateUser(user) error {
	registeredUser, err := a.userUseCase.RegisterUser()

	if err != nil {
		return err
	}
	return nil
}
