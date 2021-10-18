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

type PenerbitController interface {
	GetAll(context *gin.Context)
	GetByID(context *gin.Context)
	CreatePenerbit(context *gin.Context)
	UpdatePenerbit(context *gin.Context)
	DeletePenerbit(context *gin.Context)
}

type penerbitController struct {
	penerbitService service.PenerbitService
	jwtService      service.JWTService
}

func NewPenerbitController(penerbitService service.PenerbitService, jwtService service.JWTService) PenerbitController {
	return &penerbitController{
		penerbitService: penerbitService,
		jwtService:      jwtService,
	}
}

func (c *penerbitController) GetAll(context *gin.Context) {
	var penerbit []entity.Penerbit = c.penerbitService.GetAll()
	res := helper.BuildResponse(true, "Sukses", penerbit)
	context.JSON(http.StatusOK, res)
}

func (c *penerbitController) GetByID(context *gin.Context) {
	id_penerbit, err := strconv.Atoi(context.Param("id_penerbit"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_penerbit yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var penerbit entity.Penerbit = c.penerbitService.GetByID(id_penerbit)
	if (penerbit == entity.Penerbit{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_penerbit", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", penerbit)
		context.JSON(http.StatusOK, res)
	}
}

func (c *penerbitController) CreatePenerbit(context *gin.Context) {
	var createPenerbitDTO dto.CreatePenerbitDTO
	errDTO := context.ShouldBindJSON(&createPenerbitDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.penerbitService.CreatePenerbit(createPenerbitDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *penerbitController) UpdatePenerbit(context *gin.Context) {
	id_penerbit, err := strconv.Atoi(context.Param("id_penerbit"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_penerbit yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var updatePenerbitDTO dto.UpdatePenerbitDTO
	errDTO := context.ShouldBindJSON(&updatePenerbitDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.penerbitService.UpdatePenerbit(id_penerbit, updatePenerbitDTO)
	response := helper.BuildResponse(true, "Berhasil memperbarui penerbit", result)
	context.JSON(http.StatusOK, response)
}

func (c *penerbitController) DeletePenerbit(context *gin.Context) {
	var penerbit entity.Penerbit
	id_penerbit, err := strconv.Atoi(context.Param("id_penerbit"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_penerbit yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	penerbit.IDPenerbit = int(id_penerbit)

	c.penerbitService.DeletePenerbit(penerbit)
	response := helper.BuildResponse(true, "Data berhasil dihapus", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}
