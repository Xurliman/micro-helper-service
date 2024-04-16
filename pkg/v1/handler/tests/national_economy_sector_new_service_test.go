package tests

import (
	"context"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_new"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateNationalEconomySectorNew(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		Code:             "12345",
		Name:             "test",
		NameUz:           "asjaklsj",
		Section:          1,
		SectionCode:      "2024-02-20",
		SectionName:      "2024-03-20",
		SectionNameUz:    "kljlk",
		GroupCode:        "128",
		GroupName:        "128",
		GroupNameUz:      "128",
		Group1Code:       "128",
		Group1Name:       "128",
		Group1NameUz:     "128",
		Group2Code:       "128",
		Group2Name:       "128",
		Group2NameUz:     "128",
		ActivationDate:   "128",
		DeactivationDate: "128",
		CbuReferenceKey:  128,
		FlexFinId:        "128",
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
		t.Error("CREATE function didnot returned id as the response")
	}
}
