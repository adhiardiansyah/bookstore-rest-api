package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type PenerbitRepository interface {
	SavePenerbit(p entity.Penerbit) entity.Penerbit
	UpdatePenerbit(p entity.Penerbit) entity.Penerbit
	DeletePenerbit(p entity.Penerbit)
	FindAll() []entity.Penerbit
	FindByID(IDPenerbit int) entity.Penerbit
}

type penerbitConnection struct {
	connection *gorm.DB
}

func NewPenerbitRepository(dbConn *gorm.DB) PenerbitRepository {
	return &penerbitConnection{
		connection: dbConn,
	}
}

func (db *penerbitConnection) SavePenerbit(p entity.Penerbit) entity.Penerbit {
	db.connection.Save(&p)
	db.connection.Find(&p)
	return p
}

func (db *penerbitConnection) UpdatePenerbit(p entity.Penerbit) entity.Penerbit {
	db.connection.Save(&p)
	db.connection.Find(&p)
	return p
}

func (db *penerbitConnection) DeletePenerbit(p entity.Penerbit) {
	db.connection.Delete(&p)
}

func (db *penerbitConnection) FindByID(IDPenerbit int) entity.Penerbit {
	var penerbit entity.Penerbit
	db.connection.Where("id_penerbit = ?", IDPenerbit).Find(&penerbit)
	return penerbit
}

func (db *penerbitConnection) FindAll() []entity.Penerbit {
	var penerbit []entity.Penerbit
	db.connection.Find(&penerbit)
	return penerbit
}
