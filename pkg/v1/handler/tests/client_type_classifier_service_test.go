package tests

import (
	"context"
	proto "github.com/Xurliman/banking-microservice/proto/client_type_classifier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateClientTypeClassifier(t *testing.T) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	clientTypeClassifier := proto.NewClientTypeClassifierServiceClient(conn)

	request := &proto.CreateClientTypeClassifierRequest{
		Code:             "12345",
		Name:             "test",
		ShortName:        "test",
		ClientType:       545,
		ActivationDate:   "2024-03-20",
		DeactivationDate: "2024-03-20",
		CbuReferenceKey:  879,
		OldCode:          474,
		OldName:          "2024-03-20",
		NameUz:           "23",
		FlexFinId:        "23",
	}

	res, err := clientTypeClassifier.Create(context.Background(), request)
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
