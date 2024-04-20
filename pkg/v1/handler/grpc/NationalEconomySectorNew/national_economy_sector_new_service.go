package NationalEconomySectorNew

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_new"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type NationalEconomySectorNewService struct {
	nationalEconomySectorNewCase interfaces.NationalEconomySectorNewCaseInterface
	proto.UnimplementedNationalEconomySectorNewServiceServer
}

func NewServer(grpcServer *grpc.Server, nationalEconomySectorNewCase interfaces.NationalEconomySectorNewCaseInterface) *NationalEconomySectorNewService {
	nationalEconomySectorNewGrpc := &NationalEconomySectorNewService{nationalEconomySectorNewCase: nationalEconomySectorNewCase}
	proto.RegisterNationalEconomySectorNewServiceServer(grpcServer, nationalEconomySectorNewGrpc)
	return nationalEconomySectorNewGrpc
}

func (service *NationalEconomySectorNewService) Create(context context.Context, request *proto.CreateNationalEconomySectorNewRequest) (*proto.NationalEconomySectorNewProfileResponse, error) {
	data := service.transformNationalEconomySectorNewRPC(request)
	if data.Code == "" || data.Name == "" {
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
	if id == "" {
		return &proto.NationalEconomySectorNewProfileResponse{}, errors.New("id cannot be blank")
	}
	nationalEconomySectorNew, err := service.nationalEconomySectorNewCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.NationalEconomySectorNewProfileResponse{}, err
	}

	return service.transformNationalEconomySectorNewModel(nationalEconomySectorNew), nil
}

func (service *NationalEconomySectorNewService) transformNationalEconomySectorNewRPC(request *proto.CreateNationalEconomySectorNewRequest) models.NationalEconomySectorNew {
	return models.NationalEconomySectorNew{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		NameUz:           request.GetNameUz(),
		Section:          request.GetSection(),
		SectionCode:      request.GetSectionCode(),
		SectionName:      request.GetSectionName(),
		SectionNameUz:    request.GetSectionNameUz(),
		GroupCode:        request.GetGroupCode(),
		GroupName:        request.GetGroupName(),
		GroupNameUz:      request.GetGroupNameUz(),
		Group1Code:       request.GetGroup1Code(),
		Group1Name:       request.GetGroup1Name(),
		Group1NameUz:     request.GetGroup1NameUz(),
		Group2Code:       request.GetGroup2Code(),
		Group2Name:       request.GetGroup2Name(),
		Group2NameUz:     request.GetGroup2NameUz(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		CBUReferenceKey:  request.GetCbuReferenceKey(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *NationalEconomySectorNewService) transformNationalEconomySectorNewModel(nationalEconomySectorNew models.NationalEconomySectorNew) *proto.NationalEconomySectorNewProfileResponse {
	return &proto.NationalEconomySectorNewProfileResponse{
		Id:               nationalEconomySectorNew.ID.String(),
		Code:             nationalEconomySectorNew.Code,
		Name:             nationalEconomySectorNew.Name,
		NameUz:           nationalEconomySectorNew.NameUz,
		Section:          nationalEconomySectorNew.Section,
		SectionCode:      nationalEconomySectorNew.SectionCode,
		SectionName:      nationalEconomySectorNew.SectionName,
		SectionNameUz:    nationalEconomySectorNew.SectionNameUz,
		GroupCode:        nationalEconomySectorNew.GroupCode,
		GroupName:        nationalEconomySectorNew.GroupName,
		GroupNameUz:      nationalEconomySectorNew.GroupNameUz,
		Group1Code:       nationalEconomySectorNew.Group1Code,
		Group1Name:       nationalEconomySectorNew.Group1Name,
		Group1NameUz:     nationalEconomySectorNew.Group1NameUz,
		Group2Code:       nationalEconomySectorNew.Group2Code,
		Group2Name:       nationalEconomySectorNew.Group2Name,
		Group2NameUz:     nationalEconomySectorNew.Group2NameUz,
		ActivationDate:   nationalEconomySectorNew.ActivationDate,
		DeactivationDate: nationalEconomySectorNew.DeactivationDate,
		CbuReferenceKey:  nationalEconomySectorNew.CBUReferenceKey,
		FlexFinId:        nationalEconomySectorNew.FlexFinId,
	}
}
