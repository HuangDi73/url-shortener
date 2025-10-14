package link

import (
	"math/rand/v2"
	"time"

	"gorm.io/gorm"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Link struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Url       string         `json:"url"`
	Hash      string         `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (l *Link) GenerateHash() {
	b := make([]rune, 6)
	for i := range b {
		b[i] = symbols[rand.IntN(len(symbols))]
	}
	l.Hash = string(b)
}
