package user_use_case

type RegisterUserInput struct {
}

type RegisterUserOutput struct {
}

type UserUseCase interface {
	RegisterUser(input RegisterUserInput) (*RegisterUserOutput, error)
}

type DefaultUserUseCase struct {
}

func (d DefaultUserUseCase) RegisterUser(input RegisterUserInput) (*RegisterUserOutput, error) {
	//TODO implement me
	return &RegisterUserOutput{}, nil
}

var _ UserUseCase = (*DefaultUserUseCase)(nil)
