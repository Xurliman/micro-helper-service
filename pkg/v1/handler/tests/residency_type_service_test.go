package tests

import (
	"context"
	"fmt"
	proto "github.com/Xurliman/banking-microservice/proto/residency_type"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
)

func TestCreateResidencyType(t *testing.T) {
	port := os.Getenv("APP_PORT")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	residencyType := proto.NewResidencyTypeServiceClient(conn)

	request := &proto.CreateResidencyTypeRequest{
		Code:             "12345",
		Name:             "test",
		ActivationDate:   "2024-03-20",
		DeactivationDate: "2024-03-20",
		FlexFinId:        "23",
		NameUz:           "tt",
	}

	res, err := residencyType.Create(context.Background(), request)
	if err != nil {
		t.Fatalf("CREATE FAILED: %v", err)
	}

	if res.Code != request.Code {
		t.Errorf("CREATE returned incorrect email, expected %v got %v", request.Code, res.Code)
	}

	if res.Name != request.Name {
		t.Errorf("CREATE returned incorrect Name, expected %v got %v", request.Name, res.Name)
	}

	if res.GetId() == "" {
		t.Error("CREATE function didn't return id as the response")
	}
}
