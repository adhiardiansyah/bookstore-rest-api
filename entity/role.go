package entity

import "time"

type Role struct {
	IDRole        int    `gorm:"primary_key:auto_increment" json:"id_role"`
	NamaRole      string `gorm:"type:varchar(255)" json:"nama_role"`
	DeskripsiRole string `gorm:"type:text" json:"deskripsi_role"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          *User `json:"user,omitempty"`
}
