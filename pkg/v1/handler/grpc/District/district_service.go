package District

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/district"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type DistrictService struct {
	districtCase interfaces.DistrictCaseInterface
	proto.UnimplementedDistrictServiceServer
}

func NewServer(grpcServer *grpc.Server, districtCase interfaces.DistrictCaseInterface) {
	districtGrpc := &DistrictService{districtCase: districtCase}
	proto.RegisterDistrictServiceServer(grpcServer, districtGrpc)
}

func (service *DistrictService) Create(context context.Context, request *proto.CreateDistrictRequest) (*proto.DistrictProfileResponse, error) {
	data := service.transformDistrictRPC(request)
	if data.Code == "" || data.Name == "" {
		return &proto.DistrictProfileResponse{}, errors.New("please provide all fields")
	}

	district, err := service.districtCase.Create(data)
	if err != nil {
		return &proto.DistrictProfileResponse{}, err
	}
	return service.transformDistrictModel(district), nil
}

func (service *DistrictService) Get(context context.Context, request *proto.SingleDistrictRequest) (*proto.DistrictProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.DistrictProfileResponse{}, errors.New("id cannot be blank")
	}
	district, err := service.districtCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.DistrictProfileResponse{}, err
	}

	return service.transformDistrictModel(district), nil
}

func (service *DistrictService) transformDistrictRPC(request *proto.CreateDistrictRequest) models.District {
	return models.District{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		RegionId:         request.GetRegionId(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *DistrictService) transformDistrictModel(district models.District) *proto.DistrictProfileResponse {
	return &proto.DistrictProfileResponse{
		Id:               district.ID.String(),
		Name:             district.Name,
		Code:             district.Code,
		RegionId:         district.RegionId,
		OldCode:          district.OldCode,
		OldName:          district.OldName,
		ActivationDate:   district.ActivationDate,
		DeactivationDate: district.DeactivationDate,
		FlexFinId:        district.FlexFinId,
	}
}
