package grpc

import (
	"context"
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/bank"
	"google.golang.org/grpc"
)

type BankService struct {
	bankCase interfaces.BankCaseInterface
	proto.UnimplementedBankServiceServer
}

func NewServer(grpcServer *grpc.Server, bankCase interfaces.BankCaseInterface) {
	bankGrpc := &BankService{bankCase: bankCase}
	proto.RegisterBankServiceServer(grpcServer, bankGrpc)
}

func (service *BankService) Create(context context.Context, request *proto.CreateBankRequest) (*proto.BankProfileResponse, error) {
	data := service.transformBankRPC(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.BankProfileResponse{}, errors.New("please provide all fields")
	}

	bank, err := service.bankCase.Create(data)
	if err != nil {
		return &proto.BankProfileResponse{}, err
	}
	return service.transformBankModel(bank), nil
}

func (service *BankService) Get(context context.Context, request *proto.SingleBankRequest) (*proto.BankProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.BankProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := service.bankCase.Get(id)
	if err != nil {
		return &proto.BankProfileResponse{}, err
	}
	return service.transformBankModel(user), nil
}

func (service *BankService) transformBankRPC(request *proto.CreateBankRequest) models.Bank {
	return models.Bank{Name: request.GetName(), Code: request.GetCode()}
}

func (service *BankService) transformBankModel(bank models.Bank) *proto.BankProfileResponse {
	return &proto.BankProfileResponse{Id: int64(bank.ID), Name: bank.Name, Code: bank.Code}
}
