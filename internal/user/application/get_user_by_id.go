package application_user

import "pruebas/internal/user/domain"

func (uc *useCaseImpl) GetUserByID(id int64) (*domain_user.User, error) {
	return uc.repo.GetByID(id)
}
