package database

import (
	"fmt"
	"github.com/Xurliman/banking-microservice/config"
	cli_command "github.com/Xurliman/banking-microservice/internal/cli-command"
	"github.com/Xurliman/banking-microservice/internal/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DbConn() *gorm.DB {
	log.Printf(manageConnection())
	db, err := gorm.Open(
		mysql.Open(manageConnection()), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	DB = db
	migrations.Execute(db)
	cli_command.Organize(db)
	return DB
}

func manageConnection() string {
	switch config.GetDatabaseConnection() {
	case "mysql":
		return fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			config.GetDatabaseUser(),
			//config.GetDatabasePassword(),
			config.GetDatabaseHost(),
			config.GetDatabasePort(),
			config.GetDatabaseName())
	case "postgres":
		return fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%v sslmode=disable",
			config.GetDatabaseHost(),
			config.GetDatabaseUser(),
			config.GetDatabaseName(),
			config.GetDatabasePassword(),
			config.GetDatabasePort())
	default:
		return fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			config.GetDatabaseUser(),
			//config.GetDatabasePassword(),
			config.GetDatabaseHost(),
			config.GetDatabasePort(),
			config.GetDatabaseName())
	}
}
