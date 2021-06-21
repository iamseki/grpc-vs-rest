package grpc_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/iamseki/grpc-vs-rest/grpc/proto"
	"google.golang.org/grpc"
)

func TestTimeFunc(t *testing.T) {
	start := time.Now()

	conn, err := grpc.Dial("localhost:50502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong when connect to grpc server :50502 : %v", err)
	}
	defer conn.Close()

	client := proto.NewCryptoPricingServiceClient(conn)

	client.GetHistoricalPrice(context.TODO(), &proto.EmptyRequest{})

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func BenchmarkGrpcCall(b *testing.B) {
	conn, err := grpc.Dial("localhost:50502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong when connect to grpc server :50502 : %v", err)
	}
	defer conn.Close()

	client := proto.NewCryptoPricingServiceClient(conn)

	for i := 0; i < b.N; i++ {
		client.GetHistoricalPrice(context.TODO(), &proto.EmptyRequest{})
	}
}
