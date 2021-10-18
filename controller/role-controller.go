package controller

import (
	"net/http"
	"strconv"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/helper"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	GetAll(context *gin.Context)
	GetByID(context *gin.Context)
	CreateRole(context *gin.Context)
	UpdateRole(context *gin.Context)
	DeleteRole(context *gin.Context)
}

type roleController struct {
	roleService service.RoleService
	jwtService  service.JWTService
}

func NewRoleController(roleService service.RoleService, jwtService service.JWTService) RoleController {
	return &roleController{
		roleService: roleService,
		jwtService:  jwtService,
	}
}

func (c *roleController) GetAll(context *gin.Context) {
	var role []entity.Role = c.roleService.GetAll()
	res := helper.BuildResponse(true, "Sukses", role)
	context.JSON(http.StatusOK, res)
}

func (c *roleController) GetByID(context *gin.Context) {
	id_role, err := strconv.Atoi(context.Param("id_role"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_role yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var role entity.Role = c.roleService.GetByID(id_role)
	if (role == entity.Role{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_role", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", role)
		context.JSON(http.StatusOK, res)
	}
}

func (c *roleController) CreateRole(context *gin.Context) {
	var createRoleDTO dto.CreateRoleDTO
	errDTO := context.ShouldBindJSON(&createRoleDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.roleService.CreateRole(createRoleDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *roleController) UpdateRole(context *gin.Context) {
	id_role, err := strconv.Atoi(context.Param("id_role"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_role yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var updateRoleDTO dto.UpdateRoleDTO
	errDTO := context.ShouldBindJSON(&updateRoleDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.roleService.UpdateRole(id_role, updateRoleDTO)
	response := helper.BuildResponse(true, "Berhasil memperbarui role", result)
	context.JSON(http.StatusOK, response)
}

func (c *roleController) DeleteRole(context *gin.Context) {
	var role entity.Role
	id_role, err := strconv.Atoi(context.Param("id_role"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_role yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	role.IDRole = int(id_role)

	c.roleService.DeleteRole(role)
	response := helper.BuildResponse(true, "Data berhasil dihapus", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}
