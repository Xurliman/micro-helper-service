package tests

import (
	"context"
	"fmt"
	proto "github.com/Xurliman/banking-microservice/proto/residency_type"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestCreateResidencyType(t *testing.T) {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
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
		Code:             faker.Word(),
		Name:             faker.Name(),
		ActivationDate:   faker.Date(),
		DeactivationDate: faker.Date(),
		FlexFinId:        strconv.Itoa(rand.Int()),
		NameUz:           faker.Name(),
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
