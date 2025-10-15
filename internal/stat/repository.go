package stat

import (
	"time"
	"url-shortener/pkg/db"

	"gorm.io/datatypes"
)

type Repository struct {
	*db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{
		Db: db,
	}
}

func (repo *Repository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}
