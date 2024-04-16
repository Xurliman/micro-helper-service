package BankBranch

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/bank_branch"
	"github.com/google/uuid"
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
	if data.Code == "" || data.Name == "" {
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
	if id == "" {
		return &proto.BankBranchProfileResponse{}, errors.New("id cannot be blank")
	}

	bankBranch, err := service.bankBranchCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.BankBranchProfileResponse{}, err
	}

	return service.transformBankBranchModel(bankBranch), nil
}

func (service *BankBranchService) transformBankBranchRpc(request *proto.CreateBankBranchRequest) models.BankBranch {
	return models.BankBranch{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		BankId:           request.GetBankId(),
		RegionId:         request.GetRegionId(),
		DistrictId:       request.GetDistrictId(),
		Address:          request.GetAddress(),
		OpenDate:         request.GetOpenDate(),
		CloseDate:        request.GetCloseDate(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *BankBranchService) transformBankBranchModel(bankBranch models.BankBranch) *proto.BankBranchProfileResponse {
	return &proto.BankBranchProfileResponse{
		Id:               bankBranch.ID.String(),
		Name:             bankBranch.Name,
		Code:             bankBranch.Code,
		BankId:           bankBranch.BankId,
		RegionId:         bankBranch.RegionId,
		DistrictId:       bankBranch.DistrictId,
		Address:          bankBranch.Address,
		OpenDate:         bankBranch.OpenDate,
		CloseDate:        bankBranch.CloseDate,
		ActivationDate:   bankBranch.ActivationDate,
		DeactivationDate: bankBranch.DeactivationDate,
		FlexFinId:        bankBranch.FlexFinId,
	}
}
