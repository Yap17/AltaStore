package main

import (
	"AltaStore/api"
	authController "AltaStore/api/v1/auth"
	cateController "AltaStore/api/v1/category"
	shopController "AltaStore/api/v1/shopping"
	userController "AltaStore/api/v1/user"

	authService "AltaStore/business/auth"
	cateService "AltaStore/business/category"
	shopService "AltaStore/business/shopping"
	userService "AltaStore/business/user"

	"AltaStore/config"
	cateRepository "AltaStore/modules/category"
	"AltaStore/modules/migration"
	shopRepository "AltaStore/modules/shopping"
	shopDetailRepository "AltaStore/modules/shoppingdetail"
	userRepository "AltaStore/modules/user"

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
	categoryRepo := cateRepository.NewRepository(dbConnection)

	// Initiate Service Category
	categoryService := cateService.NewService(categoryRepo)

	// Initiate Controller Category
	controllerCategory := cateController.NewController(categoryService)

	//initiate user repository
	user := userRepository.NewDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(user)

	//initiate user controller
	userController := userController.NewController(userService)

	// Initiate Respository Category
	//_ = authRepository.NewRepository(redisConnection)

	//initiate auth service
	authService := authService.NewService(userService)

	//initiate auth controller
	authController := authController.NewController(authService)

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
	api.RegisterPath(e, controllerCategory, userController, authController, shopHandler)

	// Run server
	func() {
		address := fmt.Sprintf(":%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("Shutdown Echo Service")
		}

	}()
}
