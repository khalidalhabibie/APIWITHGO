package repository

import(
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	//"../model"
	"github.com/khalid/apiWithGO/model"
	
)


func DB() *gorm.DB{
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")


	dsn := username + ":" + password + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		log.Fatal("Error Connecting to Databse")
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Transaction{})
	return db
}
