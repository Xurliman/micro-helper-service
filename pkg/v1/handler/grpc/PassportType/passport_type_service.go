package PassportType

import (
	"context"
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	proto "github.com/Xurliman/banking-microservice/proto/passport_type"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type PassportTypeService struct {
	passportTypeCase interfaces.PassportTypeCaseInterface
	proto.UnimplementedPassportTypeServiceServer
}

func NewServer(grpcServer *grpc.Server, passportTypeCase interfaces.PassportTypeCaseInterface) *PassportTypeService {
	passportTypeGrpc := &PassportTypeService{passportTypeCase: passportTypeCase}
	proto.RegisterPassportTypeServiceServer(grpcServer, passportTypeGrpc)
	return passportTypeGrpc
}

func (service *PassportTypeService) Create(context context.Context, request *proto.CreatePassportTypeRequest) (*proto.PassportTypeProfileResponse, error) {
	data := service.transformClassifierRpc(request)
	if data.Code == "" || data.Name == "" {
		return &proto.PassportTypeProfileResponse{}, errors.New("please provide all fields")
	}

	passportType, err := service.passportTypeCase.Create(data)
	if err != nil {
		return &proto.PassportTypeProfileResponse{}, err
	}
	return service.transformClassifierModel(passportType), nil
}

func (service *PassportTypeService) Get(context context.Context, request *proto.SinglePassportTypeRequest) (*proto.PassportTypeProfileResponse, error) {
	id := request.GetId()
	if id == "" {
		return &proto.PassportTypeProfileResponse{}, errors.New("id cannot be blank")
	}

	passportType, err := service.passportTypeCase.Get(uuid.MustParse(id))
	if err != nil {
		return &proto.PassportTypeProfileResponse{}, err
	}

	return service.transformClassifierModel(passportType), nil
}

func (service *PassportTypeService) transformClassifierRpc(request *proto.CreatePassportTypeRequest) models.PassportType {
	return models.PassportType{
		Code:             request.GetCode(),
		Name:             request.GetName(),
		OldCode:          request.GetOldCode(),
		OldName:          request.GetOldName(),
		ActivationDate:   request.GetActivationDate(),
		DeactivationDate: request.GetDeactivationDate(),
		FlexFinId:        request.GetFlexFinId(),
	}
}

func (service *PassportTypeService) transformClassifierModel(passportType models.PassportType) *proto.PassportTypeProfileResponse {
	return &proto.PassportTypeProfileResponse{
		Id:               passportType.ID.String(),
		Code:             passportType.Code,
		Name:             passportType.Name,
		OldCode:          passportType.OldCode,
		OldName:          passportType.OldName,
		ActivationDate:   passportType.ActivationDate,
		DeactivationDate: passportType.DeactivationDate,
		FlexFinId:        passportType.FlexFinId,
	}
}
