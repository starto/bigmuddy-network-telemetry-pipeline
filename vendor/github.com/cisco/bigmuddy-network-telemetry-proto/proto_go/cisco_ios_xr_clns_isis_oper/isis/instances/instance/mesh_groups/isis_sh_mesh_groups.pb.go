// Code generated by protoc-gen-go.
// source: isis_sh_mesh_groups.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_mesh_groups is a generated protocol buffer package.

It is generated from these files:
	isis_sh_mesh_groups.proto

It has these top-level messages:
	IsisShMeshGroups_KEYS
	IsisShMeshGroups
	IsisShMeshEntry
	IsisShMeshEntryItem
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_mesh_groups

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// IS-IS mesh-group data
type IsisShMeshGroups_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *IsisShMeshGroups_KEYS) Reset()                    { *m = IsisShMeshGroups_KEYS{} }
func (m *IsisShMeshGroups_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShMeshGroups_KEYS) ProtoMessage()               {}
func (*IsisShMeshGroups_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShMeshGroups_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type IsisShMeshGroups struct {
	// List of mesh-group-configured interfaces
	MeshGroupConfiguredInterfaceList *IsisShMeshEntry `protobuf:"bytes,50,opt,name=mesh_group_configured_interface_list,json=meshGroupConfiguredInterfaceList" json:"mesh_group_configured_interface_list,omitempty"`
}

func (m *IsisShMeshGroups) Reset()                    { *m = IsisShMeshGroups{} }
func (m *IsisShMeshGroups) String() string            { return proto.CompactTextString(m) }
func (*IsisShMeshGroups) ProtoMessage()               {}
func (*IsisShMeshGroups) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShMeshGroups) GetMeshGroupConfiguredInterfaceList() *IsisShMeshEntry {
	if m != nil {
		return m.MeshGroupConfiguredInterfaceList
	}
	return nil
}

// Mesh-group informaiton for one interface
type IsisShMeshEntry struct {
	// Next entry in list
	IsisShMeshEntry []*IsisShMeshEntryItem `protobuf:"bytes,1,rep,name=isis_sh_mesh_entry,json=isisShMeshEntry" json:"isis_sh_mesh_entry,omitempty"`
}

func (m *IsisShMeshEntry) Reset()                    { *m = IsisShMeshEntry{} }
func (m *IsisShMeshEntry) String() string            { return proto.CompactTextString(m) }
func (*IsisShMeshEntry) ProtoMessage()               {}
func (*IsisShMeshEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisShMeshEntry) GetIsisShMeshEntry() []*IsisShMeshEntryItem {
	if m != nil {
		return m.IsisShMeshEntry
	}
	return nil
}

type IsisShMeshEntryItem struct {
	// This interface
	MeshGroupInterface string `protobuf:"bytes,1,opt,name=mesh_group_interface,json=meshGroupInterface" json:"mesh_group_interface,omitempty"`
	// Mesh-group number
	MeshGroupNumber uint32 `protobuf:"varint,2,opt,name=mesh_group_number,json=meshGroupNumber" json:"mesh_group_number,omitempty"`
}

func (m *IsisShMeshEntryItem) Reset()                    { *m = IsisShMeshEntryItem{} }
func (m *IsisShMeshEntryItem) String() string            { return proto.CompactTextString(m) }
func (*IsisShMeshEntryItem) ProtoMessage()               {}
func (*IsisShMeshEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisShMeshEntryItem) GetMeshGroupInterface() string {
	if m != nil {
		return m.MeshGroupInterface
	}
	return ""
}

func (m *IsisShMeshEntryItem) GetMeshGroupNumber() uint32 {
	if m != nil {
		return m.MeshGroupNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*IsisShMeshGroups_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_groups_KEYS")
	proto.RegisterType((*IsisShMeshGroups)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_groups")
	proto.RegisterType((*IsisShMeshEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_entry")
	proto.RegisterType((*IsisShMeshEntryItem)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.mesh_groups.isis_sh_mesh_entry_item")
}

func init() { proto.RegisterFile("isis_sh_mesh_groups.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0x15, 0x04, 0xb7, 0x96, 0xe2, 0x2a, 0x18, 0x6f, 0x21, 0x7a, 0x08, 0x1e, 0x16, 0xa9,
	0x1f, 0xd0, 0x83, 0x14, 0x11, 0xb5, 0x87, 0xf4, 0xa2, 0xa7, 0x21, 0x5d, 0xa7, 0xcd, 0x42, 0xb3,
	0x1b, 0x76, 0x36, 0xa8, 0x1f, 0xe0, 0x5f, 0x78, 0xf3, 0x33, 0xfc, 0x39, 0x49, 0x30, 0x69, 0xc4,
	0x78, 0x12, 0x6f, 0x93, 0x37, 0xef, 0xbd, 0x99, 0x37, 0x59, 0x7e, 0xac, 0x49, 0x13, 0x50, 0x06,
	0x39, 0x52, 0x06, 0x2b, 0x67, 0xcb, 0x82, 0x64, 0xe1, 0xac, 0xb7, 0x62, 0xa2, 0x34, 0x29, 0x0b,
	0xda, 0x12, 0x3c, 0x3b, 0x50, 0x6b, 0x43, 0x50, 0x93, 0x6d, 0x81, 0x4e, 0x56, 0x95, 0xd4, 0x86,
	0x7c, 0x6a, 0x14, 0x6e, 0x2a, 0xd9, 0xb1, 0x89, 0x26, 0x3c, 0xe8, 0x71, 0x87, 0x9b, 0xe9, 0xc3,
	0x5c, 0x9c, 0xf0, 0x61, 0xa3, 0x01, 0x93, 0xe6, 0x18, 0xb0, 0x90, 0xc5, 0xbb, 0xc9, 0x5e, 0x03,
	0xce, 0xd2, 0x1c, 0xa3, 0x0f, 0xc6, 0x0f, 0x7a, 0x1c, 0xc4, 0x3b, 0xe3, 0xa7, 0x9b, 0x6f, 0x50,
	0xd6, 0x2c, 0xf5, 0xaa, 0x74, 0xf8, 0x08, 0xda, 0x78, 0x74, 0xcb, 0x54, 0x21, 0xac, 0x35, 0xf9,
	0x60, 0x1c, 0xb2, 0x78, 0x30, 0x9e, 0xcb, 0x3f, 0x26, 0x91, 0xdf, 0x96, 0x40, 0xe3, 0xdd, 0x4b,
	0x12, 0x56, 0xf5, 0x55, 0xd5, 0xbe, 0x6c, 0xc7, 0x5f, 0x37, 0xd3, 0x6f, 0x35, 0xf9, 0xe8, 0x8d,
	0x71, 0xf1, 0x53, 0x28, 0x5e, 0x7b, 0xe1, 0x80, 0x85, 0xdb, 0xf1, 0x60, 0x7c, 0xff, 0x0f, 0xab,
	0x82, 0xf6, 0x98, 0x27, 0xa3, 0xaa, 0x31, 0xcf, 0xee, 0x90, 0xb2, 0x69, 0x85, 0x46, 0x4f, 0xfc,
	0xe8, 0x17, 0xae, 0x38, 0xe7, 0x87, 0x9d, 0xf3, 0xb6, 0x37, 0xfd, 0xfa, 0x47, 0xa2, 0x4d, 0xde,
	0xe6, 0x15, 0x67, 0x7c, 0xbf, 0xa3, 0x30, 0x65, 0xbe, 0x40, 0x17, 0x6c, 0x85, 0x2c, 0x1e, 0x26,
	0xa3, 0x96, 0x3e, 0xab, 0xe1, 0xc5, 0x4e, 0xfd, 0xbc, 0x2e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x61, 0x02, 0x40, 0x7f, 0x7b, 0x02, 0x00, 0x00,
}
