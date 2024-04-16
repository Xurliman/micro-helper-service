package TaxOrganisation

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/tax_organisation"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type TaxOrganisationService struct {
	taxOrganisationCase interfaces.TaxOrganisationCaseInterface
	proto.UnimplementedTaxOrganisationServiceServer
}

func NewServer(grpcServer *grpc.Server, taxOrganisationCase interfaces.TaxOrganisationCaseInterface) {
	taxOrganisationGrpc := &TaxOrganisationService{taxOrganisationCase: taxOrganisationCase}
	proto.RegisterTaxOrganisationServiceServer(grpcServer, taxOrganisationGrpc)
}

func (service *TaxOrganisationService) Create(context context.Context, request *proto.CreateTaxOrganisationRequest) (*proto.TaxOrganisationProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.TaxOrganisationProfileResponse{}, errors.New("please provide all fields")
	}

	taxOrganisation, err := service.taxOrganisationCase.Create(data)
	if err != nil {
		return &proto.TaxOrganisationProfileResponse{}, err
	}
	return service.transformClassifierModel(taxOrganisation), nil
}

func (service *TaxOrganisationService) Get(context context.Context, request *proto.SingleTaxOrganisationRequest) (*proto.TaxOrganisationProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.TaxOrganisationProfileResponse{}, errors.New("id cannot be blank")
	}

	taxOrganisation, err := service.taxOrganisationCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.TaxOrganisationProfileResponse{}, err
	}

	return service.transformClassifierModel(taxOrganisation), nil
}

func (service *TaxOrganisationService) transformClassifierRpc(request *proto.CreateTaxOrganisationRequest) models.TaxOrganisation {
	return models.TaxOrganisation{
		Code:      request.GetCode(),
		Name:      request.GetName(),
		FlexFinId: request.GetFlexFinId(),
	}
}

func (service *TaxOrganisationService) transformClassifierModel(taxOrganisation models.TaxOrganisation) *proto.TaxOrganisationProfileResponse {
	return &proto.TaxOrganisationProfileResponse{
		Id:        taxOrganisation.ID.String(),
		Code:      taxOrganisation.Code,
		Name:      taxOrganisation.Name,
		FlexFinId: taxOrganisation.FlexFinId,
	}
}
