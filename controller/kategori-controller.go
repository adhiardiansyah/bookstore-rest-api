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

type KategoriController interface {
	GetAll(context *gin.Context)
	GetByID(context *gin.Context)
	CreateKategori(context *gin.Context)
	UpdateKategori(context *gin.Context)
	DeleteKategori(context *gin.Context)
}

type kategoriController struct {
	kategoriService service.KategoriService
	jwtService      service.JWTService
}

func NewKategoriController(kategoriService service.KategoriService, jwtService service.JWTService) KategoriController {
	return &kategoriController{
		kategoriService: kategoriService,
		jwtService:      jwtService,
	}
}

func (c *kategoriController) GetAll(context *gin.Context) {
	var kategori []entity.Kategori = c.kategoriService.GetAll()
	res := helper.BuildResponse(true, "Sukses", kategori)
	context.JSON(http.StatusOK, res)
}

func (c *kategoriController) GetByID(context *gin.Context) {
	id_kategori, err := strconv.Atoi(context.Param("id_kategori"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_kategori yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var kategori entity.Kategori = c.kategoriService.GetByID(id_kategori)
	if (kategori == entity.Kategori{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_kategori", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", kategori)
		context.JSON(http.StatusOK, res)
	}
}

func (c *kategoriController) CreateKategori(context *gin.Context) {
	var createKategoriDTO dto.CreateKategoriDTO
	errDTO := context.ShouldBindJSON(&createKategoriDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.kategoriService.CreateKategori(createKategoriDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *kategoriController) UpdateKategori(context *gin.Context) {
	id_kategori, err := strconv.Atoi(context.Param("id_kategori"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_kategori yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var updateKategoriDTO dto.UpdateKategoriDTO
	errDTO := context.ShouldBindJSON(&updateKategoriDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.kategoriService.UpdateKategori(id_kategori, updateKategoriDTO)
	response := helper.BuildResponse(true, "Berhasil memperbarui kategori", result)
	context.JSON(http.StatusOK, response)
}

func (c *kategoriController) DeleteKategori(context *gin.Context) {
	var kategori entity.Kategori
	id_kategori, err := strconv.Atoi(context.Param("id_kategori"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_kategori yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	kategori.IDKategori = int(id_kategori)

	c.kategoriService.DeleteKategori(kategori)
	response := helper.BuildResponse(true, "Data berhasil dihapus", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}
