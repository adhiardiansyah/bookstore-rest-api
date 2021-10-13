package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/adhiardiansyah/bookstore-rest-api/helper"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.BuildErrorResponse("Gagal memproses permintaan", "Format token salah", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := jwtService.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])

			userID := claims["user_id"]
			user := userService.GetUserByID(int(userID.(float64)))
			c.Set("currentUser", user)
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token tidak valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
