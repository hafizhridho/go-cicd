package latihan

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID        uint 			`gorm:"primarykey"`
	CreatedAt time.Time		`gorm:"autoCreateTime"`
	UpdatedAt time.Time		`gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Buku string
	Author string
	Kategori string
}

func (latihan Books) MapFromBookAdd(book BookAdd) Books {
	var newDB Books
	newDB.Author = book.Author
	newDB.Kategori = book.Kategori
	newDB.Buku = book.Buku
	return newDB
}