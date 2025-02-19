// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta1/genesis.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisDeployment defines the basic genesis state used by deployment module
type GenesisDeployment struct {
	Deployment Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment" yaml:"deployment"`
	Groups     []Group    `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups" yaml:"groups"`
}

func (m *GenesisDeployment) Reset()         { *m = GenesisDeployment{} }
func (m *GenesisDeployment) String() string { return proto.CompactTextString(m) }
func (*GenesisDeployment) ProtoMessage()    {}
func (*GenesisDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea837e5a570e958, []int{0}
}
func (m *GenesisDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDeployment.Merge(m, src)
}
func (m *GenesisDeployment) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDeployment proto.InternalMessageInfo

func (m *GenesisDeployment) GetDeployment() Deployment {
	if m != nil {
		return m.Deployment
	}
	return Deployment{}
}

func (m *GenesisDeployment) GetGroups() []Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

// GenesisState stores slice of genesis deployment instance
type GenesisState struct {
	Deployments []GenesisDeployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments" yaml:"deployments"`
	Params      Params              `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea837e5a570e958, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetDeployments() []GenesisDeployment {
	if m != nil {
		return m.Deployments
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisDeployment)(nil), "akash.deployment.v1beta1.GenesisDeployment")
	proto.RegisterType((*GenesisState)(nil), "akash.deployment.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta1/genesis.proto", fileDescriptor_8ea837e5a570e958)
}

var fileDescriptor_8ea837e5a570e958 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x6e, 0xf2, 0x30,
	0x14, 0x85, 0x63, 0x7e, 0x89, 0xc1, 0xfc, 0x1d, 0x88, 0x3a, 0x44, 0x0c, 0x31, 0xb2, 0x68, 0x0b,
	0xaa, 0x94, 0x08, 0xba, 0x75, 0x8c, 0x2a, 0xb1, 0x74, 0xa8, 0xd2, 0xa5, 0xed, 0x66, 0xa8, 0x15,
	0x10, 0x04, 0x47, 0xb1, 0x41, 0xf0, 0x16, 0x7d, 0x2c, 0x46, 0xc6, 0x0e, 0x55, 0x54, 0x91, 0xad,
	0x23, 0x4f, 0x50, 0x61, 0xbb, 0x75, 0x84, 0xe4, 0x6e, 0x71, 0xee, 0x77, 0xee, 0xb9, 0xe7, 0xea,
	0xc2, 0x4b, 0x32, 0x23, 0x7c, 0x12, 0xbe, 0xd2, 0x6c, 0xce, 0x36, 0x29, 0x5d, 0x88, 0x70, 0xd5,
	0x1f, 0x51, 0x41, 0xfa, 0x61, 0x42, 0x17, 0x94, 0x4f, 0x79, 0x90, 0xe5, 0x4c, 0x30, 0xd7, 0x93,
	0x5c, 0x60, 0xb8, 0x40, 0x73, 0xad, 0xf3, 0x84, 0x25, 0x4c, 0x42, 0xe1, 0xf1, 0x4b, 0xf1, 0xad,
	0x9e, 0xb5, 0x6f, 0xa5, 0x85, 0x42, 0x3b, 0xf6, 0x11, 0x72, 0xb6, 0xcc, 0x34, 0x75, 0x61, 0xa5,
	0x32, 0x92, 0x93, 0x54, 0xcf, 0x89, 0x3f, 0x00, 0x6c, 0x0e, 0xd5, 0xe4, 0x77, 0xbf, 0xa8, 0x9b,
	0x42, 0x68, 0x84, 0x1e, 0x68, 0x83, 0x6e, 0x63, 0xd0, 0x09, 0x6c, 0x91, 0x02, 0xa3, 0x8c, 0xae,
	0xb6, 0x05, 0x72, 0xbe, 0x0a, 0x54, 0xd1, 0x1f, 0x0a, 0xd4, 0xdc, 0x90, 0x74, 0x7e, 0x8b, 0xcd,
	0x3f, 0x1c, 0x57, 0x00, 0xf7, 0x09, 0xd6, 0xe5, 0xe8, 0xdc, 0xab, 0xb5, 0xff, 0x75, 0x1b, 0x03,
	0x64, 0xb7, 0x1a, 0x1e, 0xb9, 0x08, 0x69, 0x17, 0x2d, 0x3b, 0x14, 0xe8, 0x4c, 0x39, 0xa8, 0x37,
	0x8e, 0x75, 0x01, 0x97, 0x00, 0xfe, 0xd7, 0xf1, 0x1e, 0x05, 0x11, 0xd4, 0x5d, 0xc3, 0x86, 0xe9,
	0xca, 0x3d, 0x20, 0xfd, 0xae, 0xff, 0xf0, 0x3b, 0xdd, 0x4d, 0xd4, 0xd3, 0xde, 0xd5, 0x3e, 0x87,
	0x02, 0xb9, 0xa7, 0x11, 0x39, 0x8e, 0xab, 0x88, 0xfb, 0x0c, 0xeb, 0x6a, 0xf3, 0x5e, 0x4d, 0xee,
	0xb3, 0x6d, 0x37, 0x7d, 0x90, 0x9c, 0x49, 0xa9, 0x74, 0x26, 0xa5, 0x7a, 0xe3, 0x58, 0x17, 0xa2,
	0xfb, 0xed, 0xde, 0x07, 0xbb, 0xbd, 0x0f, 0x3e, 0xf7, 0x3e, 0x78, 0x2b, 0x7d, 0x67, 0x57, 0xfa,
	0xce, 0x7b, 0xe9, 0x3b, 0x2f, 0x83, 0x64, 0x2a, 0x26, 0xcb, 0x51, 0x30, 0x66, 0x69, 0xc8, 0x56,
	0xf9, 0x78, 0x3e, 0x0b, 0xd5, 0x5d, 0xac, 0xab, 0x97, 0x21, 0x36, 0x19, 0xe5, 0x3f, 0xf7, 0x31,
	0xaa, 0xcb, 0xcb, 0xb8, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x5d, 0x97, 0xbc, 0xeb, 0x02,
	0x00, 0x00,
}

func (m *GenesisDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Deployment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Deployments) > 0 {
		for iNdEx := len(m.Deployments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deployments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deployment.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Deployments) > 0 {
		for _, e := range m.Deployments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deployment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, Group{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deployments = append(m.Deployments, GenesisDeployment{})
			if err := m.Deployments[len(m.Deployments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
