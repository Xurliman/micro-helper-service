package Account

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/account"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type AccountService struct {
	accountCase interfaces.AccountCaseInterface
	proto.UnimplementedAccountServiceServer
}

func NewServer(grpcServer *grpc.Server, accountCase interfaces.AccountCaseInterface) {
	accountGrpc := &AccountService{accountCase: accountCase}
	proto.RegisterAccountServiceServer(grpcServer, accountGrpc)
}

func (service *AccountService) Create(context context.Context, request *proto.CreateAccountRequest) (*proto.AccountProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.AccountProfileResponse{}, errors.New("please provide all fields")
	}

	account, err := service.accountCase.Create(data)
	if err != nil {
		return &proto.AccountProfileResponse{}, err
	}
	return service.transformClassifierModel(account), nil
}

func (service *AccountService) Get(context context.Context, request *proto.SingleAccountRequest) (*proto.AccountProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.AccountProfileResponse{}, errors.New("id cannot be blank")
	}

	account, err := service.accountCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.AccountProfileResponse{}, err
	}

	return service.transformClassifierModel(account), nil
}

func (service *AccountService) transformClassifierRpc(request *proto.CreateAccountRequest) models.Account {
	return models.Account{
		Code:        request.GetCode(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		FlexFinId:   request.GetFlexFinId(),
	}
}

func (service *AccountService) transformClassifierModel(account models.Account) *proto.AccountProfileResponse {
	return &proto.AccountProfileResponse{
		Id:          account.ID.String(),
		Code:        account.Code,
		Name:        account.Name,
		Description: account.Description,
		FlexFinId:   account.FlexFinId,
	}
}
