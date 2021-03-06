// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/cluster/redis/redis_cluster.proto

package envoy_config_cluster_redis

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// [#next-free-field: 7]
type RedisClusterConfig struct {
	// Interval between successive topology refresh requests. If not set, this defaults to 5s.
	ClusterRefreshRate *types.Duration `protobuf:"bytes,1,opt,name=cluster_refresh_rate,json=clusterRefreshRate,proto3" json:"cluster_refresh_rate,omitempty"`
	// Timeout for topology refresh request. If not set, this defaults to 3s.
	ClusterRefreshTimeout *types.Duration `protobuf:"bytes,2,opt,name=cluster_refresh_timeout,json=clusterRefreshTimeout,proto3" json:"cluster_refresh_timeout,omitempty"`
	// The minimum interval that must pass after triggering a topology refresh request before a new
	// request can possibly be triggered again. Any errors received during one of these
	// time intervals are ignored. If not set, this defaults to 5s.
	RedirectRefreshInterval *types.Duration `protobuf:"bytes,3,opt,name=redirect_refresh_interval,json=redirectRefreshInterval,proto3" json:"redirect_refresh_interval,omitempty"`
	// The number of redirection errors that must be received before
	// triggering a topology refresh request. If not set, this defaults to 5.
	// If this is set to 0, topology refresh after redirect is disabled.
	RedirectRefreshThreshold *types.UInt32Value `protobuf:"bytes,4,opt,name=redirect_refresh_threshold,json=redirectRefreshThreshold,proto3" json:"redirect_refresh_threshold,omitempty"`
	// The number of failures that must be received before triggering a topology refresh request.
	// If not set, this defaults to 0, which disables the topology refresh due to failure.
	FailureRefreshThreshold uint32 `protobuf:"varint,5,opt,name=failure_refresh_threshold,json=failureRefreshThreshold,proto3" json:"failure_refresh_threshold,omitempty"`
	// The number of hosts became degraded or unhealthy before triggering a topology refresh request.
	// If not set, this defaults to 0, which disables the topology refresh due to degraded or
	// unhealthy host.
	HostDegradedRefreshThreshold uint32   `protobuf:"varint,6,opt,name=host_degraded_refresh_threshold,json=hostDegradedRefreshThreshold,proto3" json:"host_degraded_refresh_threshold,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *RedisClusterConfig) Reset()         { *m = RedisClusterConfig{} }
func (m *RedisClusterConfig) String() string { return proto.CompactTextString(m) }
func (*RedisClusterConfig) ProtoMessage()    {}
func (*RedisClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6593a6ec218c02, []int{0}
}
func (m *RedisClusterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedisClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedisClusterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedisClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisClusterConfig.Merge(m, src)
}
func (m *RedisClusterConfig) XXX_Size() int {
	return m.Size()
}
func (m *RedisClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RedisClusterConfig proto.InternalMessageInfo

func (m *RedisClusterConfig) GetClusterRefreshRate() *types.Duration {
	if m != nil {
		return m.ClusterRefreshRate
	}
	return nil
}

func (m *RedisClusterConfig) GetClusterRefreshTimeout() *types.Duration {
	if m != nil {
		return m.ClusterRefreshTimeout
	}
	return nil
}

func (m *RedisClusterConfig) GetRedirectRefreshInterval() *types.Duration {
	if m != nil {
		return m.RedirectRefreshInterval
	}
	return nil
}

func (m *RedisClusterConfig) GetRedirectRefreshThreshold() *types.UInt32Value {
	if m != nil {
		return m.RedirectRefreshThreshold
	}
	return nil
}

func (m *RedisClusterConfig) GetFailureRefreshThreshold() uint32 {
	if m != nil {
		return m.FailureRefreshThreshold
	}
	return 0
}

func (m *RedisClusterConfig) GetHostDegradedRefreshThreshold() uint32 {
	if m != nil {
		return m.HostDegradedRefreshThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*RedisClusterConfig)(nil), "envoy.config.cluster.redis.RedisClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/redis/redis_cluster.proto", fileDescriptor_6d6593a6ec218c02)
}

var fileDescriptor_6d6593a6ec218c02 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0xeb, 0x29, 0x1d, 0x21, 0x23, 0x24, 0xb0, 0x40, 0x93, 0x19, 0x95, 0xa1, 0x62, 0x35,
	0x62, 0x61, 0x4b, 0xed, 0x8e, 0x65, 0x5a, 0x16, 0xdd, 0x55, 0xa1, 0xc3, 0x02, 0x16, 0x91, 0x3b,
	0x7e, 0xc9, 0x58, 0x0a, 0x71, 0xe4, 0x3c, 0x0f, 0x74, 0xc7, 0x79, 0x38, 0x02, 0x27, 0xe8, 0x12,
	0x6e, 0x80, 0xe6, 0x18, 0x2c, 0x10, 0x8a, 0xed, 0x20, 0x98, 0x14, 0xc1, 0x26, 0xb1, 0xfc, 0xff,
	0xff, 0xa7, 0xdf, 0x4f, 0x8f, 0x72, 0xa8, 0x37, 0xe6, 0x5a, 0xac, 0x4c, 0x5d, 0xe8, 0x52, 0xac,
	0x2a, 0xd7, 0x22, 0x58, 0x61, 0x41, 0xe9, 0x36, 0x7c, 0xf3, 0x78, 0xc7, 0x1b, 0x6b, 0xd0, 0xb0,
	0x99, 0xf7, 0xf3, 0xe0, 0xe7, 0xbd, 0xe6, 0x9d, 0xb3, 0x79, 0x69, 0x4c, 0x59, 0x81, 0xf0, 0xce,
	0x2b, 0x57, 0x08, 0xe5, 0xac, 0x44, 0x6d, 0xea, 0x90, 0x1d, 0xea, 0xef, 0xad, 0x6c, 0x1a, 0xb0,
	0x6d, 0xd4, 0x9f, 0x38, 0xd5, 0x48, 0x21, 0xeb, 0xda, 0xa0, 0x8f, 0xb5, 0xa2, 0x45, 0x89, 0xae,
	0x97, 0x27, 0x1b, 0x59, 0x69, 0x25, 0x11, 0x44, 0x7f, 0x08, 0xc2, 0xb3, 0x1f, 0xfb, 0x94, 0x65,
	0x5d, 0x83, 0xd3, 0x50, 0xe7, 0xd4, 0x97, 0x63, 0x4b, 0xfa, 0x28, 0xf6, 0xcb, 0x2d, 0x14, 0x16,
	0xda, 0x75, 0x6e, 0x25, 0x42, 0x42, 0x8e, 0xc8, 0xe2, 0xde, 0xf1, 0x94, 0x87, 0x36, 0xbc, 0x6f,
	0xc3, 0xcf, 0x62, 0xdb, 0xf4, 0xee, 0xf7, 0xf4, 0xe0, 0x13, 0x19, 0x3d, 0xdf, 0xcb, 0x58, 0x04,
	0x64, 0x21, 0x9f, 0x49, 0x04, 0xf6, 0x96, 0x4e, 0x76, 0xb1, 0xa8, 0xdf, 0x81, 0x71, 0x98, 0x8c,
	0xfe, 0x9f, 0xfc, 0xf8, 0x4f, 0xf2, 0x65, 0x20, 0xb0, 0x25, 0x9d, 0x76, 0xb3, 0xb4, 0xb0, 0xc2,
	0x5f, 0x74, 0x5d, 0x23, 0xd8, 0x8d, 0xac, 0x92, 0xfd, 0x7f, 0xe0, 0xb3, 0x49, 0x9f, 0x8d, 0xd4,
	0xf3, 0x98, 0x64, 0x6f, 0xe8, 0x6c, 0x80, 0xc5, 0x75, 0xf7, 0x33, 0x95, 0x4a, 0xee, 0x78, 0xee,
	0xe1, 0x80, 0xbb, 0x3c, 0xaf, 0xf1, 0xe4, 0xf8, 0xb5, 0xac, 0x1c, 0x64, 0xc9, 0x0e, 0xfa, 0xb2,
	0x4f, 0xb3, 0x17, 0x74, 0x5a, 0x48, 0x5d, 0x39, 0x0b, 0xb7, 0xa0, 0x0f, 0x8e, 0xc8, 0xe2, 0x7e,
	0x36, 0x89, 0x86, 0x41, 0xf6, 0x25, 0x7d, 0xba, 0x36, 0x2d, 0xe6, 0x0a, 0x4a, 0x2b, 0x15, 0xa8,
	0x5b, 0x08, 0x63, 0x4f, 0x38, 0xec, 0x6c, 0x67, 0xd1, 0xb5, 0x8b, 0x49, 0x5f, 0xdd, 0x6c, 0xe7,
	0xe4, 0xcb, 0x76, 0x4e, 0xbe, 0x6d, 0xe7, 0xe4, 0xf3, 0xc7, 0x9b, 0xaf, 0xe3, 0xd1, 0x83, 0x3d,
	0xba, 0xd0, 0x26, 0x6c, 0x77, 0x63, 0xcd, 0x87, 0x6b, 0xfe, 0xf7, 0xc5, 0x4d, 0x1f, 0xfe, 0xbe,
	0x3d, 0x17, 0xdd, 0xf3, 0x2f, 0xc8, 0xd5, 0xd8, 0xcf, 0xe1, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0xd6, 0xd9, 0xa2, 0x22, 0x03, 0x00, 0x00,
}

func (m *RedisClusterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedisClusterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedisClusterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HostDegradedRefreshThreshold != 0 {
		i = encodeVarintRedisCluster(dAtA, i, uint64(m.HostDegradedRefreshThreshold))
		i--
		dAtA[i] = 0x30
	}
	if m.FailureRefreshThreshold != 0 {
		i = encodeVarintRedisCluster(dAtA, i, uint64(m.FailureRefreshThreshold))
		i--
		dAtA[i] = 0x28
	}
	if m.RedirectRefreshThreshold != nil {
		{
			size, err := m.RedirectRefreshThreshold.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRedisCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.RedirectRefreshInterval != nil {
		{
			size, err := m.RedirectRefreshInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRedisCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ClusterRefreshTimeout != nil {
		{
			size, err := m.ClusterRefreshTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRedisCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ClusterRefreshRate != nil {
		{
			size, err := m.ClusterRefreshRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRedisCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRedisCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovRedisCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RedisClusterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterRefreshRate != nil {
		l = m.ClusterRefreshRate.Size()
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.ClusterRefreshTimeout != nil {
		l = m.ClusterRefreshTimeout.Size()
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.RedirectRefreshInterval != nil {
		l = m.RedirectRefreshInterval.Size()
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.RedirectRefreshThreshold != nil {
		l = m.RedirectRefreshThreshold.Size()
		n += 1 + l + sovRedisCluster(uint64(l))
	}
	if m.FailureRefreshThreshold != 0 {
		n += 1 + sovRedisCluster(uint64(m.FailureRefreshThreshold))
	}
	if m.HostDegradedRefreshThreshold != 0 {
		n += 1 + sovRedisCluster(uint64(m.HostDegradedRefreshThreshold))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRedisCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRedisCluster(x uint64) (n int) {
	return sovRedisCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RedisClusterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedisCluster
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
			return fmt.Errorf("proto: RedisClusterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedisClusterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterRefreshRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
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
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterRefreshRate == nil {
				m.ClusterRefreshRate = &types.Duration{}
			}
			if err := m.ClusterRefreshRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterRefreshTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
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
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterRefreshTimeout == nil {
				m.ClusterRefreshTimeout = &types.Duration{}
			}
			if err := m.ClusterRefreshTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedirectRefreshInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
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
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RedirectRefreshInterval == nil {
				m.RedirectRefreshInterval = &types.Duration{}
			}
			if err := m.RedirectRefreshInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedirectRefreshThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
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
				return ErrInvalidLengthRedisCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RedirectRefreshThreshold == nil {
				m.RedirectRefreshThreshold = &types.UInt32Value{}
			}
			if err := m.RedirectRefreshThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureRefreshThreshold", wireType)
			}
			m.FailureRefreshThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailureRefreshThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDegradedRefreshThreshold", wireType)
			}
			m.HostDegradedRefreshThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HostDegradedRefreshThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRedisCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedisCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRedisCluster
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
func skipRedisCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRedisCluster
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
					return 0, ErrIntOverflowRedisCluster
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
					return 0, ErrIntOverflowRedisCluster
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
				return 0, ErrInvalidLengthRedisCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRedisCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRedisCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRedisCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRedisCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRedisCluster = fmt.Errorf("proto: unexpected end of group")
)
