package discount

import (
	"context"
	"fmt"
	"github.com/murilosrg/checkout-api/pkg/log"
	pb "github.com/murilosrg/checkout-api/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

type Service interface {
	Get(ctx context.Context, id int) (float32, error)
}

type service struct {
	conn    *grpc.ClientConn
	timeout time.Duration
	logger  log.Logger
}

func NewService(conn *grpc.ClientConn, timeout time.Duration, logger log.Logger) Service {
	return service{
		conn:    conn,
		timeout: timeout,
		logger:  logger,
	}
}

func (s service) Get(ctx context.Context, id int) (float32, error) {
	client := pb.NewDiscountClient(s.conn)
	request := &pb.GetDiscountRequest{ProductID: int32(id)}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(s.timeout))
	defer cancel()

	response, err := client.GetDiscount(ctx, request)
	if err != nil {
		if er, ok := status.FromError(err); ok {
			s.logger.Errorf("grpc: %s, %s", er.Code(), er.Message())
			return 0, fmt.Errorf("grpc: %s, %s", er.Code(), er.Message())
		}

		s.logger.Errorf("server: %s", err.Error())
		return 0, fmt.Errorf("server: %s", err.Error())
	}

	return response.GetPercentage(), nil
}
