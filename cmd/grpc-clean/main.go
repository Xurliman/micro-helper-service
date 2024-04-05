package main

import (
	"github.com/Xurliman/banking-microservice/cmd/di"
	"log"
	"net"
)

func main() {

	//db := database.DbConn()

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	//grpcServer := grpc.NewServer()
	grpcServer := di.InitializeServer()
	log.Fatal(grpcServer.Serve(listen))
	//
	//bankCase := initBankServer(db)
	//handler.NewServer(grpcServer, bankCase)
}

//
//func initBankServer(db *gorm.DB) interfaces.BankCaseInterface {
//	bankRepo := repository.NewBank(db)
//	return usecase.NewBank(bankRepo)
//}
