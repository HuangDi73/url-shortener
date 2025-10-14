package user

import "url-shortener/pkg/db"

type Repository struct {
	Db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{
		Db: db,
	}
}

func (repo *Repository) Create(user *User) (*User, error) {
	result := repo.Db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
