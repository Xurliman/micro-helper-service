package ClientTypeClassifier

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/client_type_classifier"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ClientTypeClassifierService struct {
	classifierCase interfaces.ClientTypeClassifierCaseInterface
	proto.UnimplementedClientTypeClassifierServiceServer
}

func NewServer(grpcServer *grpc.Server, classifierCase interfaces.ClientTypeClassifierCaseInterface) *ClientTypeClassifierService {
	classifierGrpc := &ClientTypeClassifierService{classifierCase: classifierCase}
	proto.RegisterClientTypeClassifierServiceServer(grpcServer, classifierGrpc)
	return classifierGrpc
}

func (service *ClientTypeClassifierService) Create(context context.Context, request *proto.CreateClientTypeClassifierRequest) (*proto.ClientTypeClassifierProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.ClientTypeClassifierProfileResponse{}, errors.New("please provide all fields")
	}

	classifier, err := service.classifierCase.Create(data)
	if err != nil {
		return &proto.ClientTypeClassifierProfileResponse{}, err
	}
	return service.transformClassifierModel(classifier), nil
}

func (service *ClientTypeClassifierService) Get(context context.Context, request *proto.SingleClientTypeClassifierRequest) (*proto.ClientTypeClassifierProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.ClientTypeClassifierProfileResponse{}, errors.New("id cannot be blank")
	}

	classifier, err := service.classifierCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.ClientTypeClassifierProfileResponse{}, err
	}

	return service.transformClassifierModel(classifier), nil
}

func (service *ClientTypeClassifierService) transformClassifierRpc(request *proto.CreateClientTypeClassifierRequest) models.ClientTypeClassifier {
	return models.ClientTypeClassifier{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		ShortName:        request.GetShortName(),
		ClientType:       request.GetClientType(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		CBUReferenceKey:  request.GetCbuReferenceKey(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
		NameUz:           request.GetNameUz(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *ClientTypeClassifierService) transformClassifierModel(classifier models.ClientTypeClassifier) *proto.ClientTypeClassifierProfileResponse {
	return &proto.ClientTypeClassifierProfileResponse{
		Id:               classifier.ID.String(),
		Code:             classifier.Code,
		Name:             classifier.Name,
		ShortName:        classifier.ShortName,
		ClientType:       classifier.ClientType,
		ActivationDate:   classifier.ActivationDate,
		DeactivationDate: classifier.DeactivationDate,
		CbuReferenceKey:  classifier.CBUReferenceKey,
		OldCode:          classifier.OldCode,
		OldName:          classifier.OldName,
		NameUz:           classifier.NameUz,
		FlexFinId:        classifier.FlexFinId,
	}
}
