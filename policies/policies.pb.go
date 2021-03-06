// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: policies/policies.proto

package policies

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PolicyState represent the status of a single policy
type PolicyState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Types that are assignable to Deployed:
	//	*PolicyState_Success
	//	*PolicyState_Error
	Deployed isPolicyState_Deployed `protobuf_oneof:"Deployed"`
}

func (x *PolicyState) Reset() {
	*x = PolicyState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_policies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyState) ProtoMessage() {}

func (x *PolicyState) ProtoReflect() protoreflect.Message {
	mi := &file_policies_policies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyState.ProtoReflect.Descriptor instead.
func (*PolicyState) Descriptor() ([]byte, []int) {
	return file_policies_policies_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *PolicyState) GetDeployed() isPolicyState_Deployed {
	if m != nil {
		return m.Deployed
	}
	return nil
}

func (x *PolicyState) GetSuccess() bool {
	if x, ok := x.GetDeployed().(*PolicyState_Success); ok {
		return x.Success
	}
	return false
}

func (x *PolicyState) GetError() string {
	if x, ok := x.GetDeployed().(*PolicyState_Error); ok {
		return x.Error
	}
	return ""
}

type isPolicyState_Deployed interface {
	isPolicyState_Deployed()
}

type PolicyState_Success struct {
	Success bool `protobuf:"varint,2,opt,name=Success,proto3,oneof"` // is true when new policy is deployed
}

type PolicyState_Error struct {
	Error string `protobuf:"bytes,3,opt,name=Error,proto3,oneof"` // is set when a policy fail deployment
}

func (*PolicyState_Success) isPolicyState_Deployed() {}

func (*PolicyState_Error) isPolicyState_Deployed() {}

// SnapshotStatus is reported by Fetcher in response to GetSnapshots
type SnapshotStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapshotName string         `protobuf:"bytes,1,opt,name=SnapshotName,proto3" json:"SnapshotName,omitempty"`
	PolicyStatus []*PolicyState `protobuf:"bytes,2,rep,name=PolicyStatus,proto3" json:"PolicyStatus,omitempty"`
}

func (x *SnapshotStatus) Reset() {
	*x = SnapshotStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_policies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotStatus) ProtoMessage() {}

func (x *SnapshotStatus) ProtoReflect() protoreflect.Message {
	mi := &file_policies_policies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotStatus.ProtoReflect.Descriptor instead.
func (*SnapshotStatus) Descriptor() ([]byte, []int) {
	return file_policies_policies_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotStatus) GetSnapshotName() string {
	if x != nil {
		return x.SnapshotName
	}
	return ""
}

func (x *SnapshotStatus) GetPolicyStatus() []*PolicyState {
	if x != nil {
		return x.PolicyStatus
	}
	return nil
}

// Snapshot represent the latest policy dump on the server
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Data [][]byte               `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policies_policies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_policies_policies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_policies_policies_proto_rawDescGZIP(), []int{2}
}

func (x *Snapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Snapshot) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Snapshot) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_policies_policies_proto protoreflect.FileDescriptor

var file_policies_policies_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x22, 0x6f, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x56, 0x0a, 0x0d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policies_policies_proto_rawDescOnce sync.Once
	file_policies_policies_proto_rawDescData = file_policies_policies_proto_rawDesc
)

func file_policies_policies_proto_rawDescGZIP() []byte {
	file_policies_policies_proto_rawDescOnce.Do(func() {
		file_policies_policies_proto_rawDescData = protoimpl.X.CompressGZIP(file_policies_policies_proto_rawDescData)
	})
	return file_policies_policies_proto_rawDescData
}

var file_policies_policies_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_policies_policies_proto_goTypes = []interface{}{
	(*PolicyState)(nil),           // 0: policies.PolicyState
	(*SnapshotStatus)(nil),        // 1: policies.SnapshotStatus
	(*Snapshot)(nil),              // 2: policies.Snapshot
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_policies_policies_proto_depIdxs = []int32{
	0, // 0: policies.SnapshotStatus.PolicyStatus:type_name -> policies.PolicyState
	3, // 1: policies.Snapshot.Time:type_name -> google.protobuf.Timestamp
	1, // 2: policies.PolicyService.StreamSnapshots:input_type -> policies.SnapshotStatus
	2, // 3: policies.PolicyService.StreamSnapshots:output_type -> policies.Snapshot
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_policies_policies_proto_init() }
func file_policies_policies_proto_init() {
	if File_policies_policies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policies_policies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyState); i {
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
		file_policies_policies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotStatus); i {
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
		file_policies_policies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
	file_policies_policies_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PolicyState_Success)(nil),
		(*PolicyState_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policies_policies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policies_policies_proto_goTypes,
		DependencyIndexes: file_policies_policies_proto_depIdxs,
		MessageInfos:      file_policies_policies_proto_msgTypes,
	}.Build()
	File_policies_policies_proto = out.File
	file_policies_policies_proto_rawDesc = nil
	file_policies_policies_proto_goTypes = nil
	file_policies_policies_proto_depIdxs = nil
}
