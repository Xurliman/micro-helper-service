package grpc

import (
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
)

type BankService struct {
	bankCase interfaces.BankCaseInterface
}