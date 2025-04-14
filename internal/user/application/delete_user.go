package application_user

func (uc *useCaseImpl) DeleteUser(id int64) error {
	return uc.repo.Delete(id)
}
