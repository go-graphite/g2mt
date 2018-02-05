// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carbon.proto

/*
	Package carbon is a generated protocol buffer package.

	It is generated from these files:
		carbon.proto

	It has these top-level messages:
		Point
		Metric
		Payload
*/
package carbon

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	Timestamp uint32  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorCarbon, []int{0} }

func (m *Point) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Metric struct {
	Metric string            `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Points []Point           `protobuf:"bytes,2,rep,name=points" json:"points"`
	Tags   map[string]string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptorCarbon, []int{1} }

func (m *Metric) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *Metric) GetPoints() []Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Metric) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Payload struct {
	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorCarbon, []int{2} }

func (m *Payload) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "carbon.Point")
	proto.RegisterType((*Metric)(nil), "carbon.Metric")
	proto.RegisterType((*Payload)(nil), "carbon.Payload")
}
func (this *Point) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Point)
	if !ok {
		that2, ok := that.(Point)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *Metric) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Metric)
	if !ok {
		that2, ok := that.(Metric)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Metric != that1.Metric {
		return false
	}
	if len(this.Points) != len(that1.Points) {
		return false
	}
	for i := range this.Points {
		if !this.Points[i].Equal(&that1.Points[i]) {
			return false
		}
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	return true
}
func (this *Payload) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Payload)
	if !ok {
		that2, ok := that.(Payload)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Metrics) != len(that1.Metrics) {
		return false
	}
	for i := range this.Metrics {
		if !this.Metrics[i].Equal(that1.Metrics[i]) {
			return false
		}
	}
	return true
}
func (this *Point) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&carbon.Point{")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Metric) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&carbon.Metric{")
	s = append(s, "Metric: "+fmt.Sprintf("%#v", this.Metric)+",\n")
	if this.Points != nil {
		vs := make([]*Point, len(this.Points))
		for i := range vs {
			vs[i] = &this.Points[i]
		}
		s = append(s, "Points: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	keysForTags := make([]string, 0, len(this.Tags))
	for k, _ := range this.Tags {
		keysForTags = append(keysForTags, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForTags)
	mapStringForTags := "map[string]string{"
	for _, k := range keysForTags {
		mapStringForTags += fmt.Sprintf("%#v: %#v,", k, this.Tags[k])
	}
	mapStringForTags += "}"
	if this.Tags != nil {
		s = append(s, "Tags: "+mapStringForTags+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Payload) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&carbon.Payload{")
	if this.Metrics != nil {
		s = append(s, "metrics: "+fmt.Sprintf("%#v", this.Metrics)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCarbon(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCarbon(dAtA, i, uint64(m.Timestamp))
	}
	if m.Value != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	return i, nil
}

func (m *Metric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metric) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metric) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCarbon(dAtA, i, uint64(len(m.Metric)))
		i += copy(dAtA[i:], m.Metric)
	}
	if len(m.Points) > 0 {
		for _, msg := range m.Points {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCarbon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tags) > 0 {
		for k, _ := range m.Tags {
			dAtA[i] = 0x1a
			i++
			v := m.Tags[k]
			mapSize := 1 + len(k) + sovCarbon(uint64(len(k))) + 1 + len(v) + sovCarbon(uint64(len(v)))
			i = encodeVarintCarbon(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintCarbon(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintCarbon(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, msg := range m.Metrics {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCarbon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintCarbon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPoint(r randyCarbon, easy bool) *Point {
	this := &Point{}
	this.Timestamp = uint32(r.Uint32())
	this.Value = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Value *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMetric(r randyCarbon, easy bool) *Metric {
	this := &Metric{}
	this.Metric = string(randStringCarbon(r))
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Points = make([]Point, v1)
		for i := 0; i < v1; i++ {
			v2 := NewPopulatedPoint(r, easy)
			this.Points[i] = *v2
		}
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Tags = make(map[string]string)
		for i := 0; i < v3; i++ {
			this.Tags[randStringCarbon(r)] = randStringCarbon(r)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPayload(r randyCarbon, easy bool) *Payload {
	this := &Payload{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(5)
		this.Metrics = make([]*Metric, v4)
		for i := 0; i < v4; i++ {
			this.Metrics[i] = NewPopulatedMetric(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCarbon interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCarbon(r randyCarbon) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCarbon(r randyCarbon) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneCarbon(r)
	}
	return string(tmps)
}
func randUnrecognizedCarbon(r randyCarbon, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCarbon(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCarbon(dAtA []byte, r randyCarbon, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCarbon(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCarbon(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Point) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovCarbon(uint64(m.Timestamp))
	}
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *Metric) Size() (n int) {
	var l int
	_ = l
	l = len(m.Metric)
	if l > 0 {
		n += 1 + l + sovCarbon(uint64(l))
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovCarbon(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCarbon(uint64(len(k))) + 1 + len(v) + sovCarbon(uint64(len(v)))
			n += mapEntrySize + 1 + sovCarbon(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Payload) Size() (n int) {
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovCarbon(uint64(l))
		}
	}
	return n
}

func sovCarbon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCarbon(x uint64) (n int) {
	return sovCarbon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Point) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Point{`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Metric) String() string {
	if this == nil {
		return "nil"
	}
	keysForTags := make([]string, 0, len(this.Tags))
	for k, _ := range this.Tags {
		keysForTags = append(keysForTags, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForTags)
	mapStringForTags := "map[string]string{"
	for _, k := range keysForTags {
		mapStringForTags += fmt.Sprintf("%v: %v,", k, this.Tags[k])
	}
	mapStringForTags += "}"
	s := strings.Join([]string{`&Metric{`,
		`Metric:` + fmt.Sprintf("%v", this.Metric) + `,`,
		`Points:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Points), "Point", "Point", 1), `&`, ``, 1) + `,`,
		`Tags:` + mapStringForTags + `,`,
		`}`,
	}, "")
	return s
}
func (this *Payload) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Payload{`,
		`metrics:` + strings.Replace(fmt.Sprintf("%v", this.Metrics), "Metric", "Metric", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCarbon(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarbon
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
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipCarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarbon
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
func (m *Metric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarbon
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
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metric", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbon
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
				return ErrInvalidLengthCarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metric = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbon
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
				return ErrInvalidLengthCarbon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, Point{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbon
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
				return ErrInvalidLengthCarbon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCarbon
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCarbon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCarbon
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCarbon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCarbon
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCarbon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCarbon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarbon
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
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCarbon
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
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCarbon
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
				return ErrInvalidLengthCarbon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, &Metric{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCarbon
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
func skipCarbon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCarbon
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
					return 0, ErrIntOverflowCarbon
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
					return 0, ErrIntOverflowCarbon
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
				return 0, ErrInvalidLengthCarbon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCarbon
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
				next, err := skipCarbon(dAtA[start:])
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
	ErrInvalidLengthCarbon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCarbon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("carbon.proto", fileDescriptorCarbon) }

var fileDescriptorCarbon = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4a, 0xc3, 0x50,
	0x14, 0xc6, 0x73, 0xfa, 0x27, 0x25, 0x47, 0x2b, 0x72, 0x11, 0x09, 0x45, 0x8e, 0xa5, 0x53, 0x40,
	0x4d, 0xc1, 0x0e, 0x16, 0xdd, 0x0a, 0x8e, 0x42, 0x09, 0xbe, 0xc0, 0x4d, 0x8d, 0x31, 0xd8, 0xf4,
	0x96, 0xe4, 0x56, 0xc8, 0xe6, 0x23, 0xf8, 0x18, 0x3e, 0x80, 0x83, 0xa3, 0x63, 0xc7, 0x8e, 0x4e,
	0x62, 0xae, 0x8b, 0x63, 0x47, 0x47, 0xe9, 0x4d, 0xaa, 0x74, 0xfb, 0xbe, 0x73, 0xbe, 0x73, 0x7f,
	0x1f, 0x17, 0xb7, 0x47, 0x3c, 0xf1, 0xc5, 0xc4, 0x9d, 0x26, 0x42, 0x0a, 0x66, 0x16, 0xae, 0x75,
	0x12, 0x46, 0xf2, 0x6e, 0xe6, 0xbb, 0x23, 0x11, 0x77, 0x43, 0x11, 0x8a, 0xae, 0x5e, 0xfb, 0xb3,
	0x5b, 0xed, 0xb4, 0xd1, 0xaa, 0x38, 0xeb, 0x5c, 0x60, 0x7d, 0x28, 0xa2, 0x89, 0x64, 0x07, 0x68,
	0xc9, 0x28, 0x0e, 0x52, 0xc9, 0xe3, 0xa9, 0x0d, 0x6d, 0x70, 0x9a, 0xde, 0xff, 0x80, 0xed, 0x61,
	0xfd, 0x81, 0x8f, 0x67, 0x81, 0x5d, 0x69, 0x83, 0x03, 0x5e, 0x61, 0x3a, 0x2f, 0x80, 0xe6, 0x55,
	0x20, 0x93, 0x68, 0xc4, 0xf6, 0xd1, 0x8c, 0xb5, 0xd2, 0xb7, 0x96, 0x57, 0x3a, 0x76, 0x84, 0xe6,
	0x74, 0xf5, 0x7e, 0x6a, 0x57, 0xda, 0x55, 0x67, 0xeb, 0xb4, 0xe9, 0x96, 0xad, 0x35, 0x75, 0x50,
	0x9b, 0x7f, 0x1c, 0x1a, 0x5e, 0x19, 0x61, 0xc7, 0x58, 0x93, 0x3c, 0x4c, 0xed, 0xaa, 0x8e, 0xda,
	0xeb, 0x68, 0x81, 0x70, 0xaf, 0x79, 0x98, 0x5e, 0x4e, 0x64, 0x92, 0x79, 0x3a, 0xd5, 0x3a, 0x43,
	0xeb, 0x6f, 0xc4, 0x76, 0xb1, 0x7a, 0x1f, 0x64, 0x25, 0x7c, 0x25, 0x37, 0x2b, 0x5b, 0x65, 0xe5,
	0xf3, 0x4a, 0x1f, 0x3a, 0x3d, 0x6c, 0x0c, 0x79, 0x36, 0x16, 0xfc, 0x86, 0x39, 0xd8, 0x28, 0x8a,
	0xa6, 0x36, 0x68, 0xe8, 0xce, 0x26, 0xd4, 0x5b, 0xaf, 0x07, 0xfd, 0x45, 0x4e, 0xc6, 0x7b, 0x4e,
	0xc6, 0x32, 0x27, 0xf8, 0xc9, 0x09, 0x1e, 0x15, 0xc1, 0xb3, 0x22, 0x78, 0x55, 0x04, 0x6f, 0x8a,
	0x60, 0xae, 0x08, 0x16, 0x8a, 0xe0, 0x53, 0x11, 0x7c, 0x2b, 0x32, 0x96, 0x8a, 0xe0, 0xe9, 0x8b,
	0x0c, 0xdf, 0xd4, 0x3f, 0xdd, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x47, 0x8b, 0xd1, 0x9e, 0xb0,
	0x01, 0x00, 0x00,
}
