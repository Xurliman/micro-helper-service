package main

import (
	"fmt"
	database "github.com/Xurliman/banking-microservice/internal/db"
	handlerAccount "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Account"
	handlerBank "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Bank"
	handlerBankBranch "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/BankBranch"
	handlerBorrowerType "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/BorrowerType"
	handlerClientTypeClassifier "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/ClientTypeClassifier"
	handlerCountry "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Country"
	handlerCurrency "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Currency"
	handlerDirectOrgan "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/DirectOrgan"
	handlerDistrict "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/District"
	handlerEducationType "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/EducationType"
	handlerNationalEconomySectorNew "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/NationalEconomySectorNew"
	handlerNationalEconomySectorOld "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/NationalEconomySectorOld"
	handlerPassportType "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/PassportType"
	handlerPaymentType "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/PaymentType"
	handlerRegion "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Region"
	handlerRegistrationPlace "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/RegistrationPlace"
	handlerResidencyType "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/ResidencyType"
	handlerTaxOrganisation "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/TaxOrganisation"
	"github.com/Xurliman/banking-microservice/pkg/v1/repository"
	"github.com/Xurliman/banking-microservice/pkg/v1/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func main() {
	db := database.DbConn()
	port := os.Getenv("APP_PORT")
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}
	grpcServer := grpc.NewServer()
	initServers(db, grpcServer)
	log.Fatal(grpcServer.Serve(listen))
}

func initServers(db *gorm.DB, grpcServer *grpc.Server) []interface{} {
	return []interface{}{
		handlerAccount.NewServer(grpcServer, usecase.NewAccount(repository.NewAccount(db))),
		handlerBank.NewServer(grpcServer, usecase.NewBank(repository.NewBank(db))),
		handlerBankBranch.NewServer(grpcServer, usecase.NewBankBranch(repository.NewBankBranch(db))),
		handlerBorrowerType.NewServer(grpcServer, usecase.NewBorrowerType(repository.NewBorrowerType(db))),
		handlerClientTypeClassifier.NewServer(grpcServer, usecase.NewClientTypeClassifier(repository.NewClientTypeClassifier(db))),
		handlerCountry.NewServer(grpcServer, usecase.NewCountry(repository.NewCountry(db))),
		handlerCurrency.NewServer(grpcServer, usecase.NewCurrency(repository.NewCurrency(db))),
		handlerDirectOrgan.NewServer(grpcServer, usecase.NewDirectOrgan(repository.NewDirectOrgan(db))),
		handlerDistrict.NewServer(grpcServer, usecase.NewDistrict(repository.NewDistrict(db))),
		handlerEducationType.NewServer(grpcServer, usecase.NewEducationType(repository.NewEducationType(db))),
		handlerNationalEconomySectorNew.NewServer(grpcServer, usecase.NewNationalEconomySectorNew(repository.NewNationalEconomySectorNew(db))),
		handlerNationalEconomySectorOld.NewServer(grpcServer, usecase.NewNationalEconomySectorOld(repository.NewNationalEconomySectorOld(db))),
		handlerPassportType.NewServer(grpcServer, usecase.NewPassportType(repository.NewPassportType(db))),
		handlerPaymentType.NewServer(grpcServer, usecase.NewPaymentType(repository.NewPaymentType(db))),
		handlerRegion.NewServer(grpcServer, usecase.NewRegion(repository.NewRegion(db))),
		handlerRegistrationPlace.NewServer(grpcServer, usecase.NewRegistrationPlace(repository.NewRegistrationPlace(db))),
		handlerResidencyType.NewServer(grpcServer, usecase.NewResidencyType(repository.NewResidencyType(db))),
		handlerTaxOrganisation.NewServer(grpcServer, usecase.NewTaxOrganisation(repository.NewTaxOrganisation(db))),
	}
}
