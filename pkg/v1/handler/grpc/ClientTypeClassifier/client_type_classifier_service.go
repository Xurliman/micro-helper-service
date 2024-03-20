package ClientTypeClassifier

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/client_type_classifier"
	"google.golang.org/grpc"
)

type ClientTypeClassifierService struct {
	classifierCase interfaces.ClientTypeClassifierCaseInterface
	proto.UnimplementedClientTypeClassifierServiceServer
}

func NewServer(grpcServer *grpc.Server, classifierCase interfaces.ClientTypeClassifierCaseInterface) {
	classifierGrpc := &ClientTypeClassifierService{classifierCase: classifierCase}
	proto.RegisterClientTypeClassifierServiceServer(grpcServer, classifierGrpc)
}

func (service *ClientTypeClassifierService) Create(context context.Context, request *proto.CreateClientTypeClassifierRequest) (*proto.ClientTypeClassifierProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == 0 || data.Name == "" {
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
	if id == 0 {
		return &proto.ClientTypeClassifierProfileResponse{}, errors.New("id cannot be blank")
	}

	classifier, err := service.classifierCase.Get(id)
	if err != nil {
		return &proto.ClientTypeClassifierProfileResponse{}, err
	}

	return service.transformClassifierModel(classifier), nil
}

func (service *ClientTypeClassifierService) transformClassifierRpc(request *proto.CreateClientTypeClassifierRequest) models.ClientTypeClassifier {
	return models.ClientTypeClassifier{
		Code:             request.Code,
		Name:             request.Name,
		ShortName:        request.ShortName,
		ClientType:       request.ClientType,
		ActivationDate:   request.ActivationDate,
		DeactivationDate: request.DeactivationDate,
		CBUReferenceKey:  request.CbuReferenceKey,
		OldCode:          request.OldCode,
		OldName:          request.OldName,
	}
}

func (service *ClientTypeClassifierService) transformClassifierModel(classifier models.ClientTypeClassifier) *proto.ClientTypeClassifierProfileResponse {
	return &proto.ClientTypeClassifierProfileResponse{
		Code:             classifier.Code,
		Name:             classifier.Name,
		ShortName:        classifier.ShortName,
		ClientType:       classifier.ClientType,
		ActivationDate:   classifier.ActivationDate,
		DeactivationDate: classifier.DeactivationDate,
		CbuReferenceKey:  classifier.CBUReferenceKey,
		OldCode:          classifier.OldCode,
		OldName:          classifier.OldName,
	}
}
