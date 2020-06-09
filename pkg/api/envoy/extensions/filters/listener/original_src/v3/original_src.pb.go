// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/listener/original_src/v3/original_src.proto

package envoy_extensions_filters_listener_original_src_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// The Original Src filter binds upstream connections to the original source address determined
// for the connection. This address could come from something like the Proxy Protocol filter, or it
// could come from trusted http headers.
type OriginalSrc struct {
	// Whether to bind the port to the one used in the original downstream connection.
	// [#not-implemented-hide:]
	BindPort bool `protobuf:"varint,1,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// Sets the SO_MARK option on the upstream connection's socket to the provided value. Used to
	// ensure that non-local addresses may be routed back through envoy when binding to the original
	// source address. The option will not be applied if the mark is 0.
	Mark                 uint32   `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d80cee6a73e34, []int{0}
}
func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return m.Size()
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetBindPort() bool {
	if m != nil {
		return m.BindPort
	}
	return false
}

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.extensions.filters.listener.original_src.v3.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/original_src/v3/original_src.proto", fileDescriptor_0a2d80cee6a73e34)
}

var fileDescriptor_0a2d80cee6a73e34 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0xcf, 0x2f,
	0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x33, 0x46, 0xe1, 0xeb,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x19, 0x82, 0x4d, 0xd1, 0x43, 0x98, 0xa2, 0x07, 0x35, 0x45,
	0x0f, 0x66, 0x8a, 0x1e, 0x8a, 0xae, 0x32, 0x63, 0x29, 0xd9, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x12, 0xb0, 0xc5, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x10, 0x13,
	0xa5, 0x14, 0x31, 0xa4, 0xcb, 0x52, 0x8b, 0x40, 0x46, 0x67, 0xe6, 0xa5, 0x43, 0x95, 0x88, 0x97,
	0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x56, 0x46, 0x2e,
	0x6e, 0x7f, 0xa8, 0x75, 0xc1, 0x45, 0xc9, 0x42, 0xd2, 0x5c, 0x9c, 0x49, 0x99, 0x79, 0x29, 0xf1,
	0x05, 0xf9, 0x45, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x1c, 0x20, 0x81, 0x80, 0xfc,
	0xa2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xdc, 0xc4, 0xa2, 0x6c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xde,
	0x20, 0x30, 0xdb, 0xca, 0x75, 0xd6, 0xd1, 0x0e, 0x39, 0x07, 0x2e, 0x3b, 0x88, 0xaf, 0x92, 0xf3,
	0xf3, 0xd2, 0x32, 0xd3, 0xa1, 0x3e, 0xc2, 0xe5, 0x21, 0xa3, 0xc4, 0x9c, 0x82, 0x8c, 0x44, 0x43,
	0x3d, 0x24, 0x7b, 0x9d, 0x32, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x5d, 0x0d, 0x27, 0x2e, 0xb2, 0x31, 0x09, 0x30, 0x72, 0xd9, 0x67, 0xe6, 0xeb, 0x81,
	0x0d, 0x2e, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0x39, 0xe4, 0x9c, 0x04, 0x90, 0xec, 0x08, 0x00,
	0x79, 0x38, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x73, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0xd3, 0xe0, 0x54, 0xcf, 0x01, 0x00, 0x00,
}

func (m *OriginalSrc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginalSrc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginalSrc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Mark != 0 {
		i = encodeVarintOriginalSrc(dAtA, i, uint64(m.Mark))
		i--
		dAtA[i] = 0x10
	}
	if m.BindPort {
		i--
		if m.BindPort {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOriginalSrc(dAtA []byte, offset int, v uint64) int {
	offset -= sovOriginalSrc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OriginalSrc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BindPort {
		n += 2
	}
	if m.Mark != 0 {
		n += 1 + sovOriginalSrc(uint64(m.Mark))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOriginalSrc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOriginalSrc(x uint64) (n int) {
	return sovOriginalSrc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OriginalSrc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOriginalSrc
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
			return fmt.Errorf("proto: OriginalSrc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginalSrc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindPort", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginalSrc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BindPort = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mark", wireType)
			}
			m.Mark = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOriginalSrc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mark |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOriginalSrc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOriginalSrc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOriginalSrc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOriginalSrc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
					return 0, ErrIntOverflowOriginalSrc
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
				return 0, ErrInvalidLengthOriginalSrc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOriginalSrc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOriginalSrc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOriginalSrc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOriginalSrc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOriginalSrc = fmt.Errorf("proto: unexpected end of group")
)