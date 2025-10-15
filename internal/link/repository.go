package link

import (
	"url-shortener/pkg/db"

	"gorm.io/gorm/clause"
)

type Repository struct {
	Db *db.Db
}

func NewRepository(db *db.Db) *Repository {
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
	var link Link
	result := repo.Db.Delete(&link, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Repository) Count() int64 {
	var count int64
	repo.Db.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count
}

func (repo *Repository) GetAll(limit, offset int) *[]Link {
	var links []Link
	repo.Db.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	return &links
}
