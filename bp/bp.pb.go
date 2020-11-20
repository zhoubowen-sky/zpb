// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: bp.proto

package bp

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Eb int32

const (
	Eb_Success Eb = 0
	Eb_Failed  Eb = 1
	Eb_Unknown Eb = 2
)

// Enum value maps for Eb.
var (
	Eb_name = map[int32]string{
		0: "Success",
		1: "Failed",
		2: "Unknown",
	}
	Eb_value = map[string]int32{
		"Success": 0,
		"Failed":  1,
		"Unknown": 2,
	}
)

func (x Eb) Enum() *Eb {
	p := new(Eb)
	*p = x
	return p
}

func (x Eb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Eb) Descriptor() protoreflect.EnumDescriptor {
	return file_bp_proto_enumTypes[0].Descriptor()
}

func (Eb) Type() protoreflect.EnumType {
	return &file_bp_proto_enumTypes[0]
}

func (x Eb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Eb.Descriptor instead.
func (Eb) EnumDescriptor() ([]byte, []int) {
	return file_bp_proto_rawDescGZIP(), []int{0}
}

type SbReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNum int32  `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
}

func (x *SbReq) Reset() {
	*x = SbReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SbReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SbReq) ProtoMessage() {}

func (x *SbReq) ProtoReflect() protoreflect.Message {
	mi := &file_bp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SbReq.ProtoReflect.Descriptor instead.
func (*SbReq) Descriptor() ([]byte, []int) {
	return file_bp_proto_rawDescGZIP(), []int{0}
}

func (x *SbReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SbReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type SbRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Eb     Eb    `protobuf:"varint,2,opt,name=eb,proto3,enum=bp.Eb" json:"eb,omitempty"`
}

func (x *SbRsp) Reset() {
	*x = SbRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SbRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SbRsp) ProtoMessage() {}

func (x *SbRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SbRsp.ProtoReflect.Descriptor instead.
func (*SbRsp) Descriptor() ([]byte, []int) {
	return file_bp_proto_rawDescGZIP(), []int{1}
}

func (x *SbRsp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SbRsp) GetEb() Eb {
	if x != nil {
		return x.Eb
	}
	return Eb_Success
}

var File_bp_proto protoreflect.FileDescriptor

var file_bp_proto_rawDesc = []byte{
	0x0a, 0x08, 0x62, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x62, 0x70, 0x22, 0x38,
	0x0a, 0x05, 0x53, 0x62, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x37, 0x0a, 0x05, 0x53, 0x62, 0x52, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x02, 0x65, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x62, 0x70, 0x2e, 0x45, 0x62, 0x52, 0x02, 0x65,
	0x62, 0x2a, 0x2a, 0x0a, 0x02, 0x45, 0x62, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x02, 0x32, 0x21, 0x0a,
	0x02, 0x53, 0x61, 0x12, 0x1b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x09, 0x2e, 0x62, 0x70, 0x2e,
	0x53, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x62, 0x70, 0x2e, 0x53, 0x62, 0x52, 0x73, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bp_proto_rawDescOnce sync.Once
	file_bp_proto_rawDescData = file_bp_proto_rawDesc
)

func file_bp_proto_rawDescGZIP() []byte {
	file_bp_proto_rawDescOnce.Do(func() {
		file_bp_proto_rawDescData = protoimpl.X.CompressGZIP(file_bp_proto_rawDescData)
	})
	return file_bp_proto_rawDescData
}

var file_bp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bp_proto_goTypes = []interface{}{
	(Eb)(0),       // 0: bp.Eb
	(*SbReq)(nil), // 1: bp.SbReq
	(*SbRsp)(nil), // 2: bp.SbRsp
}
var file_bp_proto_depIdxs = []int32{
	0, // 0: bp.SbRsp.eb:type_name -> bp.Eb
	1, // 1: bp.Sa.Get:input_type -> bp.SbReq
	2, // 2: bp.Sa.Get:output_type -> bp.SbRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bp_proto_init() }
func file_bp_proto_init() {
	if File_bp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SbReq); i {
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
		file_bp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SbRsp); i {
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
			RawDescriptor: file_bp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bp_proto_goTypes,
		DependencyIndexes: file_bp_proto_depIdxs,
		EnumInfos:         file_bp_proto_enumTypes,
		MessageInfos:      file_bp_proto_msgTypes,
	}.Build()
	File_bp_proto = out.File
	file_bp_proto_rawDesc = nil
	file_bp_proto_goTypes = nil
	file_bp_proto_depIdxs = nil
}
