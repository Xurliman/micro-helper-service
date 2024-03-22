package Region

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/region"
	"google.golang.org/grpc"
)

type RegionService struct {
	regionCase interfaces.RegionCaseInterface
	proto.UnimplementedRegionServiceServer
}

func NewServer(grpcServer *grpc.Server, regionCase interfaces.RegionCaseInterface) {
	regionGrpc := &RegionService{regionCase: regionCase}
	proto.RegisterRegionServiceServer(grpcServer, regionGrpc)
}

func (service *RegionService) Create(context context.Context, request *proto.CreateRegionRequest) (*proto.RegionProfileResponse, error) {
	data := service.transformRegionRPC(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.RegionProfileResponse{}, errors.New("please provide all fields")
	}

	region, err := service.regionCase.Create(data)
	if err != nil {
		return &proto.RegionProfileResponse{}, err
	}
	return service.transformRegionModel(region), nil
}

func (service *RegionService) Get(context context.Context, request *proto.SingleRegionRequest) (*proto.RegionProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.RegionProfileResponse{}, errors.New("id cannot be blank")
	}
	region, err := service.regionCase.Get(id)

	if err != nil {
		return &proto.RegionProfileResponse{}, err
	}

	return service.transformRegionModel(region), nil
}

func (service *RegionService) transformRegionRPC(request *proto.CreateRegionRequest) models.Region {
	return models.Region{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		CountryID:        request.GetCountryId(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		OldName:          request.GetOldName(),
	}
}

func (service *RegionService) transformRegionModel(region models.Region) *proto.RegionProfileResponse {
	return &proto.RegionProfileResponse{
		Id:               int64(region.ID),
		Name:             region.Name,
		Code:             region.Code,
		CountryId:        region.CountryID,
		ActivationDate:   region.ActivationDate,
		DeactivationDate: region.DeactivationDate,
		OldName:          region.OldName,
	}
}
