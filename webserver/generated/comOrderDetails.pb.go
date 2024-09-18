// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: comordr/comOrderDetails.proto

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

type ComOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodClmMtchAccnt string `protobuf:"bytes,1,opt,name=cod_clm_mtch_accnt,json=codClmMtchAccnt,proto3" json:"cod_clm_mtch_accnt,omitempty"`
}

func (x *ComOrderRequest) Reset() {
	*x = ComOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comordr_comOrderDetails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComOrderRequest) ProtoMessage() {}

func (x *ComOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comordr_comOrderDetails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComOrderRequest.ProtoReflect.Descriptor instead.
func (*ComOrderRequest) Descriptor() ([]byte, []int) {
	return file_comordr_comOrderDetails_proto_rawDescGZIP(), []int{0}
}

func (x *ComOrderRequest) GetCodClmMtchAccnt() string {
	if x != nil {
		return x.CodClmMtchAccnt
	}
	return ""
}

type ComOrdrDtls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodClmMtchAccnt string  `protobuf:"bytes,1,opt,name=cod_clm_mtch_accnt,json=codClmMtchAccnt,proto3" json:"cod_clm_mtch_accnt,omitempty"`
	CodPrdctTyp     string  `protobuf:"bytes,2,opt,name=cod_prdct_typ,json=codPrdctTyp,proto3" json:"cod_prdct_typ,omitempty"`
	CodUndrlyng     string  `protobuf:"bytes,3,opt,name=cod_undrlyng,json=codUndrlyng,proto3" json:"cod_undrlyng,omitempty"`
	CodExpryDt      string  `protobuf:"bytes,4,opt,name=cod_expry_dt,json=codExpryDt,proto3" json:"cod_expry_dt,omitempty"`
	CodLmtRt        float32 `protobuf:"fixed32,5,opt,name=cod_lmt_rt,json=codLmtRt,proto3" json:"cod_lmt_rt,omitempty"`
	CodOrdrValidDt  string  `protobuf:"bytes,6,opt,name=cod_ordr_valid_dt,json=codOrdrValidDt,proto3" json:"cod_ordr_valid_dt,omitempty"`
	CodOrdrFlw      string  `protobuf:"bytes,7,opt,name=cod_ordr_flw,json=codOrdrFlw,proto3" json:"cod_ordr_flw,omitempty"`
	CodOrdrTotQty   int32   `protobuf:"varint,8,opt,name=cod_ordr_tot_qty,json=codOrdrTotQty,proto3" json:"cod_ordr_tot_qty,omitempty"`
	CodOrdrStts     string  `protobuf:"bytes,9,opt,name=cod_ordr_stts,json=codOrdrStts,proto3" json:"cod_ordr_stts,omitempty"`
	CcpOpnpstnQty   float32 `protobuf:"fixed32,10,opt,name=ccp_opnpstn_qty,json=ccpOpnpstnQty,proto3" json:"ccp_opnpstn_qty,omitempty"`
}

func (x *ComOrdrDtls) Reset() {
	*x = ComOrdrDtls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comordr_comOrderDetails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComOrdrDtls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComOrdrDtls) ProtoMessage() {}

func (x *ComOrdrDtls) ProtoReflect() protoreflect.Message {
	mi := &file_comordr_comOrderDetails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComOrdrDtls.ProtoReflect.Descriptor instead.
func (*ComOrdrDtls) Descriptor() ([]byte, []int) {
	return file_comordr_comOrderDetails_proto_rawDescGZIP(), []int{1}
}

func (x *ComOrdrDtls) GetCodClmMtchAccnt() string {
	if x != nil {
		return x.CodClmMtchAccnt
	}
	return ""
}

func (x *ComOrdrDtls) GetCodPrdctTyp() string {
	if x != nil {
		return x.CodPrdctTyp
	}
	return ""
}

func (x *ComOrdrDtls) GetCodUndrlyng() string {
	if x != nil {
		return x.CodUndrlyng
	}
	return ""
}

func (x *ComOrdrDtls) GetCodExpryDt() string {
	if x != nil {
		return x.CodExpryDt
	}
	return ""
}

func (x *ComOrdrDtls) GetCodLmtRt() float32 {
	if x != nil {
		return x.CodLmtRt
	}
	return 0
}

func (x *ComOrdrDtls) GetCodOrdrValidDt() string {
	if x != nil {
		return x.CodOrdrValidDt
	}
	return ""
}

func (x *ComOrdrDtls) GetCodOrdrFlw() string {
	if x != nil {
		return x.CodOrdrFlw
	}
	return ""
}

func (x *ComOrdrDtls) GetCodOrdrTotQty() int32 {
	if x != nil {
		return x.CodOrdrTotQty
	}
	return 0
}

func (x *ComOrdrDtls) GetCodOrdrStts() string {
	if x != nil {
		return x.CodOrdrStts
	}
	return ""
}

func (x *ComOrdrDtls) GetCcpOpnpstnQty() float32 {
	if x != nil {
		return x.CcpOpnpstnQty
	}
	return 0
}

type ComOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdDtls []*ComOrdrDtls `protobuf:"bytes,1,rep,name=ord_dtls,json=ordDtls,proto3" json:"ord_dtls,omitempty"`
}

func (x *ComOrderResponse) Reset() {
	*x = ComOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comordr_comOrderDetails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComOrderResponse) ProtoMessage() {}

func (x *ComOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comordr_comOrderDetails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComOrderResponse.ProtoReflect.Descriptor instead.
func (*ComOrderResponse) Descriptor() ([]byte, []int) {
	return file_comordr_comOrderDetails_proto_rawDescGZIP(), []int{2}
}

func (x *ComOrderResponse) GetOrdDtls() []*ComOrdrDtls {
	if x != nil {
		return x.OrdDtls
	}
	return nil
}

var File_comordr_comOrderDetails_proto protoreflect.FileDescriptor

var file_comordr_comOrderDetails_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x64, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x64, 0x72, 0x22, 0x3e, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x63,
	0x6f, 0x64, 0x5f, 0x63, 0x6c, 0x6d, 0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x64, 0x43, 0x6c, 0x6d, 0x4d,
	0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6e, 0x74, 0x22, 0x83, 0x03, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x4f, 0x72, 0x64, 0x72, 0x44, 0x74, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x6f, 0x64, 0x5f,
	0x63, 0x6c, 0x6d, 0x5f, 0x6d, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x64, 0x43, 0x6c, 0x6d, 0x4d, 0x74, 0x63, 0x68,
	0x41, 0x63, 0x63, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x64,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x64, 0x50, 0x72, 0x64, 0x63, 0x74, 0x54, 0x79, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x64,
	0x5f, 0x75, 0x6e, 0x64, 0x72, 0x6c, 0x79, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x64, 0x55, 0x6e, 0x64, 0x72, 0x6c, 0x79, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0c,
	0x63, 0x6f, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x79, 0x5f, 0x64, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x45, 0x78, 0x70, 0x72, 0x79, 0x44, 0x74, 0x12, 0x1c,
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x5f, 0x6c, 0x6d, 0x74, 0x5f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x4c, 0x6d, 0x74, 0x52, 0x74, 0x12, 0x29, 0x0a, 0x11,
	0x63, 0x6f, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x4f, 0x72, 0x64, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x5f, 0x6f,
	0x72, 0x64, 0x72, 0x5f, 0x66, 0x6c, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x46, 0x6c, 0x77, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x6f, 0x64,
	0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x64, 0x4f, 0x72, 0x64, 0x72, 0x54, 0x6f, 0x74, 0x51,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x72, 0x5f, 0x73,
	0x74, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x4f, 0x72,
	0x64, 0x72, 0x53, 0x74, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x63, 0x70, 0x5f, 0x6f, 0x70,
	0x6e, 0x70, 0x73, 0x74, 0x6e, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x63, 0x63, 0x70, 0x4f, 0x70, 0x6e, 0x70, 0x73, 0x74, 0x6e, 0x51, 0x74, 0x79, 0x22, 0x43,
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x74, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x64, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x72, 0x44, 0x74, 0x6c, 0x73, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x44,
	0x74, 0x6c, 0x73, 0x32, 0x54, 0x0a, 0x0e, 0x43, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x64, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x64, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comordr_comOrderDetails_proto_rawDescOnce sync.Once
	file_comordr_comOrderDetails_proto_rawDescData = file_comordr_comOrderDetails_proto_rawDesc
)

func file_comordr_comOrderDetails_proto_rawDescGZIP() []byte {
	file_comordr_comOrderDetails_proto_rawDescOnce.Do(func() {
		file_comordr_comOrderDetails_proto_rawDescData = protoimpl.X.CompressGZIP(file_comordr_comOrderDetails_proto_rawDescData)
	})
	return file_comordr_comOrderDetails_proto_rawDescData
}

var file_comordr_comOrderDetails_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_comordr_comOrderDetails_proto_goTypes = []any{
	(*ComOrderRequest)(nil),  // 0: comordr.ComOrderRequest
	(*ComOrdrDtls)(nil),      // 1: comordr.ComOrdrDtls
	(*ComOrderResponse)(nil), // 2: comordr.ComOrderResponse
}
var file_comordr_comOrderDetails_proto_depIdxs = []int32{
	1, // 0: comordr.ComOrderResponse.ord_dtls:type_name -> comordr.ComOrdrDtls
	0, // 1: comordr.C_OrderService.GetComOrder:input_type -> comordr.ComOrderRequest
	2, // 2: comordr.C_OrderService.GetComOrder:output_type -> comordr.ComOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comordr_comOrderDetails_proto_init() }
func file_comordr_comOrderDetails_proto_init() {
	if File_comordr_comOrderDetails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comordr_comOrderDetails_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ComOrderRequest); i {
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
		file_comordr_comOrderDetails_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ComOrdrDtls); i {
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
		file_comordr_comOrderDetails_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ComOrderResponse); i {
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
			RawDescriptor: file_comordr_comOrderDetails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comordr_comOrderDetails_proto_goTypes,
		DependencyIndexes: file_comordr_comOrderDetails_proto_depIdxs,
		MessageInfos:      file_comordr_comOrderDetails_proto_msgTypes,
	}.Build()
	File_comordr_comOrderDetails_proto = out.File
	file_comordr_comOrderDetails_proto_rawDesc = nil
	file_comordr_comOrderDetails_proto_goTypes = nil
	file_comordr_comOrderDetails_proto_depIdxs = nil
}
