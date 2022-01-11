// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scanner/api/v1/image.proto

package scannerV1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Image struct {
	Namespace            string     `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Features             []*Feature `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_505f23f17db497c4, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Image.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return m.Size()
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Image) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *Image) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Image) Clone() *Image {
	if m == nil {
		return nil
	}
	cloned := new(Image)
	*cloned = *m

	if m.Features != nil {
		cloned.Features = make([]*Feature, len(m.Features))
		for idx, v := range m.Features {
			cloned.Features[idx] = v.Clone()
		}
	}
	return cloned
}

func init() {
	proto.RegisterType((*Image)(nil), "scannerV1.Image")
}

func init() { proto.RegisterFile("scanner/api/v1/image.proto", fileDescriptor_505f23f17db497c4) }

var fileDescriptor_505f23f17db497c4 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x4e, 0xcc,
	0xcb, 0x4b, 0x2d, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd4, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xca, 0x85, 0x19, 0x4a, 0x29, 0xa1, 0x29,
	0x2b, 0x2b, 0xcd, 0xc9, 0x4b, 0x2d, 0x4a, 0x4c, 0xca, 0xcc, 0xc9, 0x2c, 0xa9, 0x84, 0x28, 0x57,
	0x0a, 0xe5, 0x62, 0xf5, 0x04, 0xe9, 0x16, 0x92, 0xe1, 0xe2, 0xcc, 0x4b, 0xcc, 0x4d, 0x2d, 0x2e,
	0x48, 0x4c, 0x4e, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0xe9, 0x71, 0x71,
	0xa4, 0xa5, 0x26, 0x96, 0x94, 0x16, 0xa5, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09,
	0xe9, 0xc1, 0x2d, 0xd2, 0x73, 0x83, 0x48, 0x05, 0xc1, 0xd5, 0x38, 0xd9, 0x9e, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x70, 0x29, 0x64,
	0xe6, 0xeb, 0x15, 0x97, 0x24, 0x26, 0x67, 0x17, 0xe5, 0x57, 0x40, 0xec, 0xd6, 0x4b, 0x2c, 0xc8,
	0x84, 0x19, 0xa2, 0x57, 0x66, 0x18, 0x85, 0x70, 0x79, 0x12, 0x1b, 0x58, 0x81, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0xd3, 0xb5, 0xfc, 0xe9, 0x00, 0x00, 0x00,
}

func (m *Image) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Image) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Image) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintImage(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Features) > 0 {
		for iNdEx := len(m.Features) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Features[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintImage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintImage(dAtA []byte, offset int, v uint64) int {
	offset -= sovImage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Image) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, e := range m.Features {
			l = e.Size()
			n += 1 + l + sovImage(uint64(l))
		}
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovImage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovImage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozImage(x uint64) (n int) {
	return sovImage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Image) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImage
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
			return fmt.Errorf("proto: Image: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Image: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Features = append(m.Features, &Feature{})
			if err := m.Features[len(m.Features)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImage
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
				return ErrInvalidLengthImage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImage
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
func skipImage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
					return 0, ErrIntOverflowImage
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
				return 0, ErrInvalidLengthImage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupImage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthImage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthImage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupImage = fmt.Errorf("proto: unexpected end of group")
)