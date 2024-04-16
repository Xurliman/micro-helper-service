package Bank

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/bank"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type BankService struct {
	bankCase interfaces.BankCaseInterface
	proto.UnimplementedBankServiceServer
}

func NewServer(grpcServer *grpc.Server, bankCase interfaces.BankCaseInterface) *BankService {
	bankGrpc := &BankService{bankCase: bankCase}
	proto.RegisterBankServiceServer(grpcServer, bankGrpc)
	return bankGrpc
}

func (service *BankService) Create(context context.Context, request *proto.CreateBankRequest) (*proto.BankProfileResponse, error) {
	data := service.transformBankRPC(request)
	if data.Code == "" || data.Name == "" {
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
	bank, err := service.bankCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.BankProfileResponse{}, err
	}

	return service.transformBankModel(bank), nil
}

func (service *BankService) transformBankRPC(request *proto.CreateBankRequest) models.Bank {
	return models.Bank{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		ShortName:        request.GetShortName(),
		CountryID:        request.GetCountryId(),
		OpenDate:         request.GetOpenDate(),
		CloseDate:        request.GetCloseDate(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *BankService) transformBankModel(bank models.Bank) *proto.BankProfileResponse {
	return &proto.BankProfileResponse{
		Id:               bank.ID.String(),
		Name:             bank.Name,
		Code:             bank.Code,
		ShortName:        bank.ShortName,
		CountryId:        bank.CountryID,
		OpenDate:         bank.OpenDate,
		CloseDate:        bank.CloseDate,
		ActivationDate:   bank.ActivationDate,
		DeactivationDate: bank.DeactivationDate,
		FlexFinId:        bank.FlexFinId,
	}
}
