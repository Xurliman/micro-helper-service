package ResidencyType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/residency_type"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ResidencyTypeService struct {
	residencyTypeCase interfaces.ResidencyTypeCaseInterface
	proto.UnimplementedResidencyTypeServiceServer
}

func NewServer(grpcServer *grpc.Server, residencyTypeCase interfaces.ResidencyTypeCaseInterface) *ResidencyTypeService {
	residencyTypeGrpc := &ResidencyTypeService{residencyTypeCase: residencyTypeCase}
	proto.RegisterResidencyTypeServiceServer(grpcServer, residencyTypeGrpc)
	return residencyTypeGrpc
}

func (service *ResidencyTypeService) Create(context context.Context, request *proto.CreateResidencyTypeRequest) (*proto.ResidencyTypeProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.ResidencyTypeProfileResponse{}, errors.New("please provide all fields")
	}

	residencyType, err := service.residencyTypeCase.Create(data)
	if err != nil {
		return &proto.ResidencyTypeProfileResponse{}, err
	}
	return service.transformClassifierModel(residencyType), nil
}

func (service *ResidencyTypeService) Get(context context.Context, request *proto.SingleResidencyTypeRequest) (*proto.ResidencyTypeProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.ResidencyTypeProfileResponse{}, errors.New("id cannot be blank")
	}

	residencyType, err := service.residencyTypeCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.ResidencyTypeProfileResponse{}, err
	}

	return service.transformClassifierModel(residencyType), nil
}

func (service *ResidencyTypeService) transformClassifierRpc(request *proto.CreateResidencyTypeRequest) models.ResidencyType {
	return models.ResidencyType{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		NameUz:           request.GetNameUz(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *ResidencyTypeService) transformClassifierModel(residencyType models.ResidencyType) *proto.ResidencyTypeProfileResponse {
	return &proto.ResidencyTypeProfileResponse{
		Id:               residencyType.ID.String(),
		Code:             residencyType.Code,
		Name:             residencyType.Name,
		ActivationDate:   residencyType.ActivationDate,
		DeactivationDate: residencyType.DeactivationDate,
		NameUz:           residencyType.NameUz,
		FlexFinId:        residencyType.FlexFinId,
	}
}
