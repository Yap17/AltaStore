package main

import (
	"AltaStore/api"
	adminController "AltaStore/api/v1/admin"
	adminAuthController "AltaStore/api/v1/adminauth"
	contrCategory "AltaStore/api/v1/category"
	productController "AltaStore/api/v1/product"
	userController "AltaStore/api/v1/user"
	userAuthController "AltaStore/api/v1/userauth"
	adminService "AltaStore/business/admin"
	adminAuthService "AltaStore/business/adminauth"
	busCategory "AltaStore/business/category"
	productService "AltaStore/business/product"
	userService "AltaStore/business/user"
	userAuthService "AltaStore/business/userauth"
	"AltaStore/config"
	adminRepository "AltaStore/modules/admin"
	repoCategory "AltaStore/modules/category"
	"AltaStore/modules/migration"
	productRepository "AltaStore/modules/product"
	userRepository "AltaStore/modules/user"

	shopController "AltaStore/api/v1/shopping"
	shopService "AltaStore/business/shopping"
	shopRepository "AltaStore/modules/shopping"
	shopDetailRepository "AltaStore/modules/shoppingdetail"
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(cfg *config.ConfigApp) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName,
	)
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func newRedisConnection(cfg *config.ConfigApp) *redis.Client {
	// stringConnection := fmt.Sprintf(
	// 	"%s:%d",
	// 	cfg.RedisHost, cfg.RedisPort,
	// )
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     stringConnection, // redis port
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })
	// _, err := client.Ping().Result()
	// if err != nil {
	// 	panic(err)
	// }
	// return client
	return nil
}

func main() {
	// retrieves application configuration and returns common values when there is a problem
	config := config.GetConfig()

	// open database server base session
	dbConnection := newDatabaseConnection(config)

	// open redis connection
	//redisConnection := newRedisConnection(config)
	_ = newRedisConnection(config)

	// Initiate Respository Category
	categoryRepo := repoCategory.NewRepository(dbConnection)

	// Initiate Service Category
	categoryService := busCategory.NewService(categoryRepo)

	// Initiate Controller Category
	controllerCategory := contrCategory.NewController(categoryService)

	//initiate user repository
	user := userRepository.NewDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(user)

	//initiate user controller
	userController := userController.NewController(userService)

	// Initiate Respository Category
	//_ = authRepository.NewRepository(redisConnection)

	//initiate admin repository
	admin := adminRepository.NewDBRepository(dbConnection)

	//initiate admin service
	adminService := adminService.NewService(admin)

	//initiate admin controller
	adminController := adminController.NewController(adminService)

	//initiate auth service
	userAuthService := userAuthService.NewService(userService)

	//initiate auth controller
	userAuthController := userAuthController.NewController(userAuthService)

	//initiate auth service
	adminAuthService := adminAuthService.NewService(adminService)

	//initiate auth controller
	adminAuthController := adminAuthController.NewController(adminAuthService)

	// Initiate Respository Product
	product := productRepository.NewRepository(dbConnection)

	// Initiate Service Product
	ProductService := productService.NewService(product)

	// Initiate Controller Product
	productController := productController.NewController(ProductService)

	// initiate shopping repository
	shopRepo := shopRepository.NewRepository(dbConnection)
	shopDetailRepo := shopDetailRepository.NewRepository(dbConnection)

	// initiate shopping service
	shopServc := shopService.NewService(shopRepo, shopDetailRepo)

	// initiate shopping controller
	shopHandler := shopController.NewController(shopServc)
	// create echo http
	e := echo.New()

	// Register API Path and Controller
	api.RegisterPath(e,
		controllerCategory,
		userController,
		adminController,
		userAuthController,
		adminAuthController,
		productController,
		shopHandler,
	)

	// Run server
	func() {
		address := fmt.Sprintf(":%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("Shutdown Echo Service")
		}

	}()
}
