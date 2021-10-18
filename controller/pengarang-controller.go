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

type PengarangController interface {
	GetAll(context *gin.Context)
	GetByID(context *gin.Context)
	CreatePengarang(context *gin.Context)
	UpdatePengarang(context *gin.Context)
	DeletePengarang(context *gin.Context)
}

type pengarangController struct {
	pengarangService service.PengarangService
	jwtService       service.JWTService
}

func NewPengarangController(pengarangService service.PengarangService, jwtService service.JWTService) PengarangController {
	return &pengarangController{
		pengarangService: pengarangService,
		jwtService:       jwtService,
	}
}

func (c *pengarangController) GetAll(context *gin.Context) {
	var pengarang []entity.Pengarang = c.pengarangService.GetAll()
	res := helper.BuildResponse(true, "Sukses", pengarang)
	context.JSON(http.StatusOK, res)
}

func (c *pengarangController) GetByID(context *gin.Context) {
	id_pengarang, err := strconv.Atoi(context.Param("id_pengarang"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_pengarang yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var pengarang entity.Pengarang = c.pengarangService.GetByID(id_pengarang)
	if (pengarang == entity.Pengarang{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_pengarang", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", pengarang)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pengarangController) CreatePengarang(context *gin.Context) {
	var createPengarangDTO dto.CreatePengarangDTO
	errDTO := context.ShouldBindJSON(&createPengarangDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.pengarangService.CreatePengarang(createPengarangDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *pengarangController) UpdatePengarang(context *gin.Context) {
	id_pengarang, err := strconv.Atoi(context.Param("id_pengarang"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_pengarang yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var updatePengarangDTO dto.UpdatePengarangDTO
	errDTO := context.ShouldBindJSON(&updatePengarangDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.pengarangService.UpdatePengarang(id_pengarang, updatePengarangDTO)
	response := helper.BuildResponse(true, "Berhasil memperbarui pengarang", result)
	context.JSON(http.StatusOK, response)
}

func (c *pengarangController) DeletePengarang(context *gin.Context) {
	var pengarang entity.Pengarang
	id_pengarang, err := strconv.Atoi(context.Param("id_pengarang"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_pengarang yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	pengarang.IDPengarang = int(id_pengarang)

	c.pengarangService.DeletePengarang(pengarang)
	response := helper.BuildResponse(true, "Data berhasil dihapus", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}
