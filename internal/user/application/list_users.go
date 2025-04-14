package application_user

import "pruebas/internal/user/domain"

func (uc *useCaseImpl) ListUsers(page, limit int) ([]*domain_user.User, error) {
	offset := (page - 1) * limit
	return uc.repo.Search("", "", offset, limit)
}
