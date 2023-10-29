package latihan

import "time"

type BookResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time        `json:"createdAT"`
	UpdatedAt time.Time `json:"updatedAT"`
	Buku      string    `json:"buku"`
	Author  string    `json:"author"`
	Kategori  string    `json:"kategori"`
}

func (bookresponse *BookResponse) MapFromDB (latihan Books)  {
	bookresponse.ID = latihan.ID
	bookresponse.CreatedAt = latihan.CreatedAt
	bookresponse.UpdatedAt = latihan.UpdatedAt
	bookresponse.Buku = latihan.Buku
	bookresponse.Author = latihan.Author
	bookresponse.Kategori = latihan.Kategori
	
}

func MapFromDBList(listNewDB []Books) []BookResponse {
	var result[] BookResponse
	var bookresponse BookResponse
	for _, books := range listNewDB {
		bookresponse.ID = books.ID
	bookresponse.CreatedAt = books.CreatedAt
	bookresponse.UpdatedAt = books.UpdatedAt
	bookresponse.Buku = books.Buku
	bookresponse.Author = books.Author
	bookresponse.Kategori = books.Kategori
	result = append(result, bookresponse)
	}
	return result
}