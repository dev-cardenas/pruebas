package domain_user

type Reader interface {
	GetByID(id int64) (*User, error)
	Search(name, lastname string, offset, limit int) ([]*User, error)
}

type Creator interface {
	Create(user *User) error
}

type Updater interface {
	Update(user *User) error
}

type Deleter interface {
	Delete(id int64) error
}

type EmailFinder interface {
	FindByEmail(email string) (*User, error)
}

type Repository interface {
	Reader
	Creator
	Updater
	Deleter
	EmailFinder
}
