// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: balance.proto

/*
	Package balance is a generated protocol buffer package.

	It is generated from these files:
		balance.proto

	It has these top-level messages:
		Balance
*/
package balance

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Balance struct {
	Type   Type   `protobuf:"varint,1,opt,name=Type,proto3,casttype=Type" json:"Type,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptorBalance, []int{0} }

func (m *Balance) GetType() Type {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Balance) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (*Balance) XXX_MessageName() string {
	return "balance.Balance"
}
func init() {
	proto.RegisterType((*Balance)(nil), "balance.Balance")
	golang_proto.RegisterType((*Balance)(nil), "balance.Balance")
}
func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBalance(dAtA, i, uint64(m.Type))
	}
	if m.Amount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBalance(dAtA, i, uint64(m.Amount))
	}
	return i, nil
}

func encodeVarintBalance(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Balance) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovBalance(uint64(m.Type))
	}
	if m.Amount != 0 {
		n += 1 + sovBalance(uint64(m.Amount))
	}
	return n
}

func sovBalance(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBalance(x uint64) (n int) {
	return sovBalance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Balance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBalance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBalance
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
func skipBalance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalance
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
					return 0, ErrIntOverflowBalance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBalance
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBalance
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBalance
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBalance(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBalance = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalance   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("balance.proto", fileDescriptorBalance) }
func init() { golang_proto.RegisterFile("balance.proto", fileDescriptorBalance) }

var fileDescriptorBalance = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0xcc, 0x49,
	0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0xf2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xf4, 0x29, 0xf9, 0x72, 0xb1, 0x3b,
	0x41, 0x74, 0x0a, 0xc9, 0x70, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x3a, 0x71, 0xfc, 0xba, 0x27, 0x0f, 0xe6, 0x07, 0x81, 0x49, 0x21, 0x31, 0x2e, 0x36, 0xc7, 0xdc,
	0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x28, 0xcf, 0x8a, 0x67, 0xc6,
	0x02, 0x79, 0x86, 0x09, 0x8b, 0xe4, 0x19, 0x66, 0x2c, 0x92, 0x67, 0x70, 0xb2, 0x3f, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x96, 0x63, 0x3c, 0xf1,
	0x58, 0x8e, 0x31, 0x4a, 0x13, 0xc9, 0x4d, 0x19, 0x95, 0x05, 0xa9, 0x45, 0x39, 0xa9, 0x29, 0xe9,
	0xa9, 0x45, 0xfa, 0x49, 0xa5, 0x45, 0x45, 0xf9, 0xe5, 0xfa, 0x89, 0xc9, 0xb9, 0xfa, 0x50, 0xe7,
	0x27, 0xb1, 0x81, 0x9d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xac, 0x60, 0x98, 0x58, 0xdf,
	0x00, 0x00, 0x00,
}
