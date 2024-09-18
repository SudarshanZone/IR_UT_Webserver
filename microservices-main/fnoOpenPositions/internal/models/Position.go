package models

type FnoPosition struct {
    Contract          string  `json:"contract" gorm:"column:Contract"`
    Position          string  `json:"position" gorm:"column:Position"`
    TotalQty          int64    `json:"totalQty" gorm:"column:TotalQty"`
    AvgCostPrice      float64 `json:"avgCostPrice" gorm:"column:AvgCostPrice"`
    ExchangeCode      string  `json:"exchangeCode" gorm:"column:ExchangeCode"`
    BuyQty            int64    `json:"buyQty" gorm:"column:BuyQty"`
    ClaimMatchAccount string  `json:"claimMatchAccount" gorm:"column:ClaimMatchAccount"`
    ProductType       string  `json:"productType" gorm:"column:ProductType"`
    IndexStock        string  `json:"indexStock" gorm:"column:IndexStock"`
    Underlying        string  `json:"underlying" gorm:"column:Underlying"`
    ExpiryDate        string  `json:"expiryDate" gorm:"column:ExpiryDate"`
    ExerciseType      string  `json:"exerciseType" gorm:"column:ExerciseType"`
    OptionType        string  `json:"optionType" gorm:"column:OptionType"`
    StrikePrice       int64 `json:"strikePrice" gorm:"column:StrikePrice"`
    UCCCode           string  `json:"uccCode" gorm:"column:UccCode"`
    OpenPstnFlow      string  `json:"openPstnFlow" gorm:"column:OpenPstnFlow"`

    //New Models Added
    BuyOrderValue         float64 `json:"buyOrderValue" gorm:"column:BuyOrderValue"`
    SellQty               int64     `json:"sellQty" gorm:"column:SellQty"`
    SellOrderValue        float64 `json:"sellOrderValue" gorm:"column:SellOrderValue"`
    ExchangeBuyQty        int64     `json:"exchangeBuyQty" gorm:"column:ExchangeBuyQty"`
    ExchangeBuyOrderValue float64 `json:"exchangeBuyOrderValue" gorm:"column:ExchangeBuyOrderValue"`
    ExchangeSellQty       int64     `json:"exchangeSellQty" gorm:"column:ExchangeSellQty"`
    ExchangeSellOrderValue float64 `json:"exchangeSellOrderValue" gorm:"column:ExchangeSellOrderValue"`
    BuyExpectedQty        int64     `json:"buyExpectedQty" gorm:"column:BuyExpectedQty"`
    SellExpectedQty       int64     `json:"sellExpectedQty" gorm:"column:SellExpectedQty"`
    OpenPositionValue     float64 `json:"openPositionValue" gorm:"column:OpenPositionValue"`
    ExQuantity            int64     `json:"exQuantity" gorm:"column:ExQuantity"`
    AssignedQuantity      int64     `json:"assignedQuantity" gorm:"column:AssignedQuantity"`
    OptionPremium         float64 `json:"optionPremium" gorm:"column:OptionPremium"`
    MtmOpenValue          float64 `json:"mtmOpenValue" gorm:"column:MtmOpenValue"`
    ImtmOpenValue         float64 `json:"imtmOpenValue" gorm:"column:ImtmOpenValue"`
    UdeliveryMargin       float64 `json:"udeliveryMargin" gorm:"column:UdeliveryMargin"`
    MtmFlag               string  `json:"mtmFlag" gorm:"column:MtmFlag"`
    TriggerPrice          float64 `json:"triggerPrice" gorm:"column:TriggerPrice"`
    MinTriggerPrice       float64 `json:"minTriggerPrice" gorm:"column:MinTriggerPrice"`
    DeliveryModeFlag      string  `json:"deliveryModeFlag" gorm:"column:DeliveryModeFlag"`
    DeliveryObligationBlocked float64 `json:"deliveryObligationBlocked" gorm:"column:DeliveryObligationBlocked"`
    DeliveryQtyBlocked    int64  `json:"deliveryQtyBlocked" gorm:"column:DeliveryQtyBlocked"`
    MarginChangeDate      string  `json:"marginChangeDate" gorm:"column:MarginChangeDate"`

}


// type FnoPosition struct {
//     Contract          string  `json:"contract" gorm:"column:Contract"`
//     Position          string  `json:"position" gorm:"column:Position"`
//     TotalQty          int32    `json:"totalQty" gorm:"column:TotalQty"`
//     AvgCostPrice      float64 `json:"avgCostPrice" gorm:"column:AvgCostPrice"`
//     ExchangeCode      string  `json:"exchangeCode" gorm:"column:ExchangeCode"`
//     BuyQty            int     `json:"buyQty" gorm:"column:BuyQty"`
//     ClaimMatchAccount string  `json:"claimMatchAccount" gorm:"column:ClaimMatchAccount"`
//     ProductType       string  `json:"productType" gorm:"column:ProductType"`
//     IndexStock        string  `json:"indexStock" gorm:"column:IndexStock"`
//     Underlying        string  `json:"underlying" gorm:"column:Underlying"`
//     ExpiryDate        string  `json:"expiryDate" gorm:"column:ExpiryDate"`
//     ExerciseType      string  `json:"exerciseType" gorm:"column:ExerciseType"`
//     OptionType        string  `json:"optionType" gorm:"column:OptionType"`
//     StrikePrice       float64 `json:"strikePrice" gorm:"column:StrikePrice"`
//     UCCCode           string  `json:"uccCode" gorm:"column:UccCode"`
//     OpenPstnFlow      string  `json:"openPstnFlow" gorm:"column:OpenPstnFlow"`

//     //New Models Added
//     BuyOrderValue         float64 `json:"buyOrderValue" gorm:"column:BuyOrderValue"`
//     SellQty               int     `json:"sellQty" gorm:"column:SellQty"`
//     SellOrderValue        float64 `json:"sellOrderValue" gorm:"column:SellOrderValue"`
//     ExchangeBuyQty        int     `json:"exchangeBuyQty" gorm:"column:ExchangeBuyQty"`
//     ExchangeBuyOrderValue float64 `json:"exchangeBuyOrderValue" gorm:"column:ExchangeBuyOrderValue"`
//     ExchangeSellQty       int     `json:"exchangeSellQty" gorm:"column:ExchangeSellQty"`
//     ExchangeSellOrderValue float64 `json:"exchangeSellOrderValue" gorm:"column:ExchangeSellOrderValue"`
//     BuyExpectedQty        int     `json:"buyExpectedQty" gorm:"column:BuyExpectedQty"`
//     SellExpectedQty       int     `json:"sellExpectedQty" gorm:"column:SellExpectedQty"`
//     OpenPositionValue     float64 `json:"openPositionValue" gorm:"column:OpenPositionValue"`
//     ExQuantity            int     `json:"exQuantity" gorm:"column:ExQuantity"`
//     AssignedQuantity      int     `json:"assignedQuantity" gorm:"column:AssignedQuantity"`
//     OptionPremium         float64 `json:"optionPremium" gorm:"column:OptionPremium"`
//     MtmOpenValue          float64 `json:"mtmOpenValue" gorm:"column:MtmOpenValue"`
//     ImtmOpenValue         float64 `json:"imtmOpenValue" gorm:"column:ImtmOpenValue"`
//     UdeliveryMargin       float64 `json:"udeliveryMargin" gorm:"column:UdeliveryMargin"`
//     MtmFlag               string  `json:"mtmFlag" gorm:"column:MtmFlag"`
//     TriggerPrice          float64 `json:"triggerPrice" gorm:"column:TriggerPrice"`
//     MinTriggerPrice       float64 `json:"minTriggerPrice" gorm:"column:MinTriggerPrice"`
//     DeliveryModeFlag      string  `json:"deliveryModeFlag" gorm:"column:DeliveryModeFlag"`
//     DeliveryObligationBlocked string `json:"deliveryObligationBlocked" gorm:"column:DeliveryObligationBlocked"`
//     DeliveryQtyBlocked    string  `json:"deliveryQtyBlocked" gorm:"column:DeliveryQtyBlocked"`
//     MarginChangeDate      string  `json:"marginChangeDate" gorm:"column:MarginChangeDate"`

// }
