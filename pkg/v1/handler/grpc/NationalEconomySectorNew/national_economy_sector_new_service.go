package NationalEconomySectorNew

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_new"
	"google.golang.org/grpc"
)

type NationalEconomySectorNewService struct {
	nationalEconomySectorNewCase interfaces.NationalEconomySectorNewCaseInterface
	proto.UnimplementedNationalEconomySectorNewServiceServer
}

func NewServer(grpcServer *grpc.Server, nationalEconomySectorNewCase interfaces.NationalEconomySectorNewCaseInterface) {
	nationalEconomySectorNewGrpc := &NationalEconomySectorNewService{nationalEconomySectorNewCase: nationalEconomySectorNewCase}
	proto.RegisterNationalEconomySectorNewServiceServer(grpcServer, nationalEconomySectorNewGrpc)
}

func (service *NationalEconomySectorNewService) Create(context context.Context, request *proto.CreateNationalEconomySectorNewRequest) (*proto.NationalEconomySectorNewProfileResponse, error) {
	data := service.transformNationalEconomySectorNewRPC(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.NationalEconomySectorNewProfileResponse{}, errors.New("please provide all fields")
	}

	nationalEconomySectorNew, err := service.nationalEconomySectorNewCase.Create(data)
	if err != nil {
		return &proto.NationalEconomySectorNewProfileResponse{}, err
	}
	return service.transformNationalEconomySectorNewModel(nationalEconomySectorNew), nil
}

func (service *NationalEconomySectorNewService) Get(context context.Context, request *proto.SingleNationalEconomySectorNewRequest) (*proto.NationalEconomySectorNewProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.NationalEconomySectorNewProfileResponse{}, errors.New("id cannot be blank")
	}
	nationalEconomySectorNew, err := service.nationalEconomySectorNewCase.Get(id)

	if err != nil {
		return &proto.NationalEconomySectorNewProfileResponse{}, err
	}

	return service.transformNationalEconomySectorNewModel(nationalEconomySectorNew), nil
}

func (service *NationalEconomySectorNewService) transformNationalEconomySectorNewRPC(request *proto.CreateNationalEconomySectorNewRequest) models.NationalEconomySectorNew {
	return models.NationalEconomySectorNew{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		CBUCode:          request.GetCbuCode(),
		CBUGroupCode:     request.GetCbuGroupCode(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		CBUReferenceKey:  request.GetCbuReferenceKey(),
	}
}

func (service *NationalEconomySectorNewService) transformNationalEconomySectorNewModel(nationalEconomySectorNew models.NationalEconomySectorNew) *proto.NationalEconomySectorNewProfileResponse {
	return &proto.NationalEconomySectorNewProfileResponse{
		Id:               int64(nationalEconomySectorNew.ID),
		Name:             nationalEconomySectorNew.Name,
		Code:             nationalEconomySectorNew.Code,
		CbuCode:          nationalEconomySectorNew.CBUCode,
		CbuGroupCode:     nationalEconomySectorNew.CBUGroupCode,
		ActivationDate:   nationalEconomySectorNew.ActivationDate,
		DeactivationDate: nationalEconomySectorNew.DeactivationDate,
		CbuReferenceKey:  nationalEconomySectorNew.CBUReferenceKey,
	}
}
