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
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	bukuRepository      repository.BukuRepository      = repository.NewBukuRepository(db)
	kategoriRepository  repository.KategoriRepository  = repository.NewKategoriRepository(db)
	pengarangRepository repository.PengarangRepository = repository.NewPengarangRepository(db)
	keranjangRepository repository.KeranjangRepository = repository.NewKeranjangRepository(db)
	transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
	jwtService          service.JWTService             = service.NewJWTService()
	userService         service.UserService            = service.NewUserService(userRepository)
	bukuService         service.BukuService            = service.NewBukuService(bukuRepository)
	kategoriService     service.KategoriService        = service.NewKategoriService(kategoriRepository)
	pengarangService    service.PengarangService       = service.NewPengarangService(pengarangRepository)
	keranjangService    service.KeranjangService       = service.NewKeranjangService(keranjangRepository)
	transaksiService    service.TransaksiService       = service.NewTransaksiService(transaksiRepository)
	authService         service.AuthService            = service.NewAuthService(userRepository)
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	userController      controller.UserController      = controller.NewUserController(userService, jwtService)
	bukuController      controller.BukuController      = controller.NewBukuController(bukuService, jwtService)
	kategoriController  controller.KategoriController  = controller.NewKategoriController(kategoriService, jwtService)
	pengarangController controller.PengarangController = controller.NewPengarangController(pengarangService, jwtService)
	keranjangController controller.KeranjangController = controller.NewKeranjangController(keranjangService, jwtService)
	transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService, jwtService)
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

	bukuRoutes := r.Group("api/book")
	{
		bukuRoutes.GET("/", bukuController.GetAll)
		bukuRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), bukuController.CreateBuku)
		bukuRoutes.POST("/image/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.UploadImageBuku)
		bukuRoutes.GET("/:id_buku", bukuController.GetByID)
		bukuRoutes.PUT("/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.UpdateBuku)
		bukuRoutes.DELETE("/:id_buku", middleware.AuthorizeJWT(jwtService, userService), bukuController.DeleteBuku)
	}

	kategoriRoutes := r.Group("api/category")
	{
		kategoriRoutes.GET("/", kategoriController.GetAll)
		kategoriRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), kategoriController.CreateKategori)
		kategoriRoutes.GET("/:id_kategori", kategoriController.GetByID)
		kategoriRoutes.PUT("/:id_kategori", middleware.AuthorizeJWT(jwtService, userService), kategoriController.UpdateKategori)
		kategoriRoutes.DELETE("/:id_kategori", middleware.AuthorizeJWT(jwtService, userService), kategoriController.DeleteKategori)
	}

	pengarangRoutes := r.Group("api/pengarang")
	{
		pengarangRoutes.GET("/", pengarangController.GetAll)
		pengarangRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), pengarangController.CreatePengarang)
		pengarangRoutes.GET("/:id_pengarang", pengarangController.GetByID)
		pengarangRoutes.PUT("/:id_pengarang", middleware.AuthorizeJWT(jwtService, userService), pengarangController.UpdatePengarang)
		pengarangRoutes.DELETE("/:id_pengarang", middleware.AuthorizeJWT(jwtService, userService), pengarangController.DeletePengarang)
	}

	keranjangRoutes := r.Group("api/cart")
	{
		keranjangRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), keranjangController.AddToCart)
		keranjangRoutes.GET("/", middleware.AuthorizeJWT(jwtService, userService), keranjangController.GetCartByUserID)
	}

	transaksiRoutes := r.Group("api/transaction")
	{
		transaksiRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), transaksiController.CreateTransaksi)
	}

	r.Run()
}
