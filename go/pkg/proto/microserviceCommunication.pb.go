// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: microserviceCommunication.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MicroserviceCommunication struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Type            string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	RequestType     string                 `protobuf:"bytes,2,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	Data            *structpb.Struct       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Metadata        map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OriginalRequest *anypb.Any             `protobuf:"bytes,5,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`
	RequestMetadata *RequestMetadata       `protobuf:"bytes,6,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	Traces          map[string][]byte      `protobuf:"bytes,7,rep,name=traces,proto3" json:"traces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Binary or textual representation of span context
	Result          []byte                 `protobuf:"bytes,8,opt,name=result,proto3" json:"result,omitempty"`
	RoutingData     []string               `protobuf:"bytes,9,rep,name=routing_data,json=routingData,proto3" json:"routing_data,omitempty"` // To be used for persistent jobs
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MicroserviceCommunication) Reset() {
	*x = MicroserviceCommunication{}
	mi := &file_microserviceCommunication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MicroserviceCommunication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MicroserviceCommunication) ProtoMessage() {}

func (x *MicroserviceCommunication) ProtoReflect() protoreflect.Message {
	mi := &file_microserviceCommunication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MicroserviceCommunication.ProtoReflect.Descriptor instead.
func (*MicroserviceCommunication) Descriptor() ([]byte, []int) {
	return file_microserviceCommunication_proto_rawDescGZIP(), []int{0}
}

func (x *MicroserviceCommunication) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MicroserviceCommunication) GetRequestType() string {
	if x != nil {
		return x.RequestType
	}
	return ""
}

func (x *MicroserviceCommunication) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MicroserviceCommunication) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MicroserviceCommunication) GetOriginalRequest() *anypb.Any {
	if x != nil {
		return x.OriginalRequest
	}
	return nil
}

func (x *MicroserviceCommunication) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

func (x *MicroserviceCommunication) GetTraces() map[string][]byte {
	if x != nil {
		return x.Traces
	}
	return nil
}

func (x *MicroserviceCommunication) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *MicroserviceCommunication) GetRoutingData() []string {
	if x != nil {
		return x.RoutingData
	}
	return nil
}

type ContinueReceiving struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ContinueReceiving bool                   `protobuf:"varint,1,opt,name=continue_receiving,json=continueReceiving,proto3" json:"continue_receiving,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ContinueReceiving) Reset() {
	*x = ContinueReceiving{}
	mi := &file_microserviceCommunication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContinueReceiving) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinueReceiving) ProtoMessage() {}

func (x *ContinueReceiving) ProtoReflect() protoreflect.Message {
	mi := &file_microserviceCommunication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinueReceiving.ProtoReflect.Descriptor instead.
func (*ContinueReceiving) Descriptor() ([]byte, []int) {
	return file_microserviceCommunication_proto_rawDescGZIP(), []int{1}
}

func (x *ContinueReceiving) GetContinueReceiving() bool {
	if x != nil {
		return x.ContinueReceiving
	}
	return false
}

var File_microserviceCommunication_proto protoreflect.FileDescriptor

const file_microserviceCommunication_proto_rawDesc = "" +
	"\n" +
	"\x1fmicroserviceCommunication.proto\x12\adynamos\x1a\x1cgoogle/protobuf/struct.proto\x1a\x19google/protobuf/any.proto\x1a\rgeneric.proto\"\xce\x04\n" +
	"\x19MicroserviceCommunication\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12!\n" +
	"\frequest_type\x18\x02 \x01(\tR\vrequestType\x12+\n" +
	"\x04data\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04data\x12L\n" +
	"\bmetadata\x18\x04 \x03(\v20.dynamos.MicroserviceCommunication.MetadataEntryR\bmetadata\x12?\n" +
	"\x10original_request\x18\x05 \x01(\v2\x14.google.protobuf.AnyR\x0foriginalRequest\x12C\n" +
	"\x10request_metadata\x18\x06 \x01(\v2\x18.dynamos.RequestMetadataR\x0frequestMetadata\x12F\n" +
	"\x06traces\x18\a \x03(\v2..dynamos.MicroserviceCommunication.TracesEntryR\x06traces\x12\x16\n" +
	"\x06result\x18\b \x01(\fR\x06result\x12!\n" +
	"\frouting_data\x18\t \x03(\tR\vroutingData\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a9\n" +
	"\vTracesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value:\x028\x01\"B\n" +
	"\x11ContinueReceiving\x12-\n" +
	"\x12continue_receiving\x18\x01 \x01(\bR\x11continueReceiving2\\\n" +
	"\fMicroservice\x12L\n" +
	"\bSendData\x12\".dynamos.MicroserviceCommunication\x1a\x1a.dynamos.ContinueReceiving\"\x00B'Z%github.com/Jorrit05/DYNAMOS/pkg/protob\x06proto3"

var (
	file_microserviceCommunication_proto_rawDescOnce sync.Once
	file_microserviceCommunication_proto_rawDescData []byte
)

func file_microserviceCommunication_proto_rawDescGZIP() []byte {
	file_microserviceCommunication_proto_rawDescOnce.Do(func() {
		file_microserviceCommunication_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_microserviceCommunication_proto_rawDesc), len(file_microserviceCommunication_proto_rawDesc)))
	})
	return file_microserviceCommunication_proto_rawDescData
}

var file_microserviceCommunication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_microserviceCommunication_proto_goTypes = []any{
	(*MicroserviceCommunication)(nil), // 0: dynamos.MicroserviceCommunication
	(*ContinueReceiving)(nil),         // 1: dynamos.ContinueReceiving
	nil,                               // 2: dynamos.MicroserviceCommunication.MetadataEntry
	nil,                               // 3: dynamos.MicroserviceCommunication.TracesEntry
	(*structpb.Struct)(nil),           // 4: google.protobuf.Struct
	(*anypb.Any)(nil),                 // 5: google.protobuf.Any
	(*RequestMetadata)(nil),           // 6: dynamos.RequestMetadata
}
var file_microserviceCommunication_proto_depIdxs = []int32{
	4, // 0: dynamos.MicroserviceCommunication.data:type_name -> google.protobuf.Struct
	2, // 1: dynamos.MicroserviceCommunication.metadata:type_name -> dynamos.MicroserviceCommunication.MetadataEntry
	5, // 2: dynamos.MicroserviceCommunication.original_request:type_name -> google.protobuf.Any
	6, // 3: dynamos.MicroserviceCommunication.request_metadata:type_name -> dynamos.RequestMetadata
	3, // 4: dynamos.MicroserviceCommunication.traces:type_name -> dynamos.MicroserviceCommunication.TracesEntry
	0, // 5: dynamos.Microservice.SendData:input_type -> dynamos.MicroserviceCommunication
	1, // 6: dynamos.Microservice.SendData:output_type -> dynamos.ContinueReceiving
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_microserviceCommunication_proto_init() }
func file_microserviceCommunication_proto_init() {
	if File_microserviceCommunication_proto != nil {
		return
	}
	file_generic_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_microserviceCommunication_proto_rawDesc), len(file_microserviceCommunication_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microserviceCommunication_proto_goTypes,
		DependencyIndexes: file_microserviceCommunication_proto_depIdxs,
		MessageInfos:      file_microserviceCommunication_proto_msgTypes,
	}.Build()
	File_microserviceCommunication_proto = out.File
	file_microserviceCommunication_proto_goTypes = nil
	file_microserviceCommunication_proto_depIdxs = nil
}
