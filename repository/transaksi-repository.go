package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	SaveTransaksi(t entity.Transaksi) entity.Transaksi
	SaveDetailTransaksi(dt entity.DetailTransaksi) entity.DetailTransaksi
	UpdateTransaksi(t entity.Transaksi) entity.Transaksi
	DeleteTransaksi(t entity.Transaksi)
	FindAll() []entity.Transaksi
	FindByID(IDTransaksi int) entity.Transaksi
}

type transaksiConnection struct {
	connection *gorm.DB
}

func NewTransaksiRepository(dbConn *gorm.DB) TransaksiRepository {
	return &transaksiConnection{
		connection: dbConn,
	}
}

func (db *transaksiConnection) SaveTransaksi(t entity.Transaksi) entity.Transaksi {
	db.connection.Create(&t)
	// db.connection.Preload("DetailTransaksi").Find(&t)
	return t
}

func (db *transaksiConnection) SaveDetailTransaksi(dt entity.DetailTransaksi) entity.DetailTransaksi {
	db.connection.Create(&dt)
	return dt
}

func (db *transaksiConnection) UpdateTransaksi(t entity.Transaksi) entity.Transaksi {
	db.connection.Save(&t)
	db.connection.Preload("DetailTransaksi").Find(&t)
	return t
}

func (db *transaksiConnection) DeleteTransaksi(t entity.Transaksi) {
	db.connection.Delete(&t)
}

func (db *transaksiConnection) FindByID(IDTransaksi int) entity.Transaksi {
	var transaksi entity.Transaksi
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Where("id_Transaksi = ?", IDTransaksi).Find(&transaksi)
	return transaksi
}

func (db *transaksiConnection) FindAll() []entity.Transaksi {
	var transaksis []entity.Transaksi
	db.connection.Preload("Kategori").Preload("Pengarang").Preload("Penerbit").Find(&transaksis)
	return transaksis
}
