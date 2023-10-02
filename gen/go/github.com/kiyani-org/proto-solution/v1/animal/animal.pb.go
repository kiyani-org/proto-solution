// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: protos/v1/animal_service/animal.proto

package animalpb

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

type Animal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classification string `protobuf:"bytes,1,opt,name=classification,proto3" json:"classification,omitempty"`
	Id             int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Foo            string `protobuf:"bytes,4,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *Animal) Reset() {
	*x = Animal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_animal_service_animal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Animal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Animal) ProtoMessage() {}

func (x *Animal) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_animal_service_animal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Animal.ProtoReflect.Descriptor instead.
func (*Animal) Descriptor() ([]byte, []int) {
	return file_protos_v1_animal_service_animal_proto_rawDescGZIP(), []int{0}
}

func (x *Animal) GetClassification() string {
	if x != nil {
		return x.Classification
	}
	return ""
}

func (x *Animal) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Animal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Animal) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

var File_protos_v1_animal_service_animal_proto protoreflect.FileDescriptor

var file_protos_v1_animal_service_animal_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x22, 0x66, 0x0a, 0x06, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x42, 0x59, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x6f, 0x6f, 0x6c,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x79, 0x61, 0x6e, 0x69, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x3b, 0x61, 0x6e, 0x69,
	0x6d, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_animal_service_animal_proto_rawDescOnce sync.Once
	file_protos_v1_animal_service_animal_proto_rawDescData = file_protos_v1_animal_service_animal_proto_rawDesc
)

func file_protos_v1_animal_service_animal_proto_rawDescGZIP() []byte {
	file_protos_v1_animal_service_animal_proto_rawDescOnce.Do(func() {
		file_protos_v1_animal_service_animal_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_animal_service_animal_proto_rawDescData)
	})
	return file_protos_v1_animal_service_animal_proto_rawDescData
}

var file_protos_v1_animal_service_animal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_v1_animal_service_animal_proto_goTypes = []interface{}{
	(*Animal)(nil), // 0: animal.v1.Animal
}
var file_protos_v1_animal_service_animal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_v1_animal_service_animal_proto_init() }
func file_protos_v1_animal_service_animal_proto_init() {
	if File_protos_v1_animal_service_animal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_animal_service_animal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Animal); i {
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
			RawDescriptor: file_protos_v1_animal_service_animal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_v1_animal_service_animal_proto_goTypes,
		DependencyIndexes: file_protos_v1_animal_service_animal_proto_depIdxs,
		MessageInfos:      file_protos_v1_animal_service_animal_proto_msgTypes,
	}.Build()
	File_protos_v1_animal_service_animal_proto = out.File
	file_protos_v1_animal_service_animal_proto_rawDesc = nil
	file_protos_v1_animal_service_animal_proto_goTypes = nil
	file_protos_v1_animal_service_animal_proto_depIdxs = nil
}
