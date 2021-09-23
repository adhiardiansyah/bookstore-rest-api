package config

import (
	"fmt"
	"os"

	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupatDatabaseConnection adalah menghubungkan koneksi ke database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Gagal terhubung ke file env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal menghubungkan ke database")
	}

	db.AutoMigrate(&entity.Buku{}, &entity.DetailTransaksi{}, &entity.Kategori{}, &entity.Keranjang{}, &entity.Penerbit{}, &entity.Pengarang{}, &entity.Role{}, &entity.Transaksi{}, &entity.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal untuk keluar koneksi database")
	}
	dbSQL.Close()
}
