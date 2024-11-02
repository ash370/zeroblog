// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: like.proto

package service

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

type ThumbupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId    string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`        // 业务id
	ObjId    int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`       // 点赞对象id
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
	LikeType int32  `protobuf:"varint,4,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *ThumbupRequest) Reset() {
	*x = ThumbupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbupRequest) ProtoMessage() {}

func (x *ThumbupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbupRequest.ProtoReflect.Descriptor instead.
func (*ThumbupRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *ThumbupRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *ThumbupRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *ThumbupRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ThumbupRequest) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

type ThumbupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId      string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`            // 业务id
	ObjId      int64  `protobuf:"varint,2,opt,name=objId,proto3" json:"objId,omitempty"`           // 点赞对象id
	LikeNum    int64  `protobuf:"varint,3,opt,name=likeNum,proto3" json:"likeNum,omitempty"`       // 点赞数
	DislikeNum int64  `protobuf:"varint,4,opt,name=dislikeNum,proto3" json:"dislikeNum,omitempty"` // 点踩数
}

func (x *ThumbupResponse) Reset() {
	*x = ThumbupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbupResponse) ProtoMessage() {}

func (x *ThumbupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbupResponse.ProtoReflect.Descriptor instead.
func (*ThumbupResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *ThumbupResponse) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *ThumbupResponse) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *ThumbupResponse) GetLikeNum() int64 {
	if x != nil {
		return x.LikeNum
	}
	return 0
}

func (x *ThumbupResponse) GetDislikeNum() int64 {
	if x != nil {
		return x.DislikeNum
	}
	return 0
}

type IsThumbupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId    string `protobuf:"bytes,1,opt,name=bizId,proto3" json:"bizId,omitempty"`        // 业务id
	TargetId int64  `protobuf:"varint,2,opt,name=targetId,proto3" json:"targetId,omitempty"` // 点赞对象id
	UserId   int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`     // 用户id
}

func (x *IsThumbupRequest) Reset() {
	*x = IsThumbupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsThumbupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsThumbupRequest) ProtoMessage() {}

func (x *IsThumbupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsThumbupRequest.ProtoReflect.Descriptor instead.
func (*IsThumbupRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *IsThumbupRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *IsThumbupRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *IsThumbupRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsThumbupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserThumbups map[int64]*UserThumbup `protobuf:"bytes,1,rep,name=userThumbups,proto3" json:"userThumbups,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IsThumbupResponse) Reset() {
	*x = IsThumbupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsThumbupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsThumbupResponse) ProtoMessage() {}

func (x *IsThumbupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsThumbupResponse.ProtoReflect.Descriptor instead.
func (*IsThumbupResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *IsThumbupResponse) GetUserThumbups() map[int64]*UserThumbup {
	if x != nil {
		return x.UserThumbups
	}
	return nil
}

type UserThumbup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ThumbupTime int64 `protobuf:"varint,2,opt,name=thumbupTime,proto3" json:"thumbupTime,omitempty"`
	LikeType    int32 `protobuf:"varint,3,opt,name=likeType,proto3" json:"likeType,omitempty"` // 类型
}

func (x *UserThumbup) Reset() {
	*x = UserThumbup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserThumbup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserThumbup) ProtoMessage() {}

func (x *UserThumbup) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserThumbup.ProtoReflect.Descriptor instead.
func (*UserThumbup) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *UserThumbup) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserThumbup) GetThumbupTime() int64 {
	if x != nil {
		return x.ThumbupTime
	}
	return 0
}

func (x *UserThumbup) GetLikeType() int32 {
	if x != nil {
		return x.LikeType
	}
	return 0
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x70, 0x0a, 0x0e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62,
	0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x77, 0x0a, 0x0f, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69,
	0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x4e, 0x75, 0x6d,
	0x22, 0x5c, 0x0a, 0x10, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xbc,
	0x01, 0x0a, 0x11, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x75, 0x70, 0x73, 0x1a, 0x55, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x75, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x32, 0x88, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x49, 0x73, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x73, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_like_proto_goTypes = []any{
	(*ThumbupRequest)(nil),    // 0: service.ThumbupRequest
	(*ThumbupResponse)(nil),   // 1: service.ThumbupResponse
	(*IsThumbupRequest)(nil),  // 2: service.IsThumbupRequest
	(*IsThumbupResponse)(nil), // 3: service.IsThumbupResponse
	(*UserThumbup)(nil),       // 4: service.UserThumbup
	nil,                       // 5: service.IsThumbupResponse.UserThumbupsEntry
}
var file_like_proto_depIdxs = []int32{
	5, // 0: service.IsThumbupResponse.userThumbups:type_name -> service.IsThumbupResponse.UserThumbupsEntry
	4, // 1: service.IsThumbupResponse.UserThumbupsEntry.value:type_name -> service.UserThumbup
	0, // 2: service.Like.Thumbup:input_type -> service.ThumbupRequest
	2, // 3: service.Like.IsThumbup:input_type -> service.IsThumbupRequest
	1, // 4: service.Like.Thumbup:output_type -> service.ThumbupResponse
	3, // 5: service.Like.IsThumbup:output_type -> service.IsThumbupResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbupRequest); i {
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
		file_like_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbupResponse); i {
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
		file_like_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IsThumbupRequest); i {
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
		file_like_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IsThumbupResponse); i {
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
		file_like_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserThumbup); i {
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
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}
