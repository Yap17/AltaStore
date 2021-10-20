package main

import (
	"AltaStore/api"
	contrCategory "AltaStore/api/v1/category"
	userController "AltaStore/api/v1/user"
	busCategory "AltaStore/business/category"
	userService "AltaStore/business/user"
	"AltaStore/config"
	repoCategory "AltaStore/modules/category"
	userRepository "AltaStore/modules/user"
	"fmt"

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

	// migration.TableMigration(db)

	return db
}

func main() {
	// retrieves application configuration and returns common values when there is a problem
	config := config.GetConfig()

	// open database server base session
	dbConnection := newDatabaseConnection(config)

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

	// create echo http
	e := echo.New()

	// Register API Path and Controller
	api.RegisterPath(e, controllerCategory, userController)

	// Run server
	func() {
		address := fmt.Sprintf(":%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("Shutdown Echo Service")
		}

	}()
}
