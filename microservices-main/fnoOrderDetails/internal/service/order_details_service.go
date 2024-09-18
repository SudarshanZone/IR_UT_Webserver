package service

import (
	"context"
	"fmt"

	pb "github.com/krishnakashyap0704/microservices/fnoOrderDetails/generated"
	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/internal/logger"
	"github.com/krishnakashyap0704/microservices/fnoOrderDetails/internal/repository"
)

type OrderDetailsService struct {
	Repo repository.OrderDetailsRepository
	pb.UnimplementedOrderDetailsServiceServer
}

// getOrderDetails
func (s *OrderDetailsService) GetOrderDetails(ctx context.Context, req *pb.OrderDetailsRequest) (*pb.OrderDetailsResponse, error) {
	orderDetails, err := s.Repo.GetOrderDetailsByClaimMatchAccount(req.FOD_CLM_MTCH_ACCNT)
	if err != nil {
		return nil, fmt.Errorf("failed to get order details: %w", err)
	}

	var ordDetails []*pb.OrdDetail

	for _, ord := range orderDetails {
		ordDtls := &pb.OrdDetail{
			ContractDescriptor: ord.ContractDescriptor,
			VTCDate:            ord.VTCDate,
			BuySell:            ord.BuySell,
			Quantity:           ord.Quantity,
			Status:             ord.Status,
			OrderPrice:         ord.OrderPrice,
			Open:               ord.Open,
		}
		ordDetails = append(ordDetails, ordDtls)
	}

	response := &pb.OrderDetailsResponse{
		OrdDetails: ordDetails,
	}
	s.logPositionDetails("OrderDetailsService", response)
	return response, nil

}

func (s *OrderDetailsService) logPositionDetails(c_ServiceName string, ordDetails *pb.OrderDetailsResponse) {
	log := logger.GetLogger()
	log.Infof("%s: Started", c_ServiceName)
	log.Infof("%s: @@@@@@@@@@@@@@@@ FETCH Orders  @@@@@@@@@@@@@@@@:", c_ServiceName)

	for _, ord := range ordDetails.OrdDetails {
		log.Infof("%s: FFO_CONTRACT: %s", c_ServiceName, ord.ContractDescriptor)
		log.Infof("%s: VTC_DATE: %s", c_ServiceName, ord.VTCDate)
		log.Infof("%s: BUY_SELL: %s", c_ServiceName, ord.BuySell)
		log.Infof("%s: QUANTITY: %d", c_ServiceName, ord.Quantity)
		log.Infof("%s: STATUS: %s", c_ServiceName, ord.Status)
		log.Infof("%s: ORDER_PRICE: %f", c_ServiceName, ord.OrderPrice)
		log.Infof("%s: OPEN: %s", c_ServiceName, ord.Open)
		log.Println()
	}
}
