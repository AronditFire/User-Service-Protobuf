// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: user_service.proto

package uservicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Логика админа
type Roles int32

const (
	Roles_UNKNOWN Roles = 0
	Roles_BUYER   Roles = 1
	Roles_ADMIN   Roles = 2
)

// Enum value maps for Roles.
var (
	Roles_name = map[int32]string{
		0: "UNKNOWN",
		1: "BUYER",
		2: "ADMIN",
	}
	Roles_value = map[string]int32{
		"UNKNOWN": 0,
		"BUYER":   1,
		"ADMIN":   2,
	}
)

func (x Roles) Enum() *Roles {
	p := new(Roles)
	*p = x
	return p
}

func (x Roles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Roles) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_proto_enumTypes[0].Descriptor()
}

func (Roles) Type() protoreflect.EnumType {
	return &file_user_service_proto_enumTypes[0]
}

func (x Roles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Roles.Descriptor instead.
func (Roles) EnumDescriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

// Логика авторизации
type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FIO           string                 `protobuf:"bytes,3,opt,name=FIO,proto3" json:"FIO,omitempty"` // Фамилия, имя, отчество
	PhoneNumber   string                 `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetFIO() string {
	if x != nil {
		return x.FIO
	}
	return ""
}

func (x *RegisterRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Tokens struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	mi := &file_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *Tokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Tokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tokens        *Tokens                `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type RefreshRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	mi := &file_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tokens        *Tokens                `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshResponse) Reset() {
	*x = RefreshResponse{}
	mi := &file_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshResponse) ProtoMessage() {}

func (x *RefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshResponse.ProtoReflect.Descriptor instead.
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshResponse) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // удаляем указанный refresh token
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_user_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *LogoutRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Логика профиля пользователя
type GetProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_user_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetProfileRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FIO           string                 `protobuf:"bytes,4,opt,name=FIO,proto3" json:"FIO,omitempty"` // Фамилия, имя, отчество
	PhoneNumber   string                 `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfileResponse) Reset() {
	*x = UserProfileResponse{}
	mi := &file_user_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileResponse) ProtoMessage() {}

func (x *UserProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileResponse.ProtoReflect.Descriptor instead.
func (*UserProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *UserProfileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserProfileResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserProfileResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfileResponse) GetFIO() string {
	if x != nil {
		return x.FIO
	}
	return ""
}

func (x *UserProfileResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type UserListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserProfileResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // Список пользователей
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	mi := &file_user_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *UserListResponse) GetUsers() []*UserProfileResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type AdminRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          Roles                  `protobuf:"varint,2,opt,name=role,proto3,enum=user_profile.Roles" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdminRoleRequest) Reset() {
	*x = AdminRoleRequest{}
	mi := &file_user_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdminRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRoleRequest) ProtoMessage() {}

func (x *AdminRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRoleRequest.ProtoReflect.Descriptor instead.
func (*AdminRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{11}
}

func (x *AdminRoleRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AdminRoleRequest) GetRole() Roles {
	if x != nil {
		return x.Role
	}
	return Roles_UNKNOWN
}

var File_user_service_proto protoreflect.FileDescriptor

const file_user_service_proto_rawDesc = "" +
	"\n" +
	"\x12user_service.proto\x12\fuser_profile\x1a\x1bgoogle/protobuf/empty.proto\"\x93\x01\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x10\n" +
	"\x03FIO\x18\x03 \x01(\tR\x03FIO\x12 \n" +
	"\vphoneNumber\x18\x04 \x01(\tR\vphoneNumber\x12\x1a\n" +
	"\bpassword\x18\x05 \x01(\tR\bpassword\"+\n" +
	"\x10RegisterResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"F\n" +
	"\fLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"P\n" +
	"\x06Tokens\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\"=\n" +
	"\rLoginResponse\x12,\n" +
	"\x06tokens\x18\x01 \x01(\v2\x14.user_profile.TokensR\x06tokens\"5\n" +
	"\x0eRefreshRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"?\n" +
	"\x0fRefreshResponse\x12,\n" +
	"\x06tokens\x18\x01 \x01(\v2\x14.user_profile.TokensR\x06tokens\"4\n" +
	"\rLogoutRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\",\n" +
	"\x11GetProfileRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"\x8b\x01\n" +
	"\x13UserProfileResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x10\n" +
	"\x03FIO\x18\x04 \x01(\tR\x03FIO\x12 \n" +
	"\vphoneNumber\x18\x05 \x01(\tR\vphoneNumber\"K\n" +
	"\x10UserListResponse\x127\n" +
	"\x05users\x18\x01 \x03(\v2!.user_profile.UserProfileResponseR\x05users\"T\n" +
	"\x10AdminRoleRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\x12'\n" +
	"\x04role\x18\x02 \x01(\x0e2\x13.user_profile.RolesR\x04role**\n" +
	"\x05Roles\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\t\n" +
	"\x05BUYER\x10\x01\x12\t\n" +
	"\x05ADMIN\x10\x022\x83\x04\n" +
	"\vUserService\x12I\n" +
	"\bRegister\x12\x1d.user_profile.RegisterRequest\x1a\x1e.user_profile.RegisterResponse\x12@\n" +
	"\x05Login\x12\x1a.user_profile.LoginRequest\x1a\x1b.user_profile.LoginResponse\x12K\n" +
	"\fRefreshToken\x12\x1c.user_profile.RefreshRequest\x1a\x1d.user_profile.RefreshResponse\x12=\n" +
	"\x06Logout\x12\x1b.user_profile.LogoutRequest\x1a\x16.google.protobuf.Empty\x12P\n" +
	"\n" +
	"GetProfile\x12\x1f.user_profile.GetProfileRequest\x1a!.user_profile.UserProfileResponse\x12C\n" +
	"\tListUsers\x12\x16.google.protobuf.Empty\x1a\x1e.user_profile.UserListResponse\x12D\n" +
	"\n" +
	"ChangeRole\x12\x1e.user_profile.AdminRoleRequest\x1a\x16.google.protobuf.EmptyB\x1bZ\x19userService.v1;uservicev1b\x06proto3"

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData []byte
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)))
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_service_proto_goTypes = []any{
	(Roles)(0),                  // 0: user_profile.Roles
	(*RegisterRequest)(nil),     // 1: user_profile.RegisterRequest
	(*RegisterResponse)(nil),    // 2: user_profile.RegisterResponse
	(*LoginRequest)(nil),        // 3: user_profile.LoginRequest
	(*Tokens)(nil),              // 4: user_profile.Tokens
	(*LoginResponse)(nil),       // 5: user_profile.LoginResponse
	(*RefreshRequest)(nil),      // 6: user_profile.RefreshRequest
	(*RefreshResponse)(nil),     // 7: user_profile.RefreshResponse
	(*LogoutRequest)(nil),       // 8: user_profile.LogoutRequest
	(*GetProfileRequest)(nil),   // 9: user_profile.GetProfileRequest
	(*UserProfileResponse)(nil), // 10: user_profile.UserProfileResponse
	(*UserListResponse)(nil),    // 11: user_profile.UserListResponse
	(*AdminRoleRequest)(nil),    // 12: user_profile.AdminRoleRequest
	(*emptypb.Empty)(nil),       // 13: google.protobuf.Empty
}
var file_user_service_proto_depIdxs = []int32{
	4,  // 0: user_profile.LoginResponse.tokens:type_name -> user_profile.Tokens
	4,  // 1: user_profile.RefreshResponse.tokens:type_name -> user_profile.Tokens
	10, // 2: user_profile.UserListResponse.users:type_name -> user_profile.UserProfileResponse
	0,  // 3: user_profile.AdminRoleRequest.role:type_name -> user_profile.Roles
	1,  // 4: user_profile.UserService.Register:input_type -> user_profile.RegisterRequest
	3,  // 5: user_profile.UserService.Login:input_type -> user_profile.LoginRequest
	6,  // 6: user_profile.UserService.RefreshToken:input_type -> user_profile.RefreshRequest
	8,  // 7: user_profile.UserService.Logout:input_type -> user_profile.LogoutRequest
	9,  // 8: user_profile.UserService.GetProfile:input_type -> user_profile.GetProfileRequest
	13, // 9: user_profile.UserService.ListUsers:input_type -> google.protobuf.Empty
	12, // 10: user_profile.UserService.ChangeRole:input_type -> user_profile.AdminRoleRequest
	2,  // 11: user_profile.UserService.Register:output_type -> user_profile.RegisterResponse
	5,  // 12: user_profile.UserService.Login:output_type -> user_profile.LoginResponse
	7,  // 13: user_profile.UserService.RefreshToken:output_type -> user_profile.RefreshResponse
	13, // 14: user_profile.UserService.Logout:output_type -> google.protobuf.Empty
	10, // 15: user_profile.UserService.GetProfile:output_type -> user_profile.UserProfileResponse
	11, // 16: user_profile.UserService.ListUsers:output_type -> user_profile.UserListResponse
	13, // 17: user_profile.UserService.ChangeRole:output_type -> google.protobuf.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		EnumInfos:         file_user_service_proto_enumTypes,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
