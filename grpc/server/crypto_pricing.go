package server

import (
	"context"

	"github.com/iamseki/grpc-vs-rest/grpc/proto"
	"github.com/iamseki/grpc-vs-rest/helper"
)

type CryptoPricingServer struct {
	proto.UnimplementedCryptoPricingServiceServer
}

func NewCryptoPricingServer() *CryptoPricingServer {
	return &CryptoPricingServer{}
}

func (c *CryptoPricingServer) GetHistoricalPrice(ctx context.Context, _ *proto.EmptyRequest) (*proto.Pricing, error) {
	pricing := &proto.Pricing{}

	helper.GetPricingFromJSONFile(pricing)

	return pricing, nil
}
