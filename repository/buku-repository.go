package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type BukuRepository interface {
	SaveBuku(b entity.Buku) entity.Buku
	UpdateBuku(b entity.Buku) entity.Buku
	DeleteBuku(b entity.Buku)
	FindAll() []entity.Buku
	FindByID(IDBuku int) entity.Buku
}

type bukuConnection struct {
	connection *gorm.DB
}

func NewBukuRepository(dbConn *gorm.DB) BukuRepository {
	return &bukuConnection{
		connection: dbConn,
	}
}

func (db *bukuConnection) SaveBuku(b entity.Buku) entity.Buku {
	db.connection.Save(&b)
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Find(&b)
	return b
}

func (db *bukuConnection) UpdateBuku(b entity.Buku) entity.Buku {
	db.connection.Save(&b)
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Find(&b)
	return b
}

func (db *bukuConnection) DeleteBuku(b entity.Buku) {
	db.connection.Delete(&b)
}

func (db *bukuConnection) FindByID(IDBuku int) entity.Buku {
	var buku entity.Buku
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Where("id_buku = ?", IDBuku).Find(&buku)
	return buku
}

func (db *bukuConnection) FindAll() []entity.Buku {
	var bukus []entity.Buku
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Find(&bukus)
	return bukus
}
