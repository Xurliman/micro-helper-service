package NationalEconomySectorOldOld

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_old"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type NationalEconomySectorOldService struct {
	nationalEconomySectorOldCase interfaces.NationalEconomySectorOldCaseInterface
	proto.UnimplementedNationalEconomySectorOldServiceServer
}

func NewServer(grpcServer *grpc.Server, nationalEconomySectorOldCase interfaces.NationalEconomySectorOldCaseInterface) {
	nationalEconomySectorOldGrpc := &NationalEconomySectorOldService{nationalEconomySectorOldCase: nationalEconomySectorOldCase}
	proto.RegisterNationalEconomySectorOldServiceServer(grpcServer, nationalEconomySectorOldGrpc)
}

func (service *NationalEconomySectorOldService) Create(context context.Context, request *proto.CreateNationalEconomySectorOldRequest) (*proto.NationalEconomySectorOldProfileResponse, error) {
	data := service.transformNationalEconomySectorOldRPC(request)
	if data.Code == "" || data.Name == "" {
		return &proto.NationalEconomySectorOldProfileResponse{}, errors.New("please provide all fields")
	}

	nationalEconomySectorOld, err := service.nationalEconomySectorOldCase.Create(data)
	if err != nil {
		return &proto.NationalEconomySectorOldProfileResponse{}, err
	}
	return service.transformNationalEconomySectorOldModel(nationalEconomySectorOld), nil
}

func (service *NationalEconomySectorOldService) Get(context context.Context, request *proto.SingleNationalEconomySectorOldRequest) (*proto.NationalEconomySectorOldProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.NationalEconomySectorOldProfileResponse{}, errors.New("id cannot be blank")
	}
	nationalEconomySectorOld, err := service.nationalEconomySectorOldCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.NationalEconomySectorOldProfileResponse{}, err
	}

	return service.transformNationalEconomySectorOldModel(nationalEconomySectorOld), nil
}

func (service *NationalEconomySectorOldService) transformNationalEconomySectorOldRPC(request *proto.CreateNationalEconomySectorOldRequest) models.NationalEconomySectorOld {
	return models.NationalEconomySectorOld{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		CBUCode:          request.GetCbuCode(),
		CBUGroupCode:     request.GetCbuGroupCode(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		CBUReferenceKey:  request.GetCbuReferenceKey(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *NationalEconomySectorOldService) transformNationalEconomySectorOldModel(nationalEconomySectorOld models.NationalEconomySectorOld) *proto.NationalEconomySectorOldProfileResponse {
	return &proto.NationalEconomySectorOldProfileResponse{
		Id:               nationalEconomySectorOld.ID.String(),
		Code:             nationalEconomySectorOld.Code,
		Name:             nationalEconomySectorOld.Name,
		CbuCode:          nationalEconomySectorOld.CBUCode,
		CbuGroupCode:     nationalEconomySectorOld.CBUGroupCode,
		ActivationDate:   nationalEconomySectorOld.ActivationDate,
		DeactivationDate: nationalEconomySectorOld.DeactivationDate,
		CbuReferenceKey:  nationalEconomySectorOld.CBUReferenceKey,
		FlexFinId:        nationalEconomySectorOld.FlexFinId,
	}
}
