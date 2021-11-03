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
	jwtService       service.JWTService
}

func NewKeranjangController(keranjangService service.KeranjangService, jwtService service.JWTService) KeranjangController {
	return &keranjangController{
		keranjangService: keranjangService,
		jwtService:       jwtService,
	}
}

func (c *keranjangController) AddToCart(context *gin.Context) {
	buku_id, err := strconv.Atoi(context.PostForm("buku_id"))
	if err != nil {
		res := helper.BuildErrorResponse("Tidak ada parameter buku_id yang ditemukan", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var addToCartDTO dto.AddToCartDTO

	errDTO := context.ShouldBindJSON(&addToCartDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		currentUser := context.MustGet("currentUser").(entity.User)
		addToCartDTO.User = currentUser
		if buku_id == 0 {
			result := c.keranjangService.AddToCart(addToCartDTO)
			response := helper.BuildResponse(true, "Sukses menambahkan data", result)
			context.JSON(http.StatusOK, response)
		} else {
			result := c.keranjangService.UpdateCart(buku_id, addToCartDTO)
			response := helper.BuildResponse(true, "Sukses menambahkan data", result)
			context.JSON(http.StatusOK, response)
		}
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
