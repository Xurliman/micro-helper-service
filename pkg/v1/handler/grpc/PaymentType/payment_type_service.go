package PaymentType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/payment_type"
	"google.golang.org/grpc"
)

type PaymentTypeService struct {
	paymentTypeCase interfaces.PaymentTypeCaseInterface
	proto.UnimplementedPaymentTypeServiceServer
}

func NewServer(grpcServer *grpc.Server, paymentTypeCase interfaces.PaymentTypeCaseInterface) {
	paymentTypeGrpc := &PaymentTypeService{paymentTypeCase: paymentTypeCase}
	proto.RegisterPaymentTypeServiceServer(grpcServer, paymentTypeGrpc)
}

func (service *PaymentTypeService) Create(context context.Context, request *proto.CreatePaymentTypeRequest) (*proto.PaymentTypeProfileResponse, error) {
	data := service.transformPaymentTypeRPC(request)
	if data.Code == 0 || data.Name == "" {
		return &proto.PaymentTypeProfileResponse{}, errors.New("please provide all fields")
	}

	paymentType, err := service.paymentTypeCase.Create(data)
	if err != nil {
		return &proto.PaymentTypeProfileResponse{}, err
	}
	return service.transformPaymentTypeModel(paymentType), nil
}

func (service *PaymentTypeService) Get(context context.Context, request *proto.SinglePaymentTypeRequest) (*proto.PaymentTypeProfileResponse, error) {
	id := request.GetId()
	if id == 0 {
		return &proto.PaymentTypeProfileResponse{}, errors.New("id cannot be blank")
	}
	paymentType, err := service.paymentTypeCase.Get(id)

	if err != nil {
		return &proto.PaymentTypeProfileResponse{}, err
	}

	return service.transformPaymentTypeModel(paymentType), nil
}

func (service *PaymentTypeService) transformPaymentTypeRPC(request *proto.CreatePaymentTypeRequest) models.PaymentType {
	return models.PaymentType{
		Name:             request.GetName(),
		Code:             request.GetCode(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
	}
}

func (service *PaymentTypeService) transformPaymentTypeModel(paymentType models.PaymentType) *proto.PaymentTypeProfileResponse {
	return &proto.PaymentTypeProfileResponse{
		Id:               int64(paymentType.ID),
		Name:             paymentType.Name,
		Code:             paymentType.Code,
		ActivationDate:   paymentType.ActivationDate,
		DeactivationDate: paymentType.DeactivationDate,
		OldCode:          paymentType.OldCode,
		OldName:          paymentType.OldName,
	}
}
