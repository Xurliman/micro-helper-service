package NationalEconomySector

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector"
	"google.golang.org/grpc"
)

type NationalEconomySectorService struct {
	nationalEconomySectorCase interfaces.NationalEconomySectorCaseInterface
	proto.UnimplementedNationalEconomySectorServiceServer
}

func NewServer(grpcServer *grpc.Server, nationalEconomySectorCase interfaces.NationalEconomySectorCaseInterface) {
	nationalEconomySectorGrpc := &NationalEconomySectorService{nationalEconomySectorCase: nationalEconomySectorCase}
	proto.RegisterNationalEconomySectorServiceServer(grpcServer, nationalEconomySectorGrpc)
}

func (service *NationalEconomySectorService) Create(context context.Context, request *proto.CreateNationalEconomySectorRequest) (*proto.NationalEconomySectorProfileResponse, error) {
	data := service.transformNationalEconomySectorRPC(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.NationalEconomySectorProfileResponse{}, errors.New("please provide all fields")
	}

	nationalEconomySector, err := service.nationalEconomySectorCase.Create(data)
	if err != nil {
		return &proto.NationalEconomySectorProfileResponse{}, err
	}
	return service.transformNationalEconomySectorModel(nationalEconomySector), nil
}

func (service *NationalEconomySectorService) Get(context context.Context, request *proto.SingleNationalEconomySectorRequest) (*proto.NationalEconomySectorProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.NationalEconomySectorProfileResponse{}, errors.New("id cannot be blank")
	}
	nationalEconomySector, err := service.nationalEconomySectorCase.Get(id)

	if err != nil {
		return &proto.NationalEconomySectorProfileResponse{}, err
	}

	return service.transformNationalEconomySectorModel(nationalEconomySector), nil
}

func (service *NationalEconomySectorService) transformNationalEconomySectorRPC(request *proto.CreateNationalEconomySectorRequest) models.NationalEconomySector {
	return models.NationalEconomySector{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		CBUCode:          request.GetCbuCode(),
		CBUGroupCode:     request.GetCbuGroupCode(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		CBUReferenceKey:  request.GetCbuReferenceKey(),
	}
}

func (service *NationalEconomySectorService) transformNationalEconomySectorModel(nationalEconomySector models.NationalEconomySector) *proto.NationalEconomySectorProfileResponse {
	return &proto.NationalEconomySectorProfileResponse{
		Id:               int64(nationalEconomySector.ID),
		Name:             nationalEconomySector.Name,
		Code:             nationalEconomySector.Code,
		CbuCode:          nationalEconomySector.CBUCode,
		CbuGroupCode:     nationalEconomySector.CBUGroupCode,
		ActivationDate:   nationalEconomySector.ActivationDate,
		DeactivationDate: nationalEconomySector.DeactivationDate,
		CbuReferenceKey:  nationalEconomySector.CBUReferenceKey,
	}
}
