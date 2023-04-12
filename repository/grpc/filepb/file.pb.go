// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: repository/grpc/filepb/file.proto

package filepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//File represents domain object for storing and retreiving files
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl   string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	LocalName     string `protobuf:"bytes,2,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	FileExtension string `protobuf:"bytes,3,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	FileSize      string `protobuf:"bytes,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	DownloadDate  string `protobuf:"bytes,5,opt,name=download_date,json=downloadDate,proto3" json:"download_date,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_grpc_filepb_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_repository_grpc_filepb_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_repository_grpc_filepb_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *File) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *File) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *File) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

func (x *File) GetDownloadDate() string {
	if x != nil {
		return x.DownloadDate
	}
	return ""
}

type DownloadFromTextFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []byte `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *DownloadFromTextFileRequest) Reset() {
	*x = DownloadFromTextFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_grpc_filepb_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFromTextFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFromTextFileRequest) ProtoMessage() {}

func (x *DownloadFromTextFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_grpc_filepb_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFromTextFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFromTextFileRequest) Descriptor() ([]byte, []int) {
	return file_repository_grpc_filepb_file_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadFromTextFileRequest) GetLinks() []byte {
	if x != nil {
		return x.Links
	}
	return nil
}

type DownloadFromTextFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadFromTextFileResponse) Reset() {
	*x = DownloadFromTextFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_grpc_filepb_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFromTextFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFromTextFileResponse) ProtoMessage() {}

func (x *DownloadFromTextFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repository_grpc_filepb_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFromTextFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFromTextFileResponse) Descriptor() ([]byte, []int) {
	return file_repository_grpc_filepb_file_proto_rawDescGZIP(), []int{2}
}

type FetchFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FetchFilesRequest) Reset() {
	*x = FetchFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_grpc_filepb_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFilesRequest) ProtoMessage() {}

func (x *FetchFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_grpc_filepb_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFilesRequest.ProtoReflect.Descriptor instead.
func (*FetchFilesRequest) Descriptor() ([]byte, []int) {
	return file_repository_grpc_filepb_file_proto_rawDescGZIP(), []int{3}
}

func (x *FetchFilesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FetchFilesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FetchFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []*File `protobuf:"bytes,1,rep,name=file,proto3" json:"file,omitempty"`
}

func (x *FetchFilesResponse) Reset() {
	*x = FetchFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_grpc_filepb_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFilesResponse) ProtoMessage() {}

func (x *FetchFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repository_grpc_filepb_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFilesResponse.ProtoReflect.Descriptor instead.
func (*FetchFilesResponse) Descriptor() ([]byte, []int) {
	return file_repository_grpc_filepb_file_proto_rawDescGZIP(), []int{4}
}

func (x *FetchFilesResponse) GetFile() []*File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_repository_grpc_filepb_file_proto protoreflect.FileDescriptor

var file_repository_grpc_filepb_file_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x33, 0x0a,
	0x1b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x78,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xb1, 0x01, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x14, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x18, 0x5a, 0x16, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_repository_grpc_filepb_file_proto_rawDescOnce sync.Once
	file_repository_grpc_filepb_file_proto_rawDescData = file_repository_grpc_filepb_file_proto_rawDesc
)

func file_repository_grpc_filepb_file_proto_rawDescGZIP() []byte {
	file_repository_grpc_filepb_file_proto_rawDescOnce.Do(func() {
		file_repository_grpc_filepb_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_grpc_filepb_file_proto_rawDescData)
	})
	return file_repository_grpc_filepb_file_proto_rawDescData
}

var file_repository_grpc_filepb_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_repository_grpc_filepb_file_proto_goTypes = []interface{}{
	(*File)(nil),                         // 0: file.File
	(*DownloadFromTextFileRequest)(nil),  // 1: file.DownloadFromTextFileRequest
	(*DownloadFromTextFileResponse)(nil), // 2: file.DownloadFromTextFileResponse
	(*FetchFilesRequest)(nil),            // 3: file.FetchFilesRequest
	(*FetchFilesResponse)(nil),           // 4: file.FetchFilesResponse
}
var file_repository_grpc_filepb_file_proto_depIdxs = []int32{
	0, // 0: file.FetchFilesResponse.file:type_name -> file.File
	1, // 1: file.FileService.DownloadFromTextFile:input_type -> file.DownloadFromTextFileRequest
	3, // 2: file.FileService.FetchFiles:input_type -> file.FetchFilesRequest
	2, // 3: file.FileService.DownloadFromTextFile:output_type -> file.DownloadFromTextFileResponse
	4, // 4: file.FileService.FetchFiles:output_type -> file.FetchFilesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_repository_grpc_filepb_file_proto_init() }
func file_repository_grpc_filepb_file_proto_init() {
	if File_repository_grpc_filepb_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repository_grpc_filepb_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_repository_grpc_filepb_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFromTextFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_repository_grpc_filepb_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFromTextFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_repository_grpc_filepb_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFilesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_repository_grpc_filepb_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFilesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_repository_grpc_filepb_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repository_grpc_filepb_file_proto_goTypes,
		DependencyIndexes: file_repository_grpc_filepb_file_proto_depIdxs,
		MessageInfos:      file_repository_grpc_filepb_file_proto_msgTypes,
	}.Build()
	File_repository_grpc_filepb_file_proto = out.File
	file_repository_grpc_filepb_file_proto_rawDesc = nil
	file_repository_grpc_filepb_file_proto_goTypes = nil
	file_repository_grpc_filepb_file_proto_depIdxs = nil
}
