package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/helper"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Profile(context *gin.Context)
	UploadImageUser(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) GetAll(context *gin.Context) {
	var user []entity.User = c.userService.GetAll()
	res := helper.BuildResponse(true, "Sukses", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Gagal untuk memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.BuildErrorResponse("Gagal memproses permintaan", "Format token salah", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, errToken := c.jwtService.ValidateToken(tokenString)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = int(id)
	u := c.userService.Update(userUpdateDTO.ID, userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.BuildErrorResponse("Gagal memproses permintaan", "Format token salah", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, err := c.jwtService.ValidateToken(tokenString)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) UploadImageUser(context *gin.Context) {
	file, err := context.FormFile("image")
	if err != nil {
		res := helper.BuildErrorResponse("Gagal untuk upload gambar", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	currentUser := context.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = context.SaveUploadedFile(file, path)
	if err != nil {
		res := helper.BuildErrorResponse("Gagal untuk upload gambar", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	u := c.userService.SaveImageUser(userID, path)
	res := helper.BuildResponse(true, "Foto profil berhasil diupload", u)
	context.JSON(http.StatusOK, res)
}
