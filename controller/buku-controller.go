package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/helper"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	"github.com/gin-gonic/gin"
)

type BukuController interface {
	GetAll(context *gin.Context)
	GetByID(context *gin.Context)
	CreateBuku(context *gin.Context)
	UploadImageBuku(context *gin.Context)
	UpdateBuku(context *gin.Context)
	DeleteBuku(context *gin.Context)
}

type bukuController struct {
	bukuService service.BukuService
	jwtService  service.JWTService
}

func NewBukuController(bukuService service.BukuService, jwtService service.JWTService) BukuController {
	return &bukuController{
		bukuService: bukuService,
		jwtService:  jwtService,
	}
}

func (c *bukuController) GetAll(context *gin.Context) {
	var buku []entity.Buku = c.bukuService.GetAll()
	res := helper.BuildResponse(true, "Sukses", buku)
	context.JSON(http.StatusOK, res)
}

func (c *bukuController) GetByID(context *gin.Context) {
	id_buku, err := strconv.Atoi(context.Param("id_buku"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_buku yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var buku entity.Buku = c.bukuService.GetByID(id_buku)
	if (buku == entity.Buku{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_buku", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", buku)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bukuController) CreateBuku(context *gin.Context) {
	var createBukuDTO dto.CreateBukuDTO
	errDTO := context.ShouldBindJSON(&createBukuDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.bukuService.CreateBuku(createBukuDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *bukuController) UploadImageBuku(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		res := helper.BuildErrorResponse("Gagal untuk upload gambar", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id_buku, err := strconv.Atoi(context.Param("id_buku"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_buku yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	path := fmt.Sprintf("images/%d-%s", id_buku, file.Filename)

	err = context.SaveUploadedFile(file, path)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal untuk upload gambar", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	b := c.bukuService.SaveImageBuku(id_buku, path)
	res := helper.BuildResponse(true, "Gambar buku berhasil diupload", b)
	context.JSON(http.StatusOK, res)
}

func (c *bukuController) UpdateBuku(context *gin.Context) {
	id_buku, err := strconv.Atoi(context.Param("id_buku"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_buku yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var updateBukuDTO dto.UpdateBukuDTO
	errDTO := context.ShouldBindJSON(&updateBukuDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.bukuService.UpdateBuku(id_buku, updateBukuDTO)
	response := helper.BuildResponse(true, "Berhasil memperbarui buku", result)
	context.JSON(http.StatusOK, response)
}

func (c *bukuController) DeleteBuku(context *gin.Context) {
	var buku entity.Buku
	id_buku, err := strconv.Atoi(context.Param("id_buku"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_buku yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	buku.IDBuku = int(id_buku)

	c.bukuService.DeleteBuku(buku)
	response := helper.BuildResponse(true, "Data berhasil dihapus", helper.EmptyObj{})
	context.JSON(http.StatusOK, response)
}

// func (c *bukuController) getUserIDByToken(token string) string {
// 	aToken, err := c.jwtService.ValidateToken(token)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	claims := aToken.Claims.(jwt.MapClaims)
// 	return fmt.Sprintf("%v", claims["user_id"])
// }
