package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type KategoriRepository interface {
	SaveKategori(k entity.Kategori) entity.Kategori
	UpdateKategori(k entity.Kategori) entity.Kategori
	DeleteKategori(k entity.Kategori)
	FindAll() []entity.Kategori
	FindByID(IDKategori int) entity.Kategori
}

type kategoriConnection struct {
	connection *gorm.DB
}

func NewKategoriRepository(dbConn *gorm.DB) KategoriRepository {
	return &kategoriConnection{
		connection: dbConn,
	}
}

func (db *kategoriConnection) SaveKategori(k entity.Kategori) entity.Kategori {
	db.connection.Save(&k)
	db.connection.Find(&k)
	return k
}

func (db *kategoriConnection) UpdateKategori(k entity.Kategori) entity.Kategori {
	db.connection.Save(&k)
	db.connection.Find(&k)
	return k
}

func (db *kategoriConnection) DeleteKategori(k entity.Kategori) {
	db.connection.Delete(&k)
}

func (db *kategoriConnection) FindByID(IDKategori int) entity.Kategori {
	var kategori entity.Kategori
	db.connection.Where("id_kategori = ?", IDKategori).Find(&kategori)
	return kategori
}

func (db *kategoriConnection) FindAll() []entity.Kategori {
	var kategori []entity.Kategori
	db.connection.Find(&kategori)
	return kategori
}
