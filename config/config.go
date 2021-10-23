package config

import (
	"os"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ConfigApp struct {
	AppHost           string `mapstructure:"app_host"`
	AppPort           int    `mapstructure:"app_port"`
	DbDriver          string `mapstructure:"db_driver"`
	DbHost            string `mapstructure:"db_host"`
	DbPort            int    `mapstructure:"db_port"`
	DbUsername        string `mapstructure:"db_username"`
	DbPassword        string `mapstructure:"db_password"`
	DbName            string `mapstructure:"db_name"`
	JwtSecretKey      string `mapstructure:"jwtsecretkey"`
	RedisHost         string `mapstructure:"redis_host"`
	RedisPort         int    `mapstructure:"redis_port"`
	MidTransServerKey string `mapstructure:"midtransserverkey"`
}

func GetConfig() *ConfigApp {
	// Set default config if error parsing file
	var defaConfig ConfigApp

	defaConfig.AppHost = "localhost"
	defaConfig.AppPort = 9000
	defaConfig.DbDriver = "pgsql"
	defaConfig.DbHost = "localhost"
	defaConfig.DbPort = 5432
	defaConfig.DbUsername = "postgres"
	defaConfig.DbPassword = "alexanderyap"
	defaConfig.DbName = "altastoredb"
	defaConfig.JwtSecretKey = "AltaStore"
	defaConfig.RedisHost = "localhost"
	defaConfig.RedisPort = 9001
	defaConfig.MidTransServerKey = "SB-Mid-server-Q_vEZuSAEJvflvzI7jvWqPZz"

	var (
		err error
		cwd string
	)
	// Geting current directory
	cwd, err = os.Getwd()
	if err != nil {
		log.Info("Failed get current directory, config set to default.")
		return &defaConfig
	}

	// Geting config in file .env
	viper.SetConfigFile(cwd + "/config/.env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Info("Failed read config, config set to default.")
		return &defaConfig
	}

	var finalConfig ConfigApp
	err = viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("Failed bind config, config set to default.")
		return &defaConfig
	}

	return &finalConfig
}
