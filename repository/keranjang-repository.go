package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type KeranjangRepository interface {
	SaveKeranjang(k entity.Keranjang) entity.Keranjang
	FindByUserID(UserID int) entity.Keranjang
	FindByBukuID(IDKeranjang int) entity.Keranjang
}

type keranjangConnection struct {
	connection *gorm.DB
}

func NewKeranjangRepository(dbConn *gorm.DB) KeranjangRepository {
	return &keranjangConnection{
		connection: dbConn,
	}
}

func (db *keranjangConnection) SaveKeranjang(k entity.Keranjang) entity.Keranjang {
	db.connection.Save(&k)
	db.connection.Preload("Buku").Preload("User").Find(&k)
	return k
}

func (db *keranjangConnection) FindByUserID(UserID int) entity.Keranjang {
	var keranjang entity.Keranjang
	db.connection.Where("user_id = ?", UserID).Preload("Buku").Preload("User").Find(&keranjang)
	return keranjang
}

func (db *keranjangConnection) FindByBukuID(BookID int) entity.Keranjang {
	var keranjang entity.Keranjang
	db.connection.Where("book_id = ?", BookID).Find(&keranjang)
	return keranjang
}
