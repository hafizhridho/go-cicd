package helper

import (
	"encoding/json"
	"latihan/configs"
	"latihan/controllers"
	"latihan/models/base"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)
func TestGetAll(t *testing.T) {
	e := echo.New()

	// Inisialisasi mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Inisialisasi objek *gorm.DB dengan driver PostgreSQL
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error initializing GORM DB: %v", err)
	}

	// Gantilah objek `controllers.DB` dengan objek *gorm.DB
	configs.DB = gormDB

	// Mock query database yang diharapkan oleh fungsi GetAll
	rows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(1, "Book 1", "Author 1").
		AddRow(2, "Book 2", "Author 2")

		mock.ExpectQuery("SELECT \\* FROM \"books\" WHERE \"books\".\"deleted_at\" IS NULL").
		WillReturnRows(rows)
	

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = controllers.GetAll(c)

	// Verifikasi hasilnya
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response base.BaseResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.True(t, response.Status)
	assert.Equal(t, "berhasil", response.Message)

	// Periksa bahwa Data tidak kosong
	assert.NotEmpty(t, response.Data)

	// Pastikan semua ekspektasi query telah dipenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}