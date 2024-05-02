package tests

import (
	"context"
	"fmt"
	"github.com/Xurliman/banking-microservice/config"
	proto "github.com/Xurliman/banking-microservice/proto/education_type"
	"github.com/go-faker/faker/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"strconv"
	"testing"
)

func TestCreateEducationType(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", config.GetAppPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		Code:      faker.Word(),
		Name:      faker.Name(),
		FlexFinId: strconv.Itoa(rand.Int()),
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
