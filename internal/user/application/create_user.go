package application_user

import "pruebas/internal/user/domain"

func (uc *useCaseImpl) CreateUser(dto CreateUserDTO) (*domain_user.User, error) {
	user := &domain_user.User{
		Name:     dto.Name,
		LastName: dto.LastName,
		Email:    dto.Email,
		Age:      dto.Age,
	}

	// Validaciones del dominio
	if err := uc.service.ValidateUser(user); err != nil {
		return nil, err
	}

	if err := uc.service.CheckEmailUniqueness(user.Email); err != nil {
		return nil, err
	}

	err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
