package service

import (
	"context"
	"fmt"

	pb "github.com/krishnakashyap0704/microservices/fnoOpenPositions/generated"
	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/logger"
	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/repository"
)

type FnoPositionService struct {
	Repo repository.FnoPositionRepository
	pb.UnimplementedFnoPositionServiceServer
}

func (s *FnoPositionService) GetFNOPosition(ctx context.Context, req *pb.FnoPositionRequest) (*pb.FcpDetailListResponse, error) {
	//log := logger.GetLogger()

	defer logger.LogFunctionExit("GetFNOPosition")
	positions, err := s.Repo.GetPositionsByClaimMatchAccount(req.FCP_CLM_MTCH_ACCNT)
	if err != nil {
		return nil, fmt.Errorf("failed to get positions: %w", err)
	}

	var fcpDetails []*pb.FcpDetail

	for _, pos := range positions {
		fcpDetail := &pb.FcpDetail{
			FFO_CONTRACT: &pos.Contract,
			FFO_PSTN:     &pos.Position,
			FCP_OPNPSTN_QTY:      &pos.TotalQty,
			FFO_AVG_PRC:  &pos.AvgCostPrice,
			FCP_XCHNG_CD:       &pos.ExchangeCode,
			FCP_IBUY_QTY:       &pos.BuyQty,
			FCP_CLM_MTCH_ACCNT: &pos.ClaimMatchAccount,
			FCP_PRDCT_TYP:      &pos.ProductType,
			FCP_INDSTK:         &pos.IndexStock,
			FCP_UNDRLYNG:       &pos.Underlying,
			FCP_EXPRY_DT:       &pos.ExpiryDate,
			FCP_EXER_TYP:       &pos.ExerciseType,
			FCP_OPT_TYP:        &pos.OptionType,
			FCP_STRK_PRC:       &pos.StrikePrice,
			FCP_UCC_CD:         &pos.UCCCode,
			FCP_OPNPSTN_FLW:    &pos.OpenPstnFlow,

			FCP_IBUY_ORD_VAL:        &pos.BuyOrderValue,
			FCP_ISELL_QTY:           &pos.SellQty,
			FCP_ISELL_ORD_VAL:       &pos.SellOrderValue,
			FCP_EXBUY_QTY:           &pos.ExchangeBuyQty,
			FCP_EXBUY_ORD_VAL:       &pos.ExchangeBuyOrderValue,
			FCP_EXSELL_QTY:          &pos.ExchangeSellQty,
			FCP_EXSELL_ORD_VAL:      &pos.ExchangeSellOrderValue,
			FCP_BUY_EXCTD_QTY:       &pos.BuyExpectedQty,
			FCP_SELL_EXCTD_QTY:      &pos.SellExpectedQty,
			FCP_OPNPSTN_VAL:         &pos.OpenPositionValue,
			FCP_EXRC_QTY:            &pos.ExQuantity,
			FCP_ASGND_QTY:           &pos.AssignedQuantity,
			FCP_OPT_PREMIUM:         &pos.OptionPremium,
			FCP_MTM_OPN_VAL:         &pos.MtmOpenValue,
			FCP_IMTM_OPN_VAL:        &pos.ImtmOpenValue,
			FCP_UDLVRY_MRGN:         &pos.UdeliveryMargin,
			FCP_MTM_FLG:             &pos.MtmFlag,
			FCP_TRG_PRC:             &pos.TriggerPrice,
			FCP_MIN_TRG_PRC:         &pos.MinTriggerPrice,
			FCP_DLVRY_MODE_FLAG:     &pos.DeliveryModeFlag,
			FCP_DLVRY_OBLGAMT_BLCKD: &pos.DeliveryObligationBlocked,
			FCP_DLVRY_QTY_BLCKD:     &pos.DeliveryQtyBlocked,
			FCP_MRGN_CHNG_DT:        &pos.MarginChangeDate,
		}

		s.logPositionDetails("FnoPositionService", fcpDetail)
		fcpDetails = append(fcpDetails, fcpDetail)
	}

	return &pb.FcpDetailListResponse{
		FcpDetails: fcpDetails,
	}, nil
}

// logPositionDetails Prints Logs on Console
func (s *FnoPositionService) logPositionDetails(c_ServiceName string, pos *pb.FcpDetail) {
	log := logger.GetLogger()
	log.Infof("%s: Started", c_ServiceName)
	log.Infof("%s: @@@@@@@@@@@@@@@@ FETCH RECORD @@@@@@@@@@@@@@@@:", c_ServiceName)
	log.Infof("%s: :FCP_CLM_MTCH_ACCNT:%s:", c_ServiceName, *pos.FCP_CLM_MTCH_ACCNT)
	log.Infof("%s: :FFO_CONTRACT :%s:", c_ServiceName, *pos.FFO_CONTRACT)
	log.Infof("%s: :FFO_PSTN :%s:", c_ServiceName, *pos.FFO_PSTN)
	log.Infof("%s: :FFO_QTY :%d:", c_ServiceName, *pos.FCP_OPNPSTN_QTY)
	log.Infof("%s: :FFO_AVG_PRC :%f:", c_ServiceName, *pos.FFO_AVG_PRC)
	log.Infof("%s: :FCP_PRDCT_TYP:%s:", c_ServiceName, *pos.FCP_PRDCT_TYP)
	log.Infof("%s: :FCP_INDSTK:%s:", c_ServiceName, *pos.FCP_INDSTK)
	log.Infof("%s: :FCP_UNDRLYNG:%s:", c_ServiceName, *pos.FCP_UNDRLYNG)
	log.Infof("%s: :FCP_EXPRY_DT:%s:", c_ServiceName, *pos.FCP_EXPRY_DT)
	log.Infof("%s: :FCP_EXER_TYP:%s:", c_ServiceName, *pos.FCP_EXER_TYP)
	log.Infof("%s: :FCP_OPT_TYP:%s:", c_ServiceName, *pos.FCP_OPT_TYP)
	log.Infof("%s: :FCP_STRK_PRC:%d:", c_ServiceName, *pos.FCP_STRK_PRC)
	log.Infof("%s: :FCP_OPNPSTN_FLW:%s:", c_ServiceName, *pos.FCP_OPNPSTN_FLW)
	log.Infof("%s: :FFO_QTY:%d:", c_ServiceName, *pos.FCP_OPNPSTN_QTY)
	log.Infof("%s: :BuyOrderValue:%f:", c_ServiceName, *pos.FCP_IBUY_ORD_VAL)
	log.Infof("%s: :SellQty:%d:", c_ServiceName, *pos.FCP_ISELL_QTY)
	log.Infof("%s: :SellOrderValue:%f:", c_ServiceName, *pos.FCP_ISELL_ORD_VAL)
	log.Infof("%s: :ExchangeBuyQty:%d:", c_ServiceName, *pos.FCP_EXBUY_QTY)
	log.Infof("%s: :ExchangeBuyOrderValue:%f:", c_ServiceName, *pos.FCP_EXBUY_ORD_VAL)
	log.Infof("%s: :ExchangeSellQty:%d:", c_ServiceName, *pos.FCP_EXSELL_QTY)
	log.Infof("%s: :ExchangeSellOrderValue:%f:", c_ServiceName, *pos.FCP_EXSELL_ORD_VAL)
	log.Infof("%s: :BuyExpectedQty:%d:", c_ServiceName, *pos.FCP_BUY_EXCTD_QTY)
	log.Infof("%s: :SellExpectedQty:%d:", c_ServiceName, *pos.FCP_SELL_EXCTD_QTY)
	log.Infof("%s: :OpenPositionValue:%f:", c_ServiceName, *pos.FCP_OPNPSTN_VAL)
	log.Infof("%s: :ExQuantity:%d:", c_ServiceName, *pos.FCP_EXRC_QTY)
	log.Infof("%s: :AssignedQuantity:%d:", c_ServiceName, *pos.FCP_ASGND_QTY)
	log.Infof("%s: :OptionPremium:%f:", c_ServiceName, *pos.FCP_OPT_PREMIUM)
	log.Infof("%s: :MtmOpenValue:%f:", c_ServiceName, *pos.FCP_MTM_OPN_VAL)
	log.Infof("%s: :ImtmOpenValue:%f:", c_ServiceName, *pos.FCP_IMTM_OPN_VAL)
	log.Infof("%s: :UdeliveryMargin:%f:", c_ServiceName, *pos.FCP_UDLVRY_MRGN)
	log.Infof("%s: :MtmFlag:%s:", c_ServiceName, *pos.FCP_MTM_FLG)
	log.Infof("%s: :TriggerPrice:%f:", c_ServiceName, *pos.FCP_TRG_PRC)
	log.Infof("%s: :MinTriggerPrice:%f:", c_ServiceName, *pos.FCP_MIN_TRG_PRC)
	log.Infof("%s: :DeliveryModeFlag:%s:", c_ServiceName, *pos.FCP_DLVRY_MODE_FLAG)
	log.Infof("%s: :DeliveryObligationBlocked:%f:", c_ServiceName, *pos.FCP_DLVRY_OBLGAMT_BLCKD)
	log.Infof("%s: :DeliveryQtyBlocked:%d:", c_ServiceName, *pos.FCP_DLVRY_QTY_BLCKD)
	log.Infof("%s: :MarginChangeDate:%s:", c_ServiceName, *pos.FCP_MRGN_CHNG_DT)
	log.Println()
}
