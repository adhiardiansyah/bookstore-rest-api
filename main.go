package main

import (
	"github.com/adhiardiansyah/bookstore-rest-api/config"
	"github.com/adhiardiansyah/bookstore-rest-api/controller"
	"github.com/adhiardiansyah/bookstore-rest-api/middleware"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/adhiardiansyah/bookstore-rest-api/service"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	bukuRepository      repository.BukuRepository      = repository.NewBukuRepository(db)
	kategoriRepository  repository.KategoriRepository  = repository.NewKategoriRepository(db)
	pengarangRepository repository.PengarangRepository = repository.NewPengarangRepository(db)
	penerbitRepository  repository.PenerbitRepository  = repository.NewPenerbitRepository(db)
	roleRepository      repository.RoleRepository      = repository.NewRoleRepository(db)
	keranjangRepository repository.KeranjangRepository = repository.NewKeranjangRepository(db)
	transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
	jwtService          service.JWTService             = service.NewJWTService()
	userService         service.UserService            = service.NewUserService(userRepository)
	bukuService         service.BukuService            = service.NewBukuService(bukuRepository)
	kategoriService     service.KategoriService        = service.NewKategoriService(kategoriRepository)
	pengarangService    service.PengarangService       = service.NewPengarangService(pengarangRepository)
	penerbitService     service.PenerbitService        = service.NewPenerbitService(penerbitRepository)
	roleService         service.RoleService            = service.NewRoleService(roleRepository)
	keranjangService    service.KeranjangService       = service.NewKeranjangService(keranjangRepository)
	transaksiService    service.TransaksiService       = service.NewTransaksiService(transaksiRepository)
	authService         service.AuthService            = service.NewAuthService(userRepository)
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	userController      controller.UserController      = controller.NewUserController(userService, jwtService)
	bukuController      controller.BukuController      = controller.NewBukuController(bukuService, jwtService)
	kategoriController  controller.KategoriController  = controller.NewKategoriController(kategoriService, jwtService)
	pengarangController controller.PengarangController = controller.NewPengarangController(pengarangService, jwtService)
	penerbitController  controller.PenerbitController  = controller.NewPenerbitController(penerbitService, jwtService)
	roleController      controller.RoleController      = controller.NewRoleController(roleService, jwtService)
	keranjangController controller.KeranjangController = controller.NewKeranjangController(keranjangService, jwtService)
	transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.Default())
	r.Static("/images", "./images")

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService, userService))
	{
		userRoutes.GET("/", userController.GetAll)
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

	penerbitRoutes := r.Group("api/penerbit")
	{
		penerbitRoutes.GET("/", penerbitController.GetAll)
		penerbitRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), penerbitController.CreatePenerbit)
		penerbitRoutes.GET("/:id_penerbit", penerbitController.GetByID)
		penerbitRoutes.PUT("/:id_penerbit", middleware.AuthorizeJWT(jwtService, userService), penerbitController.UpdatePenerbit)
		penerbitRoutes.DELETE("/:id_penerbit", middleware.AuthorizeJWT(jwtService, userService), penerbitController.DeletePenerbit)
	}

	roleRoutes := r.Group("api/role")
	{
		roleRoutes.GET("/", roleController.GetAll)
		roleRoutes.POST("/", roleController.CreateRole)
		roleRoutes.GET("/:id_role", roleController.GetByID)
		roleRoutes.PUT("/:id_role", middleware.AuthorizeJWT(jwtService, userService), roleController.UpdateRole)
		roleRoutes.DELETE("/:id_role", middleware.AuthorizeJWT(jwtService, userService), roleController.DeleteRole)
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
