package config

import (
	"auth_fiber/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



var DB *gorm.DB

func Connect() *gorm.DB {


	godotenv.Load()

	dbhost :=os.Getenv("MYSQL_HOST")
	dbPassword :=os.Getenv("MYSQL_PASSWORD")
	dbUserName :=os.Getenv("MYSQL_USER")
	dbName := os.Getenv("MYSQL_DBNAME")

	fmt.Println(`Host`+dbhost)
	fmt.Println(`dbPassword`+dbPassword)
	fmt.Println(`dbUserName`+dbUserName)
	fmt.Println(`dbName`+dbName)

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",dbUserName,dbPassword,dbhost,dbName)

	dbConfig :=gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var db,err=gorm.Open(mysql.Open(connection), &dbConfig)

	if(err !=nil){
		panic("db connection failed")
	}

	DB=db;
	fmt.Println("Db connected...")

	AutoMigrate(db)

	return db;
}



func AutoMigrate(connection *gorm.DB){
	 connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	 )
}