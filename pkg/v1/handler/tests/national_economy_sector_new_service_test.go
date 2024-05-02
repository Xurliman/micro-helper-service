package tests

import (
	"context"
	"fmt"
	"github.com/Xurliman/banking-microservice/config"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_new"
	"github.com/go-faker/faker/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	rand2 "math/rand/v2"
	"strconv"
	"testing"
)

func TestCreateNationalEconomySectorNew(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", config.GetAppPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	nationalEconomySectorNew := proto.NewNationalEconomySectorNewServiceClient(conn)

	request := &proto.CreateNationalEconomySectorNewRequest{
		Code:             faker.Word(),
		Name:             faker.Name(),
		NameUz:           faker.Name(),
		Section:          rand2.Int32(),
		SectionCode:      faker.Word(),
		SectionName:      faker.Word(),
		SectionNameUz:    faker.Word(),
		GroupCode:        faker.Word(),
		GroupName:        faker.Word(),
		GroupNameUz:      faker.Word(),
		Group1Code:       faker.Word(),
		Group1Name:       faker.Word(),
		Group1NameUz:     faker.Word(),
		Group2Code:       faker.Word(),
		Group2Name:       faker.Word(),
		Group2NameUz:     faker.Word(),
		ActivationDate:   faker.Word(),
		DeactivationDate: faker.Word(),
		CbuReferenceKey:  rand2.Int32(),
		FlexFinId:        strconv.Itoa(rand2.Int()),
	}

	res, err := nationalEconomySectorNew.Create(context.Background(), request)
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
