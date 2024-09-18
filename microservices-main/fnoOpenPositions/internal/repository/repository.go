package repository

import (
	"fmt"

	"github.com/krishnakashyap0704/microservices/fnoOpenPositions/internal/models"
	"gorm.io/gorm"
)

type FnoPositionRepository struct {
	Db *gorm.DB
}

func (repo *FnoPositionRepository) GetPositionsByClaimMatchAccount(claimMatchAccount string) ([]models.FnoPosition, error) {
	var positions []models.FnoPosition

	query := ` SELECT
    CASE
        WHEN FCP_PRDCT_TYP = 'F' THEN
            'FUT-' || TRIM(FCP_UNDRLYNG) || '-' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '')
        WHEN FCP_PRDCT_TYP = 'O' THEN
            'OPT-' || TRIM(FCP_UNDRLYNG) || '-' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '') || '-' ||
            COALESCE(FCP_STRK_PRC::text, '0') || '-' ||
            CASE
                WHEN FCP_OPT_TYP = 'C' THEN 'CE'
                WHEN FCP_OPT_TYP = 'P' THEN 'PE'
                ELSE ''
            END
        ELSE
            TRIM(FCP_UNDRLYNG) || ' ' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '')
    END AS "Contract",
    CASE
        WHEN FCP_OPNPSTN_FLW = 'B' THEN 'BUY'
        WHEN FCP_OPNPSTN_FLW = 'S' THEN 'SELL'
        ELSE ''
    END AS "Position",
    ABS(FCP_OPNPSTN_QTY) AS "TotalQty",
    COALESCE(NULLIF(FCP_AVG_PRC::text, ''), '0') AS "AvgCostPrice",
    COALESCE(FCP_XCHNG_CD, '') AS "ExchangeCode",
    COALESCE(NULLIF(FCP_IBUY_QTY::text, ''), '0') AS "BuyQty",
    FCP_CLM_MTCH_ACCNT AS "ClaimMatchAccount",
    FCP_PRDCT_TYP AS "ProductType",
    FCP_INDSTK AS "IndexStock",
    TRIM(FCP_UNDRLYNG) AS "Underlying",
    FCP_EXPRY_DT AS "ExpiryDate",
    FCP_EXER_TYP AS "ExerciseType",
    FCP_OPT_TYP AS "OptionType",
    COALESCE(NULLIF(FCP_STRK_PRC::text, ''), '0') AS "StrikePrice",
    FCP_UCC_CD AS "UccCode",
    FCP_OPNPSTN_FLW AS "OpenPstnFlow",

    -- New fields
    COALESCE(NULLIF(FCP_IBUY_ORD_VAL::text, ''), '0') AS "BuyOrderValue",
    COALESCE(NULLIF(FCP_ISELL_QTY::text, ''), '0') AS "SellQty",
    COALESCE(NULLIF(FCP_ISELL_ORD_VAL::text, ''), '0') AS "SellOrderValue",
    COALESCE(NULLIF(FCP_EXBUY_QTY::text, ''), '0') AS "ExchangeBuyQty",
    COALESCE(NULLIF(FCP_EXBUY_ORD_VAL::text, ''), '0') AS "ExchangeBuyOrderValue",
    COALESCE(NULLIF(FCP_EXSELL_QTY::text, ''), '0') AS "ExchangeSellQty",
    COALESCE(NULLIF(FCP_EXSELL_ORD_VAL::text, ''), '0') AS "ExchangeSellOrderValue",
    COALESCE(NULLIF(FCP_BUY_EXCTD_QTY::text, ''), '0') AS "BuyExpectedQty",
    COALESCE(NULLIF(FCP_SELL_EXCTD_QTY::text, ''), '0') AS "SellExpectedQty",
    COALESCE(NULLIF(FCP_OPNPSTN_VAL::text, ''), '0') AS "OpenPositionValue",
    COALESCE(NULLIF(FCP_EXRC_QTY::text, ''), '0') AS "ExQuantity",
    COALESCE(NULLIF(FCP_ASGND_QTY::text, ''), '0') AS "AssignedQuantity",
    COALESCE(NULLIF(FCP_OPT_PREMIUM::text, ''), '0') AS "OptionPremium",
    COALESCE(NULLIF(FCP_MTM_OPN_VAL::text, ''), '0') AS "MtmOpenValue",
    COALESCE(NULLIF(FCP_IMTM_OPN_VAL::text, ''), '0') AS "ImtmOpenValue",
    COALESCE(NULLIF(FCP_UDLVRY_MRGN::text, ''), '0') AS "UdeliveryMargin",
    COALESCE(FCP_MTM_FLG, '') AS "MtmFlag",
    COALESCE(NULLIF(FCP_TRG_PRC::text, ''), '0') AS "TriggerPrice",
    COALESCE(NULLIF(FCP_MIN_TRG_PRC::text, ''), '0') AS "MinTriggerPrice",
    COALESCE(FCP_DLVRY_MODE_FLAG, '') AS "DeliveryModeFlag",
    COALESCE(NULLIF(FCP_DLVRY_OBLGAMT_BLCKD::text, ''), '0')::float AS "DeliveryObligationBlocked",
    COALESCE(NULLIF(FCP_DLVRY_QTY_BLCKD::text, ''), '0')::int AS "DeliveryQtyBlocked",
    COALESCE(NULLIF(FCP_MRGN_CHNG_DT::text, ''), '') AS "MarginChangeDate"
FROM
    FCP_FO_SPN_CNTRCT_PSTN
WHERE
    FCP_CLM_MTCH_ACCNT = ?
    AND FCP_OPNPSTN_FLW != 'N';
`
	err := repo.Db.Raw(query, claimMatchAccount).Scan(&positions).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching positions: %w", err)
	}

	return positions, nil
}

//    query := `SELECT
//     CASE
//         WHEN FCP_PRDCT_TYP = 'F' THEN
//             'FUT-' || TRIM(FCP_UNDRLYNG) || '-' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '')
//         WHEN FCP_PRDCT_TYP = 'O' THEN
//             'OPT-' || TRIM(FCP_UNDRLYNG) || '-' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '') || '-' ||
//             COALESCE(FCP_STRK_PRC::text, '0') || '-' ||
//             CASE
//                 WHEN FCP_OPT_TYP = 'C' THEN 'CE'
//                 WHEN FCP_OPT_TYP = 'P' THEN 'PE'
//                 ELSE ''
//             END
//         ELSE
//             TRIM(FCP_UNDRLYNG) || ' ' || COALESCE(TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY'), '')
//     END AS "Contract",
//     CASE
//         WHEN FCP_OPNPSTN_FLW = 'B' THEN 'BUY'
//         WHEN FCP_OPNPSTN_FLW = 'S' THEN 'SELL'
//         ELSE ''
//     END AS "Position",
//     ABS(FCP_OPNPSTN_QTY) AS "TotalQty",
//     COALESCE(NULLIF(FCP_AVG_PRC::text, ''), '0') AS "AvgCostPrice",
//     COALESCE(FCP_XCHNG_CD, '') AS "ExchangeCode",
//     COALESCE(NULLIF(FCP_IBUY_QTY::text, ''), '0') AS "BuyQty",
//     FCP_CLM_MTCH_ACCNT AS "ClaimMatchAccount",
//     FCP_PRDCT_TYP AS "ProductType",
//     FCP_INDSTK AS "IndexStock",
//     TRIM(FCP_UNDRLYNG) AS "Underlying",
//     FCP_EXPRY_DT AS "ExpiryDate",
//     FCP_EXER_TYP AS "ExerciseType",
//     FCP_OPT_TYP AS "OptionType",
//     COALESCE(NULLIF(FCP_STRK_PRC::text, ''), '0') AS "StrikePrice",
//     FCP_UCC_CD AS "UccCode",
//     FCP_OPNPSTN_FLW AS "OpenPstnFlow",

//     -- new fields added
//     COALESCE(NULLIF(FCP_IBUY_ORD_VAL::text, ''), '0') AS "BuyOrderValue",
//     COALESCE(NULLIF(FCP_ISELL_QTY::text, ''), '0') AS "SellQty",
//     COALESCE(NULLIF(FCP_ISELL_ORD_VAL::text, ''), '0') AS "SellOrderValue",
//     COALESCE(NULLIF(FCP_EXBUY_QTY::text, ''), '0') AS "ExchangeBuyQty",
//     COALESCE(NULLIF(FCP_EXBUY_ORD_VAL::text, ''), '0') AS "ExchangeBuyOrderValue",
//     COALESCE(NULLIF(FCP_EXSELL_QTY::text, ''), '0') AS "ExchangeSellQty",
//     COALESCE(NULLIF(FCP_EXSELL_ORD_VAL::text, ''), '0') AS "ExchangeSellOrderValue",
//     COALESCE(NULLIF(FCP_BUY_EXCTD_QTY::text, ''), '0') AS "BuyExpectedQty",
//     COALESCE(NULLIF(FCP_SELL_EXCTD_QTY::text, ''), '0') AS "SellExpectedQty",
//     COALESCE(NULLIF(FCP_OPNPSTN_VAL::text, ''), '0') AS "OpenPositionValue",
//     COALESCE(NULLIF(FCP_EXRC_QTY::text, ''), '0') AS "ExQuantity",
//     COALESCE(NULLIF(FCP_ASGND_QTY::text, ''), '0') AS "AssignedQuantity",
//     COALESCE(NULLIF(FCP_OPT_PREMIUM::text, ''), '0') AS "OptionPremium",
//     COALESCE(NULLIF(FCP_MTM_OPN_VAL::text, ''), '0') AS "MtmOpenValue",
//     COALESCE(NULLIF(FCP_IMTM_OPN_VAL::text, ''), '0') AS "ImtmOpenValue",
//     COALESCE(NULLIF(FCP_UDLVRY_MRGN::text, ''), '0') AS "UdeliveryMargin",
//     COALESCE(FCP_MTM_FLG, '') AS "MtmFlag",
//     COALESCE(NULLIF(FCP_TRG_PRC::text, ''), '0') AS "TriggerPrice",
//     COALESCE(NULLIF(FCP_MIN_TRG_PRC::text, ''), '0') AS "MinTriggerPrice",
//     COALESCE(FCP_DLVRY_MODE_FLAG, '') AS "DeliveryModeFlag",
//     COALESCE(NULLIF(FCP_DLVRY_OBLGAMT_BLCKD::text, ''), '0')::float AS "DeliveryObligationBlocked",  -- Handle string to float conversion
//     COALESCE(NULLIF(FCP_DLVRY_QTY_BLCKD::text, ''), '0')::int AS "DeliveryQtyBlocked",  -- Handle string to int conversion
//     COALESCE(NULLIF(FCP_MRGN_CHNG_DT::text, ''), '') AS "MarginChangeDate"
// FROM
//     FCP_FO_SPN_CNTRCT_PSTN
// WHERE
//     FCP_CLM_MTCH_ACCNT = ?
//     AND FCP_OPNPSTN_FLW != 'N';
// `
