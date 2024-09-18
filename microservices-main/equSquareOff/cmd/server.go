package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/krishnakashyap0704/microservices/equSquareOff/generated"
	"github.com/krishnakashyap0704/microservices/equSquareOff/internal/database"
	logs "github.com/krishnakashyap0704/microservices/equSquareOff/utils"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50052, "gRPC server port")
)

type server struct {
	pb.UnimplementedSqureoffServer
}

func (s *server) Squareoff_Epb(ctx context.Context, req *pb.Epb_SquareoffRequest) (*pb.Epb_SquareoffResponse, error) {
	if req == nil || req.GetEpb() == nil {

		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}
	logs.InfoLogger.Printf("Squareoff Started")

	epb := req.GetEpb()
	if epb == nil {
		return nil, errors.New("data is nil")
	}

	var pipeID = 99
	ordr_rfrnc_no := GenerateOrderReference(pipeID)

	var param1 database.Clm_clnt_mstr
	if err := database.DB.WithContext(ctx).Raw("SELECT clm_bp_id, clm_clnt_cd FROM clm_clnt_mstr WHERE clm_mtch_accnt = ?", epb.EpbClmMtchAccnt).Scan(&param1).Error; err != nil {
		logs.ErrorLogger.Printf("Failed to fetch client master: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to fetch client master: %v", err)
	}

	var param database.Esp_em_systm_prmtr
	if err := database.DB.WithContext(ctx).Raw("SELECT esp_dp_id, esp_dp_clnt_id FROM esp_em_systm_prmtr WHERE esp_xchng_cd = ?", epb.GetEpbXchngCd()).Scan(&param).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch system parameters: %v", err)
	}

	trd_ref_num, err := generateTradeReference()
	if err != nil {
		logs.ErrorLogger.Printf("Failed to generate trade reference: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to fetch client master: %v", err)
	}
	logs.InfoLogger.Println("Processing trade with reference:", trd_ref_num)

	data_Epb := database.EPBEmPstnBook{
		EpbClmMtchAccnt:      epb.GetEpbClmMtchAccnt(),
		EpbXchngCd:           epb.GetEpbXchngCd(),
		EpbXchngSgmntCd:      epb.GetEpbXchngSgmntCd(),
		EpbXchngSgmntSttlmnt: int(epb.GetEpbXchngSgmntSttlmnt()),
		EpbStckCd:            epb.GetEpbStckCd(),
		EpbOrgnlPstnQty:      int(epb.GetEpbOrgnlPstnQty()),
		EpbRate:              float64(epb.GetEpbRate()),
		EpbOrgnlAmtPayble:    float64(epb.GetEpbOrgnlAmtPayble()),
		EpbOrgnlMrgnAmt:      float64(epb.EpbOrgnlMrgnAmt),
		EpbSellQty:           int(epb.GetEpbSellQty()),
		EpbCvrOrdQty:         int(epb.GetEpbOrgnlPstnQty()),
		EpbNetMrgnAmt:        float64(epb.GetEpbNetMrgnAmt()),
		EpbNetAmtPayble:      0,
		EpbNetPstnQty:        int(epb.GetEpbNetPstnQty()),
		EpbCtdQty:            int(epb.GetEpbCtdQty()),
		EpbPstnStts:          epb.GetEpbPstnStts(),
		EpbLpcCalcStts:       epb.GetEpbLpcCalcStts(),
		EpbSqroffMode:        epb.GetEpbSqroffMode(),
		EpbPstnTrdDt:         convertStringToNullString(epb.GetEpbPstnTrdDt()),
		EpbMtmPrcsFlg:        epb.GetEpbMtmPrcsFlg(),
		EpbLastMdfcnDt:       convertStringToNullString(epb.GetEpbLastMdfcnDt()),
		EpbInsDate:           convertStringToNullString(epb.GetEpbInsDate()),
		EpbCloseDate:         convertStringToNullString(epb.GetEpbCloseDate()),
		EpbSysFailFlg:        epb.GetEpbSysFailFlg(),
		EpbLastPymntDt:       convertStringToNullString(epb.GetEpbLastPymntDt()),
		EpbLpcCalcEndDt:      convertStringToNullString(epb.GetEpbLpcCalcEndDt()),
		EpbMtmCansq:          "N",
		EpbExpiryDt:          convertStringToNullString(epb.GetEpbExpiryDt()),
		EpbMinMrgn:           float64(epb.GetEpbMinMrgn()),
		EpbMrgnDbcrPrcsFlg:   "N",
		EpbDpId:              epb.GetEpbDpId(),
		EpbDpClntId:          epb.GetEpbDpClntId(),
		EpbPledgeStts:        epb.GetEpbPledgeStts(), // if Pledge="N" squareoff will not be happend.--- we have to handle this
		EpbBtstNetMrgnAmt:    float64(epb.GetEpbNetMrgnAmt()),
		EpbBtstMrgnBlckd:     float64(epb.GetEpbBtstMrgnBlckd()),
		EpbBtstMrgnDbcrFlg:   epb.GetEpbBtstMrgnDbcrFlg(),
		EpbBtstSgmntCd:       epb.GetEpbBtstSgmntCd(),
		EpbBtstStlmnt:        int(epb.GetEpbBtstStlmnt()),
		EpbBtstCshBlckd:      float64(epb.GetEpbBtstCshBlckd()),
		EpbBtstSamBlckd:      float64(epb.GetEpbBtstSamBlckd()),
		EpbBtstCalcDt:        convertStringToNullString(epb.GetEpbBtstCalcDt()),
		EpbDbcrCalcDt:        convertStringToNullString(epb.GetEpbDbcrCalcDt()),
		EpbNsdlRefNo:         epb.GetEpbNsdlRefNo(),
	}

	logs.InfoLogger.Printf("EpbClmMtchAccnt: %v", data_Epb.EpbClmMtchAccnt)
	logs.InfoLogger.Printf("EpbXchngCd: %v", data_Epb.EpbXchngCd)
	logs.InfoLogger.Printf("EpbXchngSgmntCd: %v", data_Epb.EpbXchngSgmntCd)
	logs.InfoLogger.Printf("EpbXchngSgmntSttlmnt: %v", data_Epb.EpbXchngSgmntSttlmnt)
	logs.InfoLogger.Printf("EpbStckCd: %v", data_Epb.EpbStckCd)
	logs.InfoLogger.Printf("EpbOrgnlPstnQty: %v", data_Epb.EpbOrgnlPstnQty)
	logs.InfoLogger.Printf("EpbRate: %v", data_Epb.EpbRate)
	logs.InfoLogger.Printf("EpbOrgnlAmtPayble: %v", data_Epb.EpbOrgnlAmtPayble)
	logs.InfoLogger.Printf("EpbOrgnlMrgnAmt: %v", data_Epb.EpbOrgnlMrgnAmt)
	logs.InfoLogger.Printf("EpbSellQty: %v", data_Epb.EpbSellQty)
	logs.InfoLogger.Printf("EpbCvrOrdQty: %v", data_Epb.EpbCvrOrdQty)
	logs.InfoLogger.Printf("EpbNetMrgnAmt: %v", data_Epb.EpbNetMrgnAmt)
	logs.InfoLogger.Printf("EpbNetAmtPayble: %v", data_Epb.EpbNetAmtPayble)
	logs.InfoLogger.Printf("EpbNetPstnQty: %v", data_Epb.EpbNetPstnQty)
	logs.InfoLogger.Printf("EpbCtdQty: %v", data_Epb.EpbCtdQty)
	logs.InfoLogger.Printf("EpbPstnStts: %v", data_Epb.EpbPstnStts)
	logs.InfoLogger.Printf("EpbLpcCalcStts: %v", data_Epb.EpbLpcCalcStts)
	logs.InfoLogger.Printf("EpbSqroffMode: %v", data_Epb.EpbSqroffMode)
	logs.InfoLogger.Printf("EpbPstnTrdDt: %v", data_Epb.EpbPstnTrdDt)
	logs.InfoLogger.Printf("EpbMtmPrcsFlg: %v", data_Epb.EpbMtmPrcsFlg)
	logs.InfoLogger.Printf("EpbLastMdfcnDt: %v", data_Epb.EpbLastMdfcnDt)
	logs.InfoLogger.Printf("EpbInsDate: %v", data_Epb.EpbInsDate)
	logs.InfoLogger.Printf("EpbCloseDate: %v", data_Epb.EpbCloseDate)
	logs.InfoLogger.Printf("EpbSysFailFlg: %v", data_Epb.EpbSysFailFlg)
	logs.InfoLogger.Printf("EpbLastPymntDt: %v", data_Epb.EpbLastPymntDt)
	logs.InfoLogger.Printf("EpbLpcCalcEndDt: %v", data_Epb.EpbLpcCalcEndDt)
	logs.InfoLogger.Printf("EpbMtmCansq: %v", data_Epb.EpbMtmCansq)
	logs.InfoLogger.Printf("EpbExpiryDt: %v", data_Epb.EpbExpiryDt)
	logs.InfoLogger.Printf("EpbMinMrgn: %v", data_Epb.EpbMinMrgn)
	logs.InfoLogger.Printf("EpbMrgnDbcrPrcsFlg: %v", data_Epb.EpbMrgnDbcrPrcsFlg)
	logs.InfoLogger.Printf("EpbDpId: %v", data_Epb.EpbDpId)
	logs.InfoLogger.Printf("EpbDpClntId: %v", data_Epb.EpbDpClntId)
	logs.InfoLogger.Printf("EpbPledgeStts: %v", data_Epb.EpbPledgeStts)
	logs.InfoLogger.Printf("EpbBtstNetMrgnAmt: %v", data_Epb.EpbBtstNetMrgnAmt)
	logs.InfoLogger.Printf("EpbBtstMrgnBlckd: %v", data_Epb.EpbBtstMrgnBlckd)
	logs.InfoLogger.Printf("EpbBtstMrgnDbcrFlg: %v", data_Epb.EpbBtstMrgnDbcrFlg)
	logs.InfoLogger.Printf("EpbBtstSgmntCd: %v", data_Epb.EpbBtstSgmntCd)
	logs.InfoLogger.Printf("EpbBtstStlmnt: %v", data_Epb.EpbBtstStlmnt)
	logs.InfoLogger.Printf("EpbBtstCshBlckd: %v", data_Epb.EpbBtstCshBlckd)
	logs.InfoLogger.Printf("EpbBtstSamBlckd: %v", data_Epb.EpbBtstSamBlckd)
	logs.InfoLogger.Printf("EpbBtstCalcDt: %v", data_Epb.EpbBtstCalcDt)
	logs.InfoLogger.Printf("EpbDbcrCalcDt: %v", data_Epb.EpbDbcrCalcDt)
	logs.InfoLogger.Printf("EpbNsdlRefNo: %v", data_Epb.EpbNsdlRefNo)

	logs.DebugLogger.Printf("EPB Details: %+v\n", data_Epb)
	logs.InfoLogger.Println("EPB saved successfully:", data_Epb.EpbClmMtchAccnt)

	data_Trd := database.Trd_Trd_Dtls{
		Trd_trd_ref:             trd_ref_num,                          //trd.GetTrdTrdRef(),
		Trd_clm_mtch_accnt:      epb.GetEpbClmMtchAccnt(),             //Y  user id number
		Trd_xchng_cd:            epb.GetEpbXchngCd(),                  //Y
		Trd_stck_cd:             epb.GetEpbStckCd(),                   //Y
		Trd_xchng_sgmnt_cd:      epb.GetEpbXchngSgmntCd(),             //Y
		Trd_xchng_sgmnt_sttlmnt: int64(epb.GetEpbXchngSgmntSttlmnt()), //Y
		Trd_ordr_rfrnc:          ordr_rfrnc_no,                        //Generated
		Trd_trd_dt:              today,                                //Generated
		Trd_trnsctn_typ:         "BFM",                                //trd.GetTrdTrnsctnTyp(),
		Trd_trd_flw:             "S",                                  //trd.GetTrdTrdFlw(),            //B/S
		Trd_exctd_qty:           int64(epb.GetEpbOrgnlPstnQty()),      //doubt
		Trd_exctd_rt:            0,                                    //trd.GetTrdExctdRt(),           //dont know Trade prc
		Trd_trd_vl:              0,                                    //trd.GetTrdTrdVl(),             //dont know
		Trd_brkrg_vl:            0,                                    //trd.GetTrdBrkrgVl(),           //0
		Trd_net_vl:              0,                                    //trd.GetTrdNetVl(),
		Trd_amt_blckd:           0,                                    //trd.GetTrdAmtBlckd(),
		Trd_xchng_rfrnc:         500032803,                            //trd.GetTrdXchngRfrnc(),
		Trd_upld_mtch_flg:       " ",                                  //trd.GetTrdUpldMtchFlg(),   //empty
		Trd_buy_oblg_flg:        " ",                                  //trd.GetTrdBuyOblgFlg(),    //empty
		Trd_prtfl_flg:           nil,                                  //trd.GetTrdPrtflFlg(),      //{Y}
		Trd_usr_id:              req.GetId(),                          //Y
		Trd_stmp_duty:           0,                                    //trd.GetTrdStmpDuty(),      //empty
		Trd_trnx_chrg:           0,                                    //trd.GetTrdTrnxChrg(),      //empty
		Trd_sebi_chrg_val:       0,                                    //trd.GetTrdSebiChrgVal(),   //empty
		Trd_stt:                 0,                                    //trd.GetTrdStt(),           //empty
		Trd_cntrct_nmbr:         "",                                   //trd.GetTrdCntrctNmbr(),    //empty
		Trd_brkrg_flg:           " ",                                  //trd.GetTrdBrkrgFlg(),      //empty
		Trd_srvc_tax:            0,                                    //trd.GetTrdSrvcTax(),       //empty
		Trd_brkrg_typ:           " ",                                  //trd.GetTrdBrkrgTyp(),      //empty
		Trd_cgst_amt:            0,                                    //trd.GetTrdCgstAmt(),       //empty
		Trd_sgst_amt:            0,                                    //trd.GetTrdSgstAmt(),       //empty
		Trd_ugst_amt:            0,                                    //trd.GetTrdUgstAmt(),       //empty
		Trd_igst_amt:            0,                                    //trd.GetTrdIgstAmt(),       //empty
		Trd_atm_upfront_amt:     0,                                    //trd.GetTrdAtmUpfrontAmt(), //empty
		Trd_fixed_brkg:          0,                                    //trd.GetTrdFixedBrkg(),     //empty
		Trd_variable_brkg:       0,                                    //trd.GetTrdVariableBrkg(),  //empty
		Trd_brkrg_mdl:           " ",                                  //trd.GetTrdBrkrgFlg(),      //empty
		Trd_csh_wthld_amt:       nil,                                  //trd.GetTrdCshWthldAmt(),   //empty
		Trd_ins_dt:              time.Now(),
	}
	logs.InfoLogger.Println("Creating trade details with the following values:")
	logs.InfoLogger.Printf("Trd_trd_ref: %v", data_Trd.Trd_trd_ref)
	logs.InfoLogger.Printf("Trd_clm_mtch_accnt: %v", data_Trd.Trd_clm_mtch_accnt)
	logs.InfoLogger.Printf("Trd_xchng_cd: %v", data_Trd.Trd_xchng_cd)
	logs.InfoLogger.Printf("Trd_stck_cd: %v", data_Trd.Trd_stck_cd)
	logs.InfoLogger.Printf("Trd_xchng_sgmnt_cd: %v", data_Trd.Trd_xchng_sgmnt_cd)
	logs.InfoLogger.Printf("Trd_xchng_sgmnt_sttlmnt: %v", data_Trd.Trd_xchng_sgmnt_sttlmnt)
	logs.InfoLogger.Printf("Trd_ordr_rfrnc: %v", data_Trd.Trd_ordr_rfrnc)
	logs.InfoLogger.Printf("Trd_trd_dt: %v", data_Trd.Trd_trd_dt)
	logs.InfoLogger.Printf("Trd_trnsctn_typ: %v", data_Trd.Trd_trnsctn_typ)
	logs.InfoLogger.Printf("Trd_trd_flw: %v", data_Trd.Trd_trd_flw)
	logs.InfoLogger.Printf("Trd_exctd_qty: %v", data_Trd.Trd_exctd_qty)
	logs.InfoLogger.Printf("Trd_exctd_rt: %v", data_Trd.Trd_exctd_rt)
	logs.InfoLogger.Printf("Trd_trd_vl: %v", data_Trd.Trd_trd_vl)
	logs.InfoLogger.Printf("Trd_brkrg_vl: %v", data_Trd.Trd_brkrg_vl)
	logs.InfoLogger.Printf("Trd_net_vl: %v", data_Trd.Trd_net_vl)
	logs.InfoLogger.Printf("Trd_amt_blckd: %v", data_Trd.Trd_amt_blckd)
	logs.InfoLogger.Printf("Trd_xchng_rfrnc: %v", data_Trd.Trd_xchng_rfrnc)
	logs.InfoLogger.Printf("Trd_upld_mtch_flg: %v", data_Trd.Trd_upld_mtch_flg)
	logs.InfoLogger.Printf("Trd_buy_oblg_flg: %v", data_Trd.Trd_buy_oblg_flg)
	logs.InfoLogger.Printf("Trd_prtfl_flg: %v", data_Trd.Trd_prtfl_flg)
	logs.InfoLogger.Printf("Trd_usr_id: %v", data_Trd.Trd_usr_id)
	logs.InfoLogger.Printf("Trd_stmp_duty: %v", data_Trd.Trd_stmp_duty)
	logs.InfoLogger.Printf("Trd_trnx_chrg: %v", data_Trd.Trd_trnx_chrg)
	logs.InfoLogger.Printf("Trd_sebi_chrg_val: %v", data_Trd.Trd_sebi_chrg_val)
	logs.InfoLogger.Printf("Trd_stt: %v", data_Trd.Trd_stt)
	logs.InfoLogger.Printf("Trd_cntrct_nmbr: %v", data_Trd.Trd_cntrct_nmbr)
	logs.InfoLogger.Printf("Trd_brkrg_flg: %v", data_Trd.Trd_brkrg_flg)
	logs.InfoLogger.Printf("Trd_srvc_tax: %v", data_Trd.Trd_srvc_tax)
	logs.InfoLogger.Printf("Trd_brkrg_typ: %v", data_Trd.Trd_brkrg_typ)
	logs.InfoLogger.Printf("Trd_cgst_amt: %v", data_Trd.Trd_cgst_amt)
	logs.InfoLogger.Printf("Trd_sgst_amt: %v", data_Trd.Trd_sgst_amt)
	logs.InfoLogger.Printf("Trd_ugst_amt: %v", data_Trd.Trd_ugst_amt)
	logs.InfoLogger.Printf("Trd_igst_amt: %v", data_Trd.Trd_igst_amt)
	logs.InfoLogger.Printf("Trd_atm_upfront_amt: %v", data_Trd.Trd_atm_upfront_amt)
	logs.InfoLogger.Printf("Trd_fixed_brkg: %v", data_Trd.Trd_fixed_brkg)
	logs.InfoLogger.Printf("Trd_variable_brkg: %v", data_Trd.Trd_variable_brkg)
	logs.InfoLogger.Printf("Trd_brkrg_mdl: %v", data_Trd.Trd_brkrg_mdl)
	logs.InfoLogger.Printf("Trd_csh_wthld_amt: %v", data_Trd.Trd_csh_wthld_amt)
	logs.InfoLogger.Printf("Trd_ins_dt: %v", data_Trd.Trd_ins_dt)

	logs.DebugLogger.Printf("Trade Details: %+v\n", data_Trd)
	logs.InfoLogger.Println("Trade saved successfully:", data_Trd.Trd_trd_ref)

	data_Ord := database.Ord_Ordr_Dtls{
		Ord_clm_mtch_accnt:      epb.GetEpbClmMtchAccnt(),           //Y  or userid number
		Ord_ordr_rfrnc:          ordr_rfrnc_no,                      //generated
		Ord_xchng_cd:            epb.GetEpbXchngCd(),                //Y
		Ord_stck_cd:             epb.GetEpbStckCd(),                 //Y
		Ord_xchng_sgmnt_cd:      epb.GetEpbXchngSgmntCd(),           //Y
		Ord_xchng_sgmnt_sttlmnt: int(epb.GetEpbXchngSgmntSttlmnt()), //Y
		Ord_ordr_dt:             today,                              //generated
		Ord_ordr_flw:            "S",                                //ord.GetOrdOrdrFlw(),                //  S/B
		Ord_prdct_typ:           "E",                                //ord.GetOrdPrdctTyp(),               //M
		Ord_ordr_qty:            int(epb.GetEpbOrgnlPstnQty()),      //Y
		Ord_lmt_mrkt_flg:        "M",                                //ord.GetOrdLmtMrktFlg(),             //L
		Ord_lmt_rt:              0,                                  //float64(epb.GetEpbRate()),          //Y
		Ord_dsclsd_qty:          0,                                  //ord.GetOrdDsclsdQty(),              // 0
		Ord_stp_lss_tgr:         0,                                  //ord.GetOrdStpLssTgr(),              // 0
		Ord_ordr_stts:           "R",                                //ord.GetOrdOrdrStts(),               // Q
		Ord_trd_dt:              today,                              //exg_nxt_trd_dt,------------------D
		Ord_sub_brkr_tag:        nil,                                //ord.GetOrdSubBrkrTag(),       // empty
		Ord_mdfctn_cntr:         1,                                  //int(ord.GetOrdMdfctnCntr()),  // 1
		Ord_ack_nmbr:            nil,                                //ord.GetOrdAckNmbr(),          //empty
		Ord_xchng_ack_old:       nil,                                //int(ord.GetOrdXchngAckOld()), //  empty
		Ord_exctd_qty:           0,                                  //int(ord.GetOrdExctdQty()),    // 0
		Ord_amt_blckd:           0,                                  //int(ord.GetOrdAmtBlckd()),    //0
		Ord_brkrg_val:           nil,                                //int(ord.GetOrdBrkrgVal()),    //empty
		Ord_dp_id:               param.Esp_dp_id,                    //esp_dp_id
		Ord_dp_clnt_id:          param.Esp_dp_clnt_id,               //esp_dp_clnt_id
		Ord_phy_qty:             nil,                                //int(ord.GetOrdPhyQty()),      //empty
		Ord_isin_nmbr:           nil,                                //ord.GetOrdIsinNmbr(),         //empty
		Ord_nd_flg:              nil,                                //ord.GetOrdNdFlg(),            //empty
		Ord_msc_char:            nil,
		Ord_msc_varchar:         nil,                                //ord.GetOrdMscVarchar(),       //empty
		Ord_msc_int:             nil,                                //ord.GetOrdMscInt(),           //empty
		Ord_plcd_stts:           "P",                                //ord.GetOrdPlcdStts(),         //P
		Ord_qty_blckd:           0,                                  //int(ord.GetOrdQtyBlckd()),    //0 doubt
		Ord_mrgn_prcntg:         nil,                                //doubt
		Ord_ipo_flg:             nil,                                //ord.GetOrdIpoFlg(),                //empty
		Ord_lss_amt_blckd:       0,                                  //int(ord.GetOrdLssAmtBlckd()),      //0 doubt
		Ord_lss_qty:             0,                                  //int(ord.GetOrdLssQty()),           //0 doubt
		Ord_mtm_flg:             nil,                                //ord.GetOrdMtmFlg(),                //empty
		Ord_sq_flg:              "N",                                //ord.GetOrdSqFlg(),                 //N
		Ord_schm_id:             nil,                                //ord.GetOrdSchmId(),                //empty
		Ord_pipe_id:             "99",                               //ord.GetOrdPipeId(),                //N1 or 99 doubt
		Ord_prtctn_rt:           0,                                  //ord.GetOrdPrtctnRt(),              //0
		Ord_sl_trg_flg:          "N",                                //ord.GetOrdSlTrgFlg(),              //N
		Ord_xchng_usr_id:        nil,                                //dont know
		Ord_btst_sttlmnt_nmbr:   nil,                                //int(ord.GetOrdBtstSttlmntNmbr()),  //empty
		Ord_btst_sgmnt_cd:       nil,                                //ord.GetOrdBtstSgmntCd(),           //empty
		Ord_channel:             "WEB",                              //ord.GetOrdChannel(),               //WEB
		Ord_bp_id:               param1.Clm_bp_id,                   //"",             //ord_bp_id,    //ord.GetOrdBpId(),                  //Rajvik doutb
		Ord_sltp_ordr_rfrnc:     nil,                                //ord.GetOrdSltpOrdrRfrnc(),         // empty
		Ord_ctcl_id:             "111111111111",                     //ord.GetOrdCtclId(),                //  111111111111 doubt
		Ord_usr_id:              req.GetId(),                        //y
		Ord_cnt_id:              nil,                                //int(ord.GetOrdCntId()),            //empty
		Ord_em_settlmnt_nmbr:    int(epb.GetEpbXchngSgmntSttlmnt()), //int(ord.GetOrdEmSettlmntNmbr()),   //empty
		Ord_mrgn_sqroff_mode:    "M",                                //ord.GetOrdMrgnSqroffMode(),        //S
		Ord_cncl_qty:            nil,                                //int(ord.GetOrdCnclQty()),          // empty
		Ord_ordr_typ:            "T",                                //ord.GetOrdOrdrTyp(),               //T
		Ord_valid_dt:            nil,                                //ord.GetOrdValidDt(),               //empty
		Ord_cal_flg:             "N",                                //ord.GetOrdCalFlg(),                //N
		Ord_xchng_ack:           nil,                                //ord.GetOrdXchngAck(),              //*
		Ord_em_rollovr_flg:      "N",                                //ord.GetOrdEmRollovrFlg(),          //empty
		Ord_trd_val:             nil,                                //int(ord.GetOrdTrdVal()),           //empty
		Ord_trd_cntrct_nmbr:     nil,                                //ord.GetOrdTrdCntrctNmbr(),         //empty
		Ord_avg_exctd_rt:        nil,                                //ord.GetOrdAvgExctdRt(),            //empty
		Ord_prc_imp_flg:         nil,                                //ord.GetOrdPrcImpFlg(),             //empty
		Ord_mbc_flg:             nil,                                //ord.GetOrdMbcFlg(),                //empty
		Ord_trl_amt:             nil,                                //ord.GetOrdTrlAmt(),                //empty
		Ord_lmt_offst:           nil,                                //ord.GetOrdLmtOffst(),              //empty
		Ord_source_flg:          nil,                                //ord.GetOrdSourceFlg(),             //empty
		Ord_pan_no:              "FWEPP2025A",                       //ord.GetOrdPanNo(),                 //don't know
		Ord_atm_payout_stts:     nil,                                //ord.GetOrdAtmPayoutStts(),         //empty
		Ord_esp_cd:              nil,                                //ord.GetOrdEspCd(),                 //empty
		Ord_remarks:             nil,                                //ord.GetOrdRemarks(),               //empty
		Ord_wthld_amt_stts:      nil,                                //ord.GetOrdWthldAmtStts(),          //empty
		Ord_pstn_xchng_cd:       nil,                                //epb.GetEpbXchngCd(),               //doubt
		Ord_interop_ord_flg:     "N",                                //ord.GetOrdIntEropOrdFlg(),         //N doubt
		Ord_settlement_period:   0,                                  //int(ord.GetOrdSettlementPeriod()), //empty
		Ord_algo_id:             nil,                                //ord.GetOrdAlgoId(),                //empty
		Ord_bundle_name:         nil,                                //ord.GetOrdBundleName(),            //empty
		Ord_prt_flg:             nil,                                //ord.GetOrdPrtFlg(),                //empty
		Ord_src_tag:             nil,                                //ord.GetOrdSrcTag(),                //empty
		Ord_rls_amt:             nil,                                //int(ord.GetOrdRlsAmt()),           //empty
		Ord_rls_date:            nil,                                //ord.GetOrdRlsDate(),               //empty
		Ord_mtf_unplg_sqroff:    "N",                                //ord.GetOrdMtfUnplgSqroff(),        //empty
		Ord_n_ordr_qty:          nil,                                //int(ord.GetOrdNOrdrQty()),         //empty
		Ord_ack_date:            &today,                             //ord.GetOrdAckDate(),               //empty
		Ord_last_activity_ref:   nil,                                //ord.GetOrdLastActivityRef(),       //empty
		Ord_clm_clnt_cd:         param1.Clm_clnt_cd,                 //ord.GetOrdClmClntCd(),
	}

	logs.InfoLogger.Println("Creating order details with the following values:")
	logs.InfoLogger.Printf("Ord_clm_mtch_accnt: %v", data_Ord.Ord_clm_mtch_accnt)
	logs.InfoLogger.Printf("Ord_ordr_rfrnc: %v", data_Ord.Ord_ordr_rfrnc)
	logs.InfoLogger.Printf("Ord_xchng_cd: %v", data_Ord.Ord_xchng_cd)
	logs.InfoLogger.Printf("Ord_stck_cd: %v", data_Ord.Ord_stck_cd)
	logs.InfoLogger.Printf("Ord_xchng_sgmnt_cd: %v", data_Ord.Ord_xchng_sgmnt_cd)
	logs.InfoLogger.Printf("Ord_xchng_sgmnt_sttlmnt: %v", data_Ord.Ord_xchng_sgmnt_sttlmnt)
	logs.InfoLogger.Printf("Ord_ordr_dt: %v", data_Ord.Ord_ordr_dt)
	logs.InfoLogger.Printf("Ord_ordr_flw: %v", data_Ord.Ord_ordr_flw)
	logs.InfoLogger.Printf("Ord_prdct_typ: %v", data_Ord.Ord_prdct_typ)
	logs.InfoLogger.Printf("Ord_ordr_qty: %v", data_Ord.Ord_ordr_qty)
	logs.InfoLogger.Printf("Ord_lmt_mrkt_flg: %v", data_Ord.Ord_lmt_mrkt_flg)
	logs.InfoLogger.Printf("Ord_lmt_rt: %v", data_Ord.Ord_lmt_rt)
	logs.InfoLogger.Printf("Ord_dsclsd_qty: %v", data_Ord.Ord_dsclsd_qty)
	logs.InfoLogger.Printf("Ord_stp_lss_tgr: %v", data_Ord.Ord_stp_lss_tgr)
	logs.InfoLogger.Printf("Ord_ordr_stts: %v", data_Ord.Ord_ordr_stts)
	logs.InfoLogger.Printf("Ord_trd_dt: %v", data_Ord.Ord_trd_dt)
	logs.InfoLogger.Printf("Ord_sub_brkr_tag: %v", data_Ord.Ord_sub_brkr_tag)
	logs.InfoLogger.Printf("Ord_mdfctn_cntr: %v", data_Ord.Ord_mdfctn_cntr)
	logs.InfoLogger.Printf("Ord_ack_nmbr: %v", data_Ord.Ord_ack_nmbr)
	logs.InfoLogger.Printf("Ord_xchng_ack_old: %v", data_Ord.Ord_xchng_ack_old)
	logs.InfoLogger.Printf("Ord_exctd_qty: %v", data_Ord.Ord_exctd_qty)
	logs.InfoLogger.Printf("Ord_amt_blckd: %v", data_Ord.Ord_amt_blckd)
	logs.InfoLogger.Printf("Ord_brkrg_val: %v", data_Ord.Ord_brkrg_val)
	logs.InfoLogger.Printf("Ord_dp_id: %v", data_Ord.Ord_dp_id)
	logs.InfoLogger.Printf("Ord_dp_clnt_id: %v", data_Ord.Ord_dp_clnt_id)
	logs.InfoLogger.Printf("Ord_phy_qty: %v", data_Ord.Ord_phy_qty)
	logs.InfoLogger.Printf("Ord_isin_nmbr: %v", data_Ord.Ord_isin_nmbr)
	logs.InfoLogger.Printf("Ord_nd_flg: %v", data_Ord.Ord_nd_flg)
	logs.InfoLogger.Printf("Ord_msc_char: %v", data_Ord.Ord_msc_char)
	logs.InfoLogger.Printf("Ord_msc_varchar: %v", data_Ord.Ord_msc_varchar)
	logs.InfoLogger.Printf("Ord_msc_int: %v", data_Ord.Ord_msc_int)
	logs.InfoLogger.Printf("Ord_plcd_stts: %v", data_Ord.Ord_plcd_stts)
	logs.InfoLogger.Printf("Ord_qty_blckd: %v", data_Ord.Ord_qty_blckd)
	logs.InfoLogger.Printf("Ord_mrgn_prcntg: %v", data_Ord.Ord_mrgn_prcntg)
	logs.InfoLogger.Printf("Ord_ipo_flg: %v", data_Ord.Ord_ipo_flg)
	logs.InfoLogger.Printf("Ord_lss_amt_blckd: %v", data_Ord.Ord_lss_amt_blckd)
	logs.InfoLogger.Printf("Ord_lss_qty: %v", data_Ord.Ord_lss_qty)
	logs.InfoLogger.Printf("Ord_mtm_flg: %v", data_Ord.Ord_mtm_flg)
	logs.InfoLogger.Printf("Ord_sq_flg: %v", data_Ord.Ord_sq_flg)
	logs.InfoLogger.Printf("Ord_schm_id: %v", data_Ord.Ord_schm_id)
	logs.InfoLogger.Printf("Ord_pipe_id: %v", data_Ord.Ord_pipe_id)
	logs.InfoLogger.Printf("Ord_prtctn_rt: %v", data_Ord.Ord_prtctn_rt)
	logs.InfoLogger.Printf("Ord_sl_trg_flg: %v", data_Ord.Ord_sl_trg_flg)
	logs.InfoLogger.Printf("Ord_xchng_usr_id: %v", data_Ord.Ord_xchng_usr_id)
	logs.InfoLogger.Printf("Ord_btst_sttlmnt_nmbr: %v", data_Ord.Ord_btst_sttlmnt_nmbr)
	logs.InfoLogger.Printf("Ord_btst_sgmnt_cd: %v", data_Ord.Ord_btst_sgmnt_cd)
	logs.InfoLogger.Printf("Ord_channel: %v", data_Ord.Ord_channel)
	logs.InfoLogger.Printf("Ord_bp_id: %v", data_Ord.Ord_bp_id)
	logs.InfoLogger.Printf("Ord_sltp_ordr_rfrnc: %v", data_Ord.Ord_sltp_ordr_rfrnc)
	logs.InfoLogger.Printf("Ord_ctcl_id: %v", data_Ord.Ord_ctcl_id)
	logs.InfoLogger.Printf("Ord_usr_id: %v", data_Ord.Ord_usr_id)
	logs.InfoLogger.Printf("Ord_cnt_id: %v", data_Ord.Ord_cnt_id)
	logs.InfoLogger.Printf("Ord_em_settlmnt_nmbr: %v", data_Ord.Ord_em_settlmnt_nmbr)
	logs.InfoLogger.Printf("Ord_mrgn_sqroff_mode: %v", data_Ord.Ord_mrgn_sqroff_mode)
	logs.InfoLogger.Printf("Ord_cncl_qty: %v", data_Ord.Ord_cncl_qty)
	logs.InfoLogger.Printf("Ord_ordr_typ: %v", data_Ord.Ord_ordr_typ)
	logs.InfoLogger.Printf("Ord_valid_dt: %v", data_Ord.Ord_valid_dt)
	logs.InfoLogger.Printf("Ord_cal_flg: %v", data_Ord.Ord_cal_flg)
	logs.InfoLogger.Printf("Ord_xchng_ack: %v", data_Ord.Ord_xchng_ack)
	logs.InfoLogger.Printf("Ord_em_rollovr_flg: %v", data_Ord.Ord_em_rollovr_flg)
	logs.InfoLogger.Printf("Ord_trd_val: %v", data_Ord.Ord_trd_val)
	logs.InfoLogger.Printf("Ord_trd_cntrct_nmbr: %v", data_Ord.Ord_trd_cntrct_nmbr)
	logs.InfoLogger.Printf("Ord_avg_exctd_rt: %v", data_Ord.Ord_avg_exctd_rt)
	logs.InfoLogger.Printf("Ord_prc_imp_flg: %v", data_Ord.Ord_prc_imp_flg)
	logs.InfoLogger.Printf("Ord_mbc_flg: %v", data_Ord.Ord_mbc_flg)
	logs.InfoLogger.Printf("Ord_trl_amt: %v", data_Ord.Ord_trl_amt)
	logs.InfoLogger.Printf("Ord_lmt_offst: %v", data_Ord.Ord_lmt_offst)
	logs.InfoLogger.Printf("Ord_source_flg: %v", data_Ord.Ord_source_flg)
	logs.InfoLogger.Printf("Ord_pan_no: %v", data_Ord.Ord_pan_no)
	logs.InfoLogger.Printf("Ord_atm_payout_stts: %v", data_Ord.Ord_atm_payout_stts)
	logs.InfoLogger.Printf("Ord_esp_cd: %v", data_Ord.Ord_esp_cd)
	logs.InfoLogger.Printf("Ord_remarks: %v", data_Ord.Ord_remarks)
	logs.InfoLogger.Printf("Ord_wthld_amt_stts: %v", data_Ord.Ord_wthld_amt_stts)
	logs.InfoLogger.Printf("Ord_pstn_xchng_cd: %v", data_Ord.Ord_pstn_xchng_cd)
	logs.InfoLogger.Printf("Ord_interop_ord_flg: %v", data_Ord.Ord_interop_ord_flg)
	logs.InfoLogger.Printf("Ord_settlement_period: %v", data_Ord.Ord_settlement_period)
	logs.InfoLogger.Printf("Ord_algo_id: %v", data_Ord.Ord_algo_id)
	logs.InfoLogger.Printf("Ord_bundle_name: %v", data_Ord.Ord_bundle_name)
	logs.InfoLogger.Printf("Ord_prt_flg: %v", data_Ord.Ord_prt_flg)
	logs.InfoLogger.Printf("Ord_src_tag: %v", data_Ord.Ord_src_tag)
	logs.InfoLogger.Printf("Ord_rls_amt: %v", data_Ord.Ord_rls_amt)
	logs.InfoLogger.Printf("Ord_rls_date: %v", data_Ord.Ord_rls_date)
	logs.InfoLogger.Printf("Ord_mtf_unplg_sqroff: %v", data_Ord.Ord_mtf_unplg_sqroff)
	logs.InfoLogger.Printf("Ord_n_ordr_qty: %v", data_Ord.Ord_n_ordr_qty)
	logs.InfoLogger.Printf("Ord_ack_date: %v", data_Ord.Ord_ack_date)
	logs.InfoLogger.Printf("Ord_last_activity_ref: %v", data_Ord.Ord_last_activity_ref)
	logs.InfoLogger.Printf("Ord_clm_clnt_cd: %v", data_Ord.Ord_clm_clnt_cd)

	logs.DebugLogger.Printf("Order Details: %+v\n", data_Ord)
	logs.InfoLogger.Println("Order saved successfully:", data_Ord.Ord_ordr_rfrnc)

	if err := database.DB.Create(&data_Epb).Error; err != nil {
		database.DB.Rollback()
		logs.ErrorLogger.Printf("Failed to create epb details: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create epb details: %v", err)
	}

	if err := database.DB.Create(&data_Trd).Error; err != nil {
		database.DB.Rollback()
		logs.ErrorLogger.Printf("Failed to create trade details: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create trade details: %v", err)
	}

	if err := database.DB.Create(&data_Ord).Error; err != nil {
		database.DB.Rollback()
		logs.ErrorLogger.Printf("Failed to create order details: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to create order details: %v", err)
	}

	// if err := database.DB.Commit().Error; err != nil {
	// 	logs.ErrorLogger.Printf("Failed to commit transaction: %v", err)
	// 	return nil, status.Errorf(codes.Internal, "Failed to commit transaction: %v", err)
	// }

	logs.InfoLogger.Printf("Trade and Order saved successfully: Trade Ref: %v, Order Ref: %v", trd_ref_num, ordr_rfrnc_no)

	return &pb.Epb_SquareoffResponse{Success: true}, nil
}

// func (s *server) Squareoff_Otp(ctx context.Context, req *pb.Opt_SquareoffRequest) (*pb.Otp_SquareoffResponse, error) {

// 	if req == nil {
// 		return nil, status.Errorf(codes.InvalidArgument, "Request is nil")
// 	}
// 	otp := req.GetOtp()
// 	if otp == nil {
// 		return nil, status.Errorf(codes.InvalidArgument, "Data is nil")
// 	}

// 	var pipeID = 99
// 	ordr_rfrnc_no := GenerateOrderReference(pipeID)
// 	var param1 database.Clm_clnt_mstr
// 	if err := database.DB.WithContext(ctx).Raw("SELECT clm_bp_id, clm_clnt_cd FROM clm_clnt_mstr WHERE clm_mtch_accnt = ?", req.Id).Scan(&param1).Error; err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to fetch client master: %v", err)
// 	}
// 	var param database.Esp_em_systm_prmtr
// 	if err := database.DB.WithContext(ctx).Raw("SELECT esp_dp_id, esp_dp_clnt_id FROM esp_em_systm_prmtr WHERE esp_xchng_cd = ?", otp.GetOtpXchngCd()).Scan(&param).Error; err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to fetch system parameters: %v", err)
// 	}

// 	trd_ref_num, err := generateTradeReference()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to fetch client master: %v", err)
// 	}

// 	var flw string
// 	if otp.OtpFlw == "B" {
// 		flw = "S"
// 	} else {
// 		flw = "B"
// 	}

// 	// data_Otp := database.Otp_trd_pstns{
// 	// 	Otp_clm_mtch_acct:       otp.GetOtpClmMtchAcct(), // Assuming this returns a string
// 	// 	Otp_xchng_cd:            getStringFromNull(sql.NullString{String: otp.GetOtpXchngCd(), Valid: true}),
// 	// 	Otp_xchng_sgmnt_cd:      getStringFromNull(sql.NullString{String: otp.GetOtpXchngSgmntCd(), Valid: true}),
// 	// 	Otp_xchng_sgmnt_sttlmnt: int32(otp.OtpXchngSgmntSttlmnt),
// 	// 	Otp_stck_cd:             getStringFromNull(sql.NullString{String: otp.GetOtpStckCd(), Valid: true}),
// 	// 	Otp_flw:                 "S", // Make sure this is one character
// 	// 	Otp_qty:                 otp.GetOtpQty(),
// 	// 	Otp_cnvrt_dlvry_qty:     87,// otp.GetOtpQty()
// 	// 	Otp_cvrd_qty:            8,//  o
// 	// 	Otp_rt:                  float64(otp.GetOtpRt()),
// 	// 	Otp_mrgn_amt:            900,
// 	// 	Otp_trd_val:             0.44,
// 	// 	Otp_rmrks:               "",  //getStringFromNull(sql.NullString{String: epb.GetEpbRmrks(), Valid: true}),
// 	// 	Otp_xfer_mrgn_stts:      "T", // Make sure this is one character
// 	// 	Otp_sell_opn_prccsd:     "8", // Make sure this is one character
// 	// 	Otp_buy_opn_prccsd:      "1", // Make sure this is one character
// 	// 	Otp_mrgn_sqroff_mode:    "M", // Make sure this is one character
// 	// 	Otp_em_trdsplt_prcs_flg: "I", // Make sure this is one character
// 	// 	Otp_mtm_flg:             "",  // Make sure this is one character
// 	// 	Otp_mtm_cansq:           "",  // Make sure this is one character
// 	// 	Otp_eos_can:             "",  // Make sure this is one character
// 	// 	Otp_trgr_prc:            6,   //epb.GetEpbTrgrPrc(),
// 	// 	Otp_16_trgr_prc:         6,   //epb.GetEpb16TrgrPrc(),
// 	// 	Otp_min_mrgn:            6,   //epb.GetEpbMinMrgn(),
// 	// }

// 	data_Trd := database.Trd_Trd_Dtls{
// 		Trd_trd_ref:             trd_ref_num,                   //trd.GetTrdTrdRef(),
// 		Trd_clm_mtch_accnt:      otp.GetOtpClmMtchAcct(),       //Y  user id number
// 		Trd_xchng_cd:            otp.GetOtpXchngCd(),           //Y
// 		Trd_stck_cd:             otp.GetOtpStckCd(),            //Y
// 		Trd_xchng_sgmnt_cd:      otp.GetOtpXchngSgmntCd(),      //Y
// 		Trd_xchng_sgmnt_sttlmnt: otp.GetOtpXchngSgmntSttlmnt(), //Y
// 		Trd_ordr_rfrnc:          ordr_rfrnc_no,                 //Generated
// 		Trd_trd_dt:              today,                         //Generated
// 		Trd_trnsctn_typ:         "BFM",                         //trd.GetTrdTrnsctnTyp(),        // BFM //doubt
// 		Trd_trd_flw:             flw,                           //trd.GetTrdTrdFlw(),            //B/S
// 		Trd_exctd_qty:           otp.GetOtpQty(),               //doubt
// 		Trd_exctd_rt:            0,                             //trd.GetTrdExctdRt(),           //dont know Trade prc
// 		Trd_trd_vl:              0,                             //trd.GetTrdTrdVl(),             //dont know
// 		Trd_brkrg_vl:            0,                             //trd.GetTrdBrkrgVl(),           //0
// 		Trd_net_vl:              1,                             //trd.GetTrdNetVl(),
// 		Trd_amt_blckd:           0,                             //trd.GetTrdAmtBlckd(),
// 		Trd_xchng_rfrnc:         500032803,                     //trd.GetTrdXchngRfrnc(),
// 		Trd_upld_mtch_flg:       " ",                           //trd.GetTrdUpldMtchFlg(),   //empty
// 		Trd_buy_oblg_flg:        " ",                           //trd.GetTrdBuyOblgFlg(),    //empty
// 		Trd_prtfl_flg:           nil,                           //trd.GetTrdPrtflFlg(),      //{Y}
// 		Trd_usr_id:              req.GetId(),                   //Y
// 		Trd_stmp_duty:           0,                             //trd.GetTrdStmpDuty(),      //empty
// 		Trd_trnx_chrg:           0,                             //trd.GetTrdTrnxChrg(),      //empty
// 		Trd_sebi_chrg_val:       0,                             //trd.GetTrdSebiChrgVal(),   //empty
// 		Trd_stt:                 0,                             //trd.GetTrdStt(),           //empty
// 		Trd_cntrct_nmbr:         " ",                           //trd.GetTrdCntrctNmbr(),    //empty
// 		Trd_brkrg_flg:           " ",                           //trd.GetTrdBrkrgFlg(),      //empty
// 		Trd_srvc_tax:            0,                             //trd.GetTrdSrvcTax(),       //empty
// 		Trd_brkrg_typ:           " ",                           //trd.GetTrdBrkrgTyp(),      //empty
// 		Trd_cgst_amt:            0,                             //trd.GetTrdCgstAmt(),       //empty
// 		Trd_sgst_amt:            0,                             //trd.GetTrdSgstAmt(),       //empty
// 		Trd_ugst_amt:            0,                             //trd.GetTrdUgstAmt(),       //empty
// 		Trd_igst_amt:            0,                             //trd.GetTrdIgstAmt(),       //empty
// 		Trd_atm_upfront_amt:     0,                             //trd.GetTrdAtmUpfrontAmt(), //empty
// 		Trd_fixed_brkg:          0,                             //trd.GetTrdFixedBrkg(),     //empty
// 		Trd_variable_brkg:       0,                             //trd.GetTrdVariableBrkg(),  //empty
// 		Trd_brkrg_mdl:           " ",                           //trd.GetTrdBrkrgFlg(),      //empty
// 		Trd_csh_wthld_amt:       nil,                           //trd.GetTrdCshWthldAmt(),   //empty
// 		Trd_ins_dt:              time.Now(),
// 	}
// 	logs.InfoLogger.Printf("Trade Reference Number: %v", data_Trd.Trd_trd_ref)
// 	logs.InfoLogger.Printf("Claim Match Account: %v", data_Trd.Trd_clm_mtch_accnt)
// 	logs.InfoLogger.Printf("Exchange Code: %v", data_Trd.Trd_xchng_cd)
// 	logs.InfoLogger.Printf("Stock Code: %v", data_Trd.Trd_stck_cd)
// 	logs.InfoLogger.Printf("Exchange Segment Code: %v", data_Trd.Trd_xchng_sgmnt_cd)
// 	logs.InfoLogger.Printf("Exchange Segment Settlement: %v", data_Trd.Trd_xchng_sgmnt_sttlmnt)
// 	logs.InfoLogger.Printf("Order Reference Number: %v", data_Trd.Trd_ordr_rfrnc)
// 	logs.InfoLogger.Printf("Trade Date: %v", data_Trd.Trd_trd_dt)
// 	logs.InfoLogger.Printf("Transaction Type: %v", data_Trd.Trd_trnsctn_typ)
// 	logs.InfoLogger.Printf("Trade Flow: %v", data_Trd.Trd_trd_flw)
// 	logs.InfoLogger.Printf("Executed Quantity: %v", data_Trd.Trd_exctd_qty)
// 	logs.InfoLogger.Printf("Executed Rate: %v", data_Trd.Trd_exctd_rt)
// 	logs.InfoLogger.Printf("Trade Value: %v", data_Trd.Trd_trd_vl)
// 	logs.InfoLogger.Printf("Brokerage Value: %v", data_Trd.Trd_brkrg_vl)
// 	logs.InfoLogger.Printf("Net Value: %v", data_Trd.Trd_net_vl)
// 	logs.InfoLogger.Printf("Amount Blocked: %v", data_Trd.Trd_amt_blckd)
// 	logs.InfoLogger.Printf("Exchange Reference: %v", data_Trd.Trd_xchng_rfrnc)
// 	logs.InfoLogger.Printf("Upload Match Flag: %v", data_Trd.Trd_upld_mtch_flg)
// 	logs.InfoLogger.Printf("Buy Obligation Flag: %v", data_Trd.Trd_buy_oblg_flg)
// 	logs.InfoLogger.Printf("Portfolio Flag: %v", data_Trd.Trd_prtfl_flg)
// 	logs.InfoLogger.Printf("User ID: %v", data_Trd.Trd_usr_id)
// 	logs.InfoLogger.Printf("Stamp Duty: %v", data_Trd.Trd_stmp_duty)
// 	logs.InfoLogger.Printf("Transaction Charge: %v", data_Trd.Trd_trnx_chrg)
// 	logs.InfoLogger.Printf("SEBI Charge Value: %v", data_Trd.Trd_sebi_chrg_val)
// 	logs.InfoLogger.Printf("STT: %v", data_Trd.Trd_stt)
// 	logs.InfoLogger.Printf("Contract Number: %v", data_Trd.Trd_cntrct_nmbr)
// 	logs.InfoLogger.Printf("Brokerage Flag: %v", data_Trd.Trd_brkrg_flg)
// 	logs.InfoLogger.Printf("Service Tax: %v", data_Trd.Trd_srvc_tax)
// 	logs.InfoLogger.Printf("Brokerage Type: %v", data_Trd.Trd_brkrg_typ)
// 	logs.InfoLogger.Printf("CGST Amount: %v", data_Trd.Trd_cgst_amt)
// 	logs.InfoLogger.Printf("SGST Amount: %v", data_Trd.Trd_sgst_amt)
// 	logs.InfoLogger.Printf("UGST Amount: %v", data_Trd.Trd_ugst_amt)
// 	logs.InfoLogger.Printf("IGST Amount: %v", data_Trd.Trd_igst_amt)
// 	logs.InfoLogger.Printf("ATM Upfront Amount: %v", data_Trd.Trd_atm_upfront_amt)
// 	logs.InfoLogger.Printf("Fixed Brokerage: %v", data_Trd.Trd_fixed_brkg)
// 	logs.InfoLogger.Printf("Variable Brokerage: %v", data_Trd.Trd_variable_brkg)
// 	logs.InfoLogger.Printf("Brokerage Model: %v", data_Trd.Trd_brkrg_mdl)
// 	logs.InfoLogger.Printf("Cash Withheld Amount: %v", data_Trd.Trd_csh_wthld_amt)
// 	logs.InfoLogger.Printf("Inserted Date: %v", data_Trd.Trd_ins_dt)

// 	data_Ord := database.Ord_Ordr_Dtls{
// 		Ord_clm_mtch_accnt:      otp.GetOtpClmMtchAcct(),            //Y  or userid number
// 		Ord_ordr_rfrnc:          ordr_rfrnc_no,                      //generated
// 		Ord_xchng_cd:            otp.GetOtpXchngCd(),                //Y
// 		Ord_stck_cd:             otp.GetOtpStckCd(),                 //Y
// 		Ord_xchng_sgmnt_cd:      otp.GetOtpXchngSgmntCd(),           //Y
// 		Ord_xchng_sgmnt_sttlmnt: int(otp.GetOtpXchngSgmntSttlmnt()), //Y
// 		Ord_ordr_dt:             today,                              //generated
// 		Ord_ordr_flw:            flw,                                //ord.GetOrdOrdrFlw(),                //  S/B
// 		Ord_prdct_typ:           "M",                                //ord.GetOrdPrdctTyp(),               //M
// 		Ord_ordr_qty:            int(otp.GetOtpQty()),               //Y
// 		Ord_lmt_mrkt_flg:        "L",                                //ord.GetOrdLmtMrktFlg(),             //L
// 		Ord_lmt_rt:              0,                                  //float64(otp.GetOtpRt()),            //Y
// 		Ord_dsclsd_qty:          0,                                  //ord.GetOrdDsclsdQty(),              // 0
// 		Ord_stp_lss_tgr:         0,                                  //ord.GetOrdStpLssTgr(),              // 0
// 		Ord_ordr_stts:           "R",                                //ord.GetOrdOrdrStts(),               // Q
// 		Ord_trd_dt:              today,                              //epb.GetEpbPstnTrdDt(),------------------D
// 		Ord_sub_brkr_tag:        nil,                                //ord.GetOrdSubBrkrTag(),       // empty
// 		Ord_mdfctn_cntr:         1,                                  //int(ord.GetOrdMdfctnCntr()),  // 1
// 		Ord_ack_nmbr:            nil,                                //ord.GetOrdAckNmbr(),          //empty
// 		Ord_xchng_ack_old:       nil,                                //int(ord.GetOrdXchngAckOld()), //  empty
// 		Ord_exctd_qty:           0,                                  //int(ord.GetOrdExctdQty()),    // 0
// 		Ord_amt_blckd:           0,                                  //int(ord.GetOrdAmtBlckd()),    //0
// 		Ord_brkrg_val:           nil,                                //int(ord.GetOrdBrkrgVal()),    //empty
// 		Ord_dp_id:               string(param.Esp_dp_id),            //ord.GetOrdDpId(),             //trd table
// 		Ord_dp_clnt_id:          string(param.Esp_dp_clnt_id),       //ord.GetOrdClmClntCd(),        //trd table
// 		Ord_phy_qty:             nil,                                //int(ord.GetOrdPhyQty()),      //empty
// 		Ord_isin_nmbr:           nil,                                //ord.GetOrdIsinNmbr(),         //empty
// 		Ord_nd_flg:              nil,                                //ord.GetOrdNdFlg(),            //empty
// 		Ord_msc_char:            "N",                                //ord.GetOrdMscChar(),          //empty
// 		Ord_msc_varchar:         nil,                                //ord.GetOrdMscVarchar(),       //empty
// 		Ord_msc_int:             nil,                                //ord.GetOrdMscInt(),           //empty
// 		Ord_plcd_stts:           "N",                                //ord.GetOrdPlcdStts(),         //P
// 		Ord_qty_blckd:           0,                                  //int(ord.GetOrdQtyBlckd()),    //0 doubt
// 		Ord_mrgn_prcntg:         nil,                                //int(ord.GetOrdMrgnPrcntg()),
// 		Ord_ipo_flg:             nil,                                //ord.GetOrdIpoFlg(),                //empty
// 		Ord_lss_amt_blckd:       0,                                  //int(ord.GetOrdLssAmtBlckd()),      //0 doubt
// 		Ord_lss_qty:             0,                                  //int(ord.GetOrdLssQty()),           //0 doubt
// 		Ord_mtm_flg:             " ",                                //ord.GetOrdMtmFlg(),                //empty
// 		Ord_sq_flg:              "N",                                //ord.GetOrdSqFlg(),                 //N
// 		Ord_schm_id:             nil,                                //ord.GetOrdSchmId(),                //empty
// 		Ord_pipe_id:             "99",                               //ord.GetOrdPipeId(),                //N1 or 99 doubt
// 		Ord_prtctn_rt:           0,                                  //ord.GetOrdPrtctnRt(),              //0
// 		Ord_sl_trg_flg:          "N",                                //ord.GetOrdSlTrgFlg(),              //N
// 		Ord_xchng_usr_id:        nil,                                //int(ord.GetOrdXchngUsrId()),       //empty
// 		Ord_btst_sttlmnt_nmbr:   nil,                                //int(ord.GetOrdBtstSttlmntNmbr()),  //empty
// 		Ord_btst_sgmnt_cd:       nil,                                //ord.GetOrdBtstSgmntCd(),           //empty
// 		Ord_channel:             "WEB",                              //ord.GetOrdChannel(),               //WEB
// 		Ord_bp_id:               param1.Clm_bp_id,                   //ord.GetOrdBpId(),                  //Rajvik doutb
// 		Ord_sltp_ordr_rfrnc:     nil,                                //ord.GetOrdSltpOrdrRfrnc(),         // empty
// 		Ord_ctcl_id:             "111111111111",                     //ord.GetOrdCtclId(),                //  111111111111 doubt
// 		Ord_usr_id:              req.GetId(),                        //y
// 		Ord_cnt_id:              nil,                                //int(ord.GetOrdCntId()),            //empty
// 		Ord_em_settlmnt_nmbr:    nil,                                //int(ord.GetOrdEmSettlmntNmbr()),   //empty
// 		Ord_mrgn_sqroff_mode:    "S",                                //ord.GetOrdMrgnSqroffMode(),        //S
// 		Ord_cncl_qty:            nil,                                //int(ord.GetOrdCnclQty()),          // empty
// 		Ord_ordr_typ:            "T",                                //ord.GetOrdOrdrTyp(),               //T
// 		Ord_valid_dt:            nil,                                //ord.GetOrdValidDt(),               //empty
// 		Ord_cal_flg:             nil,                                //ord.GetOrdCalFlg(),                //N
// 		Ord_xchng_ack:           "*",                                //ord.GetOrdXchngAck(),              //*
// 		Ord_em_rollovr_flg:      nil,                                //ord.GetOrdEmRollovrFlg(),          //empty
// 		Ord_trd_val:             nil,                                //int(ord.GetOrdTrdVal()),           //empty
// 		Ord_trd_cntrct_nmbr:     nil,                                //ord.GetOrdTrdCntrctNmbr(),         //empty
// 		Ord_avg_exctd_rt:        nil,                                //ord.GetOrdAvgExctdRt(),            //empty
// 		Ord_prc_imp_flg:         nil,                                //ord.GetOrdPrcImpFlg(),             //empty
// 		Ord_mbc_flg:             nil,                                //ord.GetOrdMbcFlg(),                //empty
// 		Ord_trl_amt:             nil,                                //ord.GetOrdTrlAmt(),                //empty
// 		Ord_lmt_offst:           nil,                                //ord.GetOrdLmtOffst(),              //empty
// 		Ord_source_flg:          nil,                                //ord.GetOrdSourceFlg(),             //empty
// 		Ord_pan_no:              "FWEPP2025A",                       //ord.GetOrdPanNo(),                 //don't know
// 		Ord_atm_payout_stts:     nil,                                //ord.GetOrdAtmPayoutStts(),         //empty
// 		Ord_esp_cd:              nil,                                //ord.GetOrdEspCd(),                 //empty
// 		Ord_remarks:             nil,                                //ord.GetOrdRemarks(),               //empty
// 		Ord_wthld_amt_stts:      nil,                                //ord.GetOrdWthldAmtStts(),          //empty
// 		Ord_pstn_xchng_cd:       nil,                                //epb.GetEpbXchngCd(),               //doubt
// 		Ord_interop_ord_flg:     "N",                                //ord.GetOrdIntEropOrdFlg(),         //N doubt
// 		Ord_settlement_period:   nil,                                //int(ord.GetOrdSettlementPeriod()), //empty
// 		Ord_algo_id:             nil,                                //ord.GetOrdAlgoId(),                //empty
// 		Ord_bundle_name:         nil,                                //ord.GetOrdBundleName(),            //empty
// 		Ord_prt_flg:             nil,                                //ord.GetOrdPrtFlg(),                //empty
// 		Ord_src_tag:             nil,                                //ord.GetOrdSrcTag(),                //empty
// 		Ord_rls_amt:             nil,                                //int(ord.GetOrdRlsAmt()),           //empty
// 		Ord_rls_date:            &today,                             //ord.GetOrdRlsDate(),               //empty
// 		Ord_mtf_unplg_sqroff:    "N",                                //ord.GetOrdMtfUnplgSqroff(),        //empty
// 		Ord_n_ordr_qty:          nil,                                //int(ord.GetOrdNOrdrQty()),         //empty
// 		Ord_ack_date:            nil,                                //ord.GetOrdAckDate(),               //empty
// 		Ord_last_activity_ref:   nil,                                //ord.GetOrdLastActivityRef(),       //empty
// 		Ord_clm_clnt_cd:         param1.Clm_clnt_cd,                 //ord.GetOrdClmClntCd(),
// 	}

// 	logs.InfoLogger.Println("Creating order details with the following values:")
// 	logs.InfoLogger.Printf("Ord_clm_mtch_accnt: %v", data_Ord.Ord_clm_mtch_accnt)
// 	logs.InfoLogger.Printf("Ord_ordr_rfrnc: %v", data_Ord.Ord_ordr_rfrnc)
// 	logs.InfoLogger.Printf("Ord_xchng_cd: %v", data_Ord.Ord_xchng_cd)
// 	logs.InfoLogger.Printf("Ord_stck_cd: %v", data_Ord.Ord_stck_cd)
// 	logs.InfoLogger.Printf("Ord_xchng_sgmnt_cd: %v", data_Ord.Ord_xchng_sgmnt_cd)
// 	logs.InfoLogger.Printf("Ord_xchng_sgmnt_sttlmnt: %v", data_Ord.Ord_xchng_sgmnt_sttlmnt)
// 	logs.InfoLogger.Printf("Ord_ordr_dt: %v", data_Ord.Ord_ordr_dt)
// 	logs.InfoLogger.Printf("Ord_ordr_flw: %v", data_Ord.Ord_ordr_flw)
// 	logs.InfoLogger.Printf("Ord_prdct_typ: %v", data_Ord.Ord_prdct_typ)
// 	logs.InfoLogger.Printf("Ord_ordr_qty: %v", data_Ord.Ord_ordr_qty)
// 	logs.InfoLogger.Printf("Ord_lmt_mrkt_flg: %v", data_Ord.Ord_lmt_mrkt_flg)
// 	logs.InfoLogger.Printf("Ord_lmt_rt: %v", data_Ord.Ord_lmt_rt)
// 	logs.InfoLogger.Printf("Ord_dsclsd_qty: %v", data_Ord.Ord_dsclsd_qty)
// 	logs.InfoLogger.Printf("Ord_stp_lss_tgr: %v", data_Ord.Ord_stp_lss_tgr)
// 	logs.InfoLogger.Printf("Ord_ordr_stts: %v", data_Ord.Ord_ordr_stts)
// 	logs.InfoLogger.Printf("Ord_trd_dt: %v", data_Ord.Ord_trd_dt)
// 	logs.InfoLogger.Printf("Ord_sub_brkr_tag: %v", data_Ord.Ord_sub_brkr_tag)
// 	logs.InfoLogger.Printf("Ord_mdfctn_cntr: %v", data_Ord.Ord_mdfctn_cntr)
// 	logs.InfoLogger.Printf("Ord_ack_nmbr: %v", data_Ord.Ord_ack_nmbr)
// 	logs.InfoLogger.Printf("Ord_xchng_ack_old: %v", data_Ord.Ord_xchng_ack_old)
// 	logs.InfoLogger.Printf("Ord_exctd_qty: %v", data_Ord.Ord_exctd_qty)
// 	logs.InfoLogger.Printf("Ord_amt_blckd: %v", data_Ord.Ord_amt_blckd)
// 	logs.InfoLogger.Printf("Ord_brkrg_val: %v", data_Ord.Ord_brkrg_val)
// 	logs.InfoLogger.Printf("Ord_dp_id: %v", data_Ord.Ord_dp_id)
// 	logs.InfoLogger.Printf("Ord_dp_clnt_id: %v", data_Ord.Ord_dp_clnt_id)
// 	logs.InfoLogger.Printf("Ord_phy_qty: %v", data_Ord.Ord_phy_qty)
// 	logs.InfoLogger.Printf("Ord_isin_nmbr: %v", data_Ord.Ord_isin_nmbr)
// 	logs.InfoLogger.Printf("Ord_nd_flg: %v", data_Ord.Ord_nd_flg)
// 	logs.InfoLogger.Printf("Ord_msc_char: %v", data_Ord.Ord_msc_char)
// 	logs.InfoLogger.Printf("Ord_msc_varchar: %v", data_Ord.Ord_msc_varchar)
// 	logs.InfoLogger.Printf("Ord_msc_int: %v", data_Ord.Ord_msc_int)
// 	logs.InfoLogger.Printf("Ord_plcd_stts: %v", data_Ord.Ord_plcd_stts)
// 	logs.InfoLogger.Printf("Ord_qty_blckd: %v", data_Ord.Ord_qty_blckd)
// 	logs.InfoLogger.Printf("Ord_mrgn_prcntg: %v", data_Ord.Ord_mrgn_prcntg)
// 	logs.InfoLogger.Printf("Ord_ipo_flg: %v", data_Ord.Ord_ipo_flg)
// 	logs.InfoLogger.Printf("Ord_lss_amt_blckd: %v", data_Ord.Ord_lss_amt_blckd)
// 	logs.InfoLogger.Printf("Ord_lss_qty: %v", data_Ord.Ord_lss_qty)
// 	logs.InfoLogger.Printf("Ord_mtm_flg: %v", data_Ord.Ord_mtm_flg)
// 	logs.InfoLogger.Printf("Ord_sq_flg: %v", data_Ord.Ord_sq_flg)
// 	logs.InfoLogger.Printf("Ord_schm_id: %v", data_Ord.Ord_schm_id)
// 	logs.InfoLogger.Printf("Ord_pipe_id: %v", data_Ord.Ord_pipe_id)
// 	logs.InfoLogger.Printf("Ord_prtctn_rt: %v", data_Ord.Ord_prtctn_rt)
// 	logs.InfoLogger.Printf("Ord_sl_trg_flg: %v", data_Ord.Ord_sl_trg_flg)
// 	logs.InfoLogger.Printf("Ord_xchng_usr_id: %v", data_Ord.Ord_xchng_usr_id)
// 	logs.InfoLogger.Printf("Ord_btst_sttlmnt_nmbr: %v", data_Ord.Ord_btst_sttlmnt_nmbr)
// 	logs.InfoLogger.Printf("Ord_btst_sgmnt_cd: %v", data_Ord.Ord_btst_sgmnt_cd)
// 	logs.InfoLogger.Printf("Ord_channel: %v", data_Ord.Ord_channel)
// 	logs.InfoLogger.Printf("Ord_bp_id: %v", data_Ord.Ord_bp_id)
// 	logs.InfoLogger.Printf("Ord_sltp_ordr_rfrnc: %v", data_Ord.Ord_sltp_ordr_rfrnc)
// 	logs.InfoLogger.Printf("Ord_ctcl_id: %v", data_Ord.Ord_ctcl_id)
// 	logs.InfoLogger.Printf("Ord_usr_id: %v", data_Ord.Ord_usr_id)
// 	logs.InfoLogger.Printf("Ord_cnt_id: %v", data_Ord.Ord_cnt_id)
// 	logs.InfoLogger.Printf("Ord_em_settlmnt_nmbr: %v", data_Ord.Ord_em_settlmnt_nmbr)
// 	logs.InfoLogger.Printf("Ord_mrgn_sqroff_mode: %v", data_Ord.Ord_mrgn_sqroff_mode)
// 	logs.InfoLogger.Printf("Ord_cncl_qty: %v", data_Ord.Ord_cncl_qty)
// 	logs.InfoLogger.Printf("Ord_ordr_typ: %v", data_Ord.Ord_ordr_typ)
// 	logs.InfoLogger.Printf("Ord_valid_dt: %v", data_Ord.Ord_valid_dt)
// 	logs.InfoLogger.Printf("Ord_cal_flg: %v", data_Ord.Ord_cal_flg)
// 	logs.InfoLogger.Printf("Ord_xchng_ack: %v", data_Ord.Ord_xchng_ack)
// 	logs.InfoLogger.Printf("Ord_em_rollovr_flg: %v", data_Ord.Ord_em_rollovr_flg)
// 	logs.InfoLogger.Printf("Ord_trd_val: %v", data_Ord.Ord_trd_val)
// 	logs.InfoLogger.Printf("Ord_trd_cntrct_nmbr: %v", data_Ord.Ord_trd_cntrct_nmbr)
// 	logs.InfoLogger.Printf("Ord_avg_exctd_rt: %v", data_Ord.Ord_avg_exctd_rt)
// 	logs.InfoLogger.Printf("Ord_prc_imp_flg: %v", data_Ord.Ord_prc_imp_flg)
// 	logs.InfoLogger.Printf("Ord_mbc_flg: %v", data_Ord.Ord_mbc_flg)
// 	logs.InfoLogger.Printf("Ord_trl_amt: %v", data_Ord.Ord_trl_amt)
// 	logs.InfoLogger.Printf("Ord_lmt_offst: %v", data_Ord.Ord_lmt_offst)
// 	logs.InfoLogger.Printf("Ord_source_flg: %v", data_Ord.Ord_source_flg)
// 	logs.InfoLogger.Printf("Ord_pan_no: %v", data_Ord.Ord_pan_no)
// 	logs.InfoLogger.Printf("Ord_atm_payout_stts: %v", data_Ord.Ord_atm_payout_stts)
// 	logs.InfoLogger.Printf("Ord_esp_cd: %v", data_Ord.Ord_esp_cd)
// 	logs.InfoLogger.Printf("Ord_remarks: %v", data_Ord.Ord_remarks)
// 	logs.InfoLogger.Printf("Ord_wthld_amt_stts: %v", data_Ord.Ord_wthld_amt_stts)
// 	logs.InfoLogger.Printf("Ord_pstn_xchng_cd: %v", data_Ord.Ord_pstn_xchng_cd)
// 	logs.InfoLogger.Printf("Ord_interop_ord_flg: %v", data_Ord.Ord_interop_ord_flg)
// 	logs.InfoLogger.Printf("Ord_settlement_period: %v", data_Ord.Ord_settlement_period)
// 	logs.InfoLogger.Printf("Ord_algo_id: %v", data_Ord.Ord_algo_id)
// 	logs.InfoLogger.Printf("Ord_bundle_name: %v", data_Ord.Ord_bundle_name)
// 	logs.InfoLogger.Printf("Ord_prt_flg: %v", data_Ord.Ord_prt_flg)
// 	logs.InfoLogger.Printf("Ord_src_tag: %v", data_Ord.Ord_src_tag)
// 	logs.InfoLogger.Printf("Ord_rls_amt: %v", data_Ord.Ord_rls_amt)
// 	logs.InfoLogger.Printf("Ord_rls_date: %v", data_Ord.Ord_rls_date)
// 	logs.InfoLogger.Printf("Ord_mtf_unplg_sqroff: %v", data_Ord.Ord_mtf_unplg_sqroff)
// 	logs.InfoLogger.Printf("Ord_n_ordr_qty: %v", data_Ord.Ord_n_ordr_qty)
// 	logs.InfoLogger.Printf("Ord_ack_date: %v", data_Ord.Ord_ack_date)
// 	logs.InfoLogger.Printf("Ord_last_activity_ref: %v", data_Ord.Ord_last_activity_ref)
// 	logs.InfoLogger.Printf("Ord_clm_clnt_cd: %v", data_Ord.Ord_clm_clnt_cd)
// 	logs.DebugLogger.Printf("Trade Details: %+v\n", data_Trd)
// 	logs.InfoLogger.Println("Trade saved successfully:", trd_ref_num)

// 	if err := database.DB.Create(&data_Trd).Error; err != nil {
// 		database.DB.Rollback()
// 		logs.ErrorLogger.Printf("Failed to create order details: %v", err)
// 		return nil, status.Errorf(codes.Internal, "Failed to create trade details: %v", err)
// 	}

// 	if err := database.DB.Create(&data_Ord).Error; err != nil {
// 		database.DB.Rollback()
// 		logs.ErrorLogger.Printf("Failed to create trade details: %v", err)
// 		return nil, status.Errorf(codes.Internal, "Failed to create order details: %v", err)
// 	}

// 	// if err := database.DB.Commit().Error; err != nil {
// 	// 	logs.ErrorLogger.Printf("Failed to commit transaction: %v", err)
// 	// 	return nil, status.Errorf(codes.Internal, "Failed to commit transaction: %v", err)
// 	// }

// 	logs.InfoLogger.Printf("Trade and Order saved successfully: Trade Ref: %v, Order Ref: %v", trd_ref_num, ordr_rfrnc_no)

// 	return &pb.Otp_SquareoffResponse{Success: true}, nil
// }

var orderCounter = 0

func GenerateOrderReference(pipeID int) string {
	currentDate := time.Now().Format("20060102")
	pipeIDStr := fmt.Sprintf("%02d", pipeID)
	orderCounter += 1
	counterStr := fmt.Sprintf("%08d", orderCounter)
	orderReference := currentDate + pipeIDStr + counterStr

	return orderReference
}

// func GenerateOrderReference(db *sql.DB, pipeID int) (string, error) {

// 	currentDate := time.Now().Format("20060102")
// 	pipeIDStr := fmt.Sprintf("%02d", pipeID)

// 	query := `
// 		SELECT MAX(fxb_ordr_rfrnc)
// 		FROM fxb_fo_xchng_book
// 		WHERE fxb_ordr_rfrnc LIKE $1;
// `
// 	pattern := currentDate + "%"
// 	var maxOrderRef sql.NullString
// 	err := db.QueryRow(query, pattern).Scan(&maxOrderRef)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to query max order reference: %v", err)
// 	}
// 	if maxOrderRef.Valid {
// 		lastCounter := maxOrderRef.String[len(maxOrderRef.String)-8:] // Extract the last 8 digits
// 		fmt.Sscanf(lastCounter, "%d", &orderCounter)
// 		orderCounter += 1
// 	} else {
// 		// If no previous order reference, start with 1
// 		orderCounter += 1
// 	}
// 	counterStr := fmt.Sprintf("%08d", orderCounter)
// 	orderReference := currentDate + pipeIDStr + counterStr

// 	return orderReference, nil
// }

var today time.Time

func init() {
	dateString := time.Now().Format("2006-01-02")
	today, _ = time.Parse("2006-01-02", dateString)
}

func generateTradeReference() (string, error) {
	var nextVal int64
	if err := database.DB.Raw("SELECT nextval('trd_seq_num')").Scan(&nextVal).Error; err != nil {
		return "", err
	}
	formattedSeqNum := fmt.Sprintf("%08d", nextVal)
	formattedDate := time.Now().Format("2006/0102/")
	return formattedDate + formattedSeqNum, nil
}

func fetchClientMaster(ctx context.Context, matchAccount string) (database.Clm_clnt_mstr, error) {
	var param database.Clm_clnt_mstr
	err := database.DB.WithContext(ctx).Raw("SELECT clm_bp_id, clm_clnt_cd FROM clm_clnt_mstr WHERE clm_mtch_accnt = ?", matchAccount).Scan(&param).Error
	return param, err
}

func fetchSystemParameters(ctx context.Context, exchangeCode string) (database.Esp_em_systm_prmtr, error) {
	var param database.Esp_em_systm_prmtr
	err := database.DB.WithContext(ctx).Raw("SELECT esp_dp_id, esp_dp_clnt_id FROM esp_em_systm_prmtr WHERE esp_xchng_cd = ?", exchangeCode).Scan(&param).Error
	return param, err
}

func convertStringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func main() {

	logs.InitLogger()
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterSqureoffServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
