package tests

import (
	"context"
	proto "github.com/Xurliman/banking-microservice/proto/national_economy_sector_old"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateNationalEconomySector(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	country := proto.NewNationalEconomySectorServiceClient(conn)

	request := &proto.CreateNationalEconomySectorRequest{
		Code:             12345,
		Name:             "test",
		CbuCode:          1,
		CbuGroupCode:     1,
		ActivationDate:   "2024-02-20",
		DeactivationDate: "2024-03-20",
		CbuReferenceKey:  1,
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

	if res.GetId() == 0 {
		t.Error("CREATE function didnot returned id as the response")
	}
}
