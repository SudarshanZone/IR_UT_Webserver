// protoc --proto_path=proto --go_out=generated --go-grpc_out=generated proto/commo.proto
// protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal/positions/fno_pos.proto
//protoc --proto_path=proto --go_out=generated --go-grpc_out=generated proto/commo.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: comsquoff/comSquareOff.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message for CCP table
type CcpCodSpnCntrctPstn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CcpClmMtchAccnt     string  `protobuf:"bytes,1,opt,name=ccp_clm_mtch_accnt,json=ccpClmMtchAccnt,proto3" json:"ccp_clm_mtch_accnt,omitempty"`                  // character(10)
	CcpXchngCd          string  `protobuf:"bytes,2,opt,name=ccp_xchng_cd,json=ccpXchngCd,proto3" json:"ccp_xchng_cd,omitempty"`                                   // character(3)
	CcpPrdctTyp         string  `protobuf:"bytes,3,opt,name=ccp_prdct_typ,json=ccpPrdctTyp,proto3" json:"ccp_prdct_typ,omitempty"`                                // character(1)
	CcpIndstk           string  `protobuf:"bytes,4,opt,name=ccp_indstk,json=ccpIndstk,proto3" json:"ccp_indstk,omitempty"`                                        // character(1)
	CcpUndrlyng         string  `protobuf:"bytes,5,opt,name=ccp_undrlyng,json=ccpUndrlyng,proto3" json:"ccp_undrlyng,omitempty"`                                  // character(6)
	CcpExpryDt          string  `protobuf:"bytes,6,opt,name=ccp_expry_dt,json=ccpExpryDt,proto3" json:"ccp_expry_dt,omitempty"`                                   // timestamp without time zone
	CcpExerTyp          string  `protobuf:"bytes,7,opt,name=ccp_exer_typ,json=ccpExerTyp,proto3" json:"ccp_exer_typ,omitempty"`                                   // character(1)
	CcpStrkPrc          int64   `protobuf:"varint,8,opt,name=ccp_strk_prc,json=ccpStrkPrc,proto3" json:"ccp_strk_prc,omitempty"`                                  // bigint
	CcpOptTyp           string  `protobuf:"bytes,9,opt,name=ccp_opt_typ,json=ccpOptTyp,proto3" json:"ccp_opt_typ,omitempty"`                                      // character(1)
	CcpIbuyQty          int64   `protobuf:"varint,10,opt,name=ccp_ibuy_qty,json=ccpIbuyQty,proto3" json:"ccp_ibuy_qty,omitempty"`                                 // bigint
	CcpIbuyOrdVal       float64 `protobuf:"fixed64,11,opt,name=ccp_ibuy_ord_val,json=ccpIbuyOrdVal,proto3" json:"ccp_ibuy_ord_val,omitempty"`                     // numeric(24,2)
	CcpIsellQty         int64   `protobuf:"varint,12,opt,name=ccp_isell_qty,json=ccpIsellQty,proto3" json:"ccp_isell_qty,omitempty"`                              // bigint
	CcpIsellOrdVal      float64 `protobuf:"fixed64,13,opt,name=ccp_isell_ord_val,json=ccpIsellOrdVal,proto3" json:"ccp_isell_ord_val,omitempty"`                  // numeric(24,2)
	CcpExbuyQty         int64   `protobuf:"varint,14,opt,name=ccp_exbuy_qty,json=ccpExbuyQty,proto3" json:"ccp_exbuy_qty,omitempty"`                              // bigint
	CcpExbuyOrdVal      float64 `protobuf:"fixed64,15,opt,name=ccp_exbuy_ord_val,json=ccpExbuyOrdVal,proto3" json:"ccp_exbuy_ord_val,omitempty"`                  // numeric(24,2)
	CcpExsellQty        int64   `protobuf:"varint,16,opt,name=ccp_exsell_qty,json=ccpExsellQty,proto3" json:"ccp_exsell_qty,omitempty"`                           // bigint
	CcpExsellOrdVal     float64 `protobuf:"fixed64,17,opt,name=ccp_exsell_ord_val,json=ccpExsellOrdVal,proto3" json:"ccp_exsell_ord_val,omitempty"`               // numeric(24,2)
	CcpBuyExctdQty      int64   `protobuf:"varint,18,opt,name=ccp_buy_exctd_qty,json=ccpBuyExctdQty,proto3" json:"ccp_buy_exctd_qty,omitempty"`                   // bigint
	CcpSellExctdQty     int64   `protobuf:"varint,19,opt,name=ccp_sell_exctd_qty,json=ccpSellExctdQty,proto3" json:"ccp_sell_exctd_qty,omitempty"`                // bigint
	CcpOpnpstnFlw       string  `protobuf:"bytes,20,opt,name=ccp_opnpstn_flw,json=ccpOpnpstnFlw,proto3" json:"ccp_opnpstn_flw,omitempty"`                         // character(1)
	CcpOpnpstnQty       int64   `protobuf:"varint,21,opt,name=ccp_opnpstn_qty,json=ccpOpnpstnQty,proto3" json:"ccp_opnpstn_qty,omitempty"`                        // bigint
	CcpOpnpstnVal       float64 `protobuf:"fixed64,22,opt,name=ccp_opnpstn_val,json=ccpOpnpstnVal,proto3" json:"ccp_opnpstn_val,omitempty"`                       // numeric(24,2)
	CcpExrcQty          int64   `protobuf:"varint,23,opt,name=ccp_exrc_qty,json=ccpExrcQty,proto3" json:"ccp_exrc_qty,omitempty"`                                 // bigint
	CcpAsgndQty         int64   `protobuf:"varint,24,opt,name=ccp_asgnd_qty,json=ccpAsgndQty,proto3" json:"ccp_asgnd_qty,omitempty"`                              // bigint
	CcpOptPremium       float64 `protobuf:"fixed64,25,opt,name=ccp_opt_premium,json=ccpOptPremium,proto3" json:"ccp_opt_premium,omitempty"`                       // numeric(24,2)
	CcpMtmOpnVal        float64 `protobuf:"fixed64,26,opt,name=ccp_mtm_opn_val,json=ccpMtmOpnVal,proto3" json:"ccp_mtm_opn_val,omitempty"`                        // numeric(24,2)
	CcpImtmOpnVal       float64 `protobuf:"fixed64,27,opt,name=ccp_imtm_opn_val,json=ccpImtmOpnVal,proto3" json:"ccp_imtm_opn_val,omitempty"`                     // numeric(24,2)
	CcpExtrmlossMrgn    float64 `protobuf:"fixed64,28,opt,name=ccp_extrmloss_mrgn,json=ccpExtrmlossMrgn,proto3" json:"ccp_extrmloss_mrgn,omitempty"`              // numeric(24,2)
	CcpAddnlMrgn        float64 `protobuf:"fixed64,29,opt,name=ccp_addnl_mrgn,json=ccpAddnlMrgn,proto3" json:"ccp_addnl_mrgn,omitempty"`                          // numeric(24,2)
	CcpSpclMrgn         float64 `protobuf:"fixed64,30,opt,name=ccp_spcl_mrgn,json=ccpSpclMrgn,proto3" json:"ccp_spcl_mrgn,omitempty"`                             // numeric(24,2)
	CcpTndrMrgn         float64 `protobuf:"fixed64,31,opt,name=ccp_tndr_mrgn,json=ccpTndrMrgn,proto3" json:"ccp_tndr_mrgn,omitempty"`                             // numeric(24,2)
	CcpDlvryMrgn        float64 `protobuf:"fixed64,32,opt,name=ccp_dlvry_mrgn,json=ccpDlvryMrgn,proto3" json:"ccp_dlvry_mrgn,omitempty"`                          // numeric(24,2)
	CcpExtrmMinLossMrgn float64 `protobuf:"fixed64,33,opt,name=ccp_extrm_min_loss_mrgn,json=ccpExtrmMinLossMrgn,proto3" json:"ccp_extrm_min_loss_mrgn,omitempty"` // numeric(24,2)
	CcpMtmFlg           string  `protobuf:"bytes,34,opt,name=ccp_mtm_flg,json=ccpMtmFlg,proto3" json:"ccp_mtm_flg,omitempty"`                                     // character(1)
	CcpExtLossMrgn      float64 `protobuf:"fixed64,35,opt,name=ccp_ext_loss_mrgn,json=ccpExtLossMrgn,proto3" json:"ccp_ext_loss_mrgn,omitempty"`                  // Renamed field to avoid conflict
	CcpFlatValMrgn      float64 `protobuf:"fixed64,36,opt,name=ccp_flat_val_mrgn,json=ccpFlatValMrgn,proto3" json:"ccp_flat_val_mrgn,omitempty"`                  // numeric(24,2)
	CcpTrgPrc           float64 `protobuf:"fixed64,37,opt,name=ccp_trg_prc,json=ccpTrgPrc,proto3" json:"ccp_trg_prc,omitempty"`                                   // double precision
	CcpMinTrgPrc        float64 `protobuf:"fixed64,38,opt,name=ccp_min_trg_prc,json=ccpMinTrgPrc,proto3" json:"ccp_min_trg_prc,omitempty"`                        // double precision
	CcpDevolmntMrgn     float64 `protobuf:"fixed64,39,opt,name=ccp_devolmnt_mrgn,json=ccpDevolmntMrgn,proto3" json:"ccp_devolmnt_mrgn,omitempty"`                 // numeric(24,2)
	CcpMtmsqOrdcnt      int32   `protobuf:"varint,40,opt,name=ccp_mtmsq_ordcnt,json=ccpMtmsqOrdcnt,proto3" json:"ccp_mtmsq_ordcnt,omitempty"`                     // integer
	CcpAvgPrc           float64 `protobuf:"fixed64,41,opt,name=ccp_avg_prc,json=ccpAvgPrc,proto3" json:"ccp_avg_prc,omitempty"`
}

func (x *CcpCodSpnCntrctPstn) Reset() {
	*x = CcpCodSpnCntrctPstn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comsquoff_comSquareOff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CcpCodSpnCntrctPstn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CcpCodSpnCntrctPstn) ProtoMessage() {}

func (x *CcpCodSpnCntrctPstn) ProtoReflect() protoreflect.Message {
	mi := &file_comsquoff_comSquareOff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CcpCodSpnCntrctPstn.ProtoReflect.Descriptor instead.
func (*CcpCodSpnCntrctPstn) Descriptor() ([]byte, []int) {
	return file_comsquoff_comSquareOff_proto_rawDescGZIP(), []int{0}
}

func (x *CcpCodSpnCntrctPstn) GetCcpClmMtchAccnt() string {
	if x != nil {
		return x.CcpClmMtchAccnt
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpXchngCd() string {
	if x != nil {
		return x.CcpXchngCd
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpPrdctTyp() string {
	if x != nil {
		return x.CcpPrdctTyp
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpIndstk() string {
	if x != nil {
		return x.CcpIndstk
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpUndrlyng() string {
	if x != nil {
		return x.CcpUndrlyng
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpExpryDt() string {
	if x != nil {
		return x.CcpExpryDt
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpExerTyp() string {
	if x != nil {
		return x.CcpExerTyp
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpStrkPrc() int64 {
	if x != nil {
		return x.CcpStrkPrc
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpOptTyp() string {
	if x != nil {
		return x.CcpOptTyp
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpIbuyQty() int64 {
	if x != nil {
		return x.CcpIbuyQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpIbuyOrdVal() float64 {
	if x != nil {
		return x.CcpIbuyOrdVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpIsellQty() int64 {
	if x != nil {
		return x.CcpIsellQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpIsellOrdVal() float64 {
	if x != nil {
		return x.CcpIsellOrdVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExbuyQty() int64 {
	if x != nil {
		return x.CcpExbuyQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExbuyOrdVal() float64 {
	if x != nil {
		return x.CcpExbuyOrdVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExsellQty() int64 {
	if x != nil {
		return x.CcpExsellQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExsellOrdVal() float64 {
	if x != nil {
		return x.CcpExsellOrdVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpBuyExctdQty() int64 {
	if x != nil {
		return x.CcpBuyExctdQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpSellExctdQty() int64 {
	if x != nil {
		return x.CcpSellExctdQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpOpnpstnFlw() string {
	if x != nil {
		return x.CcpOpnpstnFlw
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpOpnpstnQty() int64 {
	if x != nil {
		return x.CcpOpnpstnQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpOpnpstnVal() float64 {
	if x != nil {
		return x.CcpOpnpstnVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExrcQty() int64 {
	if x != nil {
		return x.CcpExrcQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpAsgndQty() int64 {
	if x != nil {
		return x.CcpAsgndQty
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpOptPremium() float64 {
	if x != nil {
		return x.CcpOptPremium
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpMtmOpnVal() float64 {
	if x != nil {
		return x.CcpMtmOpnVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpImtmOpnVal() float64 {
	if x != nil {
		return x.CcpImtmOpnVal
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExtrmlossMrgn() float64 {
	if x != nil {
		return x.CcpExtrmlossMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpAddnlMrgn() float64 {
	if x != nil {
		return x.CcpAddnlMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpSpclMrgn() float64 {
	if x != nil {
		return x.CcpSpclMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpTndrMrgn() float64 {
	if x != nil {
		return x.CcpTndrMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpDlvryMrgn() float64 {
	if x != nil {
		return x.CcpDlvryMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpExtrmMinLossMrgn() float64 {
	if x != nil {
		return x.CcpExtrmMinLossMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpMtmFlg() string {
	if x != nil {
		return x.CcpMtmFlg
	}
	return ""
}

func (x *CcpCodSpnCntrctPstn) GetCcpExtLossMrgn() float64 {
	if x != nil {
		return x.CcpExtLossMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpFlatValMrgn() float64 {
	if x != nil {
		return x.CcpFlatValMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpTrgPrc() float64 {
	if x != nil {
		return x.CcpTrgPrc
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpMinTrgPrc() float64 {
	if x != nil {
		return x.CcpMinTrgPrc
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpDevolmntMrgn() float64 {
	if x != nil {
		return x.CcpDevolmntMrgn
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpMtmsqOrdcnt() int32 {
	if x != nil {
		return x.CcpMtmsqOrdcnt
	}
	return 0
}

func (x *CcpCodSpnCntrctPstn) GetCcpAvgPrc() float64 {
	if x != nil {
		return x.CcpAvgPrc
	}
	return 0
}

type SquareoffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ccp           *CcpCodSpnCntrctPstn `protobuf:"bytes,1,opt,name=ccp,proto3" json:"ccp,omitempty"`
	CCP_USR_ID    string               `protobuf:"bytes,2,opt,name=CCP_USR_ID,json=CCPUSRID,proto3" json:"CCP_USR_ID,omitempty"`
	CCP_PRDCT_TYP string               `protobuf:"bytes,3,opt,name=CCP_PRDCT_TYP,json=CCPPRDCTTYP,proto3" json:"CCP_PRDCT_TYP,omitempty"` // repeated FnoData Cdata = 3;
}

func (x *SquareoffRequest) Reset() {
	*x = SquareoffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comsquoff_comSquareOff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareoffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareoffRequest) ProtoMessage() {}

func (x *SquareoffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comsquoff_comSquareOff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareoffRequest.ProtoReflect.Descriptor instead.
func (*SquareoffRequest) Descriptor() ([]byte, []int) {
	return file_comsquoff_comSquareOff_proto_rawDescGZIP(), []int{1}
}

func (x *SquareoffRequest) GetCcp() *CcpCodSpnCntrctPstn {
	if x != nil {
		return x.Ccp
	}
	return nil
}

func (x *SquareoffRequest) GetCCP_USR_ID() string {
	if x != nil {
		return x.CCP_USR_ID
	}
	return ""
}

func (x *SquareoffRequest) GetCCP_PRDCT_TYP() string {
	if x != nil {
		return x.CCP_PRDCT_TYP
	}
	return ""
}

type SquareoffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SquareoffResponse) Reset() {
	*x = SquareoffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comsquoff_comSquareOff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareoffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareoffResponse) ProtoMessage() {}

func (x *SquareoffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comsquoff_comSquareOff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareoffResponse.ProtoReflect.Descriptor instead.
func (*SquareoffResponse) Descriptor() ([]byte, []int) {
	return file_comsquoff_comSquareOff_proto_rawDescGZIP(), []int{2}
}

func (x *SquareoffResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_comsquoff_comSquareOff_proto protoreflect.FileDescriptor

var file_comsquoff_comSquareOff_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x73, 0x71, 0x75, 0x6f, 0x66, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x4f, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x63, 0x6f, 0x6d, 0x73, 0x71, 0x75, 0x6f, 0x66, 0x66, 0x22, 0xc9, 0x0c, 0x0a, 0x17, 0x63, 0x63,
	0x70, 0x5f, 0x63, 0x6f, 0x64, 0x5f, 0x73, 0x70, 0x6e, 0x5f, 0x63, 0x6e, 0x74, 0x72, 0x63, 0x74,
	0x5f, 0x70, 0x73, 0x74, 0x6e, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x6d,
	0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x63, 0x70, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x63, 0x70, 0x5f, 0x78, 0x63, 0x68, 0x6e, 0x67, 0x5f,
	0x63, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x63, 0x70, 0x58, 0x63, 0x68,
	0x6e, 0x67, 0x43, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x64, 0x63,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x63, 0x70,
	0x50, 0x72, 0x64, 0x63, 0x74, 0x54, 0x79, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x63, 0x70, 0x5f,
	0x69, 0x6e, 0x64, 0x73, 0x74, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x63,
	0x70, 0x49, 0x6e, 0x64, 0x73, 0x74, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x63, 0x70, 0x5f, 0x75,
	0x6e, 0x64, 0x72, 0x6c, 0x79, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x63, 0x70, 0x55, 0x6e, 0x64, 0x72, 0x6c, 0x79, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x63,
	0x70, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x79, 0x5f, 0x64, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x63, 0x70, 0x45, 0x78, 0x70, 0x72, 0x79, 0x44, 0x74, 0x12, 0x20, 0x0a, 0x0c,
	0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x63, 0x70, 0x45, 0x78, 0x65, 0x72, 0x54, 0x79, 0x70, 0x12, 0x20,
	0x0a, 0x0c, 0x63, 0x63, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x6b, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x63, 0x70, 0x53, 0x74, 0x72, 0x6b, 0x50, 0x72, 0x63,
	0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x54, 0x79, 0x70,
	0x12, 0x20, 0x0a, 0x0c, 0x63, 0x63, 0x70, 0x5f, 0x69, 0x62, 0x75, 0x79, 0x5f, 0x71, 0x74, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x63, 0x70, 0x49, 0x62, 0x75, 0x79, 0x51,
	0x74, 0x79, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x63, 0x70, 0x5f, 0x69, 0x62, 0x75, 0x79, 0x5f, 0x6f,
	0x72, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x63,
	0x70, 0x49, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x63,
	0x63, 0x70, 0x5f, 0x69, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x63, 0x70, 0x49, 0x73, 0x65, 0x6c, 0x6c, 0x51, 0x74, 0x79, 0x12,
	0x29, 0x0a, 0x11, 0x63, 0x63, 0x70, 0x5f, 0x69, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x63, 0x70, 0x49,
	0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x63,
	0x70, 0x5f, 0x65, 0x78, 0x62, 0x75, 0x79, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x63, 0x70, 0x45, 0x78, 0x62, 0x75, 0x79, 0x51, 0x74, 0x79, 0x12, 0x29,
	0x0a, 0x11, 0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x63, 0x70, 0x45, 0x78,
	0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x63, 0x70,
	0x5f, 0x65, 0x78, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x63, 0x70, 0x45, 0x78, 0x73, 0x65, 0x6c, 0x6c, 0x51, 0x74, 0x79, 0x12,
	0x2b, 0x0a, 0x12, 0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x63, 0x70,
	0x45, 0x78, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x11,
	0x63, 0x63, 0x70, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x65, 0x78, 0x63, 0x74, 0x64, 0x5f, 0x71, 0x74,
	0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x63, 0x70, 0x42, 0x75, 0x79, 0x45,
	0x78, 0x63, 0x74, 0x64, 0x51, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x63, 0x70, 0x5f, 0x73,
	0x65, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x74, 0x64, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x63, 0x70, 0x53, 0x65, 0x6c, 0x6c, 0x45, 0x78, 0x63, 0x74,
	0x64, 0x51, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x6e, 0x70,
	0x73, 0x74, 0x6e, 0x5f, 0x66, 0x6c, 0x77, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x63, 0x70, 0x4f, 0x70, 0x6e, 0x70, 0x73, 0x74, 0x6e, 0x46, 0x6c, 0x77, 0x12, 0x26, 0x0a, 0x0f,
	0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x6e, 0x70, 0x73, 0x74, 0x6e, 0x5f, 0x71, 0x74, 0x79, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x63, 0x70, 0x4f, 0x70, 0x6e, 0x70, 0x73, 0x74,
	0x6e, 0x51, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x6e, 0x70,
	0x73, 0x74, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63,
	0x63, 0x70, 0x4f, 0x70, 0x6e, 0x70, 0x73, 0x74, 0x6e, 0x56, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c,
	0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x72, 0x63, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x63, 0x70, 0x45, 0x78, 0x72, 0x63, 0x51, 0x74, 0x79, 0x12, 0x22,
	0x0a, 0x0d, 0x63, 0x63, 0x70, 0x5f, 0x61, 0x73, 0x67, 0x6e, 0x64, 0x5f, 0x71, 0x74, 0x79, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x63, 0x70, 0x41, 0x73, 0x67, 0x6e, 0x64, 0x51,
	0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x74, 0x5f, 0x70, 0x72,
	0x65, 0x6d, 0x69, 0x75, 0x6d, 0x18, 0x19, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x63, 0x70,
	0x4f, 0x70, 0x74, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0f, 0x63, 0x63,
	0x70, 0x5f, 0x6d, 0x74, 0x6d, 0x5f, 0x6f, 0x70, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x63, 0x63, 0x70, 0x4d, 0x74, 0x6d, 0x4f, 0x70, 0x6e, 0x56, 0x61,
	0x6c, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x63, 0x70, 0x5f, 0x69, 0x6d, 0x74, 0x6d, 0x5f, 0x6f, 0x70,
	0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x63, 0x70,
	0x49, 0x6d, 0x74, 0x6d, 0x4f, 0x70, 0x6e, 0x56, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x63,
	0x70, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x6d, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x6d, 0x72, 0x67, 0x6e,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x63, 0x63, 0x70, 0x45, 0x78, 0x74, 0x72, 0x6d,
	0x6c, 0x6f, 0x73, 0x73, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x63, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x6e, 0x6c, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x63, 0x63, 0x70, 0x41, 0x64, 0x64, 0x6e, 0x6c, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x22,
	0x0a, 0x0d, 0x63, 0x63, 0x70, 0x5f, 0x73, 0x70, 0x63, 0x6c, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x63, 0x70, 0x53, 0x70, 0x63, 0x6c, 0x4d, 0x72,
	0x67, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x63, 0x70, 0x5f, 0x74, 0x6e, 0x64, 0x72, 0x5f, 0x6d,
	0x72, 0x67, 0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x63, 0x70, 0x54, 0x6e,
	0x64, 0x72, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x63, 0x70, 0x5f, 0x64, 0x6c,
	0x76, 0x72, 0x79, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x63, 0x63, 0x70, 0x44, 0x6c, 0x76, 0x72, 0x79, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x34, 0x0a, 0x17,
	0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x6d, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f,
	0x73, 0x73, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x63,
	0x63, 0x70, 0x45, 0x78, 0x74, 0x72, 0x6d, 0x4d, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x72,
	0x67, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x63, 0x70, 0x5f, 0x6d, 0x74, 0x6d, 0x5f, 0x66, 0x6c,
	0x67, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x63, 0x70, 0x4d, 0x74, 0x6d, 0x46,
	0x6c, 0x67, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x63, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x6f,
	0x73, 0x73, 0x5f, 0x6d, 0x72, 0x67, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63,
	0x63, 0x70, 0x45, 0x78, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x29, 0x0a,
	0x11, 0x63, 0x63, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x72,
	0x67, 0x6e, 0x18, 0x24, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x63, 0x70, 0x46, 0x6c, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x63, 0x70, 0x5f,
	0x74, 0x72, 0x67, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x25, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63,
	0x63, 0x70, 0x54, 0x72, 0x67, 0x50, 0x72, 0x63, 0x12, 0x25, 0x0a, 0x0f, 0x63, 0x63, 0x70, 0x5f,
	0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x67, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x26, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x63, 0x63, 0x70, 0x4d, 0x69, 0x6e, 0x54, 0x72, 0x67, 0x50, 0x72, 0x63, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x63, 0x70, 0x5f, 0x64, 0x65, 0x76, 0x6f, 0x6c, 0x6d, 0x6e, 0x74, 0x5f,
	0x6d, 0x72, 0x67, 0x6e, 0x18, 0x27, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x63, 0x70, 0x44,
	0x65, 0x76, 0x6f, 0x6c, 0x6d, 0x6e, 0x74, 0x4d, 0x72, 0x67, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x63,
	0x63, 0x70, 0x5f, 0x6d, 0x74, 0x6d, 0x73, 0x71, 0x5f, 0x6f, 0x72, 0x64, 0x63, 0x6e, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x63, 0x70, 0x4d, 0x74, 0x6d, 0x73, 0x71, 0x4f,
	0x72, 0x64, 0x63, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x63, 0x70, 0x5f, 0x61, 0x76, 0x67,
	0x5f, 0x70, 0x72, 0x63, 0x18, 0x29, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x63, 0x70, 0x41,
	0x76, 0x67, 0x50, 0x72, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x6f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x63, 0x63,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x73, 0x71, 0x75,
	0x6f, 0x66, 0x66, 0x2e, 0x63, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x5f, 0x73, 0x70, 0x6e, 0x5f,
	0x63, 0x6e, 0x74, 0x72, 0x63, 0x74, 0x5f, 0x70, 0x73, 0x74, 0x6e, 0x52, 0x03, 0x63, 0x63, 0x70,
	0x12, 0x1c, 0x0a, 0x0a, 0x43, 0x43, 0x50, 0x5f, 0x55, 0x53, 0x52, 0x5f, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x43, 0x50, 0x55, 0x53, 0x52, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0d, 0x43, 0x43, 0x50, 0x5f, 0x50, 0x52, 0x44, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x43, 0x50, 0x50, 0x52, 0x44, 0x43, 0x54, 0x54,
	0x59, 0x50, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x6f, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x53, 0x0a, 0x09, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x6f, 0x66, 0x66, 0x12, 0x46,
	0x0a, 0x09, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x6f, 0x66, 0x66, 0x12, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x73, 0x71, 0x75, 0x6f, 0x66, 0x66, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x6f, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x73, 0x71,
	0x75, 0x6f, 0x66, 0x66, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x6f, 0x66, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_comsquoff_comSquareOff_proto_rawDescOnce sync.Once
	file_comsquoff_comSquareOff_proto_rawDescData = file_comsquoff_comSquareOff_proto_rawDesc
)

func file_comsquoff_comSquareOff_proto_rawDescGZIP() []byte {
	file_comsquoff_comSquareOff_proto_rawDescOnce.Do(func() {
		file_comsquoff_comSquareOff_proto_rawDescData = protoimpl.X.CompressGZIP(file_comsquoff_comSquareOff_proto_rawDescData)
	})
	return file_comsquoff_comSquareOff_proto_rawDescData
}

var file_comsquoff_comSquareOff_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_comsquoff_comSquareOff_proto_goTypes = []any{
	(*CcpCodSpnCntrctPstn)(nil), // 0: comsquoff.ccp_cod_spn_cntrct_pstn
	(*SquareoffRequest)(nil),    // 1: comsquoff.SquareoffRequest
	(*SquareoffResponse)(nil),   // 2: comsquoff.SquareoffResponse
}
var file_comsquoff_comSquareOff_proto_depIdxs = []int32{
	0, // 0: comsquoff.SquareoffRequest.ccp:type_name -> comsquoff.ccp_cod_spn_cntrct_pstn
	1, // 1: comsquoff.Squareoff.Squareoff:input_type -> comsquoff.SquareoffRequest
	2, // 2: comsquoff.Squareoff.Squareoff:output_type -> comsquoff.SquareoffResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comsquoff_comSquareOff_proto_init() }
func file_comsquoff_comSquareOff_proto_init() {
	if File_comsquoff_comSquareOff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comsquoff_comSquareOff_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CcpCodSpnCntrctPstn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_comsquoff_comSquareOff_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SquareoffRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_comsquoff_comSquareOff_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SquareoffResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comsquoff_comSquareOff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comsquoff_comSquareOff_proto_goTypes,
		DependencyIndexes: file_comsquoff_comSquareOff_proto_depIdxs,
		MessageInfos:      file_comsquoff_comSquareOff_proto_msgTypes,
	}.Build()
	File_comsquoff_comSquareOff_proto = out.File
	file_comsquoff_comSquareOff_proto_rawDesc = nil
	file_comsquoff_comSquareOff_proto_goTypes = nil
	file_comsquoff_comSquareOff_proto_depIdxs = nil
}
