package repository

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	SaveRole(r entity.Role) entity.Role
	UpdateRole(r entity.Role) entity.Role
	DeleteRole(r entity.Role)
	FindAll() []entity.Role
	FindByID(IDRole int) entity.Role
}

type roleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(dbConn *gorm.DB) RoleRepository {
	return &roleConnection{
		connection: dbConn,
	}
}

func (db *roleConnection) SaveRole(r entity.Role) entity.Role {
	db.connection.Save(&r)
	db.connection.Find(&r)
	return r
}

func (db *roleConnection) UpdateRole(r entity.Role) entity.Role {
	db.connection.Save(&r)
	db.connection.Find(&r)
	return r
}

func (db *roleConnection) DeleteRole(r entity.Role) {
	db.connection.Delete(&r)
}

func (db *roleConnection) FindByID(IDRole int) entity.Role {
	var role entity.Role
	db.connection.Where("id_role = ?", IDRole).Find(&role)
	return role
}

func (db *roleConnection) FindAll() []entity.Role {
	var role []entity.Role
	db.connection.Find(&role)
	return role
}
