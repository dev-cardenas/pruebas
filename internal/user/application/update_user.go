package application_user

import "pruebas/internal/user/domain"

func (uc *useCaseImpl) UpdateUser(u *domain_user.User) error {
	return uc.repo.Update(u)
}
