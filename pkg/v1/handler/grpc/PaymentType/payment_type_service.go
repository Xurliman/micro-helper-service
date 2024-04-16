package PaymentType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/payment_type"
	"github.com/google/uuid"
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
	if data.Code == "" || data.Name == "" {
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
	if id == "" {
		return &proto.PaymentTypeProfileResponse{}, errors.New("id cannot be blank")
	}
	paymentType, err := service.paymentTypeCase.Get(uuid.MustParse(id))

	if err != nil {
		return &proto.PaymentTypeProfileResponse{}, err
	}

	return service.transformPaymentTypeModel(paymentType), nil
}

func (service *PaymentTypeService) transformPaymentTypeRPC(request *proto.CreatePaymentTypeRequest) models.PaymentType {
	return models.PaymentType{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
		NameUz:           request.GetNameUz(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *PaymentTypeService) transformPaymentTypeModel(paymentType models.PaymentType) *proto.PaymentTypeProfileResponse {
	return &proto.PaymentTypeProfileResponse{
		Id:               paymentType.ID.String(),
		Name:             paymentType.Name,
		Code:             paymentType.Code,
		ActivationDate:   paymentType.ActivationDate,
		DeactivationDate: paymentType.DeactivationDate,
		OldCode:          paymentType.OldCode,
		OldName:          paymentType.OldName,
		NameUz:           paymentType.NameUz,
		FlexFinId:        paymentType.FlexFinId,
	}
}
