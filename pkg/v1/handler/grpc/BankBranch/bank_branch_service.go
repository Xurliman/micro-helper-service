package BankBranch

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/bank_branch"
	"google.golang.org/grpc"
)

type BankBranchService struct {
	bankBranchCase interfaces.BankBranchCaseInterface
	proto.UnimplementedBankBranchServiceServer
}

func NewServer(grpcServer *grpc.Server, bankBranchCase interfaces.BankBranchCaseInterface) {
	bankBranchGrpc := &BankBranchService{bankBranchCase: bankBranchCase}
	proto.RegisterBankBranchServiceServer(grpcServer, bankBranchGrpc)
}

func (service *BankBranchService) Create(context context.Context, request *proto.CreateBankBranchRequest) (*proto.BankBranchProfileResponse, error) {
	data := service.transformBankBranchRpc(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.BankBranchProfileResponse{}, errors.New("please provide all fields")
	}
	bankBranch, err := service.bankBranchCase.Create(data)
	if err != nil {
		return &proto.BankBranchProfileResponse{}, err
	}
	return service.transformBankBranchModel(bankBranch), nil
}

func (service *BankBranchService) Get(context context.Context, request *proto.SingleBankBranchRequest) (*proto.BankBranchProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.BankBranchProfileResponse{}, errors.New("id cannot be blank")
	}

	bankBranch, err := service.bankBranchCase.Get(id)
	if err != nil {
		return &proto.BankBranchProfileResponse{}, err
	}

	return service.transformBankBranchModel(bankBranch), nil
}

func (service *BankBranchService) transformBankBranchRpc(request *proto.CreateBankBranchRequest) models.BankBranch {
	return models.BankBranch{
		Code:       request.GetCode(),
		Name:       request.GetName(),
		BankID:     request.GetBankId(),
		RegionID:   request.GetRegionId(),
		DistrictID: request.GetDistrictId(),
		Address:    request.GetAddress(),
		OpenDate:   request.GetOpenDate(),
		CloseDate:  request.GetCloseDate(),
		CrudDates:  request.GetCrudDates(),
	}
}

func (service *BankBranchService) transformBankBranchModel(bankBranch models.BankBranch) *proto.BankBranchProfileResponse {
	return &proto.BankBranchProfileResponse{
		Id:         int64(bankBranch.ID),
		Name:       bankBranch.Name,
		Code:       bankBranch.Code,
		BankId:     bankBranch.BankID,
		RegionId:   bankBranch.RegionID,
		DistrictId: bankBranch.DistrictID,
		Address:    bankBranch.Address,
		OpenDate:   bankBranch.OpenDate,
		CloseDate:  bankBranch.CloseDate,
		CrudDates:  bankBranch.CrudDates,
	}
}
