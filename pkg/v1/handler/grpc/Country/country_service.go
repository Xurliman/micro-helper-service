package Country

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/country"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CountryService struct {
	countryCase interfaces.CountryCaseInterface
	proto.UnimplementedCountryServiceServer
}

func NewServer(grpcServer *grpc.Server, countryCase interfaces.CountryCaseInterface) {
	countryGrpc := &CountryService{countryCase: countryCase}
	proto.RegisterCountryServiceServer(grpcServer, countryGrpc)
}

func (service *CountryService) Create(context context.Context, request *proto.CreateCountryRequest) (*proto.CountryProfileResponse, error) {
	data := service.transformCountryRpc(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.CountryProfileResponse{}, errors.New("please provide all fields")
	}

	country, err := service.countryCase.Create(data)
	if err != nil {
		return &proto.CountryProfileResponse{}, err
	}
	return service.transformCountryModel(country), nil
}

func (service *CountryService) Get(context context.Context, request *proto.SingleCountryRequest) (*proto.CountryProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.CountryProfileResponse{}, errors.New("id cannot be blank")
	}
	bank, err := service.countryCase.Get(id)

	if err != nil {
		return &proto.CountryProfileResponse{}, err
	}

	return service.transformCountryModel(bank), nil
}

func (service *CountryService) transformCountryRpc(request *proto.CreateCountryRequest) models.Country {
	return models.Country{
		Code:          request.GetCode(),
		Name:          request.GetName(),
		ShortName:     request.GetShortName(),
		CurrencyID:    request.GetCurrencyId(),
		CodeAlpha2:    request.GetCodeAlpha2(),
		CodeAlpha3:    request.GetCodeAlpha3(),
		TerritoryCode: request.GetTerritoryCode(),
		CrudDates:     request.GetCrudDates(),
	}
}

func (service *CountryService) transformCountryModel(country models.Country) *proto.CountryProfileResponse {
	return &proto.CountryProfileResponse{
		Id:            int64(country.ID),
		Code:          country.Code,
		Name:          country.Name,
		ShortName:     country.ShortName,
		CurrencyId:    country.CurrencyID,
		CodeAlpha2:    country.CodeAlpha2,
		CodeAlpha3:    country.CodeAlpha3,
		TerritoryCode: country.TerritoryCode,
		CrudDates:     country.CrudDates,
	}
}
