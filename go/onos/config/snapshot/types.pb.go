// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/config/snapshot/types.proto

package snapshot

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Phase is a snapshot phase
type Phase int32

const (
	// MARK is the first phase in which changes are marked for deletion
	Phase_MARK Phase = 0
	// DELETE is the second phase in which changes are deleted from stores
	Phase_DELETE Phase = 1
)

var Phase_name = map[int32]string{
	0: "MARK",
	1: "DELETE",
}

var Phase_value = map[string]int32{
	"MARK":   0,
	"DELETE": 1,
}

func (x Phase) String() string {
	return proto.EnumName(Phase_name, int32(x))
}

func (Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9914b040a386c191, []int{0}
}

// State is the state of a snapshot within a phase
type State int32

const (
	// PENDING indicates the snapshot is pending
	State_PENDING State = 0
	// RUNNING indicates the snapshot is in progress
	State_RUNNING State = 1
	// COMPLETE indicates the snapshot is complete
	State_COMPLETE State = 2
)

var State_name = map[int32]string{
	0: "PENDING",
	1: "RUNNING",
	2: "COMPLETE",
}

var State_value = map[string]int32{
	"PENDING":  0,
	"RUNNING":  1,
	"COMPLETE": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9914b040a386c191, []int{1}
}

// Status is the status of a snapshot
type Status struct {
	// 'phase' is the snapshot phase
	Phase Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=onos.config.snapshot.Phase" json:"phase,omitempty"`
	// 'state' is the state of a snapshot
	State State `protobuf:"varint,2,opt,name=state,proto3,enum=onos.config.snapshot.State" json:"state,omitempty"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_9914b040a386c191, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_MARK
}

func (m *Status) GetState() State {
	if m != nil {
		return m.State
	}
	return State_PENDING
}

// RetentionOptions specifies the retention policy for a change log
type RetentionOptions struct {
	// 'retain_window' is the duration for which to retain network changes
	RetainWindow *time.Duration `protobuf:"bytes,1,opt,name=retain_window,json=retainWindow,proto3,stdduration" json:"retain_window,omitempty"`
}

func (m *RetentionOptions) Reset()         { *m = RetentionOptions{} }
func (m *RetentionOptions) String() string { return proto.CompactTextString(m) }
func (*RetentionOptions) ProtoMessage()    {}
func (*RetentionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_9914b040a386c191, []int{1}
}
func (m *RetentionOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetentionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetentionOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetentionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetentionOptions.Merge(m, src)
}
func (m *RetentionOptions) XXX_Size() int {
	return m.Size()
}
func (m *RetentionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RetentionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RetentionOptions proto.InternalMessageInfo

func (m *RetentionOptions) GetRetainWindow() *time.Duration {
	if m != nil {
		return m.RetainWindow
	}
	return nil
}

func init() {
	proto.RegisterEnum("onos.config.snapshot.Phase", Phase_name, Phase_value)
	proto.RegisterEnum("onos.config.snapshot.State", State_name, State_value)
	proto.RegisterType((*Status)(nil), "onos.config.snapshot.Status")
	proto.RegisterType((*RetentionOptions)(nil), "onos.config.snapshot.RetentionOptions")
}

func init() { proto.RegisterFile("onos/config/snapshot/types.proto", fileDescriptor_9914b040a386c191) }

var fileDescriptor_9914b040a386c191 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x86, 0xa7, 0x04, 0xb8, 0xa4, 0x70, 0xcd, 0xa4, 0x61, 0x81, 0x18, 0x2b, 0x61, 0x65, 0x58,
	0xb4, 0x11, 0x9f, 0x40, 0x9c, 0x89, 0x31, 0xca, 0x40, 0x46, 0x8d, 0xee, 0x4c, 0x91, 0x32, 0x4c,
	0x62, 0x7a, 0x26, 0xb4, 0x84, 0xf8, 0x16, 0x2e, 0x7d, 0x24, 0x97, 0x2c, 0xdd, 0x69, 0xe0, 0x45,
	0x4c, 0x5b, 0xd9, 0x19, 0x37, 0xcd, 0x39, 0xcd, 0xf7, 0xb5, 0xff, 0x8f, 0x3b, 0xa0, 0x40, 0xf3,
	0x27, 0x50, 0xb3, 0x3c, 0xe3, 0x5a, 0x89, 0x42, 0xcf, 0xc1, 0x70, 0xf3, 0x52, 0x48, 0xcd, 0x8a,
	0x05, 0x18, 0x20, 0x4d, 0x4b, 0x30, 0x4f, 0xb0, 0x1d, 0xd1, 0xa6, 0x19, 0x40, 0xf6, 0x2c, 0xb9,
	0x63, 0x26, 0xcb, 0x19, 0x9f, 0x2e, 0x17, 0xc2, 0xe4, 0xa0, 0xbc, 0xd5, 0x6e, 0x66, 0x90, 0x81,
	0x1b, 0xb9, 0x9d, 0xfc, 0x6d, 0x57, 0xe1, 0xea, 0x8d, 0x11, 0x66, 0xa9, 0xc9, 0x09, 0xae, 0x14,
	0x73, 0xa1, 0x65, 0x0b, 0x75, 0xd0, 0xf1, 0x5e, 0xff, 0x80, 0xfd, 0xf6, 0x0b, 0x1b, 0x5b, 0x24,
	0xf5, 0xa4, 0x55, 0xb4, 0x11, 0x46, 0xb6, 0x4a, 0x7f, 0x29, 0xf6, 0x7d, 0x99, 0x7a, 0xb2, 0xfb,
	0x80, 0xc3, 0x54, 0x1a, 0xa9, 0x6c, 0xb0, 0x51, 0x61, 0x4f, 0x4d, 0x22, 0xfc, 0x7f, 0x21, 0x8d,
	0xc8, 0xd5, 0xe3, 0x2a, 0x57, 0x53, 0x58, 0xb9, 0x04, 0xf5, 0xfe, 0x3e, 0xf3, 0x8d, 0xd8, 0xae,
	0x11, 0x8b, 0x7e, 0x1a, 0x0d, 0xca, 0x6f, 0x9f, 0x47, 0x28, 0x6d, 0x78, 0xeb, 0xde, 0x49, 0xbd,
	0x43, 0x5c, 0x71, 0xe1, 0x48, 0x0d, 0x97, 0x87, 0x67, 0xe9, 0x55, 0x18, 0x10, 0x8c, 0xab, 0x51,
	0x7c, 0x1d, 0xdf, 0xc6, 0x21, 0xea, 0x71, 0x5c, 0x71, 0x41, 0x48, 0x1d, 0xff, 0x1b, 0xc7, 0x49,
	0x74, 0x99, 0x5c, 0x84, 0x81, 0x5d, 0xd2, 0xbb, 0x24, 0xb1, 0x0b, 0x22, 0x0d, 0x5c, 0x3b, 0x1f,
	0x0d, 0xc7, 0x4e, 0x28, 0x0d, 0x5a, 0xef, 0x1b, 0x8a, 0xd6, 0x1b, 0x8a, 0xbe, 0x36, 0x14, 0xbd,
	0x6e, 0x69, 0xb0, 0xde, 0xd2, 0xe0, 0x63, 0x4b, 0x83, 0x49, 0xd5, 0x05, 0x3a, 0xfd, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x38, 0x33, 0x50, 0xaa, 0x01, 0x00, 0x00,
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.Phase != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RetentionOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetentionOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetentionOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RetainWindow != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.RetainWindow, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RetainWindow):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTypes(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Phase != 0 {
		n += 1 + sovTypes(uint64(m.Phase))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	return n
}

func (m *RetentionOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetainWindow != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RetainWindow)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= Phase(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RetentionOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RetentionOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetentionOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetainWindow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RetainWindow == nil {
				m.RetainWindow = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.RetainWindow, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)