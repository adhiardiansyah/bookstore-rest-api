package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type PengarangRepository interface {
	SavePengarang(p entity.Pengarang) entity.Pengarang
	UpdatePengarang(p entity.Pengarang) entity.Pengarang
	DeletePengarang(p entity.Pengarang)
	FindAll() []entity.Pengarang
	FindByID(IDPengarang int) entity.Pengarang
}

type pengarangConnection struct {
	connection *gorm.DB
}

func NewPengarangRepository(dbConn *gorm.DB) PengarangRepository {
	return &pengarangConnection{
		connection: dbConn,
	}
}

func (db *pengarangConnection) SavePengarang(p entity.Pengarang) entity.Pengarang {
	db.connection.Save(&p)
	db.connection.Find(&p)
	return p
}

func (db *pengarangConnection) UpdatePengarang(p entity.Pengarang) entity.Pengarang {
	db.connection.Save(&p)
	db.connection.Find(&p)
	return p
}

func (db *pengarangConnection) DeletePengarang(p entity.Pengarang) {
	db.connection.Delete(&p)
}

func (db *pengarangConnection) FindByID(IDPengarang int) entity.Pengarang {
	var pengarang entity.Pengarang
	db.connection.Where("id_pengarang = ?", IDPengarang).Find(&pengarang)
	return pengarang
}

func (db *pengarangConnection) FindAll() []entity.Pengarang {
	var pengarang []entity.Pengarang
	db.connection.Find(&pengarang)
	return pengarang
}
