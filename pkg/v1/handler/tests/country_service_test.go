package tests

import (
	"context"
	"fmt"
	proto "github.com/Xurliman/banking-microservice/proto/country"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"testing"
)

func TestCreateCountry(t *testing.T) {
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

	country := proto.NewCountryServiceClient(conn)

	request := &proto.CreateCountryRequest{
		Code:             "12345",
		Name:             "test",
		ShortName:        "tt",
		CurrencyId:       1,
		CodeAlpha2:       "Kolner Str. 200",
		CodeAlpha3:       "2024-02-20",
		TerritoryCode:    2024,
		ActivationDate:   "2024-03-20",
		DeactivationDate: "2024-03-20",
		FlexFinId:        "23",
	}

	res, err := country.Create(context.Background(), request)
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
