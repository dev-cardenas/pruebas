package domain_user

import "errors"

type Service interface {
	ValidateUser(u *User) error
	CheckEmailUniqueness(email string) error
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) ValidateUser(u *User) error {
	if u.Age < 18 {
		return errors.New("user must be at least 18 years old")
	}
	return nil
}

func (serviceRepository *serviceImpl) CheckEmailUniqueness(email string) error {
	existingUser, _ := serviceRepository.repo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("email already in use")
	}
	return nil
}
