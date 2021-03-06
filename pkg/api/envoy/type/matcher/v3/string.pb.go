// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/v3/string.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
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

// Specifies the way to match a string.
// [#next-free-field: 7]
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	// If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no
	// effect for the safe_regex match.
	// For example, the matcher *data* will match both input string *Data* and *data* if set to true.
	IgnoreCase           bool     `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{0}
}
func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
}
type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}
type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof" json:"suffix,omitempty"`
}
type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof" json:"safe_regex,omitempty"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern()     {}
func (*StringMatcher_Prefix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_Suffix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (m *StringMatcher) GetIgnoreCase() bool {
	if m != nil {
		return m.IgnoreCase
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{1}
}
func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.v3.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.v3.ListStringMatcher")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3/string.proto", fileDescriptor_e33cffa01bf36e0e)
}

var fileDescriptor_e33cffa01bf36e0e = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0xef, 0x49, 0x9b, 0x9a, 0x3b, 0xe5, 0x42, 0x0d, 0xfe, 0x09, 0x05, 0x6b, 0xda, 0x5e,
	0xb4, 0x20, 0x24, 0x70, 0xbb, 0xbb, 0xcb, 0xb8, 0x29, 0x17, 0x95, 0x4b, 0x7c, 0x80, 0x32, 0xb6,
	0xa7, 0x75, 0x40, 0x67, 0xc2, 0xcc, 0x34, 0xb4, 0x3b, 0x97, 0x22, 0xae, 0x5c, 0xfa, 0x08, 0x3e,
	0x82, 0x7b, 0xe1, 0x2e, 0xf5, 0x0d, 0xa4, 0x4f, 0x21, 0x77, 0x25, 0x33, 0x13, 0x0b, 0xa1, 0x71,
	0x97, 0xc9, 0xf7, 0xfb, 0xce, 0xf9, 0xce, 0xe1, 0x90, 0x11, 0xf2, 0x52, 0xec, 0x52, 0xbd, 0x2b,
	0x30, 0x7d, 0x4f, 0xf5, 0xe2, 0x2d, 0xca, 0xb4, 0x9c, 0xa6, 0x4a, 0x4b, 0xc6, 0xd7, 0x49, 0x21,
	0x85, 0x16, 0xe1, 0x7d, 0xcb, 0x24, 0x86, 0x49, 0x2a, 0x26, 0x29, 0xa7, 0xfd, 0x61, 0xb3, 0x55,
	0xe2, 0x1a, 0xb7, 0xce, 0xd9, 0x1f, 0x3b, 0x84, 0x72, 0x2e, 0x34, 0xd5, 0x4c, 0x70, 0x95, 0x2e,
	0xb1, 0x90, 0xb8, 0xb0, 0x8f, 0x0a, 0x7a, 0xb4, 0x59, 0x16, 0xb4, 0xc6, 0x28, 0x4d, 0xf5, 0x46,
	0x55, 0xf2, 0xf0, 0x48, 0x2e, 0x51, 0x2a, 0x26, 0xf8, 0x21, 0x60, 0xff, 0x61, 0x49, 0xdf, 0xb1,
	0x25, 0xd5, 0x98, 0xfe, 0xfb, 0x70, 0xc2, 0xe8, 0x9b, 0x47, 0xce, 0x5e, 0xdb, 0x51, 0x5e, 0xba,
	0x80, 0xe1, 0x03, 0xe2, 0xe3, 0x96, 0x2e, 0x74, 0x04, 0x31, 0x4c, 0x4e, 0x67, 0x27, 0xb9, 0x7b,
	0x86, 0x43, 0xd2, 0x29, 0x24, 0xae, 0xd8, 0x36, 0xf2, 0x8c, 0x90, 0xdd, 0xb9, 0xcd, 0xda, 0xd2,
	0x8b, 0x61, 0x76, 0x92, 0x57, 0x82, 0x41, 0xd4, 0x66, 0x65, 0x90, 0xd6, 0x11, 0xe2, 0x84, 0xf0,
	0x15, 0x21, 0x8a, 0xae, 0x70, 0x6e, 0x77, 0x10, 0xf9, 0x31, 0x4c, 0xba, 0x17, 0xe3, 0xa4, 0x71,
	0x7d, 0x49, 0x6e, 0x98, 0x2a, 0x56, 0x16, 0xdc, 0x66, 0xfe, 0x27, 0xf0, 0x7a, 0xa6, 0xd8, 0xa9,
	0x29, 0x61, 0xd5, 0xf0, 0x31, 0xe9, 0xb2, 0x35, 0x17, 0x12, 0xe7, 0x0b, 0xaa, 0x30, 0xea, 0xc4,
	0x30, 0x09, 0x72, 0xe2, 0x7e, 0x3d, 0xa7, 0x0a, 0x2f, 0x9f, 0x7e, 0xfd, 0xf1, 0x71, 0x30, 0x22,
	0x71, 0x43, 0x8b, 0xda, 0xdc, 0xd9, 0x3d, 0x72, 0x66, 0x85, 0x79, 0x41, 0xb5, 0x46, 0xc9, 0xc3,
	0xd6, 0x9f, 0x0c, 0xae, 0xda, 0x41, 0xbb, 0xe7, 0xe7, 0xbe, 0x8d, 0x3b, 0xfa, 0x0c, 0xe4, 0xee,
	0x0b, 0xa6, 0x74, 0x7d, 0x61, 0x57, 0x24, 0xa8, 0x2c, 0x2a, 0x82, 0xb8, 0x35, 0xe9, 0x5e, 0x9c,
	0xff, 0x67, 0xa0, 0x7a, 0x43, 0x33, 0xd1, 0x17, 0xf0, 0x02, 0xc8, 0x0f, 0xfe, 0xcb, 0x67, 0x26,
	0xed, 0x13, 0x72, 0xde, 0xe0, 0x3f, 0x6a, 0x9c, 0xcd, 0x6e, 0xf6, 0x03, 0xf8, 0xb9, 0x1f, 0xc0,
	0xef, 0xfd, 0x00, 0xbe, 0x7f, 0xb8, 0xf9, 0xd5, 0xf1, 0x7a, 0x40, 0xc6, 0x4c, 0xb8, 0xf6, 0x85,
	0x14, 0xdb, 0x5d, 0x73, 0x92, 0xac, 0xeb, 0x2a, 0x5d, 0x9b, 0x1b, 0xb8, 0x86, 0x37, 0x1d, 0x7b,
	0x0c, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0x9b, 0xef, 0xdb, 0xec, 0x02, 0x00, 0x00,
}

func (m *StringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IgnoreCase {
		i--
		if m.IgnoreCase {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MatchPattern != nil {
		{
			size := m.MatchPattern.Size()
			i -= size
			if _, err := m.MatchPattern.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *StringMatcher_Exact) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Exact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Exact)
	copy(dAtA[i:], m.Exact)
	i = encodeVarintString(dAtA, i, uint64(len(m.Exact)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Prefix) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Prefix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Prefix)
	copy(dAtA[i:], m.Prefix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Prefix)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Suffix) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Suffix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Suffix)
	copy(dAtA[i:], m.Suffix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Suffix)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *StringMatcher_SafeRegex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_SafeRegex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SafeRegex != nil {
		{
			size, err := m.SafeRegex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintString(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *ListStringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListStringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListStringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Patterns) > 0 {
		for iNdEx := len(m.Patterns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Patterns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintString(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintString(dAtA []byte, offset int, v uint64) int {
	offset -= sovString(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.IgnoreCase {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StringMatcher_Exact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Exact)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Prefix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Suffix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Suffix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_SafeRegex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SafeRegex != nil {
		l = m.SafeRegex.Size()
		n += 1 + l + sovString(uint64(l))
	}
	return n
}
func (m *ListStringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Patterns) > 0 {
		for _, e := range m.Patterns {
			l = e.Size()
			n += 1 + l + sovString(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovString(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozString(x uint64) (n int) {
	return sovString(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: StringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Exact{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Prefix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Suffix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeRegex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegexMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &StringMatcher_SafeRegex{v}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreCase", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
			m.IgnoreCase = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func (m *ListStringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: ListStringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListStringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patterns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patterns = append(m.Patterns, &StringMatcher{})
			if err := m.Patterns[len(m.Patterns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func skipString(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
				return 0, ErrInvalidLengthString
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupString
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthString
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthString        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowString          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupString = fmt.Errorf("proto: unexpected end of group")
)
