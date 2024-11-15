package tests

import (
	"context"
	"fmt"
	"github.com/Xurliman/banking-microservice/config"
	proto "github.com/Xurliman/banking-microservice/proto/country"
	"github.com/go-faker/faker/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	rand2 "math/rand/v2"
	"strconv"
	"testing"
)

func TestCreateCountry(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", config.GetAppPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		Code:             faker.Word(),
		Name:             faker.Name(),
		ShortName:        faker.Name(),
		CurrencyId:       rand2.Int64(),
		CodeAlpha2:       faker.Word(),
		CodeAlpha3:       faker.Word(),
		TerritoryCode:    rand2.Int64(),
		ActivationDate:   faker.Date(),
		DeactivationDate: faker.Date(),
		FlexFinId:        strconv.Itoa(rand.Int()),
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
