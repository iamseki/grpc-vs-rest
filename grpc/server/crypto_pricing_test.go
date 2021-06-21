package server_test

import (
	"context"
	"testing"

	"github.com/iamseki/grpc-vs-rest/grpc/proto"
	"github.com/iamseki/grpc-vs-rest/grpc/server"
	"github.com/stretchr/testify/assert"
)

func TestCryptoPricingLogic(t *testing.T) {
	sut := server.NewCryptoPricingServer()

	pricing, err := sut.GetHistoricalPrice(context.TODO(), &proto.EmptyRequest{})

	assert.Nil(t, err)
	assert.Equal(t, "ETH", pricing.Symbol)
}
