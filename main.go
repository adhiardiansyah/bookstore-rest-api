package main

import (
	"github.com/adhiardiansyah/bookstore-rest-api/config"
	"github.com/adhiardiansyah/bookstore-rest-api/controller"
	"github.com/adhiardiansyah/bookstore-rest-api/middleware"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/adhiardiansyah/bookstore-rest-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bukuRepository repository.BukuRepository = repository.NewBukuRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	bukuService    service.BukuService       = service.NewBukuService(bukuRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bukuController controller.BukuController = controller.NewBukuController(bukuService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Static("/images", "./images")

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService, userService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.POST("/image", userController.UploadImageUser)
	}

	bukuRoutes := r.Group("api/books")
	{
		bukuRoutes.GET("/", bukuController.GetAll)
		bukuRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), bukuController.CreateBuku)
		bukuRoutes.POST("/image/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.UploadImageBuku)
		bukuRoutes.GET("/:id_buku", bukuController.GetByID)
		bukuRoutes.PUT("/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.UpdateBuku)
		bukuRoutes.DELETE("/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.DeleteBuku)
	}

	r.Run()
}
