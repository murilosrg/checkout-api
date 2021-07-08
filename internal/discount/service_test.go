package discount

import (
	"context"
	"errors"
	"fmt"
	"github.com/murilosrg/checkout-api/pkg/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

type mockDiscountServer struct {
	proto.UnimplementedDiscountServer
}

func (*mockDiscountServer) GetDiscount(_ context.Context, req *proto.GetDiscountRequest) (*proto.GetDiscountResponse, error) {
	if req.GetProductID() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid product %v", req.GetProductID())
	}

	return &proto.GetDiscountResponse{Percentage: 0.05}, nil
}

func TestService_Discount(t *testing.T) {
	tests := []struct {
		name string
		id   int
		res  float32
		err  error
	}{
		{
			"invalid request with invalid product",
			-1,
			0,
			fmt.Errorf("grpc: invalid product %v", -1),
		},
		{
			"valid request",
			1,
			0.05,
			nil,
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := NewService(conn, time.Second, logrus.New()).Get(context.Background(), tt.id)

			if response != tt.res {
				t.Error("error: expected", tt.res, "received", response)
			}

			if err != nil && errors.Is(err, tt.err) {
				t.Error("error: expected", tt.err, "received", err)
			}
		})
	}
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterDiscountServer(server, &mockDiscountServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
