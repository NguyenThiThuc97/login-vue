package models

import (
	"fmt"
	"log"
	"login/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

var DB *gorm.DB

func ConnectDataBase() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	viper_user := config.PostgresUser
	viper_password := config.PostgresPassword
	viper_db := config.PostgresDb
	viper_host := config.PostgresHost
	viper_port := config.PostgresPort
	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{})
	// Initialise value
	// m := User{Username: "user1", Password: "123"}
	// db.Create(&m)
	DB = db
}
