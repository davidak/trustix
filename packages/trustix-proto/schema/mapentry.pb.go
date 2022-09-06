// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: schema/mapentry.proto

package schema

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

type MapEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value digest of tree node
	Digest []byte `protobuf:"bytes,1,req,name=Digest" json:"Digest,omitempty"`
	// Index of value in log
	Index *uint64 `protobuf:"varint,2,req,name=Index" json:"Index,omitempty"`
}

func (x *MapEntry) Reset() {
	*x = MapEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mapentry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapEntry) ProtoMessage() {}

func (x *MapEntry) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mapentry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapEntry.ProtoReflect.Descriptor instead.
func (*MapEntry) Descriptor() ([]byte, []int) {
	return file_schema_mapentry_proto_rawDescGZIP(), []int{0}
}

func (x *MapEntry) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *MapEntry) GetIndex() uint64 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

var File_schema_mapentry_proto protoreflect.FileDescriptor

var file_schema_mapentry_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x69, 0x78, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x69, 0x78, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61,
}

var (
	file_schema_mapentry_proto_rawDescOnce sync.Once
	file_schema_mapentry_proto_rawDescData = file_schema_mapentry_proto_rawDesc
)

func file_schema_mapentry_proto_rawDescGZIP() []byte {
	file_schema_mapentry_proto_rawDescOnce.Do(func() {
		file_schema_mapentry_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_mapentry_proto_rawDescData)
	})
	return file_schema_mapentry_proto_rawDescData
}

var file_schema_mapentry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_mapentry_proto_goTypes = []interface{}{
	(*MapEntry)(nil), // 0: MapEntry
}
var file_schema_mapentry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_mapentry_proto_init() }
func file_schema_mapentry_proto_init() {
	if File_schema_mapentry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_mapentry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapEntry); i {
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
			RawDescriptor: file_schema_mapentry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_mapentry_proto_goTypes,
		DependencyIndexes: file_schema_mapentry_proto_depIdxs,
		MessageInfos:      file_schema_mapentry_proto_msgTypes,
	}.Build()
	File_schema_mapentry_proto = out.File
	file_schema_mapentry_proto_rawDesc = nil
	file_schema_mapentry_proto_goTypes = nil
	file_schema_mapentry_proto_depIdxs = nil
}
