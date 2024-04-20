package EducationType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/education_type"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type EducationTypeService struct {
	educationTypeCase interfaces.EducationTypeCaseInterface
	proto.UnimplementedEducationTypeServiceServer
}

func NewServer(grpcServer *grpc.Server, educationTypeCase interfaces.EducationTypeCaseInterface) *EducationTypeService {
	educationTypeGrpc := &EducationTypeService{educationTypeCase: educationTypeCase}
	proto.RegisterEducationTypeServiceServer(grpcServer, educationTypeGrpc)
	return educationTypeGrpc
}

func (service *EducationTypeService) Create(context context.Context, request *proto.CreateEducationTypeRequest) (*proto.EducationTypeProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.EducationTypeProfileResponse{}, errors.New("please provide all fields")
	}

	educationType, err := service.educationTypeCase.Create(data)
	if err != nil {
		return &proto.EducationTypeProfileResponse{}, err
	}
	return service.transformClassifierModel(educationType), nil
}

func (service *EducationTypeService) Get(context context.Context, request *proto.SingleEducationTypeRequest) (*proto.EducationTypeProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.EducationTypeProfileResponse{}, errors.New("id cannot be blank")
	}

	educationType, err := service.educationTypeCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.EducationTypeProfileResponse{}, err
	}

	return service.transformClassifierModel(educationType), nil
}

func (service *EducationTypeService) transformClassifierRpc(request *proto.CreateEducationTypeRequest) models.EducationType {
	return models.EducationType{
		Code:      request.GetCode(),
		Name:      request.GetName(),
		FlexFinId: request.GetFlexFinId(),
	}
}

func (service *EducationTypeService) transformClassifierModel(educationType models.EducationType) *proto.EducationTypeProfileResponse {
	return &proto.EducationTypeProfileResponse{
		Id:        educationType.ID.String(),
		Code:      educationType.Code,
		Name:      educationType.Name,
		FlexFinId: educationType.FlexFinId,
	}
}
