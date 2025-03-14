// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/api_gateway.proto

package meshv2beta1

import (
	pbproxystate "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1/pbproxystate"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
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

type APIGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GatewayClassName is the name of the GatewayClass used by the APIGateway
	GatewayClassName string `protobuf:"bytes,1,opt,name=gateway_class_name,json=gatewayClassName,proto3" json:"gateway_class_name,omitempty"`
	// +kubebuilder:validation:MinItems=1
	Listeners []*APIGatewayListener `protobuf:"bytes,2,rep,name=listeners,proto3" json:"listeners,omitempty"`
}

func (x *APIGateway) Reset() {
	*x = APIGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIGateway) ProtoMessage() {}

func (x *APIGateway) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIGateway.ProtoReflect.Descriptor instead.
func (*APIGateway) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_api_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *APIGateway) GetGatewayClassName() string {
	if x != nil {
		return x.GatewayClassName
	}
	return ""
}

func (x *APIGateway) GetListeners() []*APIGatewayListener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

type APIGatewayListener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the listener in a given gateway. This must be
	// unique within a gateway.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=0
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Protocol is the protocol that a listener should use, it must
	// either be "http" or "tcp"
	// +kubebuilder:validation≈ftg6:Enum=tcp,http
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Hostname is the host name that a listener should be bound to, if
	// unspecified, the listener accepts requests for all hostnames.
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// TLS is the TLS settings for the listener.
	Tls *APIGatewayTLSConfiguration `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty"` //TODO NET-7311 @Gateway-Management
}

func (x *APIGatewayListener) Reset() {
	*x = APIGatewayListener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIGatewayListener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIGatewayListener) ProtoMessage() {}

func (x *APIGatewayListener) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIGatewayListener.ProtoReflect.Descriptor instead.
func (*APIGatewayListener) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_api_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *APIGatewayListener) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *APIGatewayListener) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *APIGatewayListener) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *APIGatewayListener) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *APIGatewayListener) GetTls() *APIGatewayTLSConfiguration {
	if x != nil {
		return x.Tls
	}
	return nil
}

// APIGatewayTLSConfiguration specifies the configuration of a listener’s
// TLS settings.
type APIGatewayTLSConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Certificates is a set of references to certificates
	// that a gateway listener uses for TLS termination.
	Certificates []*pbresource.Reference `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	// TLSParameters contains optional configuration for running TLS termination.
	TlsParameters *pbproxystate.TLSParameters `protobuf:"bytes,2,opt,name=tls_parameters,json=tlsParameters,proto3,oneof" json:"tls_parameters,omitempty"`
}

func (x *APIGatewayTLSConfiguration) Reset() {
	*x = APIGatewayTLSConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIGatewayTLSConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIGatewayTLSConfiguration) ProtoMessage() {}

func (x *APIGatewayTLSConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_api_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIGatewayTLSConfiguration.ProtoReflect.Descriptor instead.
func (*APIGatewayTLSConfiguration) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_api_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *APIGatewayTLSConfiguration) GetCertificates() []*pbresource.Reference {
	if x != nil {
		return x.Certificates
	}
	return nil
}

func (x *APIGatewayTLSConfiguration) GetTlsParameters() *pbproxystate.TLSParameters {
	if x != nil {
		return x.TlsParameters
	}
	return nil
}

var File_pbmesh_v2beta1_api_gateway_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_api_gateway_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x32, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93,
	0x01, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41,
	0x50, 0x49, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x3a, 0x06, 0xa2, 0x93,
	0x04, 0x02, 0x08, 0x03, 0x22, 0xc1, 0x01, 0x0a, 0x12, 0x41, 0x50, 0x49, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x03, 0x74,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x1a, 0x41, 0x50, 0x49,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x65, 0x0a, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x4c, 0x53, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x74, 0x6c, 0x73,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42, 0x90, 0x02, 0x0a, 0x21,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70,
	0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x65,
	0x73, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa,
	0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x29, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a,
	0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_api_gateway_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_api_gateway_proto_rawDescData = file_pbmesh_v2beta1_api_gateway_proto_rawDesc
)

func file_pbmesh_v2beta1_api_gateway_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_api_gateway_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_api_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_api_gateway_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_api_gateway_proto_rawDescData
}

var file_pbmesh_v2beta1_api_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pbmesh_v2beta1_api_gateway_proto_goTypes = []interface{}{
	(*APIGateway)(nil),                 // 0: hashicorp.consul.mesh.v2beta1.APIGateway
	(*APIGatewayListener)(nil),         // 1: hashicorp.consul.mesh.v2beta1.APIGatewayListener
	(*APIGatewayTLSConfiguration)(nil), // 2: hashicorp.consul.mesh.v2beta1.APIGatewayTLSConfiguration
	(*pbresource.Reference)(nil),       // 3: hashicorp.consul.resource.Reference
	(*pbproxystate.TLSParameters)(nil), // 4: hashicorp.consul.mesh.v2beta1.pbproxystate.TLSParameters
}
var file_pbmesh_v2beta1_api_gateway_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.mesh.v2beta1.APIGateway.listeners:type_name -> hashicorp.consul.mesh.v2beta1.APIGatewayListener
	2, // 1: hashicorp.consul.mesh.v2beta1.APIGatewayListener.tls:type_name -> hashicorp.consul.mesh.v2beta1.APIGatewayTLSConfiguration
	3, // 2: hashicorp.consul.mesh.v2beta1.APIGatewayTLSConfiguration.certificates:type_name -> hashicorp.consul.resource.Reference
	4, // 3: hashicorp.consul.mesh.v2beta1.APIGatewayTLSConfiguration.tls_parameters:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.TLSParameters
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_api_gateway_proto_init() }
func file_pbmesh_v2beta1_api_gateway_proto_init() {
	if File_pbmesh_v2beta1_api_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_api_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIGateway); i {
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
		file_pbmesh_v2beta1_api_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIGatewayListener); i {
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
		file_pbmesh_v2beta1_api_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIGatewayTLSConfiguration); i {
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
	file_pbmesh_v2beta1_api_gateway_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_api_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_api_gateway_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_api_gateway_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_api_gateway_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_api_gateway_proto = out.File
	file_pbmesh_v2beta1_api_gateway_proto_rawDesc = nil
	file_pbmesh_v2beta1_api_gateway_proto_goTypes = nil
	file_pbmesh_v2beta1_api_gateway_proto_depIdxs = nil
}
