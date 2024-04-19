package tests

import (
	"context"
	"fmt"
	proto "github.com/Xurliman/banking-microservice/proto/bank"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"testing"
)

func TestCreateBank(t *testing.T) {
	err := godotenv.Load()
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

	bank := proto.NewBankServiceClient(conn)

	request := &proto.CreateBankRequest{
		Code:             "12345",
		Name:             "test",
		ShortName:        "tt",
		CountryId:        1,
		OpenDate:         "2024-02-20",
		CloseDate:        "2024-03-20",
		ActivationDate:   "2024-02-20",
		DeactivationDate: "2024-03-20",
		FlexFinId:        "2024-03-20",
	}

	res, err := bank.Create(context.Background(), request)
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
