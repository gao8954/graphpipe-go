// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape" json:"default_graph_output_tensor_shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()         { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()    {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptorRemoteFusedGraphExecuteInfo, []int{0}
}

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptorRemoteFusedGraphExecuteInfo, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
}
func (m *RemoteFusedGraphExecuteInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteFusedGraphExecuteInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RemoteGraph != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.RemoteGraph.Size()))
		n1, err := m.RemoteGraph.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.GraphInputNodeName) > 0 {
		for _, s := range m.GraphInputNodeName {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.GraphOutputNodeName) > 0 {
		for _, s := range m.GraphOutputNodeName {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ExecutorName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(len(m.ExecutorName)))
		i += copy(dAtA[i:], m.ExecutorName)
	}
	if len(m.SerializedExecutorParameters) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(len(m.SerializedExecutorParameters)))
		i += copy(dAtA[i:], m.SerializedExecutorParameters)
	}
	if len(m.DefaultGraphInputTensorShape) > 0 {
		for _, msg := range m.DefaultGraphInputTensorShape {
			dAtA[i] = 0x32
			i++
			i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DefaultGraphOutputTensorShape) > 0 {
		for _, msg := range m.DefaultGraphOutputTensorShape {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dtype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.Dtype))
	}
	if m.Shape != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.Shape.Size()))
		n2, err := m.Shape.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintRemoteFusedGraphExecuteInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteFusedGraphExecuteInfo) Size() (n int) {
	var l int
	_ = l
	if m.RemoteGraph != nil {
		l = m.RemoteGraph.Size()
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	if len(m.GraphInputNodeName) > 0 {
		for _, s := range m.GraphInputNodeName {
			l = len(s)
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	if len(m.GraphOutputNodeName) > 0 {
		for _, s := range m.GraphOutputNodeName {
			l = len(s)
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	l = len(m.ExecutorName)
	if l > 0 {
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	l = len(m.SerializedExecutorParameters)
	if l > 0 {
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	if len(m.DefaultGraphInputTensorShape) > 0 {
		for _, e := range m.DefaultGraphInputTensorShape {
			l = e.Size()
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	if len(m.DefaultGraphOutputTensorShape) > 0 {
		for _, e := range m.DefaultGraphOutputTensorShape {
			l = e.Size()
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	return n
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Size() (n int) {
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovRemoteFusedGraphExecuteInfo(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	return n
}

func sovRemoteFusedGraphExecuteInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRemoteFusedGraphExecuteInfo(x uint64) (n int) {
	return sovRemoteFusedGraphExecuteInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteFusedGraphExecuteInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
			return fmt.Errorf("proto: RemoteFusedGraphExecuteInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteFusedGraphExecuteInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteGraph", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemoteGraph == nil {
				m.RemoteGraph = &GraphDef{}
			}
			if err := m.RemoteGraph.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GraphInputNodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GraphInputNodeName = append(m.GraphInputNodeName, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GraphOutputNodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GraphOutputNodeName = append(m.GraphOutputNodeName, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerializedExecutorParameters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerializedExecutorParameters = append(m.SerializedExecutorParameters[:0], dAtA[iNdEx:postIndex]...)
			if m.SerializedExecutorParameters == nil {
				m.SerializedExecutorParameters = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGraphInputTensorShape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultGraphInputTensorShape = append(m.DefaultGraphInputTensorShape, &RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{})
			if err := m.DefaultGraphInputTensorShape[len(m.DefaultGraphInputTensorShape)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGraphOutputTensorShape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultGraphOutputTensorShape = append(m.DefaultGraphOutputTensorShape, &RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{})
			if err := m.DefaultGraphOutputTensorShape[len(m.DefaultGraphOutputTensorShape)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteFusedGraphExecuteInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
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
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
			return fmt.Errorf("proto: TensorShapeTypeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorShapeTypeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= (DataType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteFusedGraphExecuteInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
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
func skipRemoteFusedGraphExecuteInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
					return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
					return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return 0, ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				next, err := skipRemoteFusedGraphExecuteInfo(dAtA[start:])
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
	ErrInvalidLengthRemoteFusedGraphExecuteInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemoteFusedGraphExecuteInfo   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptorRemoteFusedGraphExecuteInfo)
}

var fileDescriptorRemoteFusedGraphExecuteInfo = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x75, 0x0d, 0x29, 0xea, 0x25, 0x30, 0x1c, 0x05, 0x59, 0x21, 0x58, 0x06, 0x84, 0x64,
	0x21, 0xe4, 0x88, 0x74, 0x60, 0x01, 0x21, 0x55, 0x29, 0x55, 0x97, 0x12, 0x99, 0x4e, 0x2c, 0xa7,
	0x6b, 0xfc, 0x9c, 0x58, 0xc4, 0x7e, 0xd6, 0xf9, 0x4c, 0x29, 0x33, 0x62, 0xe6, 0xab, 0xf0, 0x2d,
	0x18, 0x19, 0x19, 0x51, 0x3e, 0x05, 0x23, 0xf2, 0x9d, 0x71, 0xce, 0xa8, 0xc9, 0xd4, 0xcd, 0xf6,
	0xfb, 0xfd, 0xdf, 0xfd, 0xdf, 0xdf, 0xef, 0xe8, 0x6b, 0x05, 0x59, 0x81, 0x32, 0x5e, 0xe2, 0xc5,
	0x68, 0x86, 0x12, 0x46, 0xb1, 0x14, 0x29, 0x5c, 0xa0, 0xfc, 0x30, 0x92, 0x90, 0xa2, 0x02, 0x1e,
	0x97, 0x05, 0x44, 0x7c, 0x2e, 0x45, 0xbe, 0xe0, 0xf0, 0x09, 0x66, 0xa5, 0x02, 0x9e, 0x64, 0x31,
	0x06, 0xb9, 0x44, 0x85, 0x8c, 0xae, 0x1b, 0x0c, 0x9e, 0x6c, 0x6e, 0xa6, 0xf5, 0x46, 0x32, 0x78,
	0xb6, 0x19, 0x33, 0x15, 0x5e, 0x2c, 0x44, 0x0e, 0x35, 0xbd, 0xa5, 0xa9, 0xba, 0xcc, 0xa1, 0x30,
	0xd8, 0xa3, 0xef, 0x5d, 0x7a, 0x3f, 0xd4, 0x8e, 0xdf, 0x54, 0x86, 0x8f, 0xab, 0xf3, 0x8e, 0x8c,
	0xdd, 0x93, 0x2c, 0x46, 0xf6, 0x82, 0xf6, 0xeb, 0x81, 0xb4, 0x15, 0x87, 0x78, 0xc4, 0xef, 0x8d,
	0xf7, 0x83, 0x75, 0xf7, 0x40, 0x6b, 0x26, 0x10, 0x87, 0x3d, 0x43, 0xea, 0x77, 0xf6, 0x9c, 0xde,
	0x35, 0xc3, 0x27, 0x59, 0x5e, 0x2a, 0x9e, 0x61, 0x04, 0x3c, 0x13, 0x29, 0x38, 0x3b, 0x5e, 0xc7,
	0xdf, 0x0b, 0x99, 0x2e, 0x9e, 0x54, 0xb5, 0x53, 0x8c, 0xe0, 0x54, 0xa4, 0xc0, 0x0e, 0xe8, 0x3d,
	0x23, 0xc1, 0x52, 0xb5, 0x35, 0x1d, 0xad, 0xb9, 0xa3, 0xab, 0x6f, 0x75, 0xb1, 0x11, 0x3d, 0xa6,
	0xb7, 0x4c, 0xbc, 0x28, 0x0d, 0x7b, 0xc3, 0x23, 0xfe, 0x5e, 0xd8, 0xff, 0xf7, 0x51, 0x43, 0x13,
	0xea, 0x16, 0x20, 0x13, 0xb1, 0x4c, 0x3e, 0x43, 0xc4, 0x1b, 0x3e, 0x17, 0x55, 0x26, 0x0a, 0x64,
	0xe1, 0x74, 0x3d, 0xe2, 0xf7, 0xc3, 0xe1, 0x9a, 0x3a, 0xaa, 0xa1, 0x69, 0xc3, 0xb0, 0x2f, 0x84,
	0x7a, 0x11, 0xc4, 0xa2, 0x5c, 0x2a, 0x6e, 0xcf, 0x66, 0xa7, 0xef, 0xec, 0x7a, 0x1d, 0xbf, 0x37,
	0x7e, 0x69, 0x07, 0xb4, 0x25, 0xdf, 0xe0, 0x4c, 0x63, 0xef, 0x2a, 0xe9, 0xd9, 0x65, 0x0e, 0xd3,
	0xea, 0xa7, 0x84, 0xc3, 0xfa, 0x94, 0xe3, 0x26, 0x23, 0x0b, 0x63, 0x5f, 0x09, 0x7d, 0xd8, 0xb6,
	0x51, 0xe7, 0xd5, 0xf2, 0x71, 0xf3, 0x1a, 0x7c, 0x3c, 0xb0, 0x7d, 0x98, 0xdc, 0x2d, 0x6e, 0xf0,
	0x91, 0xee, 0x5f, 0x25, 0x63, 0x4f, 0x69, 0x37, 0xaa, 0x76, 0x4c, 0x2f, 0xcb, 0xed, 0xf6, 0xb2,
	0x4c, 0x84, 0x12, 0x15, 0x19, 0x1a, 0x84, 0x8d, 0x69, 0xd7, 0xf8, 0xdd, 0xd1, 0x8b, 0x35, 0xb4,
	0x59, 0xab, 0xb9, 0xf1, 0x63, 0xd0, 0xc3, 0x6f, 0xe4, 0xc7, 0xca, 0x25, 0x3f, 0x57, 0x2e, 0xf9,
	0xb5, 0x72, 0xc9, 0xef, 0x95, 0x4b, 0xa8, 0x83, 0x72, 0x6e, 0x4b, 0x9b, 0x65, 0x3f, 0xf4, 0xb6,
	0x4c, 0xad, 0xbb, 0x4e, 0xc9, 0xfb, 0x57, 0xf3, 0x44, 0x2d, 0xca, 0xf3, 0x60, 0x86, 0xe9, 0xc8,
	0xba, 0x36, 0x57, 0x3f, 0xce, 0xf1, 0xbf, 0xfb, 0xf4, 0x87, 0x90, 0xf3, 0x5d, 0x7d, 0x9b, 0x0e,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x21, 0x2f, 0x9d, 0x7a, 0x18, 0x04, 0x00, 0x00,
}