package controller

import (
	"net/http"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/helper"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
	CreateTransaksi(context *gin.Context)
}

type transaksiController struct {
	transaksiService service.TransaksiService
	jwtService       service.JWTService
}

func NewTransaksiController(transaksiService service.TransaksiService, jwtService service.JWTService) TransaksiController {
	return &transaksiController{
		transaksiService: transaksiService,
		jwtService:       jwtService,
	}
}

func (c *transaksiController) CreateTransaksi(context *gin.Context) {
	var createTransaksiDTO dto.CreateTransaksiDTO
	var createDetailTransaksiDTO dto.CreateDetailTransaksiDTO
	errDTO := context.ShouldBindJSON(&createTransaksiDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		currentUser := context.MustGet("currentUser").(entity.User)
		createTransaksiDTO.User = currentUser
		result1, _ := c.transaksiService.CreateTransaksi(createTransaksiDTO, createDetailTransaksiDTO)
		response := helper.BuildResponse(true, "Sukses menambahkan data", result1)
		context.JSON(http.StatusOK, response)
	}
}
