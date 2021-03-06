/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kmodules.xyz/prober/api/v1/generated.proto

package v1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	k8s_io_api_core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
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

func (m *FormEntry) Reset()      { *m = FormEntry{} }
func (*FormEntry) ProtoMessage() {}
func (*FormEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c9649438138bbb, []int{0}
}
func (m *FormEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FormEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FormEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormEntry.Merge(m, src)
}
func (m *FormEntry) XXX_Size() int {
	return m.Size()
}
func (m *FormEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FormEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FormEntry proto.InternalMessageInfo

func (m *HTTPPostAction) Reset()      { *m = HTTPPostAction{} }
func (*HTTPPostAction) ProtoMessage() {}
func (*HTTPPostAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c9649438138bbb, []int{1}
}
func (m *HTTPPostAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPPostAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HTTPPostAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPPostAction.Merge(m, src)
}
func (m *HTTPPostAction) XXX_Size() int {
	return m.Size()
}
func (m *HTTPPostAction) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPPostAction.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPPostAction proto.InternalMessageInfo

func (m *Handler) Reset()      { *m = Handler{} }
func (*Handler) ProtoMessage() {}
func (*Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_90c9649438138bbb, []int{2}
}
func (m *Handler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handler.Merge(m, src)
}
func (m *Handler) XXX_Size() int {
	return m.Size()
}
func (m *Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Handler proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FormEntry)(nil), "kmodules.xyz.prober.api.v1.FormEntry")
	proto.RegisterType((*HTTPPostAction)(nil), "kmodules.xyz.prober.api.v1.HTTPPostAction")
	proto.RegisterType((*Handler)(nil), "kmodules.xyz.prober.api.v1.Handler")
}

func init() {
	proto.RegisterFile("kmodules.xyz/prober/api/v1/generated.proto", fileDescriptor_90c9649438138bbb)
}

var fileDescriptor_90c9649438138bbb = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x9b, 0x34, 0x69, 0xd6, 0x2d, 0x87, 0x45, 0x48, 0x56, 0x04, 0x76, 0x08, 0x42, 0x8a,
	0x2a, 0xb1, 0xa6, 0xe1, 0x82, 0x04, 0x07, 0x70, 0xd5, 0x36, 0x15, 0x52, 0x89, 0xb6, 0x29, 0x07,
	0x6e, 0x8e, 0xb3, 0xb5, 0xad, 0xc4, 0x5e, 0x6b, 0xbd, 0x89, 0x6a, 0x4e, 0x7c, 0x02, 0xdf, 0xc2,
	0x57, 0xf4, 0xd8, 0x63, 0x4f, 0x16, 0x35, 0xff, 0xc0, 0x81, 0x13, 0xda, 0xb5, 0x93, 0x38, 0xd0,
	0x70, 0xf3, 0xbe, 0x79, 0xf3, 0x26, 0x6f, 0xde, 0x28, 0x60, 0x7f, 0x12, 0xd0, 0xf1, 0x6c, 0x4a,
	0x62, 0x74, 0x95, 0x7c, 0x31, 0x23, 0x46, 0x47, 0x84, 0x99, 0x76, 0xe4, 0x9b, 0xf3, 0x03, 0xd3,
	0x25, 0x21, 0x61, 0x36, 0x27, 0x63, 0x14, 0x31, 0xca, 0x29, 0x6c, 0x95, 0xb9, 0x28, 0xe7, 0x22,
	0x3b, 0xf2, 0xd1, 0xfc, 0xa0, 0xf5, 0xc2, 0xf5, 0xb9, 0x37, 0x1b, 0x21, 0x87, 0x06, 0xa6, 0x4b,
	0x5d, 0x6a, 0xca, 0x96, 0xd1, 0xec, 0x52, 0xbe, 0xe4, 0x43, 0x7e, 0xe5, 0x52, 0xad, 0xce, 0xe4,
	0x75, 0x8c, 0x7c, 0x2a, 0x27, 0x39, 0x94, 0x91, 0x7b, 0xc6, 0xb5, 0x5e, 0xad, 0x38, 0x81, 0xed,
	0x78, 0x7e, 0x48, 0x58, 0x62, 0x46, 0x13, 0xd7, 0x9c, 0x71, 0x7f, 0x6a, 0xfa, 0x21, 0x8f, 0x39,
	0xfb, 0xbb, 0xa9, 0x73, 0x06, 0x9a, 0xc7, 0x94, 0x05, 0x47, 0x21, 0x67, 0x09, 0x7c, 0x02, 0xaa,
	0x13, 0x92, 0x68, 0x4a, 0x5b, 0xe9, 0x36, 0x2d, 0xf5, 0x3a, 0x35, 0x2a, 0x59, 0x6a, 0x54, 0x3f,
	0x90, 0x04, 0x0b, 0x1c, 0x76, 0x40, 0x7d, 0x6e, 0x4f, 0x67, 0x24, 0xd6, 0xb6, 0xda, 0xd5, 0x6e,
	0xd3, 0x02, 0x59, 0x6a, 0xd4, 0x3f, 0x49, 0x04, 0x17, 0x95, 0xce, 0xf7, 0x2a, 0x78, 0xd0, 0x1f,
	0x0e, 0x07, 0x03, 0x1a, 0xf3, 0xf7, 0x0e, 0xf7, 0x69, 0x08, 0xdb, 0xa0, 0x16, 0xd9, 0xdc, 0x2b,
	0x64, 0x77, 0x0b, 0xd9, 0xda, 0xc0, 0xe6, 0x1e, 0x96, 0x15, 0x88, 0x41, 0x2d, 0xa2, 0x8c, 0x6b,
	0x5b, 0x6d, 0xa5, 0xab, 0xf6, 0x5e, 0xa2, 0xdc, 0x08, 0x2a, 0x1b, 0x41, 0xd1, 0xc4, 0x45, 0xc2,
	0x08, 0xca, 0x8d, 0xa0, 0xd3, 0x90, 0x7f, 0x64, 0xe7, 0x9c, 0xf9, 0xa1, 0x5b, 0xd2, 0xa4, 0x8c,
	0x63, 0xa9, 0x25, 0xa6, 0x7a, 0x34, 0xe6, 0x5a, 0x75, 0x7d, 0x6a, 0x9f, 0xc6, 0x1c, 0xcb, 0x0a,
	0x3c, 0x06, 0xf5, 0xd8, 0xf1, 0x48, 0x40, 0xb4, 0x9a, 0xe4, 0xa0, 0x82, 0x53, 0x3f, 0x97, 0xe8,
	0xef, 0xd4, 0x78, 0xfc, 0xef, 0xd6, 0xd1, 0x05, 0x3e, 0xcd, 0xeb, 0xb8, 0xe8, 0x86, 0x17, 0x40,
	0xf5, 0x38, 0x8f, 0xfa, 0xc4, 0x1e, 0x13, 0x16, 0x6b, 0xdb, 0xed, 0x6a, 0x57, 0xed, 0xe9, 0x25,
	0x13, 0x48, 0xf4, 0xa2, 0xf9, 0x01, 0x12, 0x8b, 0xc9, 0x69, 0xd6, 0xc3, 0x62, 0x98, 0xba, 0xc2,
	0x62, 0x5c, 0xd6, 0x11, 0x06, 0x46, 0x74, 0x9c, 0x68, 0xf5, 0x75, 0x03, 0x16, 0x1d, 0x27, 0x58,
	0x56, 0xe0, 0x09, 0xa8, 0x5d, 0x52, 0x16, 0x68, 0x0d, 0x39, 0xf1, 0x39, 0xda, 0x7c, 0x6e, 0x68,
	0x99, 0xf1, 0x4a, 0x48, 0x40, 0x58, 0x0a, 0x74, 0x7e, 0x6d, 0x81, 0x46, 0xdf, 0x0e, 0xc7, 0x53,
	0xc2, 0xe0, 0x5b, 0x50, 0x23, 0x57, 0xc4, 0x91, 0x69, 0x6d, 0xb0, 0x71, 0x74, 0x45, 0x9c, 0x3c,
	0x5b, 0x6b, 0x47, 0x28, 0x89, 0x37, 0x96, 0x5d, 0xb0, 0x0f, 0x1a, 0xc2, 0xc3, 0x09, 0x59, 0x84,
	0xf9, 0x74, 0xd3, 0x1e, 0x4e, 0x48, 0x71, 0x1f, 0x96, 0x9a, 0xa5, 0x46, 0xa3, 0x80, 0xf0, 0xa2,
	0x1d, 0x0e, 0xc1, 0x8e, 0xf8, 0x1c, 0x2c, 0x32, 0x54, 0x7b, 0xfb, 0xff, 0x33, 0xb8, 0x7e, 0x73,
	0xd6, 0x6e, 0x96, 0x1a, 0x3b, 0x0b, 0x0c, 0x2f, 0x95, 0xe0, 0x00, 0x34, 0xb9, 0x13, 0x9d, 0x53,
	0x67, 0x42, 0xb8, 0x8c, 0x5d, 0xed, 0x3d, 0xbb, 0xef, 0x17, 0x0e, 0x0f, 0x07, 0x39, 0xa9, 0xd0,
	0xdb, 0xcb, 0x52, 0xa3, 0xb9, 0x04, 0xf1, 0x4a, 0x04, 0xbe, 0x01, 0x7b, 0x0e, 0x0d, 0xb9, 0x2d,
	0xae, 0xf4, 0xcc, 0x0e, 0x88, 0xb6, 0x2d, 0xf3, 0x7a, 0x54, 0xac, 0x79, 0xef, 0xb0, 0x5c, 0xc4,
	0xeb, 0x5c, 0xeb, 0xdd, 0xf5, 0x9d, 0x5e, 0xb9, 0xb9, 0xd3, 0x2b, 0xb7, 0x77, 0x7a, 0xe5, 0x6b,
	0xa6, 0x2b, 0xd7, 0x99, 0xae, 0xdc, 0x64, 0xba, 0x72, 0x9b, 0xe9, 0xca, 0x8f, 0x4c, 0x57, 0xbe,
	0xfd, 0xd4, 0x2b, 0x9f, 0x5b, 0x9b, 0xff, 0x73, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x76, 0x37,
	0xcb, 0x78, 0x90, 0x04, 0x00, 0x00,
}

func (m *FormEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FormEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FormEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	i -= len(m.Key)
	copy(dAtA[i:], m.Key)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Key)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HTTPPostAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPPostAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTTPPostAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Form) > 0 {
		for iNdEx := len(m.Form) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Form[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	i -= len(m.Body)
	copy(dAtA[i:], m.Body)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Body)))
	i--
	dAtA[i] = 0x32
	if len(m.HTTPHeaders) > 0 {
		for iNdEx := len(m.HTTPHeaders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HTTPHeaders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	i -= len(m.Scheme)
	copy(dAtA[i:], m.Scheme)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Scheme)))
	i--
	dAtA[i] = 0x22
	i -= len(m.Host)
	copy(dAtA[i:], m.Host)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Host)))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Port.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i -= len(m.Path)
	copy(dAtA[i:], m.Path)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Path)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Handler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Handler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.ContainerName)
	copy(dAtA[i:], m.ContainerName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ContainerName)))
	i--
	dAtA[i] = 0x2a
	if m.TCPSocket != nil {
		{
			size, err := m.TCPSocket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.HTTPPost != nil {
		{
			size, err := m.HTTPPost.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HTTPGet != nil {
		{
			size, err := m.HTTPGet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Exec != nil {
		{
			size, err := m.Exec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FormEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *HTTPPostAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Port.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Host)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Scheme)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.HTTPHeaders) > 0 {
		for _, e := range m.HTTPHeaders {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.Body)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Form) > 0 {
		for _, e := range m.Form {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Handler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exec != nil {
		l = m.Exec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.HTTPGet != nil {
		l = m.HTTPGet.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.HTTPPost != nil {
		l = m.HTTPPost.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.TCPSocket != nil {
		l = m.TCPSocket.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.ContainerName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FormEntry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FormEntry{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPPostAction) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForHTTPHeaders := "[]HTTPHeader{"
	for _, f := range this.HTTPHeaders {
		repeatedStringForHTTPHeaders += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForHTTPHeaders += "}"
	repeatedStringForForm := "[]FormEntry{"
	for _, f := range this.Form {
		repeatedStringForForm += strings.Replace(strings.Replace(f.String(), "FormEntry", "FormEntry", 1), `&`, ``, 1) + ","
	}
	repeatedStringForForm += "}"
	s := strings.Join([]string{`&HTTPPostAction{`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`Port:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Port), "IntOrString", "intstr.IntOrString", 1), `&`, ``, 1) + `,`,
		`Host:` + fmt.Sprintf("%v", this.Host) + `,`,
		`Scheme:` + fmt.Sprintf("%v", this.Scheme) + `,`,
		`HTTPHeaders:` + repeatedStringForHTTPHeaders + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`Form:` + repeatedStringForForm + `,`,
		`}`,
	}, "")
	return s
}
func (this *Handler) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Handler{`,
		`Exec:` + strings.Replace(fmt.Sprintf("%v", this.Exec), "ExecAction", "v1.ExecAction", 1) + `,`,
		`HTTPGet:` + strings.Replace(fmt.Sprintf("%v", this.HTTPGet), "HTTPGetAction", "v1.HTTPGetAction", 1) + `,`,
		`HTTPPost:` + strings.Replace(this.HTTPPost.String(), "HTTPPostAction", "HTTPPostAction", 1) + `,`,
		`TCPSocket:` + strings.Replace(fmt.Sprintf("%v", this.TCPSocket), "TCPSocketAction", "v1.TCPSocketAction", 1) + `,`,
		`ContainerName:` + fmt.Sprintf("%v", this.ContainerName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FormEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: FormEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FormEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *HTTPPostAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: HTTPPostAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPPostAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Port.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scheme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scheme = k8s_io_api_core_v1.URIScheme(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTPHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HTTPHeaders = append(m.HTTPHeaders, v1.HTTPHeader{})
			if err := m.HTTPHeaders[len(m.HTTPHeaders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Form", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Form = append(m.Form, FormEntry{})
			if err := m.Form[len(m.Form)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Handler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Handler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Exec == nil {
				m.Exec = &v1.ExecAction{}
			}
			if err := m.Exec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTPGet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HTTPGet == nil {
				m.HTTPGet = &v1.HTTPGetAction{}
			}
			if err := m.HTTPGet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTPPost", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HTTPPost == nil {
				m.HTTPPost = &HTTPPostAction{}
			}
			if err := m.HTTPPost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TCPSocket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TCPSocket == nil {
				m.TCPSocket = &v1.TCPSocketAction{}
			}
			if err := m.TCPSocket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
