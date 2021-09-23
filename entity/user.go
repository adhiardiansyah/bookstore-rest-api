package entity

import "time"

type User struct {
	ID         int    `gorm:"primary_key:auto_increment" json:"id"`
	RoleID     int    `gorm:"not null;default:2" json:"-"`
	Nama       string `gorm:"type:varchar(255)" json:"nama"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	GambarUser string `gorm:"type:varchar(255)" json:"gambar_user"`
	Telp       string `gorm:"type:varchar(20)" json:"telp"`
	Alamat     string `gorm:"type:text" json:"alamat"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Token      string `gorm:"-" json:"token,omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Role       Role `gorm:"foreignkey:RoleID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role"`
}
