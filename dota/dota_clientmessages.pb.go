// Code generated by protoc-gen-go.
// source: dota_clientmessages.proto
// DO NOT EDIT!

package dota

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EDotaClientMessages int32

const (
	EDotaClientMessages_DOTA_CM_MapLine                                EDotaClientMessages = 1
	EDotaClientMessages_DOTA_CM_AspectRatio                            EDotaClientMessages = 2
	EDotaClientMessages_DOTA_CM_MapPing                                EDotaClientMessages = 3
	EDotaClientMessages_DOTA_CM_UnitsAutoAttack                        EDotaClientMessages = 4
	EDotaClientMessages_DOTA_CM_AutoPurchaseItems                      EDotaClientMessages = 5
	EDotaClientMessages_DOTA_CM_TestItems                              EDotaClientMessages = 6
	EDotaClientMessages_DOTA_CM_SearchString                           EDotaClientMessages = 7
	EDotaClientMessages_DOTA_CM_Pause                                  EDotaClientMessages = 8
	EDotaClientMessages_DOTA_CM_ShopViewMode                           EDotaClientMessages = 9
	EDotaClientMessages_DOTA_CM_SetUnitShareFlag                       EDotaClientMessages = 10
	EDotaClientMessages_DOTA_CM_SwapRequest                            EDotaClientMessages = 11
	EDotaClientMessages_DOTA_CM_SwapAccept                             EDotaClientMessages = 12
	EDotaClientMessages_DOTA_CM_WorldLine                              EDotaClientMessages = 13
	EDotaClientMessages_DOTA_CM_RequestGraphUpdate                     EDotaClientMessages = 14
	EDotaClientMessages_DOTA_CM_ItemAlert                              EDotaClientMessages = 15
	EDotaClientMessages_DOTA_CM_ChatWheel                              EDotaClientMessages = 16
	EDotaClientMessages_DOTA_CM_SendStatPopup                          EDotaClientMessages = 17
	EDotaClientMessages_DOTA_CM_BeginLastHitChallenge                  EDotaClientMessages = 18
	EDotaClientMessages_DOTA_CM_UpdateQuickBuy                         EDotaClientMessages = 19
	EDotaClientMessages_DOTA_CM_UpdateCoachListen                      EDotaClientMessages = 20
	EDotaClientMessages_DOTA_CM_CoachHUDPing                           EDotaClientMessages = 21
	EDotaClientMessages_DOTA_CM_RecordVote                             EDotaClientMessages = 22
	EDotaClientMessages_DOTA_CM_UnitsAutoAttackAfterSpell              EDotaClientMessages = 23
	EDotaClientMessages_DOTA_CM_WillPurchaseAlert                      EDotaClientMessages = 24
	EDotaClientMessages_DOTA_CM_PlayerShowCase                         EDotaClientMessages = 25
	EDotaClientMessages_DOTA_CM_TeleportRequiresHalt                   EDotaClientMessages = 26
	EDotaClientMessages_DOTA_CM_CameraZoomAmount                       EDotaClientMessages = 27
	EDotaClientMessages_DOTA_CM_BroadcasterUsingCamerman               EDotaClientMessages = 28
	EDotaClientMessages_DOTA_CM_BroadcasterUsingAssistedCameraOperator EDotaClientMessages = 29
	EDotaClientMessages_DOTA_CM_EnemyItemAlert                         EDotaClientMessages = 30
	EDotaClientMessages_DOTA_CM_FreeInventory                          EDotaClientMessages = 31
	EDotaClientMessages_DOTA_CM_BuyBackStateAlert                      EDotaClientMessages = 32
	EDotaClientMessages_DOTA_CM_QuickBuyAlert                          EDotaClientMessages = 33
	EDotaClientMessages_DOTA_CM_HeroStatueLike                         EDotaClientMessages = 34
	EDotaClientMessages_DOTA_CM_ModifierAlert                          EDotaClientMessages = 35
	EDotaClientMessages_DOTA_CM_TeamShowcaseEditor                     EDotaClientMessages = 36
	EDotaClientMessages_DOTA_CM_HPManaAlert                            EDotaClientMessages = 37
	EDotaClientMessages_DOTA_CM_GlyphAlert                             EDotaClientMessages = 38
	EDotaClientMessages_DOTA_CM_TeamShowcaseClientData                 EDotaClientMessages = 39
	EDotaClientMessages_DOTA_CM_PlayTeamShowcase                       EDotaClientMessages = 40
)

var EDotaClientMessages_name = map[int32]string{
	1:  "DOTA_CM_MapLine",
	2:  "DOTA_CM_AspectRatio",
	3:  "DOTA_CM_MapPing",
	4:  "DOTA_CM_UnitsAutoAttack",
	5:  "DOTA_CM_AutoPurchaseItems",
	6:  "DOTA_CM_TestItems",
	7:  "DOTA_CM_SearchString",
	8:  "DOTA_CM_Pause",
	9:  "DOTA_CM_ShopViewMode",
	10: "DOTA_CM_SetUnitShareFlag",
	11: "DOTA_CM_SwapRequest",
	12: "DOTA_CM_SwapAccept",
	13: "DOTA_CM_WorldLine",
	14: "DOTA_CM_RequestGraphUpdate",
	15: "DOTA_CM_ItemAlert",
	16: "DOTA_CM_ChatWheel",
	17: "DOTA_CM_SendStatPopup",
	18: "DOTA_CM_BeginLastHitChallenge",
	19: "DOTA_CM_UpdateQuickBuy",
	20: "DOTA_CM_UpdateCoachListen",
	21: "DOTA_CM_CoachHUDPing",
	22: "DOTA_CM_RecordVote",
	23: "DOTA_CM_UnitsAutoAttackAfterSpell",
	24: "DOTA_CM_WillPurchaseAlert",
	25: "DOTA_CM_PlayerShowCase",
	26: "DOTA_CM_TeleportRequiresHalt",
	27: "DOTA_CM_CameraZoomAmount",
	28: "DOTA_CM_BroadcasterUsingCamerman",
	29: "DOTA_CM_BroadcasterUsingAssistedCameraOperator",
	30: "DOTA_CM_EnemyItemAlert",
	31: "DOTA_CM_FreeInventory",
	32: "DOTA_CM_BuyBackStateAlert",
	33: "DOTA_CM_QuickBuyAlert",
	34: "DOTA_CM_HeroStatueLike",
	35: "DOTA_CM_ModifierAlert",
	36: "DOTA_CM_TeamShowcaseEditor",
	37: "DOTA_CM_HPManaAlert",
	38: "DOTA_CM_GlyphAlert",
	39: "DOTA_CM_TeamShowcaseClientData",
	40: "DOTA_CM_PlayTeamShowcase",
}
var EDotaClientMessages_value = map[string]int32{
	"DOTA_CM_MapLine":                                1,
	"DOTA_CM_AspectRatio":                            2,
	"DOTA_CM_MapPing":                                3,
	"DOTA_CM_UnitsAutoAttack":                        4,
	"DOTA_CM_AutoPurchaseItems":                      5,
	"DOTA_CM_TestItems":                              6,
	"DOTA_CM_SearchString":                           7,
	"DOTA_CM_Pause":                                  8,
	"DOTA_CM_ShopViewMode":                           9,
	"DOTA_CM_SetUnitShareFlag":                       10,
	"DOTA_CM_SwapRequest":                            11,
	"DOTA_CM_SwapAccept":                             12,
	"DOTA_CM_WorldLine":                              13,
	"DOTA_CM_RequestGraphUpdate":                     14,
	"DOTA_CM_ItemAlert":                              15,
	"DOTA_CM_ChatWheel":                              16,
	"DOTA_CM_SendStatPopup":                          17,
	"DOTA_CM_BeginLastHitChallenge":                  18,
	"DOTA_CM_UpdateQuickBuy":                         19,
	"DOTA_CM_UpdateCoachListen":                      20,
	"DOTA_CM_CoachHUDPing":                           21,
	"DOTA_CM_RecordVote":                             22,
	"DOTA_CM_UnitsAutoAttackAfterSpell":              23,
	"DOTA_CM_WillPurchaseAlert":                      24,
	"DOTA_CM_PlayerShowCase":                         25,
	"DOTA_CM_TeleportRequiresHalt":                   26,
	"DOTA_CM_CameraZoomAmount":                       27,
	"DOTA_CM_BroadcasterUsingCamerman":               28,
	"DOTA_CM_BroadcasterUsingAssistedCameraOperator": 29,
	"DOTA_CM_EnemyItemAlert":                         30,
	"DOTA_CM_FreeInventory":                          31,
	"DOTA_CM_BuyBackStateAlert":                      32,
	"DOTA_CM_QuickBuyAlert":                          33,
	"DOTA_CM_HeroStatueLike":                         34,
	"DOTA_CM_ModifierAlert":                          35,
	"DOTA_CM_TeamShowcaseEditor":                     36,
	"DOTA_CM_HPManaAlert":                            37,
	"DOTA_CM_GlyphAlert":                             38,
	"DOTA_CM_TeamShowcaseClientData":                 39,
	"DOTA_CM_PlayTeamShowcase":                       40,
}

func (x EDotaClientMessages) Enum() *EDotaClientMessages {
	p := new(EDotaClientMessages)
	*p = x
	return p
}
func (x EDotaClientMessages) String() string {
	return proto.EnumName(EDotaClientMessages_name, int32(x))
}
func (x *EDotaClientMessages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDotaClientMessages_value, data, "EDotaClientMessages")
	if err != nil {
		return err
	}
	*x = EDotaClientMessages(value)
	return nil
}

type CDOTAClientMsg_MapPing struct {
	LocationPing     *CDOTAMsg_LocationPing `protobuf:"bytes,1,opt,name=location_ping" json:"location_ping,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CDOTAClientMsg_MapPing) Reset()         { *m = CDOTAClientMsg_MapPing{} }
func (m *CDOTAClientMsg_MapPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_MapPing) ProtoMessage()    {}

func (m *CDOTAClientMsg_MapPing) GetLocationPing() *CDOTAMsg_LocationPing {
	if m != nil {
		return m.LocationPing
	}
	return nil
}

type CDOTAClientMsg_ItemAlert struct {
	ItemAlert        *CDOTAMsg_ItemAlert `protobuf:"bytes,1,opt,name=item_alert" json:"item_alert,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAClientMsg_ItemAlert) Reset()         { *m = CDOTAClientMsg_ItemAlert{} }
func (m *CDOTAClientMsg_ItemAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_ItemAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_ItemAlert) GetItemAlert() *CDOTAMsg_ItemAlert {
	if m != nil {
		return m.ItemAlert
	}
	return nil
}

type CDOTAClientMsg_EnemyItemAlert struct {
	ItemEntindex     *uint32 `protobuf:"varint,1,opt,name=item_entindex" json:"item_entindex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_EnemyItemAlert) Reset()         { *m = CDOTAClientMsg_EnemyItemAlert{} }
func (m *CDOTAClientMsg_EnemyItemAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_EnemyItemAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_EnemyItemAlert) GetItemEntindex() uint32 {
	if m != nil && m.ItemEntindex != nil {
		return *m.ItemEntindex
	}
	return 0
}

type CDOTAClientMsg_ModifierAlert struct {
	BuffInternalIndex *int32 `protobuf:"varint,1,opt,name=buff_internal_index" json:"buff_internal_index,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *CDOTAClientMsg_ModifierAlert) Reset()         { *m = CDOTAClientMsg_ModifierAlert{} }
func (m *CDOTAClientMsg_ModifierAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_ModifierAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_ModifierAlert) GetBuffInternalIndex() int32 {
	if m != nil && m.BuffInternalIndex != nil {
		return *m.BuffInternalIndex
	}
	return 0
}

type CDOTAClientMsg_HPManaAlert struct {
	TargetEntindex   *uint32 `protobuf:"varint,1,opt,name=target_entindex" json:"target_entindex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_HPManaAlert) Reset()         { *m = CDOTAClientMsg_HPManaAlert{} }
func (m *CDOTAClientMsg_HPManaAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_HPManaAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_HPManaAlert) GetTargetEntindex() uint32 {
	if m != nil && m.TargetEntindex != nil {
		return *m.TargetEntindex
	}
	return 0
}

type CDOTAClientMsg_GlyphAlert struct {
	Negative         *bool  `protobuf:"varint,1,opt,name=negative" json:"negative,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_GlyphAlert) Reset()         { *m = CDOTAClientMsg_GlyphAlert{} }
func (m *CDOTAClientMsg_GlyphAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_GlyphAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_GlyphAlert) GetNegative() bool {
	if m != nil && m.Negative != nil {
		return *m.Negative
	}
	return false
}

type CDOTAClientMsg_MapLine struct {
	Mapline          *CDOTAMsg_MapLine `protobuf:"bytes,1,opt,name=mapline" json:"mapline,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CDOTAClientMsg_MapLine) Reset()         { *m = CDOTAClientMsg_MapLine{} }
func (m *CDOTAClientMsg_MapLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_MapLine) ProtoMessage()    {}

func (m *CDOTAClientMsg_MapLine) GetMapline() *CDOTAMsg_MapLine {
	if m != nil {
		return m.Mapline
	}
	return nil
}

type CDOTAClientMsg_AspectRatio struct {
	Ratio            *float32 `protobuf:"fixed32,1,opt,name=ratio" json:"ratio,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAClientMsg_AspectRatio) Reset()         { *m = CDOTAClientMsg_AspectRatio{} }
func (m *CDOTAClientMsg_AspectRatio) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_AspectRatio) ProtoMessage()    {}

func (m *CDOTAClientMsg_AspectRatio) GetRatio() float32 {
	if m != nil && m.Ratio != nil {
		return *m.Ratio
	}
	return 0
}

type CDOTAClientMsg_UnitsAutoAttack struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_UnitsAutoAttack) Reset()         { *m = CDOTAClientMsg_UnitsAutoAttack{} }
func (m *CDOTAClientMsg_UnitsAutoAttack) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_UnitsAutoAttack) ProtoMessage()    {}

func (m *CDOTAClientMsg_UnitsAutoAttack) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

type CDOTAClientMsg_UnitsAutoAttackAfterSpell struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_UnitsAutoAttackAfterSpell) Reset() {
	*m = CDOTAClientMsg_UnitsAutoAttackAfterSpell{}
}
func (m *CDOTAClientMsg_UnitsAutoAttackAfterSpell) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_UnitsAutoAttackAfterSpell) ProtoMessage()    {}

func (m *CDOTAClientMsg_UnitsAutoAttackAfterSpell) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

type CDOTAClientMsg_TeleportRequiresHalt struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_TeleportRequiresHalt) Reset()         { *m = CDOTAClientMsg_TeleportRequiresHalt{} }
func (m *CDOTAClientMsg_TeleportRequiresHalt) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_TeleportRequiresHalt) ProtoMessage()    {}

func (m *CDOTAClientMsg_TeleportRequiresHalt) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

type CDOTAClientMsg_AutoPurchaseItems struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_AutoPurchaseItems) Reset()         { *m = CDOTAClientMsg_AutoPurchaseItems{} }
func (m *CDOTAClientMsg_AutoPurchaseItems) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_AutoPurchaseItems) ProtoMessage()    {}

func (m *CDOTAClientMsg_AutoPurchaseItems) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

type CDOTAClientMsg_TestItems struct {
	KeyValues        *string `protobuf:"bytes,1,opt,name=key_values" json:"key_values,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_TestItems) Reset()         { *m = CDOTAClientMsg_TestItems{} }
func (m *CDOTAClientMsg_TestItems) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_TestItems) ProtoMessage()    {}

func (m *CDOTAClientMsg_TestItems) GetKeyValues() string {
	if m != nil && m.KeyValues != nil {
		return *m.KeyValues
	}
	return ""
}

type CDOTAClientMsg_SearchString struct {
	Search           *string `protobuf:"bytes,1,opt,name=search" json:"search,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_SearchString) Reset()         { *m = CDOTAClientMsg_SearchString{} }
func (m *CDOTAClientMsg_SearchString) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_SearchString) ProtoMessage()    {}

func (m *CDOTAClientMsg_SearchString) GetSearch() string {
	if m != nil && m.Search != nil {
		return *m.Search
	}
	return ""
}

type CDOTAClientMsg_Pause struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_Pause) Reset()         { *m = CDOTAClientMsg_Pause{} }
func (m *CDOTAClientMsg_Pause) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_Pause) ProtoMessage()    {}

type CDOTAClientMsg_ShopViewMode struct {
	Mode             *uint32 `protobuf:"varint,1,opt,name=mode" json:"mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_ShopViewMode) Reset()         { *m = CDOTAClientMsg_ShopViewMode{} }
func (m *CDOTAClientMsg_ShopViewMode) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_ShopViewMode) ProtoMessage()    {}

func (m *CDOTAClientMsg_ShopViewMode) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

type CDOTAClientMsg_SetUnitShareFlag struct {
	PlayerID         *uint32 `protobuf:"varint,1,opt,name=playerID" json:"playerID,omitempty"`
	Flag             *uint32 `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	State            *bool   `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_SetUnitShareFlag) Reset()         { *m = CDOTAClientMsg_SetUnitShareFlag{} }
func (m *CDOTAClientMsg_SetUnitShareFlag) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_SetUnitShareFlag) ProtoMessage()    {}

func (m *CDOTAClientMsg_SetUnitShareFlag) GetPlayerID() uint32 {
	if m != nil && m.PlayerID != nil {
		return *m.PlayerID
	}
	return 0
}

func (m *CDOTAClientMsg_SetUnitShareFlag) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *CDOTAClientMsg_SetUnitShareFlag) GetState() bool {
	if m != nil && m.State != nil {
		return *m.State
	}
	return false
}

type CDOTAClientMsg_SwapRequest struct {
	PlayerId         *uint32 `protobuf:"varint,1,opt,name=player_id" json:"player_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_SwapRequest) Reset()         { *m = CDOTAClientMsg_SwapRequest{} }
func (m *CDOTAClientMsg_SwapRequest) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_SwapRequest) ProtoMessage()    {}

func (m *CDOTAClientMsg_SwapRequest) GetPlayerId() uint32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type CDOTAClientMsg_SwapAccept struct {
	PlayerId         *uint32 `protobuf:"varint,1,opt,name=player_id" json:"player_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_SwapAccept) Reset()         { *m = CDOTAClientMsg_SwapAccept{} }
func (m *CDOTAClientMsg_SwapAccept) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_SwapAccept) ProtoMessage()    {}

func (m *CDOTAClientMsg_SwapAccept) GetPlayerId() uint32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type CDOTAClientMsg_WorldLine struct {
	Worldline        *CDOTAMsg_WorldLine `protobuf:"bytes,1,opt,name=worldline" json:"worldline,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAClientMsg_WorldLine) Reset()         { *m = CDOTAClientMsg_WorldLine{} }
func (m *CDOTAClientMsg_WorldLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_WorldLine) ProtoMessage()    {}

func (m *CDOTAClientMsg_WorldLine) GetWorldline() *CDOTAMsg_WorldLine {
	if m != nil {
		return m.Worldline
	}
	return nil
}

type CDOTAClientMsg_RequestGraphUpdate struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_RequestGraphUpdate) Reset()         { *m = CDOTAClientMsg_RequestGraphUpdate{} }
func (m *CDOTAClientMsg_RequestGraphUpdate) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_RequestGraphUpdate) ProtoMessage()    {}

type CDOTAClientMsg_ChatWheel struct {
	ChatMessage      *EDOTAChatWheelMessage `protobuf:"varint,1,opt,name=chat_message,enum=dota.EDOTAChatWheelMessage,def=0" json:"chat_message,omitempty"`
	ParamHeroId      *uint32                `protobuf:"varint,2,opt,name=param_hero_id" json:"param_hero_id,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CDOTAClientMsg_ChatWheel) Reset()         { *m = CDOTAClientMsg_ChatWheel{} }
func (m *CDOTAClientMsg_ChatWheel) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_ChatWheel) ProtoMessage()    {}

const Default_CDOTAClientMsg_ChatWheel_ChatMessage EDOTAChatWheelMessage = EDOTAChatWheelMessage_k_EDOTA_CW_Ok

func (m *CDOTAClientMsg_ChatWheel) GetChatMessage() EDOTAChatWheelMessage {
	if m != nil && m.ChatMessage != nil {
		return *m.ChatMessage
	}
	return Default_CDOTAClientMsg_ChatWheel_ChatMessage
}

func (m *CDOTAClientMsg_ChatWheel) GetParamHeroId() uint32 {
	if m != nil && m.ParamHeroId != nil {
		return *m.ParamHeroId
	}
	return 0
}

type CDOTAClientMsg_SendStatPopup struct {
	Statpopup        *CDOTAMsg_SendStatPopup `protobuf:"bytes,1,opt,name=statpopup" json:"statpopup,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *CDOTAClientMsg_SendStatPopup) Reset()         { *m = CDOTAClientMsg_SendStatPopup{} }
func (m *CDOTAClientMsg_SendStatPopup) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_SendStatPopup) ProtoMessage()    {}

func (m *CDOTAClientMsg_SendStatPopup) GetStatpopup() *CDOTAMsg_SendStatPopup {
	if m != nil {
		return m.Statpopup
	}
	return nil
}

type CDOTAClientMsg_BeginLastHitChallenge struct {
	ChosenLane       *uint32 `protobuf:"varint,1,opt,name=chosen_lane" json:"chosen_lane,omitempty"`
	HelperEnabled    *bool   `protobuf:"varint,2,opt,name=helper_enabled" json:"helper_enabled,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_BeginLastHitChallenge) Reset()         { *m = CDOTAClientMsg_BeginLastHitChallenge{} }
func (m *CDOTAClientMsg_BeginLastHitChallenge) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_BeginLastHitChallenge) ProtoMessage()    {}

func (m *CDOTAClientMsg_BeginLastHitChallenge) GetChosenLane() uint32 {
	if m != nil && m.ChosenLane != nil {
		return *m.ChosenLane
	}
	return 0
}

func (m *CDOTAClientMsg_BeginLastHitChallenge) GetHelperEnabled() bool {
	if m != nil && m.HelperEnabled != nil {
		return *m.HelperEnabled
	}
	return false
}

type CDOTAClientMsg_UpdateQuickBuyItem struct {
	ItemType         *int32 `protobuf:"varint,1,opt,name=item_type" json:"item_type,omitempty"`
	Purchasable      *bool  `protobuf:"varint,2,opt,name=purchasable" json:"purchasable,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_UpdateQuickBuyItem) Reset()         { *m = CDOTAClientMsg_UpdateQuickBuyItem{} }
func (m *CDOTAClientMsg_UpdateQuickBuyItem) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_UpdateQuickBuyItem) ProtoMessage()    {}

func (m *CDOTAClientMsg_UpdateQuickBuyItem) GetItemType() int32 {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return 0
}

func (m *CDOTAClientMsg_UpdateQuickBuyItem) GetPurchasable() bool {
	if m != nil && m.Purchasable != nil {
		return *m.Purchasable
	}
	return false
}

type CDOTAClientMsg_UpdateQuickBuy struct {
	Items            []*CDOTAClientMsg_UpdateQuickBuyItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CDOTAClientMsg_UpdateQuickBuy) Reset()         { *m = CDOTAClientMsg_UpdateQuickBuy{} }
func (m *CDOTAClientMsg_UpdateQuickBuy) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_UpdateQuickBuy) ProtoMessage()    {}

func (m *CDOTAClientMsg_UpdateQuickBuy) GetItems() []*CDOTAClientMsg_UpdateQuickBuyItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type CDOTAClientMsg_UpdateCoachListen struct {
	PlayerMask       *uint32 `protobuf:"varint,1,opt,name=player_mask" json:"player_mask,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_UpdateCoachListen) Reset()         { *m = CDOTAClientMsg_UpdateCoachListen{} }
func (m *CDOTAClientMsg_UpdateCoachListen) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_UpdateCoachListen) ProtoMessage()    {}

func (m *CDOTAClientMsg_UpdateCoachListen) GetPlayerMask() uint32 {
	if m != nil && m.PlayerMask != nil {
		return *m.PlayerMask
	}
	return 0
}

type CDOTAClientMsg_CoachHUDPing struct {
	HudPing          *CDOTAMsg_CoachHUDPing `protobuf:"bytes,1,opt,name=hud_ping" json:"hud_ping,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CDOTAClientMsg_CoachHUDPing) Reset()         { *m = CDOTAClientMsg_CoachHUDPing{} }
func (m *CDOTAClientMsg_CoachHUDPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_CoachHUDPing) ProtoMessage()    {}

func (m *CDOTAClientMsg_CoachHUDPing) GetHudPing() *CDOTAMsg_CoachHUDPing {
	if m != nil {
		return m.HudPing
	}
	return nil
}

type CDOTAClientMsg_RecordVote struct {
	ChoiceIndex      *int32 `protobuf:"varint,1,opt,name=choice_index" json:"choice_index,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_RecordVote) Reset()         { *m = CDOTAClientMsg_RecordVote{} }
func (m *CDOTAClientMsg_RecordVote) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_RecordVote) ProtoMessage()    {}

func (m *CDOTAClientMsg_RecordVote) GetChoiceIndex() int32 {
	if m != nil && m.ChoiceIndex != nil {
		return *m.ChoiceIndex
	}
	return 0
}

type CDOTAClientMsg_WillPurchaseAlert struct {
	Itemid           *int32  `protobuf:"varint,1,opt,name=itemid" json:"itemid,omitempty"`
	GoldRemaining    *uint32 `protobuf:"varint,2,opt,name=gold_remaining" json:"gold_remaining,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_WillPurchaseAlert) Reset()         { *m = CDOTAClientMsg_WillPurchaseAlert{} }
func (m *CDOTAClientMsg_WillPurchaseAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_WillPurchaseAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_WillPurchaseAlert) GetItemid() int32 {
	if m != nil && m.Itemid != nil {
		return *m.Itemid
	}
	return 0
}

func (m *CDOTAClientMsg_WillPurchaseAlert) GetGoldRemaining() uint32 {
	if m != nil && m.GoldRemaining != nil {
		return *m.GoldRemaining
	}
	return 0
}

type CDOTAClientMsg_BuyBackStateAlert struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_BuyBackStateAlert) Reset()         { *m = CDOTAClientMsg_BuyBackStateAlert{} }
func (m *CDOTAClientMsg_BuyBackStateAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_BuyBackStateAlert) ProtoMessage()    {}

type CDOTAClientMsg_QuickBuyAlert struct {
	Itemid           *int32 `protobuf:"varint,1,opt,name=itemid" json:"itemid,omitempty"`
	GoldRequired     *int32 `protobuf:"varint,2,opt,name=gold_required" json:"gold_required,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_QuickBuyAlert) Reset()         { *m = CDOTAClientMsg_QuickBuyAlert{} }
func (m *CDOTAClientMsg_QuickBuyAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_QuickBuyAlert) ProtoMessage()    {}

func (m *CDOTAClientMsg_QuickBuyAlert) GetItemid() int32 {
	if m != nil && m.Itemid != nil {
		return *m.Itemid
	}
	return 0
}

func (m *CDOTAClientMsg_QuickBuyAlert) GetGoldRequired() int32 {
	if m != nil && m.GoldRequired != nil {
		return *m.GoldRequired
	}
	return 0
}

type CDOTAClientMsg_PlayerShowCase struct {
	Showcase         *bool  `protobuf:"varint,1,opt,name=showcase" json:"showcase,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_PlayerShowCase) Reset()         { *m = CDOTAClientMsg_PlayerShowCase{} }
func (m *CDOTAClientMsg_PlayerShowCase) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_PlayerShowCase) ProtoMessage()    {}

func (m *CDOTAClientMsg_PlayerShowCase) GetShowcase() bool {
	if m != nil && m.Showcase != nil {
		return *m.Showcase
	}
	return false
}

type CDOTAClientMsg_CameraZoomAmount struct {
	ZoomAmount       *float32 `protobuf:"fixed32,1,opt,name=zoom_amount" json:"zoom_amount,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAClientMsg_CameraZoomAmount) Reset()         { *m = CDOTAClientMsg_CameraZoomAmount{} }
func (m *CDOTAClientMsg_CameraZoomAmount) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_CameraZoomAmount) ProtoMessage()    {}

func (m *CDOTAClientMsg_CameraZoomAmount) GetZoomAmount() float32 {
	if m != nil && m.ZoomAmount != nil {
		return *m.ZoomAmount
	}
	return 0
}

type CDOTAClientMsg_BroadcasterUsingCameraman struct {
	Cameraman        *bool  `protobuf:"varint,1,opt,name=cameraman" json:"cameraman,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_BroadcasterUsingCameraman) Reset() {
	*m = CDOTAClientMsg_BroadcasterUsingCameraman{}
}
func (m *CDOTAClientMsg_BroadcasterUsingCameraman) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_BroadcasterUsingCameraman) ProtoMessage()    {}

func (m *CDOTAClientMsg_BroadcasterUsingCameraman) GetCameraman() bool {
	if m != nil && m.Cameraman != nil {
		return *m.Cameraman
	}
	return false
}

type CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator) Reset() {
	*m = CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator{}
}
func (m *CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator) String() string {
	return proto.CompactTextString(m)
}
func (*CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator) ProtoMessage() {}

func (m *CDOTAClientMsg_BroadcasterUsingAssistedCameraOperator) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

type CAdditionalEquipSlotClientMsg struct {
	ClassId          *uint32 `protobuf:"varint,1,opt,name=class_id" json:"class_id,omitempty"`
	SlotId           *uint32 `protobuf:"varint,2,opt,name=slot_id" json:"slot_id,omitempty"`
	DefIndex         *uint32 `protobuf:"varint,3,opt,name=def_index" json:"def_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CAdditionalEquipSlotClientMsg) Reset()         { *m = CAdditionalEquipSlotClientMsg{} }
func (m *CAdditionalEquipSlotClientMsg) String() string { return proto.CompactTextString(m) }
func (*CAdditionalEquipSlotClientMsg) ProtoMessage()    {}

func (m *CAdditionalEquipSlotClientMsg) GetClassId() uint32 {
	if m != nil && m.ClassId != nil {
		return *m.ClassId
	}
	return 0
}

func (m *CAdditionalEquipSlotClientMsg) GetSlotId() uint32 {
	if m != nil && m.SlotId != nil {
		return *m.SlotId
	}
	return 0
}

func (m *CAdditionalEquipSlotClientMsg) GetDefIndex() uint32 {
	if m != nil && m.DefIndex != nil {
		return *m.DefIndex
	}
	return 0
}

type CDOTAClientMsg_FreeInventory struct {
	Equips           []*CAdditionalEquipSlotClientMsg `protobuf:"bytes,1,rep,name=equips" json:"equips,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *CDOTAClientMsg_FreeInventory) Reset()         { *m = CDOTAClientMsg_FreeInventory{} }
func (m *CDOTAClientMsg_FreeInventory) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_FreeInventory) ProtoMessage()    {}

func (m *CDOTAClientMsg_FreeInventory) GetEquips() []*CAdditionalEquipSlotClientMsg {
	if m != nil {
		return m.Equips
	}
	return nil
}

type CDOTAClientMsg_HeroStatueLike struct {
	OwnerPlayerId    *uint32 `protobuf:"varint,1,opt,name=owner_player_id" json:"owner_player_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAClientMsg_HeroStatueLike) Reset()         { *m = CDOTAClientMsg_HeroStatueLike{} }
func (m *CDOTAClientMsg_HeroStatueLike) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_HeroStatueLike) ProtoMessage()    {}

func (m *CDOTAClientMsg_HeroStatueLike) GetOwnerPlayerId() uint32 {
	if m != nil && m.OwnerPlayerId != nil {
		return *m.OwnerPlayerId
	}
	return 0
}

type CDOTAClientMsg_TeamShowcaseEditor struct {
	Data             []byte `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_TeamShowcaseEditor) Reset()         { *m = CDOTAClientMsg_TeamShowcaseEditor{} }
func (m *CDOTAClientMsg_TeamShowcaseEditor) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_TeamShowcaseEditor) ProtoMessage()    {}

func (m *CDOTAClientMsg_TeamShowcaseEditor) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDOTAClientMsg_TeamShowcaseClientData struct {
	Data             []byte `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_TeamShowcaseClientData) Reset()         { *m = CDOTAClientMsg_TeamShowcaseClientData{} }
func (m *CDOTAClientMsg_TeamShowcaseClientData) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_TeamShowcaseClientData) ProtoMessage()    {}

func (m *CDOTAClientMsg_TeamShowcaseClientData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CDOTAClientMsg_PlayTeamShowcase struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDOTAClientMsg_PlayTeamShowcase) Reset()         { *m = CDOTAClientMsg_PlayTeamShowcase{} }
func (m *CDOTAClientMsg_PlayTeamShowcase) String() string { return proto.CompactTextString(m) }
func (*CDOTAClientMsg_PlayTeamShowcase) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("dota.EDotaClientMessages", EDotaClientMessages_name, EDotaClientMessages_value)
}
