// Code generated by protoc-gen-gogo.
// source: srcstore.proto
// DO NOT EDIT!

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		srcstore.proto

	It has these top-level messages:
		ImportOp
		IndexOp
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import unit "sourcegraph.com/sourcegraph/srclib/unit"
import graph3 "sourcegraph.com/sourcegraph/srclib/graph"
import pbtypes "sourcegraph.com/sqs/pbtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImportOp struct {
	Repo     string               `protobuf:"bytes,1,opt,name=Repo,proto3" json:"Repo,omitempty"`
	CommitID string               `protobuf:"bytes,2,opt,name=CommitID,proto3" json:"CommitID,omitempty"`
	Unit     *unit.RepoSourceUnit `protobuf:"bytes,3,opt,name=Unit" json:"Unit,omitempty"`
	Data     *graph3.Output       `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *ImportOp) Reset()         { *m = ImportOp{} }
func (m *ImportOp) String() string { return proto.CompactTextString(m) }
func (*ImportOp) ProtoMessage()    {}

type IndexOp struct {
	Repo     string `protobuf:"bytes,1,opt,name=Repo,proto3" json:"Repo,omitempty"`
	CommitID string `protobuf:"bytes,2,opt,name=CommitID,proto3" json:"CommitID,omitempty"`
}

func (m *IndexOp) Reset()         { *m = IndexOp{} }
func (m *IndexOp) String() string { return proto.CompactTextString(m) }
func (*IndexOp) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for MultiRepoImporter service

type MultiRepoImporterClient interface {
	// Import imports srclib build data for a source unit at a
	// specific version into the store.
	Import(ctx context.Context, in *ImportOp, opts ...grpc.CallOption) (*pbtypes.Void, error)
	// Index builds indexes for a specific repo at a specific version.
	Index(ctx context.Context, in *IndexOp, opts ...grpc.CallOption) (*pbtypes.Void, error)
}

type multiRepoImporterClient struct {
	cc *grpc.ClientConn
}

func NewMultiRepoImporterClient(cc *grpc.ClientConn) MultiRepoImporterClient {
	return &multiRepoImporterClient{cc}
}

func (c *multiRepoImporterClient) Import(ctx context.Context, in *ImportOp, opts ...grpc.CallOption) (*pbtypes.Void, error) {
	out := new(pbtypes.Void)
	err := grpc.Invoke(ctx, "/pb.MultiRepoImporter/Import", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiRepoImporterClient) Index(ctx context.Context, in *IndexOp, opts ...grpc.CallOption) (*pbtypes.Void, error) {
	out := new(pbtypes.Void)
	err := grpc.Invoke(ctx, "/pb.MultiRepoImporter/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MultiRepoImporter service

type MultiRepoImporterServer interface {
	// Import imports srclib build data for a source unit at a
	// specific version into the store.
	Import(context.Context, *ImportOp) (*pbtypes.Void, error)
	// Index builds indexes for a specific repo at a specific version.
	Index(context.Context, *IndexOp) (*pbtypes.Void, error)
}

func RegisterMultiRepoImporterServer(s *grpc.Server, srv MultiRepoImporterServer) {
	s.RegisterService(&_MultiRepoImporter_serviceDesc, srv)
}

func _MultiRepoImporter_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ImportOp)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MultiRepoImporterServer).Import(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _MultiRepoImporter_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IndexOp)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MultiRepoImporterServer).Index(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _MultiRepoImporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MultiRepoImporter",
	HandlerType: (*MultiRepoImporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Import",
			Handler:    _MultiRepoImporter_Import_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _MultiRepoImporter_Index_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *ImportOp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ImportOp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Repo) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintSrcstore(data, i, uint64(len(m.Repo)))
		i += copy(data[i:], m.Repo)
	}
	if len(m.CommitID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintSrcstore(data, i, uint64(len(m.CommitID)))
		i += copy(data[i:], m.CommitID)
	}
	if m.Unit != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintSrcstore(data, i, uint64(m.Unit.Size()))
		n1, err := m.Unit.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Data != nil {
		data[i] = 0x22
		i++
		i = encodeVarintSrcstore(data, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *IndexOp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IndexOp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Repo) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintSrcstore(data, i, uint64(len(m.Repo)))
		i += copy(data[i:], m.Repo)
	}
	if len(m.CommitID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintSrcstore(data, i, uint64(len(m.CommitID)))
		i += copy(data[i:], m.CommitID)
	}
	return i, nil
}

func encodeFixed64Srcstore(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Srcstore(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSrcstore(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ImportOp) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovSrcstore(uint64(l))
	}
	l = len(m.CommitID)
	if l > 0 {
		n += 1 + l + sovSrcstore(uint64(l))
	}
	if m.Unit != nil {
		l = m.Unit.Size()
		n += 1 + l + sovSrcstore(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovSrcstore(uint64(l))
	}
	return n
}

func (m *IndexOp) Size() (n int) {
	var l int
	_ = l
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovSrcstore(uint64(l))
	}
	l = len(m.CommitID)
	if l > 0 {
		n += 1 + l + sovSrcstore(uint64(l))
	}
	return n
}

func sovSrcstore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSrcstore(x uint64) (n int) {
	return sovSrcstore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ImportOp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrcstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImportOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImportOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Unit == nil {
				m.Unit = &unit.RepoSourceUnit{}
			}
			if err := m.Unit.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &graph3.Output{}
			}
			if err := m.Data.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrcstore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrcstore
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
func (m *IndexOp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrcstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSrcstore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSrcstore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrcstore
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
func skipSrcstore(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSrcstore
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSrcstore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSrcstore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSrcstore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipSrcstore(data[start:])
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
	ErrInvalidLengthSrcstore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSrcstore   = fmt.Errorf("proto: integer overflow")
)
