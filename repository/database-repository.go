package repository

import(
	"os"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/khalidalhabibie/APIWITHGO/model"
	"fmt"

	
)


func DB() *gorm.DB{
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")



	//dsn := username + ":" + password + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username,password, host, port, dbname)
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		log.Fatal("Error Connecting to Databse")
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Transaction{})
	return db
}
