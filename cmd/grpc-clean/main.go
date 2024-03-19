package main

import (
	"fmt"
	database "github.com/Xurliman/banking-microservice/internal/db"
	"github.com/Xurliman/banking-microservice/internal/models"
	handler "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/Xurliman/banking-microservice/pkg/v1/repository"
	"github.com/Xurliman/banking-microservice/pkg/v1/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	db := database.DbConn()
	migrations(db)

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	bankCase := initBankServer(db)
	handler.NewServer(grpcServer, bankCase)
	log.Fatal(grpcServer.Serve(listen))
}

func initBankServer(db *gorm.DB) interfaces.BankCaseInterface {
	bankRepo := repository.NewBank(db)
	return usecase.NewBank(bankRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Bank{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrateeeed!")
	}
}
