// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: tweet/v1/tweet.proto

package tweetv1

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

type GetTweetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTweetsRequest) Reset() {
	*x = GetTweetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_v1_tweet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetsRequest) ProtoMessage() {}

func (x *GetTweetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_v1_tweet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetsRequest.ProtoReflect.Descriptor instead.
func (*GetTweetsRequest) Descriptor() ([]byte, []int) {
	return file_tweet_v1_tweet_proto_rawDescGZIP(), []int{0}
}

func (x *GetTweetsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetTweetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tweets []*GetTweetsResponse_Tweet `protobuf:"bytes,1,rep,name=tweets,proto3" json:"tweets,omitempty"`
}

func (x *GetTweetsResponse) Reset() {
	*x = GetTweetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_v1_tweet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetsResponse) ProtoMessage() {}

func (x *GetTweetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_v1_tweet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetsResponse.ProtoReflect.Descriptor instead.
func (*GetTweetsResponse) Descriptor() ([]byte, []int) {
	return file_tweet_v1_tweet_proto_rawDescGZIP(), []int{1}
}

func (x *GetTweetsResponse) GetTweets() []*GetTweetsResponse_Tweet {
	if x != nil {
		return x.Tweets
	}
	return nil
}

type SetTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SetTweetRequest) Reset() {
	*x = SetTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_v1_tweet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTweetRequest) ProtoMessage() {}

func (x *SetTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_v1_tweet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTweetRequest.ProtoReflect.Descriptor instead.
func (*SetTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweet_v1_tweet_proto_rawDescGZIP(), []int{2}
}

func (x *SetTweetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SetTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SetTweetResponse) Reset() {
	*x = SetTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_v1_tweet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTweetResponse) ProtoMessage() {}

func (x *SetTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_v1_tweet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTweetResponse.ProtoReflect.Descriptor instead.
func (*SetTweetResponse) Descriptor() ([]byte, []int) {
	return file_tweet_v1_tweet_proto_rawDescGZIP(), []int{3}
}

func (x *SetTweetResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetTweetsResponse_Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId  string `protobuf:"bytes,1,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *GetTweetsResponse_Tweet) Reset() {
	*x = GetTweetsResponse_Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_v1_tweet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetsResponse_Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetsResponse_Tweet) ProtoMessage() {}

func (x *GetTweetsResponse_Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_v1_tweet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetsResponse_Tweet.ProtoReflect.Descriptor instead.
func (*GetTweetsResponse_Tweet) Descriptor() ([]byte, []int) {
	return file_tweet_v1_tweet_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetTweetsResponse_Tweet) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *GetTweetsResponse_Tweet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetTweetsResponse_Tweet) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

var File_tweet_v1_tweet_proto protoreflect.FileDescriptor

var file_tweet_v1_tweet_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa3, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x06, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x1a, 0x53,
	0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0d, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65,
	0x74, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x08, 0x53, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tweet_v1_tweet_proto_rawDescOnce sync.Once
	file_tweet_v1_tweet_proto_rawDescData = file_tweet_v1_tweet_proto_rawDesc
)

func file_tweet_v1_tweet_proto_rawDescGZIP() []byte {
	file_tweet_v1_tweet_proto_rawDescOnce.Do(func() {
		file_tweet_v1_tweet_proto_rawDescData = protoimpl.X.CompressGZIP(file_tweet_v1_tweet_proto_rawDescData)
	})
	return file_tweet_v1_tweet_proto_rawDescData
}

var file_tweet_v1_tweet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tweet_v1_tweet_proto_goTypes = []interface{}{
	(*GetTweetsRequest)(nil),        // 0: tweet.v1.GetTweetsRequest
	(*GetTweetsResponse)(nil),       // 1: tweet.v1.GetTweetsResponse
	(*SetTweetRequest)(nil),         // 2: tweet.v1.SetTweetRequest
	(*SetTweetResponse)(nil),        // 3: tweet.v1.SetTweetResponse
	(*GetTweetsResponse_Tweet)(nil), // 4: tweet.v1.GetTweetsResponse.Tweet
}
var file_tweet_v1_tweet_proto_depIdxs = []int32{
	4, // 0: tweet.v1.GetTweetsResponse.tweets:type_name -> tweet.v1.GetTweetsResponse.Tweet
	0, // 1: tweet.v1.TweetsService.GetTweets:input_type -> tweet.v1.GetTweetsRequest
	2, // 2: tweet.v1.TweetsService.SetTweet:input_type -> tweet.v1.SetTweetRequest
	1, // 3: tweet.v1.TweetsService.GetTweets:output_type -> tweet.v1.GetTweetsResponse
	3, // 4: tweet.v1.TweetsService.SetTweet:output_type -> tweet.v1.SetTweetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tweet_v1_tweet_proto_init() }
func file_tweet_v1_tweet_proto_init() {
	if File_tweet_v1_tweet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tweet_v1_tweet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetsRequest); i {
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
		file_tweet_v1_tweet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetsResponse); i {
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
		file_tweet_v1_tweet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTweetRequest); i {
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
		file_tweet_v1_tweet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTweetResponse); i {
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
		file_tweet_v1_tweet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetsResponse_Tweet); i {
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
			RawDescriptor: file_tweet_v1_tweet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tweet_v1_tweet_proto_goTypes,
		DependencyIndexes: file_tweet_v1_tweet_proto_depIdxs,
		MessageInfos:      file_tweet_v1_tweet_proto_msgTypes,
	}.Build()
	File_tweet_v1_tweet_proto = out.File
	file_tweet_v1_tweet_proto_rawDesc = nil
	file_tweet_v1_tweet_proto_goTypes = nil
	file_tweet_v1_tweet_proto_depIdxs = nil
}
