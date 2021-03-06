// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_network_thrift_proxy_filters_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#next-free-field: 6]
type RateLimit struct {
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configuration stage. Each configured rate limit filter performs a
	// rate limit check using descriptors configured in the
	// :ref:`envoy_api_msg_extensions.filters.network.thrift_proxy.v3.RouteAction` for the request.
	// Only those entries with a matching stage number are used for a given filter. If not set, the
	// default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *types.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v3.RateLimitServiceConfig `protobuf:"bytes,5,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_920e335ffb05829c, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetTimeout() *types.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.extensions.filters.network.thrift_proxy.filters.ratelimit.v3.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3/rate_limit.proto", fileDescriptor_920e335ffb05829c)
}

var fileDescriptor_920e335ffb05829c = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0xd5, 0x9a, 0xdc, 0x25, 0x59, 0x04, 0x1c, 0x6e, 0x30, 0x91, 0x62, 0x0c, 0x34, 0xa7, 0x14,
	0xbb, 0x4a, 0xae, 0x41, 0x69, 0x90, 0x4c, 0x4a, 0x90, 0x4e, 0xce, 0x07, 0x58, 0x9b, 0x78, 0xec,
	0xac, 0xf0, 0xed, 0x9e, 0xd6, 0x63, 0x73, 0xd7, 0x51, 0x22, 0x3e, 0x81, 0x4f, 0xa1, 0x47, 0x4a,
	0x09, 0x1d, 0x25, 0xba, 0xcf, 0x48, 0x85, 0x6e, 0x77, 0xcf, 0x97, 0x84, 0x92, 0x6e, 0x77, 0xde,
	0xcc, 0x7b, 0x33, 0xf3, 0x86, 0x9e, 0x83, 0xea, 0xf4, 0x92, 0xc3, 0x02, 0x41, 0x35, 0x52, 0xab,
	0x86, 0x97, 0xb2, 0x46, 0x30, 0x0d, 0x57, 0x80, 0x9f, 0xb4, 0xf9, 0xc8, 0xf1, 0xca, 0xc8, 0x12,
	0xf3, 0xb9, 0xd1, 0x8b, 0x65, 0x0f, 0x1a, 0x81, 0x50, 0xcb, 0x99, 0x44, 0xde, 0x4d, 0xec, 0x27,
	0xb7, 0x3f, 0x36, 0x37, 0x1a, 0x75, 0x98, 0x5a, 0x52, 0xb6, 0x25, 0x65, 0xbe, 0x8e, 0x79, 0x52,
	0x76, 0x9b, 0xb4, 0x07, 0x7b, 0x52, 0xd6, 0x4d, 0x0e, 0x5e, 0xbb, 0xc6, 0x2e, 0xb5, 0x2a, 0x65,
	0x75, 0x4f, 0xaf, 0x6e, 0x9c, 0xd0, 0x41, 0x5c, 0x69, 0x5d, 0xd5, 0xc0, 0xed, 0xef, 0xa2, 0x2d,
	0x79, 0xd1, 0x1a, 0x81, 0x52, 0x2b, 0x8f, 0x1f, 0xb6, 0xc5, 0x5c, 0x70, 0xa1, 0x94, 0x46, 0x1b,
	0x6e, 0x78, 0x83, 0x02, 0xdb, 0x4d, 0xf9, 0xcb, 0x7f, 0xe0, 0x0e, 0xcc, 0xba, 0x61, 0xa9, 0x2a,
	0x9f, 0xf2, 0xac, 0x13, 0xb5, 0x2c, 0x04, 0x02, 0xdf, 0x3c, 0x1c, 0xf0, 0xea, 0x77, 0x40, 0xf7,
	0x33, 0x81, 0xf0, 0x7e, 0xdd, 0x55, 0xf8, 0x82, 0x0e, 0x0b, 0x3d, 0x13, 0x52, 0x45, 0x24, 0x21,
	0xe3, 0xfd, 0x74, 0xf7, 0x26, 0xdd, 0x31, 0x41, 0x42, 0x32, 0x1f, 0x0e, 0x0f, 0xe9, 0xa0, 0x41,
	0x51, 0x41, 0x14, 0x24, 0x64, 0xfc, 0xc8, 0xe2, 0x47, 0x41, 0x44, 0x33, 0x17, 0x0d, 0x27, 0x74,
	0x17, 0xe5, 0x0c, 0x74, 0x8b, 0xd1, 0x83, 0x84, 0x8c, 0x1f, 0x9e, 0x3c, 0x67, 0x6e, 0x34, 0xb6,
	0x19, 0x8d, 0x9d, 0xf9, 0xd1, 0xb2, 0x4d, 0x66, 0x78, 0x44, 0x9f, 0x96, 0x42, 0xd6, 0xad, 0x81,
	0x7c, 0xa6, 0x0b, 0xc8, 0x0b, 0x50, 0xcb, 0x68, 0x27, 0x21, 0xe3, 0xbd, 0xec, 0x89, 0x07, 0x3e,
	0xe8, 0x02, 0xce, 0x40, 0x2d, 0x43, 0x49, 0xc3, 0xad, 0x4d, 0x79, 0x03, 0xa6, 0x93, 0x97, 0x10,
	0x0d, 0xac, 0xd6, 0x31, 0x73, 0x7e, 0xb9, 0x5d, 0xdf, 0xb1, 0x81, 0xf5, 0x23, 0x9e, 0xbb, 0x92,
	0x77, 0x36, 0x27, 0xdd, 0xbb, 0x49, 0x07, 0x5f, 0x49, 0x30, 0x22, 0xd9, 0xc8, 0xdc, 0xcb, 0x38,
	0x7d, 0xfb, 0xed, 0xc7, 0x97, 0xf8, 0x94, 0xbe, 0xb9, 0x43, 0xea, 0x3c, 0xf6, 0xbe, 0xb3, 0x5b,
	0x17, 0xd3, 0x9d, 0x88, 0x7a, 0x7e, 0x25, 0x8e, 0xb7, 0x4a, 0xe9, 0xe2, 0x7a, 0x15, 0x93, 0x9f,
	0xab, 0x98, 0xfc, 0x59, 0xc5, 0xe4, 0xfb, 0xe7, 0xeb, 0x5f, 0xc3, 0x60, 0x44, 0xe8, 0x54, 0x6a,
	0xd7, 0xa7, 0x3b, 0x9a, 0xff, 0x3f, 0xb1, 0xf4, 0x71, 0x2f, 0x39, 0x5d, 0x2f, 0x79, 0x4a, 0x2e,
	0x86, 0x76, 0xdb, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe7, 0x09, 0x1e, 0x16, 0x03,
	0x00, 0x00,
}

func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RateLimitService != nil {
		{
			size, err := m.RateLimitService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.FailureModeDeny {
		i--
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Timeout != nil {
		{
			size, err := m.Timeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Stage != 0 {
		i = encodeVarintRateLimit(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovRateLimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.Stage != 0 {
		n += 1 + sovRateLimit(uint64(m.Stage))
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitService != nil {
		l = m.RateLimitService.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
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
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &types.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeDeny", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
			m.FailureModeDeny = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RateLimitService == nil {
				m.RateLimitService = &v3.RateLimitServiceConfig{}
			}
			if err := m.RateLimitService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRateLimit
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
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
				return 0, ErrInvalidLengthRateLimit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRateLimit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRateLimit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRateLimit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRateLimit = fmt.Errorf("proto: unexpected end of group")
)
