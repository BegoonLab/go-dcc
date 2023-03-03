// controller.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: controller.proto

package controller

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

type Controller_Locomotive_Direction int32

const (
	Controller_Locomotive_Backward Controller_Locomotive_Direction = 0
	Controller_Locomotive_Forward  Controller_Locomotive_Direction = 1
)

// Enum value maps for Controller_Locomotive_Direction.
var (
	Controller_Locomotive_Direction_name = map[int32]string{
		0: "Backward",
		1: "Forward",
	}
	Controller_Locomotive_Direction_value = map[string]int32{
		"Backward": 0,
		"Forward":  1,
	}
)

func (x Controller_Locomotive_Direction) Enum() *Controller_Locomotive_Direction {
	p := new(Controller_Locomotive_Direction)
	*p = x
	return p
}

func (x Controller_Locomotive_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Controller_Locomotive_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_proto_enumTypes[0].Descriptor()
}

func (Controller_Locomotive_Direction) Type() protoreflect.EnumType {
	return &file_controller_proto_enumTypes[0]
}

func (x Controller_Locomotive_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Controller_Locomotive_Direction.Descriptor instead.
func (Controller_Locomotive_Direction) EnumDescriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0, 2, 0}
}

type Controller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locomotives    map[string]*Controller_Locomotive    `protobuf:"bytes,1,rep,name=locomotives,proto3" json:"locomotives,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RailwayModules map[string]*Controller_RailwayModule `protobuf:"bytes,2,rep,name=railway_modules,json=railwayModules,proto3" json:"railway_modules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Started        bool                                 `protobuf:"varint,3,opt,name=started,proto3" json:"started,omitempty"`
	Reboot         bool                                 `protobuf:"varint,4,opt,name=reboot,proto3" json:"reboot,omitempty"`
	Poweroff       bool                                 `protobuf:"varint,5,opt,name=poweroff,proto3" json:"poweroff,omitempty"`
	Id             string                               `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Controller) Reset() {
	*x = Controller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller) ProtoMessage() {}

func (x *Controller) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller.ProtoReflect.Descriptor instead.
func (*Controller) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *Controller) GetLocomotives() map[string]*Controller_Locomotive {
	if x != nil {
		return x.Locomotives
	}
	return nil
}

func (x *Controller) GetRailwayModules() map[string]*Controller_RailwayModule {
	if x != nil {
		return x.RailwayModules
	}
	return nil
}

func (x *Controller) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *Controller) GetReboot() bool {
	if x != nil {
		return x.Reboot
	}
	return false
}

func (x *Controller) GetPoweroff() bool {
	if x != nil {
		return x.Poweroff
	}
	return false
}

func (x *Controller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Controller_Locomotive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address   uint32                          `protobuf:"varint,2,opt,name=address,proto3" json:"address,omitempty"`
	Speed     uint32                          `protobuf:"varint,3,opt,name=speed,proto3" json:"speed,omitempty"`
	Direction Controller_Locomotive_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=controller.Controller_Locomotive_Direction" json:"direction,omitempty"`
	Enabled   bool                            `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Fl        bool                            `protobuf:"varint,6,opt,name=fl,proto3" json:"fl,omitempty"`
	F1        bool                            `protobuf:"varint,7,opt,name=f1,proto3" json:"f1,omitempty"`
	F2        bool                            `protobuf:"varint,8,opt,name=f2,proto3" json:"f2,omitempty"`
	F3        bool                            `protobuf:"varint,9,opt,name=f3,proto3" json:"f3,omitempty"`
	F4        bool                            `protobuf:"varint,10,opt,name=f4,proto3" json:"f4,omitempty"`
	F5        bool                            `protobuf:"varint,11,opt,name=f5,proto3" json:"f5,omitempty"`
	F6        bool                            `protobuf:"varint,12,opt,name=f6,proto3" json:"f6,omitempty"`
	F7        bool                            `protobuf:"varint,13,opt,name=f7,proto3" json:"f7,omitempty"`
	F8        bool                            `protobuf:"varint,14,opt,name=f8,proto3" json:"f8,omitempty"`
	F9        bool                            `protobuf:"varint,15,opt,name=f9,proto3" json:"f9,omitempty"`
	F10       bool                            `protobuf:"varint,16,opt,name=f10,proto3" json:"f10,omitempty"`
	F11       bool                            `protobuf:"varint,17,opt,name=f11,proto3" json:"f11,omitempty"`
	F12       bool                            `protobuf:"varint,18,opt,name=f12,proto3" json:"f12,omitempty"`
	F13       bool                            `protobuf:"varint,19,opt,name=f13,proto3" json:"f13,omitempty"`
	F14       bool                            `protobuf:"varint,20,opt,name=f14,proto3" json:"f14,omitempty"`
	F15       bool                            `protobuf:"varint,21,opt,name=f15,proto3" json:"f15,omitempty"`
	F16       bool                            `protobuf:"varint,22,opt,name=f16,proto3" json:"f16,omitempty"`
	F17       bool                            `protobuf:"varint,23,opt,name=f17,proto3" json:"f17,omitempty"`
	F18       bool                            `protobuf:"varint,24,opt,name=f18,proto3" json:"f18,omitempty"`
	F19       bool                            `protobuf:"varint,25,opt,name=f19,proto3" json:"f19,omitempty"`
	F20       bool                            `protobuf:"varint,26,opt,name=f20,proto3" json:"f20,omitempty"`
	F21       bool                            `protobuf:"varint,27,opt,name=f21,proto3" json:"f21,omitempty"`
	F22       bool                            `protobuf:"varint,28,opt,name=f22,proto3" json:"f22,omitempty"`
	F23       bool                            `protobuf:"varint,29,opt,name=f23,proto3" json:"f23,omitempty"`
	F24       bool                            `protobuf:"varint,30,opt,name=f24,proto3" json:"f24,omitempty"`
	F25       bool                            `protobuf:"varint,31,opt,name=f25,proto3" json:"f25,omitempty"`
	F26       bool                            `protobuf:"varint,32,opt,name=f26,proto3" json:"f26,omitempty"`
	F27       bool                            `protobuf:"varint,33,opt,name=f27,proto3" json:"f27,omitempty"`
	F28       bool                            `protobuf:"varint,34,opt,name=f28,proto3" json:"f28,omitempty"`
}

func (x *Controller_Locomotive) Reset() {
	*x = Controller_Locomotive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller_Locomotive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller_Locomotive) ProtoMessage() {}

func (x *Controller_Locomotive) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller_Locomotive.ProtoReflect.Descriptor instead.
func (*Controller_Locomotive) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Controller_Locomotive) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Controller_Locomotive) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *Controller_Locomotive) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Controller_Locomotive) GetDirection() Controller_Locomotive_Direction {
	if x != nil {
		return x.Direction
	}
	return Controller_Locomotive_Backward
}

func (x *Controller_Locomotive) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Controller_Locomotive) GetFl() bool {
	if x != nil {
		return x.Fl
	}
	return false
}

func (x *Controller_Locomotive) GetF1() bool {
	if x != nil {
		return x.F1
	}
	return false
}

func (x *Controller_Locomotive) GetF2() bool {
	if x != nil {
		return x.F2
	}
	return false
}

func (x *Controller_Locomotive) GetF3() bool {
	if x != nil {
		return x.F3
	}
	return false
}

func (x *Controller_Locomotive) GetF4() bool {
	if x != nil {
		return x.F4
	}
	return false
}

func (x *Controller_Locomotive) GetF5() bool {
	if x != nil {
		return x.F5
	}
	return false
}

func (x *Controller_Locomotive) GetF6() bool {
	if x != nil {
		return x.F6
	}
	return false
}

func (x *Controller_Locomotive) GetF7() bool {
	if x != nil {
		return x.F7
	}
	return false
}

func (x *Controller_Locomotive) GetF8() bool {
	if x != nil {
		return x.F8
	}
	return false
}

func (x *Controller_Locomotive) GetF9() bool {
	if x != nil {
		return x.F9
	}
	return false
}

func (x *Controller_Locomotive) GetF10() bool {
	if x != nil {
		return x.F10
	}
	return false
}

func (x *Controller_Locomotive) GetF11() bool {
	if x != nil {
		return x.F11
	}
	return false
}

func (x *Controller_Locomotive) GetF12() bool {
	if x != nil {
		return x.F12
	}
	return false
}

func (x *Controller_Locomotive) GetF13() bool {
	if x != nil {
		return x.F13
	}
	return false
}

func (x *Controller_Locomotive) GetF14() bool {
	if x != nil {
		return x.F14
	}
	return false
}

func (x *Controller_Locomotive) GetF15() bool {
	if x != nil {
		return x.F15
	}
	return false
}

func (x *Controller_Locomotive) GetF16() bool {
	if x != nil {
		return x.F16
	}
	return false
}

func (x *Controller_Locomotive) GetF17() bool {
	if x != nil {
		return x.F17
	}
	return false
}

func (x *Controller_Locomotive) GetF18() bool {
	if x != nil {
		return x.F18
	}
	return false
}

func (x *Controller_Locomotive) GetF19() bool {
	if x != nil {
		return x.F19
	}
	return false
}

func (x *Controller_Locomotive) GetF20() bool {
	if x != nil {
		return x.F20
	}
	return false
}

func (x *Controller_Locomotive) GetF21() bool {
	if x != nil {
		return x.F21
	}
	return false
}

func (x *Controller_Locomotive) GetF22() bool {
	if x != nil {
		return x.F22
	}
	return false
}

func (x *Controller_Locomotive) GetF23() bool {
	if x != nil {
		return x.F23
	}
	return false
}

func (x *Controller_Locomotive) GetF24() bool {
	if x != nil {
		return x.F24
	}
	return false
}

func (x *Controller_Locomotive) GetF25() bool {
	if x != nil {
		return x.F25
	}
	return false
}

func (x *Controller_Locomotive) GetF26() bool {
	if x != nil {
		return x.F26
	}
	return false
}

func (x *Controller_Locomotive) GetF27() bool {
	if x != nil {
		return x.F27
	}
	return false
}

func (x *Controller_Locomotive) GetF28() bool {
	if x != nil {
		return x.F28
	}
	return false
}

type Controller_RailwayModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address     uint32                                     `protobuf:"varint,2,opt,name=address,proto3" json:"address,omitempty"`
	Enabled     bool                                       `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Routes      map[string]*Controller_RailwayModule_Route `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ActiveRoute string                                     `protobuf:"bytes,5,opt,name=activeRoute,proto3" json:"activeRoute,omitempty"`
}

func (x *Controller_RailwayModule) Reset() {
	*x = Controller_RailwayModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller_RailwayModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller_RailwayModule) ProtoMessage() {}

func (x *Controller_RailwayModule) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller_RailwayModule.ProtoReflect.Descriptor instead.
func (*Controller_RailwayModule) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Controller_RailwayModule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Controller_RailwayModule) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *Controller_RailwayModule) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Controller_RailwayModule) GetRoutes() map[string]*Controller_RailwayModule_Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Controller_RailwayModule) GetActiveRoute() string {
	if x != nil {
		return x.ActiveRoute
	}
	return ""
}

type Controller_RailwayModule_Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	F0   bool   `protobuf:"varint,2,opt,name=f0,proto3" json:"f0,omitempty"`
	F1   bool   `protobuf:"varint,3,opt,name=f1,proto3" json:"f1,omitempty"`
	F2   bool   `protobuf:"varint,4,opt,name=f2,proto3" json:"f2,omitempty"`
	F3   bool   `protobuf:"varint,5,opt,name=f3,proto3" json:"f3,omitempty"`
	F4   bool   `protobuf:"varint,6,opt,name=f4,proto3" json:"f4,omitempty"`
	F5   bool   `protobuf:"varint,7,opt,name=f5,proto3" json:"f5,omitempty"`
	F6   bool   `protobuf:"varint,8,opt,name=f6,proto3" json:"f6,omitempty"`
	F7   bool   `protobuf:"varint,9,opt,name=f7,proto3" json:"f7,omitempty"`
	F8   bool   `protobuf:"varint,10,opt,name=f8,proto3" json:"f8,omitempty"`
	F9   bool   `protobuf:"varint,11,opt,name=f9,proto3" json:"f9,omitempty"`
	F10  bool   `protobuf:"varint,12,opt,name=f10,proto3" json:"f10,omitempty"`
	F11  bool   `protobuf:"varint,13,opt,name=f11,proto3" json:"f11,omitempty"`
	F12  bool   `protobuf:"varint,14,opt,name=f12,proto3" json:"f12,omitempty"`
	F13  bool   `protobuf:"varint,15,opt,name=f13,proto3" json:"f13,omitempty"`
	F14  bool   `protobuf:"varint,16,opt,name=f14,proto3" json:"f14,omitempty"`
	F15  bool   `protobuf:"varint,17,opt,name=f15,proto3" json:"f15,omitempty"`
	F16  bool   `protobuf:"varint,18,opt,name=f16,proto3" json:"f16,omitempty"`
	F17  bool   `protobuf:"varint,19,opt,name=f17,proto3" json:"f17,omitempty"`
	F18  bool   `protobuf:"varint,20,opt,name=f18,proto3" json:"f18,omitempty"`
	F19  bool   `protobuf:"varint,21,opt,name=f19,proto3" json:"f19,omitempty"`
	F20  bool   `protobuf:"varint,22,opt,name=f20,proto3" json:"f20,omitempty"`
	F21  bool   `protobuf:"varint,23,opt,name=f21,proto3" json:"f21,omitempty"`
	F22  bool   `protobuf:"varint,24,opt,name=f22,proto3" json:"f22,omitempty"`
	F23  bool   `protobuf:"varint,25,opt,name=f23,proto3" json:"f23,omitempty"`
	F24  bool   `protobuf:"varint,26,opt,name=f24,proto3" json:"f24,omitempty"`
	F25  bool   `protobuf:"varint,27,opt,name=f25,proto3" json:"f25,omitempty"`
	F26  bool   `protobuf:"varint,28,opt,name=f26,proto3" json:"f26,omitempty"`
	F27  bool   `protobuf:"varint,29,opt,name=f27,proto3" json:"f27,omitempty"`
	F28  bool   `protobuf:"varint,30,opt,name=f28,proto3" json:"f28,omitempty"`
	F29  bool   `protobuf:"varint,31,opt,name=f29,proto3" json:"f29,omitempty"`
	F30  bool   `protobuf:"varint,32,opt,name=f30,proto3" json:"f30,omitempty"`
	F31  bool   `protobuf:"varint,33,opt,name=f31,proto3" json:"f31,omitempty"`
}

func (x *Controller_RailwayModule_Route) Reset() {
	*x = Controller_RailwayModule_Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller_RailwayModule_Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller_RailwayModule_Route) ProtoMessage() {}

func (x *Controller_RailwayModule_Route) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller_RailwayModule_Route.ProtoReflect.Descriptor instead.
func (*Controller_RailwayModule_Route) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0, 3, 1}
}

func (x *Controller_RailwayModule_Route) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Controller_RailwayModule_Route) GetF0() bool {
	if x != nil {
		return x.F0
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF1() bool {
	if x != nil {
		return x.F1
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF2() bool {
	if x != nil {
		return x.F2
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF3() bool {
	if x != nil {
		return x.F3
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF4() bool {
	if x != nil {
		return x.F4
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF5() bool {
	if x != nil {
		return x.F5
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF6() bool {
	if x != nil {
		return x.F6
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF7() bool {
	if x != nil {
		return x.F7
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF8() bool {
	if x != nil {
		return x.F8
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF9() bool {
	if x != nil {
		return x.F9
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF10() bool {
	if x != nil {
		return x.F10
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF11() bool {
	if x != nil {
		return x.F11
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF12() bool {
	if x != nil {
		return x.F12
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF13() bool {
	if x != nil {
		return x.F13
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF14() bool {
	if x != nil {
		return x.F14
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF15() bool {
	if x != nil {
		return x.F15
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF16() bool {
	if x != nil {
		return x.F16
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF17() bool {
	if x != nil {
		return x.F17
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF18() bool {
	if x != nil {
		return x.F18
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF19() bool {
	if x != nil {
		return x.F19
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF20() bool {
	if x != nil {
		return x.F20
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF21() bool {
	if x != nil {
		return x.F21
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF22() bool {
	if x != nil {
		return x.F22
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF23() bool {
	if x != nil {
		return x.F23
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF24() bool {
	if x != nil {
		return x.F24
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF25() bool {
	if x != nil {
		return x.F25
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF26() bool {
	if x != nil {
		return x.F26
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF27() bool {
	if x != nil {
		return x.F27
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF28() bool {
	if x != nil {
		return x.F28
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF29() bool {
	if x != nil {
		return x.F29
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF30() bool {
	if x != nil {
		return x.F30
	}
	return false
}

func (x *Controller_RailwayModule_Route) GetF31() bool {
	if x != nil {
		return x.F31
	}
	return false
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0xa3,
	0x10, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x49, 0x0a,
	0x0b, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6c, 0x6f, 0x63,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x0f, 0x72, 0x61, 0x69, 0x6c,
	0x77, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x69, 0x6c, 0x77, 0x61,
	0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72,
	0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x62, 0x6f, 0x6f,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x6f, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x6f, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x61, 0x0a, 0x10, 0x4c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x67,
	0x0a, 0x13, 0x52, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x52,
	0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xd3, 0x05, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x76, 0x65,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x31, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x32, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x33, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x33, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x34, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x35, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x36, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x36, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x37, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x37, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x38, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x38, 0x12,
	0x0e, 0x0a, 0x02, 0x66, 0x39, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x39, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x31, 0x30, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31,
	0x30, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x31, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x66, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x32, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x66, 0x31, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x33, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x34, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x35,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x31, 0x36, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x36, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x31, 0x37, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x37, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x31, 0x38, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31,
	0x38, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x39, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x66, 0x31, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x30, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x66, 0x32, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x31, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x32, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x33,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x32, 0x34, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x34, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x32, 0x35, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x35, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x32, 0x36, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32,
	0x36, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x37, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x66, 0x32, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x38, 0x18, 0x22, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x66, 0x32, 0x38, 0x22, 0x26, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x10, 0x01, 0x1a, 0xf4, 0x06,
	0x0a, 0x0d, 0x52, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x52, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x1a, 0x65, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x69, 0x6c,
	0x77, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xc7, 0x04, 0x0a, 0x05, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x30, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x32, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x33, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x34, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x36, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x37, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x37, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x38, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x39, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x66, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x30, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31,
	0x31, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x31, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x32, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x31, 0x33, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x33,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x34, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66,
	0x31, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x35, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x66, 0x31, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x36, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x66, 0x31, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x37, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x37, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x38, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31,
	0x39, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x31, 0x39, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x32, 0x30, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x30, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x32, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x31,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x32, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66,
	0x32, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x33, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x66, 0x32, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x34, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x66, 0x32, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x35, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32, 0x36, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x32,
	0x37, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x37, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x32, 0x38, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x38, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x32, 0x39, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x32, 0x39,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x33, 0x30, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66,
	0x33, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x33, 0x31, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x66, 0x33, 0x31, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_controller_proto_goTypes = []interface{}{
	(Controller_Locomotive_Direction)(0),   // 0: controller.Controller.Locomotive.Direction
	(*Controller)(nil),                     // 1: controller.Controller
	nil,                                    // 2: controller.Controller.LocomotivesEntry
	nil,                                    // 3: controller.Controller.RailwayModulesEntry
	(*Controller_Locomotive)(nil),          // 4: controller.Controller.Locomotive
	(*Controller_RailwayModule)(nil),       // 5: controller.Controller.RailwayModule
	nil,                                    // 6: controller.Controller.RailwayModule.RoutesEntry
	(*Controller_RailwayModule_Route)(nil), // 7: controller.Controller.RailwayModule.Route
}
var file_controller_proto_depIdxs = []int32{
	2, // 0: controller.Controller.locomotives:type_name -> controller.Controller.LocomotivesEntry
	3, // 1: controller.Controller.railway_modules:type_name -> controller.Controller.RailwayModulesEntry
	4, // 2: controller.Controller.LocomotivesEntry.value:type_name -> controller.Controller.Locomotive
	5, // 3: controller.Controller.RailwayModulesEntry.value:type_name -> controller.Controller.RailwayModule
	0, // 4: controller.Controller.Locomotive.direction:type_name -> controller.Controller.Locomotive.Direction
	6, // 5: controller.Controller.RailwayModule.routes:type_name -> controller.Controller.RailwayModule.RoutesEntry
	7, // 6: controller.Controller.RailwayModule.RoutesEntry.value:type_name -> controller.Controller.RailwayModule.Route
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller); i {
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
		file_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller_Locomotive); i {
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
		file_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller_RailwayModule); i {
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
		file_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller_RailwayModule_Route); i {
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
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		EnumInfos:         file_controller_proto_enumTypes,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}
