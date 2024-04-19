package tests

import (
	"context"
	"fmt"
	proto "github.com/Xurliman/banking-microservice/proto/education_type"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
)

func TestCreateEducationType(t *testing.T) {
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

	educationType := proto.NewEducationTypeServiceClient(conn)

	request := &proto.CreateEducationTypeRequest{
		Code:      "12345",
		Name:      "test",
		FlexFinId: "23",
	}

	res, err := educationType.Create(context.Background(), request)
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
