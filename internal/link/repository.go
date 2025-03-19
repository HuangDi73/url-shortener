package link

import "url-shortener/pkg/db"

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

func (repo *Repository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Db.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
