// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: flash_sale_submodule/auth.proto

package genproto

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

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	FullName    string `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	DateOfBirth string `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *RegisterReq) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role     string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Role         string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRes) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GetByEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetByEmail) Reset() {
	*x = GetByEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByEmail) ProtoMessage() {}

func (x *GetByEmail) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByEmail.ProtoReflect.Descriptor instead.
func (*GetByEmail) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{4}
}

func (x *GetByEmail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResetToken  string `protobuf:"bytes,1,opt,name=reset_token,json=resetToken,proto3" json:"reset_token,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ResetPassReq) Reset() {
	*x = ResetPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPassReq) ProtoMessage() {}

func (x *ResetPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPassReq.ProtoReflect.Descriptor instead.
func (*ResetPassReq) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ResetPassReq) GetResetToken() string {
	if x != nil {
		return x.ResetToken
	}
	return ""
}

func (x *ResetPassReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetPassReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ResetPassReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResetToken  string `protobuf:"bytes,1,opt,name=reset_token,json=resetToken,proto3" json:"reset_token,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ResetPassReqBody) Reset() {
	*x = ResetPassReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPassReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPassReqBody) ProtoMessage() {}

func (x *ResetPassReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPassReqBody.ProtoReflect.Descriptor instead.
func (*ResetPassReqBody) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ResetPassReqBody) GetResetToken() string {
	if x != nil {
		return x.ResetToken
	}
	return ""
}

func (x *ResetPassReqBody) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Code     string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{7}
}

func (x *Params) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Params) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Params) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Params) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Params) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RefToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *RefToken) Reset() {
	*x = RefToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefToken) ProtoMessage() {}

func (x *RefToken) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefToken.ProtoReflect.Descriptor instead.
func (*RefToken) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{8}
}

func (x *RefToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RefToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RefToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RefToken) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RefToken) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RefToken) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ListUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string      `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName   string      `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Pagination *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListUserReq) Reset() {
	*x = ListUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReq) ProtoMessage() {}

func (x *ListUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReq.ProtoReflect.Descriptor instead.
func (*ListUserReq) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{9}
}

func (x *ListUserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListUserReq) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *ListUserReq) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserRes `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Count int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListUserRes) Reset() {
	*x = ListUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_sale_submodule_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRes) ProtoMessage() {}

func (x *ListUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_flash_sale_submodule_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRes.ProtoReflect.Descriptor instead.
func (*ListUserRes) Descriptor() ([]byte, []int) {
	return file_flash_sale_submodule_auth_proto_rawDescGZIP(), []int{10}
}

func (x *ListUserRes) GetUsers() []*UserRes {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUserRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_flash_sale_submodule_auth_proto protoreflect.FileDescriptor

var file_flash_sale_submodule_auth_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x5f,
	0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x6c, 0x61,
	0x73, 0x68, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x22, 0x5c, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x66, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x68, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x56, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x76, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x79, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xde, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x10, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x35, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x45, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flash_sale_submodule_auth_proto_rawDescOnce sync.Once
	file_flash_sale_submodule_auth_proto_rawDescData = file_flash_sale_submodule_auth_proto_rawDesc
)

func file_flash_sale_submodule_auth_proto_rawDescGZIP() []byte {
	file_flash_sale_submodule_auth_proto_rawDescOnce.Do(func() {
		file_flash_sale_submodule_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_flash_sale_submodule_auth_proto_rawDescData)
	})
	return file_flash_sale_submodule_auth_proto_rawDescData
}

var file_flash_sale_submodule_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_flash_sale_submodule_auth_proto_goTypes = []any{
	(*RegisterReq)(nil),      // 0: proto.RegisterReq
	(*User)(nil),             // 1: proto.User
	(*LoginReq)(nil),         // 2: proto.LoginReq
	(*LoginRes)(nil),         // 3: proto.LoginRes
	(*GetByEmail)(nil),       // 4: proto.GetByEmail
	(*ResetPassReq)(nil),     // 5: proto.ResetPassReq
	(*ResetPassReqBody)(nil), // 6: proto.ResetPassReqBody
	(*Params)(nil),           // 7: proto.Params
	(*RefToken)(nil),         // 8: proto.RefToken
	(*ListUserReq)(nil),      // 9: proto.ListUserReq
	(*ListUserRes)(nil),      // 10: proto.ListUserRes
	(*Pagination)(nil),       // 11: proto.Pagination
	(*UserRes)(nil),          // 12: proto.UserRes
	(*GetById)(nil),          // 13: proto.GetById
	(*Void)(nil),             // 14: proto.Void
}
var file_flash_sale_submodule_auth_proto_depIdxs = []int32{
	11, // 0: proto.ListUserReq.pagination:type_name -> proto.Pagination
	12, // 1: proto.ListUserRes.users:type_name -> proto.UserRes
	0,  // 2: proto.AuthService.Register:input_type -> proto.RegisterReq
	2,  // 3: proto.AuthService.Login:input_type -> proto.LoginReq
	4,  // 4: proto.AuthService.ForgotPassword:input_type -> proto.GetByEmail
	5,  // 5: proto.AuthService.ResetPassword:input_type -> proto.ResetPassReq
	8,  // 6: proto.AuthService.SaveRefreshToken:input_type -> proto.RefToken
	9,  // 7: proto.AuthService.GetAllUsers:input_type -> proto.ListUserReq
	13, // 8: proto.AuthService.GEtUserById:input_type -> proto.GetById
	14, // 9: proto.AuthService.Register:output_type -> proto.Void
	1,  // 10: proto.AuthService.Login:output_type -> proto.User
	14, // 11: proto.AuthService.ForgotPassword:output_type -> proto.Void
	14, // 12: proto.AuthService.ResetPassword:output_type -> proto.Void
	14, // 13: proto.AuthService.SaveRefreshToken:output_type -> proto.Void
	10, // 14: proto.AuthService.GetAllUsers:output_type -> proto.ListUserRes
	12, // 15: proto.AuthService.GEtUserById:output_type -> proto.UserRes
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_flash_sale_submodule_auth_proto_init() }
func file_flash_sale_submodule_auth_proto_init() {
	if File_flash_sale_submodule_auth_proto != nil {
		return
	}
	file_flash_sale_submodule_common_proto_init()
	file_flash_sale_submodule_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flash_sale_submodule_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterReq); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoginReq); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRes); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetByEmail); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ResetPassReq); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ResetPassReqBody); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Params); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RefToken); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserReq); i {
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
		file_flash_sale_submodule_auth_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserRes); i {
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
			RawDescriptor: file_flash_sale_submodule_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flash_sale_submodule_auth_proto_goTypes,
		DependencyIndexes: file_flash_sale_submodule_auth_proto_depIdxs,
		MessageInfos:      file_flash_sale_submodule_auth_proto_msgTypes,
	}.Build()
	File_flash_sale_submodule_auth_proto = out.File
	file_flash_sale_submodule_auth_proto_rawDesc = nil
	file_flash_sale_submodule_auth_proto_goTypes = nil
	file_flash_sale_submodule_auth_proto_depIdxs = nil
}
