package link

import (
	"url-shortener/pkg/db"

	"gorm.io/gorm/clause"
)

type IRepository interface {
	Create(*Link) (*Link, error)
	FindByHash(string) (*Link, error)
	FindById(uint) (*Link, error)
	Update(*Link) (*Link, error)
	Delete(uint) error
}

type Repository struct {
	Db *db.Db
}

func NewRepository(db *db.Db) IRepository {
	return &Repository{
		Db: db,
	}
}

func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.Db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) FindByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Db.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *Repository) FindById(id uint) (*Link, error) {
	var link Link
	result := repo.Db.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *Repository) Update(link *Link) (*Link, error) {
	result := repo.Db.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) Delete(id uint) error {
	result := repo.Db.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
