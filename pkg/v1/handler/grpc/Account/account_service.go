package Account

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/account"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

type AccountService struct {
	accountCase interfaces.AccountCaseInterface
	proto.UnimplementedAccountServiceServer
}

func NewServer(grpcServer *grpc.Server, accountCase interfaces.AccountCaseInterface) *AccountService {
	accountGrpc := &AccountService{accountCase: accountCase}
	proto.RegisterAccountServiceServer(grpcServer, accountGrpc)
	return accountGrpc
}

//func (services *AccountService) List(context context.Context, request *proto.ListAccountRequest) (*proto.ListAccountResponse, error) {
//
//}

func (service *AccountService) Create(context context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	account := models.Account{
		Code:        request.GetCode(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		FlexFinId:   request.GetFlexFinId(),
	}

	validated := service.validated(account)
	if !validated {
		return &proto.CreateAccountResponse{}, errors.New("please provide all fields")
	}

	data, err := service.accountCase.Create(account)
	if err != nil {
		return &proto.CreateAccountResponse{}, err
	}

	return &proto.CreateAccountResponse{
		Id:          data.ID.String(),
		Code:        data.Code,
		Name:        data.Name,
		Description: data.Description,
		FlexFinId:   data.FlexFinId,
	}, nil
}

func (service *AccountService) Get(context context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	code := request.GetCode()
	log.Printf("account code: %s", code)
	if code == "" {
		return &proto.GetAccountResponse{}, errors.New("id cannot be blank")
	}

	data, err := service.accountCase.Get(uuid.MustParse(request.GetId()))
	if err != nil {
		return &proto.GetAccountResponse{}, err
	}

	return &proto.GetAccountResponse{
		Id:          data.ID.String(),
		Code:        data.Code,
		Name:        data.Name,
		Description: data.Description,
		FlexFinId:   data.FlexFinId,
	}, nil
}

//func (services *AccountService) Update(context context.Context, request *proto.UpdateAccountRequest) (*proto.UpdateAccountResponse, error) {
//
//}

func (service *AccountService) validated(account models.Account) bool {
	validate := validator.New()

	err := validate.Struct(account)
	if err != nil {
		log.Fatalf("Validation error: %s", err.(validator.ValidationErrors))
		return false
	}
	return true
}
