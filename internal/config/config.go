package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB      DB
	Logging Logging
	JWT     JWT
}

type DB struct {
	DbHost string
	DbUser string
	DbPass string
	DbPort string
	DbName string
}

type Logging struct {
	GormLevel   string
	LogrusLevel string
}

type JWT struct {
	SecretKey string
}

var lock = &sync.Mutex{}
var defaultConfig Configuration

func Get() *Configuration {
	lock.Lock()
	defer lock.Unlock()
	return &defaultConfig
}

func Init() *Configuration {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Production")
	} else {
		fmt.Println("Development")
	}

	defaultConfig.DB.DbHost = os.Getenv("DB_HOST")
	defaultConfig.DB.DbUser = os.Getenv("DB_USER")
	defaultConfig.DB.DbPass = os.Getenv("DB_PASS")
	defaultConfig.DB.DbPort = os.Getenv("DB_PORT")
	defaultConfig.DB.DbName = os.Getenv("DB_NAME")
	defaultConfig.Logging.GormLevel = os.Getenv("GORM_LEVEL")
	defaultConfig.Logging.LogrusLevel = os.Getenv("LOGRUS_LEVEL")
	defaultConfig.JWT.SecretKey = os.Getenv("SECRET_KEY")

	return &defaultConfig
}
