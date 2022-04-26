// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: home/home.proto

package home

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	person "internal/grpc_class/pb/person"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Home struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*person.Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	Visitor *Home_Visitor    `protobuf:"bytes,2,opt,name=visitor,proto3" json:"visitor,omitempty"`
}

func (x *Home) Reset() {
	*x = Home{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home) ProtoMessage() {}

func (x *Home) ProtoReflect() protoreflect.Message {
	mi := &file_home_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home.ProtoReflect.Descriptor instead.
func (*Home) Descriptor() ([]byte, []int) {
	return file_home_home_proto_rawDescGZIP(), []int{0}
}

func (x *Home) GetPersons() []*person.Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *Home) GetVisitor() *Home_Visitor {
	if x != nil {
		return x.Visitor
	}
	return nil
}

type Home_Visitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Home_Visitor) Reset() {
	*x = Home_Visitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home_Visitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home_Visitor) ProtoMessage() {}

func (x *Home_Visitor) ProtoReflect() protoreflect.Message {
	mi := &file_home_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home_Visitor.ProtoReflect.Descriptor instead.
func (*Home_Visitor) Descriptor() ([]byte, []int) {
	return file_home_home_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Home_Visitor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_home_home_proto protoreflect.FileDescriptor

var file_home_home_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x1a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x04,
	0x48, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x2c,
	0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x0a, 0x07,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x3b, 0x68, 0x6f, 0x6d, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_home_proto_rawDescOnce sync.Once
	file_home_home_proto_rawDescData = file_home_home_proto_rawDesc
)

func file_home_home_proto_rawDescGZIP() []byte {
	file_home_home_proto_rawDescOnce.Do(func() {
		file_home_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_home_proto_rawDescData)
	})
	return file_home_home_proto_rawDescData
}

var file_home_home_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_home_home_proto_goTypes = []interface{}{
	(*Home)(nil),          // 0: home.Home
	(*Home_Visitor)(nil),  // 1: home.Home.Visitor
	(*person.Person)(nil), // 2: person.Person
}
var file_home_home_proto_depIdxs = []int32{
	2, // 0: home.Home.persons:type_name -> person.Person
	1, // 1: home.Home.visitor:type_name -> home.Home.Visitor
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_home_home_proto_init() }
func file_home_home_proto_init() {
	if File_home_home_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_home_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home); i {
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
		file_home_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home_Visitor); i {
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
			RawDescriptor: file_home_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_home_home_proto_goTypes,
		DependencyIndexes: file_home_home_proto_depIdxs,
		MessageInfos:      file_home_home_proto_msgTypes,
	}.Build()
	File_home_home_proto = out.File
	file_home_home_proto_rawDesc = nil
	file_home_home_proto_goTypes = nil
	file_home_home_proto_depIdxs = nil
}
