// Code generated by protoc-gen-go.
// source: steammessages_parental.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import math "math"

// discarding unused import google_protobuf "github.com/golang/protobuf/protoc-gen-go//descriptor"
// discarding unused import steammessages_unified_base_steamclient "steammessages_unified_base.steamclient.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ParentalApp struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	IsAllowed        *bool   `protobuf:"varint,2,opt,name=is_allowed" json:"is_allowed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ParentalApp) Reset()         { *m = ParentalApp{} }
func (m *ParentalApp) String() string { return proto.CompactTextString(m) }
func (*ParentalApp) ProtoMessage()    {}

func (m *ParentalApp) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *ParentalApp) GetIsAllowed() bool {
	if m != nil && m.IsAllowed != nil {
		return *m.IsAllowed
	}
	return false
}

type ParentalSettings struct {
	Steamid                *uint64        `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	ApplistBaseId          *uint32        `protobuf:"varint,2,opt,name=applist_base_id" json:"applist_base_id,omitempty"`
	ApplistBaseDescription *string        `protobuf:"bytes,3,opt,name=applist_base_description" json:"applist_base_description,omitempty"`
	ApplistBase            []*ParentalApp `protobuf:"bytes,4,rep,name=applist_base" json:"applist_base,omitempty"`
	ApplistCustom          []*ParentalApp `protobuf:"bytes,5,rep,name=applist_custom" json:"applist_custom,omitempty"`
	Passwordhashtype       *uint32        `protobuf:"varint,6,opt,name=passwordhashtype" json:"passwordhashtype,omitempty"`
	Salt                   []byte         `protobuf:"bytes,7,opt,name=salt" json:"salt,omitempty"`
	Passwordhash           []byte         `protobuf:"bytes,8,opt,name=passwordhash" json:"passwordhash,omitempty"`
	IsEnabled              *bool          `protobuf:"varint,9,opt,name=is_enabled" json:"is_enabled,omitempty"`
	EnabledFeatures        *uint32        `protobuf:"varint,10,opt,name=enabled_features" json:"enabled_features,omitempty"`
	XXX_unrecognized       []byte         `json:"-"`
}

func (m *ParentalSettings) Reset()         { *m = ParentalSettings{} }
func (m *ParentalSettings) String() string { return proto.CompactTextString(m) }
func (*ParentalSettings) ProtoMessage()    {}

func (m *ParentalSettings) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *ParentalSettings) GetApplistBaseId() uint32 {
	if m != nil && m.ApplistBaseId != nil {
		return *m.ApplistBaseId
	}
	return 0
}

func (m *ParentalSettings) GetApplistBaseDescription() string {
	if m != nil && m.ApplistBaseDescription != nil {
		return *m.ApplistBaseDescription
	}
	return ""
}

func (m *ParentalSettings) GetApplistBase() []*ParentalApp {
	if m != nil {
		return m.ApplistBase
	}
	return nil
}

func (m *ParentalSettings) GetApplistCustom() []*ParentalApp {
	if m != nil {
		return m.ApplistCustom
	}
	return nil
}

func (m *ParentalSettings) GetPasswordhashtype() uint32 {
	if m != nil && m.Passwordhashtype != nil {
		return *m.Passwordhashtype
	}
	return 0
}

func (m *ParentalSettings) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ParentalSettings) GetPasswordhash() []byte {
	if m != nil {
		return m.Passwordhash
	}
	return nil
}

func (m *ParentalSettings) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *ParentalSettings) GetEnabledFeatures() uint32 {
	if m != nil && m.EnabledFeatures != nil {
		return *m.EnabledFeatures
	}
	return 0
}

type CParental_EnableParentalSettings_Request struct {
	Password         *string           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Settings         *ParentalSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
	Sessionid        *string           `protobuf:"bytes,3,opt,name=sessionid" json:"sessionid,omitempty"`
	Steamid          *uint64           `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_EnableParentalSettings_Request) Reset() {
	*m = CParental_EnableParentalSettings_Request{}
}
func (m *CParental_EnableParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_EnableParentalSettings_Request) ProtoMessage()    {}

func (m *CParental_EnableParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_EnableParentalSettings_Request) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *CParental_EnableParentalSettings_Request) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func (m *CParental_EnableParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_EnableParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_EnableParentalSettings_Response) Reset() {
	*m = CParental_EnableParentalSettings_Response{}
}
func (m *CParental_EnableParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_EnableParentalSettings_Response) ProtoMessage()    {}

type CParental_DisableParentalSettings_Request struct {
	Password         *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_DisableParentalSettings_Request) Reset() {
	*m = CParental_DisableParentalSettings_Request{}
}
func (m *CParental_DisableParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_DisableParentalSettings_Request) ProtoMessage()    {}

func (m *CParental_DisableParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_DisableParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_DisableParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_DisableParentalSettings_Response) Reset() {
	*m = CParental_DisableParentalSettings_Response{}
}
func (m *CParental_DisableParentalSettings_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_DisableParentalSettings_Response) ProtoMessage() {}

type CParental_GetParentalSettings_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_GetParentalSettings_Request) Reset()         { *m = CParental_GetParentalSettings_Request{} }
func (m *CParental_GetParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_GetParentalSettings_Request) ProtoMessage()    {}

func (m *CParental_GetParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_GetParentalSettings_Response struct {
	Settings         *ParentalSettings `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_GetParentalSettings_Response) Reset() {
	*m = CParental_GetParentalSettings_Response{}
}
func (m *CParental_GetParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_GetParentalSettings_Response) ProtoMessage()    {}

func (m *CParental_GetParentalSettings_Response) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type CParental_GetSignedParentalSettings_Request struct {
	Priority         *uint32 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_GetSignedParentalSettings_Request) Reset() {
	*m = CParental_GetSignedParentalSettings_Request{}
}
func (m *CParental_GetSignedParentalSettings_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_GetSignedParentalSettings_Request) ProtoMessage() {}

func (m *CParental_GetSignedParentalSettings_Request) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

type CParental_GetSignedParentalSettings_Response struct {
	SerializedSettings []byte `protobuf:"bytes,1,opt,name=serialized_settings" json:"serialized_settings,omitempty"`
	Signature          []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *CParental_GetSignedParentalSettings_Response) Reset() {
	*m = CParental_GetSignedParentalSettings_Response{}
}
func (m *CParental_GetSignedParentalSettings_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_GetSignedParentalSettings_Response) ProtoMessage() {}

func (m *CParental_GetSignedParentalSettings_Response) GetSerializedSettings() []byte {
	if m != nil {
		return m.SerializedSettings
	}
	return nil
}

func (m *CParental_GetSignedParentalSettings_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CParental_SetParentalSettings_Request struct {
	Password         *string           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Settings         *ParentalSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
	NewPassword      *string           `protobuf:"bytes,3,opt,name=new_password" json:"new_password,omitempty"`
	Sessionid        *string           `protobuf:"bytes,4,opt,name=sessionid" json:"sessionid,omitempty"`
	Steamid          *uint64           `protobuf:"fixed64,10,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CParental_SetParentalSettings_Request) Reset()         { *m = CParental_SetParentalSettings_Request{} }
func (m *CParental_SetParentalSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_SetParentalSettings_Request) ProtoMessage()    {}

func (m *CParental_SetParentalSettings_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSettings() *ParentalSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *CParental_SetParentalSettings_Request) GetNewPassword() string {
	if m != nil && m.NewPassword != nil {
		return *m.NewPassword
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func (m *CParental_SetParentalSettings_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CParental_SetParentalSettings_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_SetParentalSettings_Response) Reset() {
	*m = CParental_SetParentalSettings_Response{}
}
func (m *CParental_SetParentalSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_SetParentalSettings_Response) ProtoMessage()    {}

type CParental_ValidateToken_Request struct {
	UnlockToken      *string `protobuf:"bytes,1,opt,name=unlock_token" json:"unlock_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ValidateToken_Request) Reset()         { *m = CParental_ValidateToken_Request{} }
func (m *CParental_ValidateToken_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidateToken_Request) ProtoMessage()    {}

func (m *CParental_ValidateToken_Request) GetUnlockToken() string {
	if m != nil && m.UnlockToken != nil {
		return *m.UnlockToken
	}
	return ""
}

type CParental_ValidateToken_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_ValidateToken_Response) Reset()         { *m = CParental_ValidateToken_Response{} }
func (m *CParental_ValidateToken_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidateToken_Response) ProtoMessage()    {}

type CParental_ValidatePassword_Request struct {
	Password            *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Session             *string `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	SendUnlockOnSuccess *bool   `protobuf:"varint,3,opt,name=send_unlock_on_success" json:"send_unlock_on_success,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *CParental_ValidatePassword_Request) Reset()         { *m = CParental_ValidatePassword_Request{} }
func (m *CParental_ValidatePassword_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidatePassword_Request) ProtoMessage()    {}

func (m *CParental_ValidatePassword_Request) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ValidatePassword_Request) GetSession() string {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return ""
}

func (m *CParental_ValidatePassword_Request) GetSendUnlockOnSuccess() bool {
	if m != nil && m.SendUnlockOnSuccess != nil {
		return *m.SendUnlockOnSuccess
	}
	return false
}

type CParental_ValidatePassword_Response struct {
	Token            *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ValidatePassword_Response) Reset()         { *m = CParental_ValidatePassword_Response{} }
func (m *CParental_ValidatePassword_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_ValidatePassword_Response) ProtoMessage()    {}

func (m *CParental_ValidatePassword_Response) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type CParental_LockClient_Request struct {
	Session          *string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_LockClient_Request) Reset()         { *m = CParental_LockClient_Request{} }
func (m *CParental_LockClient_Request) String() string { return proto.CompactTextString(m) }
func (*CParental_LockClient_Request) ProtoMessage()    {}

func (m *CParental_LockClient_Request) GetSession() string {
	if m != nil && m.Session != nil {
		return *m.Session
	}
	return ""
}

type CParental_LockClient_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CParental_LockClient_Response) Reset()         { *m = CParental_LockClient_Response{} }
func (m *CParental_LockClient_Response) String() string { return proto.CompactTextString(m) }
func (*CParental_LockClient_Response) ProtoMessage()    {}

type CParental_ParentalSettingsChange_Notification struct {
	SerializedSettings []byte  `protobuf:"bytes,1,opt,name=serialized_settings" json:"serialized_settings,omitempty"`
	Signature          []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Password           *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Sessionid          *string `protobuf:"bytes,4,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CParental_ParentalSettingsChange_Notification) Reset() {
	*m = CParental_ParentalSettingsChange_Notification{}
}
func (m *CParental_ParentalSettingsChange_Notification) String() string {
	return proto.CompactTextString(m)
}
func (*CParental_ParentalSettingsChange_Notification) ProtoMessage() {}

func (m *CParental_ParentalSettingsChange_Notification) GetSerializedSettings() []byte {
	if m != nil {
		return m.SerializedSettings
	}
	return nil
}

func (m *CParental_ParentalSettingsChange_Notification) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CParental_ParentalSettingsChange_Notification) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ParentalSettingsChange_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

type CParental_ParentalUnlock_Notification struct {
	Password         *string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Sessionid        *string `protobuf:"bytes,2,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ParentalUnlock_Notification) Reset()         { *m = CParental_ParentalUnlock_Notification{} }
func (m *CParental_ParentalUnlock_Notification) String() string { return proto.CompactTextString(m) }
func (*CParental_ParentalUnlock_Notification) ProtoMessage()    {}

func (m *CParental_ParentalUnlock_Notification) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CParental_ParentalUnlock_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

type CParental_ParentalLock_Notification struct {
	Sessionid        *string `protobuf:"bytes,1,opt,name=sessionid" json:"sessionid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CParental_ParentalLock_Notification) Reset()         { *m = CParental_ParentalLock_Notification{} }
func (m *CParental_ParentalLock_Notification) String() string { return proto.CompactTextString(m) }
func (*CParental_ParentalLock_Notification) ProtoMessage()    {}

func (m *CParental_ParentalLock_Notification) GetSessionid() string {
	if m != nil && m.Sessionid != nil {
		return *m.Sessionid
	}
	return ""
}

func init() {
}
