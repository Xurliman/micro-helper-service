package District

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/district"
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
	if data.Code == 0 || data.Name == "" {
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
	if id == 0 {
		return &proto.DistrictProfileResponse{}, errors.New("id cannot be blank")
	}
	district, err := service.districtCase.Get(id)

	if err != nil {
		return &proto.DistrictProfileResponse{}, err
	}

	return service.transformDistrictModel(district), nil
}

func (service *DistrictService) transformDistrictRPC(request *proto.CreateDistrictRequest) models.District {
	return models.District{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		RegionID:         request.GetRegionId(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		OldCode:          request.GetOldCode(),
	}
}

func (service *DistrictService) transformDistrictModel(district models.District) *proto.DistrictProfileResponse {
	return &proto.DistrictProfileResponse{
		Id:               int64(district.ID),
		Name:             district.Name,
		Code:             district.Code,
		RegionId:         district.RegionID,
		ActivationDate:   district.ActivationDate,
		DeactivationDate: district.DeactivationDate,
		OldCode:          district.OldCode,
	}
}
