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

type KeranjangController interface {
	AddToCart(context *gin.Context)
	GetCartByUserID(context *gin.Context)
}

type keranjangController struct {
	keranjangService service.KeranjangService
	bukuService      service.BukuService
	jwtService       service.JWTService
}

func NewKeranjangController(keranjangService service.KeranjangService, jwtService service.JWTService) KeranjangController {
	return &keranjangController{
		keranjangService: keranjangService,
		jwtService:       jwtService,
	}
}

func (c *keranjangController) AddToCart(context *gin.Context) {
	var addToCartDTO dto.AddToCartDTO
	buku_id, err := strconv.Atoi(context.Query("buku_id"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter id_buku yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var buku entity.Buku = c.bukuService.GetByID(buku_id)

	errDTO := context.ShouldBindJSON(&addToCartDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		currentUser := context.MustGet("currentUser").(entity.User)
		addToCartDTO.User = currentUser
		result := c.keranjangService.AddToCart(addToCartDTO, buku)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *keranjangController) GetCartByUserID(context *gin.Context) {
	currentUser := context.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	var keranjang entity.Keranjang = c.keranjangService.GetCartByUserID(userID)
	if (keranjang == entity.Keranjang{}) {
		res := helper.BuildErrorResponse("Data tidak ditemukan", "Tidak ada data yang diberikan id_user", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "Sukses", keranjang)
		context.JSON(http.StatusOK, res)
	}
}
