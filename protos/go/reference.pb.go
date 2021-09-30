// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: reference.proto

package didar

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

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousDocumentId string `protobuf:"bytes,1,opt,name=previous_document_id,json=previousDocumentId,proto3" json:"previous_document_id,omitempty"`
	DataHash           string `protobuf:"bytes,2,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	SigningKey         *Key   `protobuf:"bytes,3,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	Signature          string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_reference_proto_rawDescGZIP(), []int{0}
}

func (x *Reference) GetPreviousDocumentId() string {
	if x != nil {
		return x.PreviousDocumentId
	}
	return ""
}

func (x *Reference) GetDataHash() string {
	if x != nil {
		return x.DataHash
	}
	return ""
}

func (x *Reference) GetSigningKey() *Key {
	if x != nil {
		return x.SigningKey
	}
	return nil
}

func (x *Reference) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_reference_proto protoreflect.FileDescriptor

var file_reference_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x64, 0x69, 0x64, 0x61, 0x72, 0x1a, 0x09, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x2b, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x69, 0x64, 0x61, 0x72, 0x2e, 0x6b, 0x65,
	0x79, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x3b, 0x64, 0x69, 0x64, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reference_proto_rawDescOnce sync.Once
	file_reference_proto_rawDescData = file_reference_proto_rawDesc
)

func file_reference_proto_rawDescGZIP() []byte {
	file_reference_proto_rawDescOnce.Do(func() {
		file_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_reference_proto_rawDescData)
	})
	return file_reference_proto_rawDescData
}

var file_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reference_proto_goTypes = []interface{}{
	(*Reference)(nil), // 0: didar.reference
	(*Key)(nil),       // 1: didar.key
}
var file_reference_proto_depIdxs = []int32{
	1, // 0: didar.reference.signing_key:type_name -> didar.key
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reference_proto_init() }
func file_reference_proto_init() {
	if File_reference_proto != nil {
		return
	}
	file_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reference_proto_goTypes,
		DependencyIndexes: file_reference_proto_depIdxs,
		MessageInfos:      file_reference_proto_msgTypes,
	}.Build()
	File_reference_proto = out.File
	file_reference_proto_rawDesc = nil
	file_reference_proto_goTypes = nil
	file_reference_proto_depIdxs = nil
}
