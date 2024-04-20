package RegistrationPlace

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/registration_place"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type RegistrationPlaceService struct {
	registrationPlaceCase interfaces.RegistrationPlaceCaseInterface
	proto.UnimplementedRegistrationPlaceServiceServer
}

func NewServer(grpcServer *grpc.Server, registrationPlaceCase interfaces.RegistrationPlaceCaseInterface) *RegistrationPlaceService {
	registrationPlaceGrpc := &RegistrationPlaceService{registrationPlaceCase: registrationPlaceCase}
	proto.RegisterRegistrationPlaceServiceServer(grpcServer, registrationPlaceGrpc)
	return registrationPlaceGrpc
}

func (service *RegistrationPlaceService) Create(context context.Context, request *proto.CreateRegistrationPlaceRequest) (*proto.RegistrationPlaceProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.RegistrationPlaceProfileResponse{}, errors.New("please provide all fields")
	}

	registrationPlace, err := service.registrationPlaceCase.Create(data)
	if err != nil {
		return &proto.RegistrationPlaceProfileResponse{}, err
	}
	return service.transformClassifierModel(registrationPlace), nil
}

func (service *RegistrationPlaceService) Get(context context.Context, request *proto.SingleRegistrationPlaceRequest) (*proto.RegistrationPlaceProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.RegistrationPlaceProfileResponse{}, errors.New("id cannot be blank")
	}

	registrationPlace, err := service.registrationPlaceCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.RegistrationPlaceProfileResponse{}, err
	}

	return service.transformClassifierModel(registrationPlace), nil
}

func (service *RegistrationPlaceService) transformClassifierRpc(request *proto.CreateRegistrationPlaceRequest) models.RegistrationPlace {
	return models.RegistrationPlace{
		Code:      request.GetCode(),
		Name:      request.GetName(),
		FlexFinId: request.GetFlexFinId(),
	}
}

func (service *RegistrationPlaceService) transformClassifierModel(registrationPlace models.RegistrationPlace) *proto.RegistrationPlaceProfileResponse {
	return &proto.RegistrationPlaceProfileResponse{
		Id:        registrationPlace.ID.String(),
		Code:      registrationPlace.Code,
		Name:      registrationPlace.Name,
		FlexFinId: registrationPlace.FlexFinId,
	}
}
