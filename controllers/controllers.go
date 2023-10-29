package controllers

import (
	"latihan/configs"
	"latihan/models/base"
	"latihan/models/latihan"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error  {
	var data []latihan.Books
	result := configs.DB.Find(&data)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	
	var bookresponse []latihan.BookResponse
	bookresponse = latihan.MapFromDBList(data)

	
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status: true,
		Message: "berhasil",
		Data: bookresponse,
	})
	
}


func CreateBook(c echo.Context) error {
	var book latihan.BookAdd
   c.Bind(&book)

   if book.Buku == "" {
	return c.JSON(http.StatusBadRequest, base.BaseResponse{
		Status: false,
		Message: "buku kosong",
		Data: nil,
	})
   }
   if book.Author  == "" {
	return c.JSON(http.StatusBadRequest, base.BaseResponse{
		Status: false,
		Message: "author kosong",
		Data: nil,
	})
   }
   if book.Kategori == "" {
	return c.JSON(http.StatusBadRequest, base.BaseResponse {
		Status: false,
		Message: "kategori kosong",
		Data: nil,
	})
   }
   var newDB latihan.Books
   newDB = newDB.MapFromBookAdd(book)
   result := configs.DB.Create(&newDB)
   
   if result.Error != nil {
	return c.JSON(http.StatusInternalServerError, base.BaseResponse{
		Status: false,
		Message: "error",
		Data: nil,
	})
   }
   var response latihan.BookResponse
   response.MapFromDB(newDB)
   return c.JSON(http.StatusOK, base.BaseResponse{
	Status: false,
	Message: "succes",
	Data: newDB,
})
}

func DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status: false,
			Message: "id salah",
			Data: nil,
		})
	}
	result := configs.DB.Delete(&latihan.Books{}, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusNoContent, base.BaseResponse{
		Status: true,
		Message: "berhasil menghapus",
		Data: nil,
	})
}

func GetByID(c echo.Context) error {
    // Dapatkan ID dari URL.
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "ID tidak valid",
            Data: nil,
        })
    }

    // Cari buku berdasarkan ID di database.
    var book latihan.Books
    result := configs.DB.First(&book, id)
    if result.Error != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Buku tidak ditemukan",
            Data: nil,
        })
    }

    var bookResponse latihan.BookResponse
    bookResponse.MapFromDB(book)

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Berhasil",
        Data: bookResponse,
    })
}

func UpdateBook(c echo.Context) error {
    // Dapatkan ID dari URL.
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "ID tidak valid",
            Data: nil,
        })
    }

    // Temukan buku berdasarkan ID di database.
    var book latihan.Books
    result := configs.DB.First(&book, id)
    if result.Error != nil {
        return c.JSON(http.StatusNotFound, base.BaseResponse{
            Status: false,
            Message: "Buku tidak ditemukan",
            Data: nil,
        })
    }

    // Bind data yang dikirim dengan permintaan PUT ke struct buku yang ada.
    if err := c.Bind(&book); err != nil {
        return c.JSON(http.StatusBadRequest, base.BaseResponse{
            Status: false,
            Message: "Permintaan tidak valid",
            Data: nil,
        })
    }

    // Perbarui buku dalam database.
    result = configs.DB.Save(&book)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, base.BaseResponse{
            Status: false,
            Message: "Gagal memperbarui buku",
            Data: nil,
        })
    }

    var bookResponse latihan.BookResponse
    bookResponse.MapFromDB(book)

    return c.JSON(http.StatusOK, base.BaseResponse{
        Status: true,
        Message: "Buku berhasil diperbarui",
        Data: bookResponse,
    })
}
