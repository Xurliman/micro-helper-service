package DirectOrgan

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/direct_organ"
	"google.golang.org/grpc"
)

type DirectOrganService struct {
	directOrganCase interfaces.DirectOrganCaseInterface
	proto.UnimplementedDirectOrganServiceServer
}

func NewServer(grpcServer *grpc.Server, directOrganCase interfaces.DirectOrganCaseInterface) {
	directOrganGrpc := &DirectOrganService{directOrganCase: directOrganCase}
	proto.RegisterDirectOrganServiceServer(grpcServer, directOrganGrpc)
}

func (service *DirectOrganService) Create(context context.Context, request *proto.CreateDirectOrganRequest) (*proto.DirectOrganProfileResponse, error) {
	data := service.transformDirectOrganRpc(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.DirectOrganProfileResponse{}, errors.New("please provide all fields")
	}

	directOrgan, err := service.directOrganCase.Create(data)
	if err != nil {
		return &proto.DirectOrganProfileResponse{}, err
	}
	return service.transformDirectOrganModel(directOrgan), nil
}

func (service *DirectOrganService) Get(context context.Context, request *proto.SingleDirectOrganRequest) (*proto.DirectOrganProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.DirectOrganProfileResponse{}, errors.New("id cannot be blank")
	}
	directOrgan, err := service.directOrganCase.Get(id)

	if err != nil {
		return &proto.DirectOrganProfileResponse{}, err
	}
	return service.transformDirectOrganModel(directOrgan), nil
}

func (service *DirectOrganService) transformDirectOrganRpc(request *proto.CreateDirectOrganRequest) models.DirectOrgan {
	return models.DirectOrgan{
		Code:      request.GetCode(),
		Name:      request.GetName(),
		ShortName: request.GetShortName(),
		CrudDates: request.GetCrudDates(),
		CBUCode:   request.GetCbuCode(),
	}
}

func (service *DirectOrganService) transformDirectOrganModel(directOrgan models.DirectOrgan) *proto.DirectOrganProfileResponse {
	return &proto.DirectOrganProfileResponse{
		Id:        int64(directOrgan.ID),
		Code:      directOrgan.Code,
		Name:      directOrgan.Name,
		ShortName: directOrgan.ShortName,
		CbuCode:   directOrgan.CBUCode,
		CrudDates: directOrgan.CrudDates,
	}
}
