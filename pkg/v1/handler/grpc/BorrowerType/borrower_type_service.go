package BorrowerType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/borrower_type"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type BorrowerTypeService struct {
	borrowerTypeCase interfaces.BorrowerTypeCaseInterface
	proto.UnimplementedBorrowerTypeServiceServer
}

func NewServer(grpcServer *grpc.Server, borrowerTypeCase interfaces.BorrowerTypeCaseInterface) *BorrowerTypeService {
	borrowerTypeGrpc := &BorrowerTypeService{borrowerTypeCase: borrowerTypeCase}
	proto.RegisterBorrowerTypeServiceServer(grpcServer, borrowerTypeGrpc)
	return borrowerTypeGrpc
}

func (service *BorrowerTypeService) Create(context context.Context, request *proto.CreateBorrowerTypeRequest) (*proto.BorrowerTypeProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.BorrowerTypeProfileResponse{}, errors.New("please provide all fields")
	}

	borrowerType, err := service.borrowerTypeCase.Create(data)
	if err != nil {
		return &proto.BorrowerTypeProfileResponse{}, err
	}
	return service.transformClassifierModel(borrowerType), nil
}

func (service *BorrowerTypeService) Get(context context.Context, request *proto.SingleBorrowerTypeRequest) (*proto.BorrowerTypeProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.BorrowerTypeProfileResponse{}, errors.New("id cannot be blank")
	}

	borrowerType, err := service.borrowerTypeCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.BorrowerTypeProfileResponse{}, err
	}

	return service.transformClassifierModel(borrowerType), nil
}

func (service *BorrowerTypeService) transformClassifierRpc(request *proto.CreateBorrowerTypeRequest) models.BorrowerType {
	return models.BorrowerType{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		NameUz:           request.GetNameUz(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *BorrowerTypeService) transformClassifierModel(borrowerType models.BorrowerType) *proto.BorrowerTypeProfileResponse {
	return &proto.BorrowerTypeProfileResponse{
		Id:               borrowerType.ID.String(),
		Code:             borrowerType.Code,
		Name:             borrowerType.Name,
		OldCode:          borrowerType.OldCode,
		OldName:          borrowerType.OldName,
		ActivationDate:   borrowerType.ActivationDate,
		DeactivationDate: borrowerType.DeactivationDate,
		NameUz:           borrowerType.NameUz,
		FlexFinId:        borrowerType.FlexFinId,
	}
}
