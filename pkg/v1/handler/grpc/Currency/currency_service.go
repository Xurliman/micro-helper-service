package Currency

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/currency"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type CurrencyServer struct {
	currencyCase interfaces.CurrencyCaseInterface
	proto.UnimplementedCurrencyServiceServer
}

func NewServer(grpcServer *grpc.Server, currencyCase interfaces.CurrencyCaseInterface) {
	currencyGrpc := &CurrencyServer{currencyCase: currencyCase}
	proto.RegisterCurrencyServiceServer(grpcServer, currencyGrpc)
}

func (service *CurrencyServer) Create(context context.Context, request *proto.CreateCurrencyRequest) (*proto.CurrencyProfileResponse, error) {
	data := service.transformCurrencyRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.CurrencyProfileResponse{}, errors.New("please provide all fields")
	}

	currency, err := service.currencyCase.Create(data)
	if err != nil {
		return &proto.CurrencyProfileResponse{}, err
	}
	return service.transformCurrencyModel(currency), nil
}

func (service *CurrencyServer) Get(context context.Context, request *proto.SingleCurrencyRequest) (*proto.CurrencyProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.CurrencyProfileResponse{}, errors.New("id cannot be blank")
	}
	currency, err := service.currencyCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.CurrencyProfileResponse{}, err
	}

	return service.transformCurrencyModel(currency), nil
}

func (service *CurrencyServer) transformCurrencyRpc(request *proto.CreateCurrencyRequest) models.Currency {
	return models.Currency{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		ShortName:        request.GetShortName(),
		Scale:            request.GetScale(),
		ScaleName:        request.GetScaleName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *CurrencyServer) transformCurrencyModel(currency models.Currency) *proto.CurrencyProfileResponse {
	return &proto.CurrencyProfileResponse{
		Id:               currency.ID.String(),
		Code:             currency.Code,
		Name:             currency.Name,
		ShortName:        currency.ShortName,
		Scale:            currency.Scale,
		ScaleName:        currency.ScaleName,
		ActivationDate:   currency.ActivationDate,
		DeactivationDate: currency.DeactivationDate,
		FlexFinId:        currency.FlexFinId,
	}
}
