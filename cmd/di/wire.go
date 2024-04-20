package di

import (
	database "github.com/Xurliman/banking-microservice/internal/db"
	handler "github.com/Xurliman/banking-microservice/pkg/v1/handler/grpc/Bank"
	"github.com/Xurliman/banking-microservice/pkg/v1/repository"
	"github.com/Xurliman/banking-microservice/pkg/v1/usecase"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeServer() grpc.Server {
	wire.Build(
		database.DbConn,
		grpc.NewServer,
		repository.NewBank,
		handler.NewServer,
		usecase.NewBank,
	)
	return grpc.Server{}
}
