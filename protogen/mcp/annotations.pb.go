// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: protogen/mcp/annotations.proto

package mcp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_protogen_mcp_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "mcp.mcp_version",
		Tag:           "bytes,50000,opt,name=mcp_version",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "mcp.mcp_type",
		Tag:           "bytes,50001,opt,name=mcp_type",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "mcp.mcp_resource_service",
		Tag:           "varint,50002,opt,name=mcp_resource_service",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "mcp.mcp_resource_base_uri",
		Tag:           "bytes,50003,opt,name=mcp_resource_base_uri",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50004,
		Name:          "mcp.mcp_server_description",
		Tag:           "bytes,50004,opt,name=mcp_server_description",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "mcp.mcp_tool_description",
		Tag:           "bytes,50000,opt,name=mcp_tool_description",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "mcp.mcp_tool_title",
		Tag:           "bytes,50001,opt,name=mcp_tool_title",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "mcp.mcp_tool_read_only",
		Tag:           "varint,50002,opt,name=mcp_tool_read_only",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50003,
		Name:          "mcp.mcp_tool_destructive",
		Tag:           "varint,50003,opt,name=mcp_tool_destructive",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50004,
		Name:          "mcp.mcp_tool_idempotent",
		Tag:           "varint,50004,opt,name=mcp_tool_idempotent",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50005,
		Name:          "mcp.mcp_tool_open_world",
		Tag:           "varint,50005,opt,name=mcp_tool_open_world",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50006,
		Name:          "mcp.mcp_expose_as_tool",
		Tag:           "varint,50006,opt,name=mcp_expose_as_tool",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "mcp.mcp_path_param",
		Tag:           "varint,50000,opt,name=mcp_path_param",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50001,
		Name:          "mcp.mcp_query_param",
		Tag:           "varint,50001,opt,name=mcp_query_param",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50002,
		Name:          "mcp.mcp_resource_uri",
		Tag:           "varint,50002,opt,name=mcp_resource_uri",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "mcp.mcp_field_description",
		Tag:           "bytes,50003,opt,name=mcp_field_description",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50004,
		Name:          "mcp.mcp_required",
		Tag:           "varint,50004,opt,name=mcp_required",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50005,
		Name:          "mcp.mcp_validation_pattern",
		Tag:           "bytes,50005,opt,name=mcp_validation_pattern",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50006,
		Name:          "mcp.mcp_validation_message",
		Tag:           "bytes,50006,opt,name=mcp_validation_message",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "mcp.mcp_resource_template_message",
		Tag:           "varint,50000,opt,name=mcp_resource_template_message",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "mcp.mcp_uri_template",
		Tag:           "bytes,50001,opt,name=mcp_uri_template",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50002,
		Name:          "mcp.mcp_template_description",
		Tag:           "bytes,50002,opt,name=mcp_template_description",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "mcp.mcp_message_description",
		Tag:           "bytes,50003,opt,name=mcp_message_description",
		Filename:      "protogen/mcp/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50004,
		Name:          "mcp.mcp_expose_as_resource",
		Tag:           "varint,50004,opt,name=mcp_expose_as_resource",
		Filename:      "protogen/mcp/annotations.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// Basic MCP service configuration
	//
	// optional string mcp_version = 50000;
	E_McpVersion = &file_protogen_mcp_annotations_proto_extTypes[0]
	// optional string mcp_type = 50001;
	E_McpType = &file_protogen_mcp_annotations_proto_extTypes[1]
	// Resource configuration
	//
	// optional bool mcp_resource_service = 50002;
	E_McpResourceService = &file_protogen_mcp_annotations_proto_extTypes[2]
	// optional string mcp_resource_base_uri = 50003;
	E_McpResourceBaseUri = &file_protogen_mcp_annotations_proto_extTypes[3]
	// Description
	//
	// optional string mcp_server_description = 50004;
	E_McpServerDescription = &file_protogen_mcp_annotations_proto_extTypes[4]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Description
	//
	// optional string mcp_tool_description = 50000;
	E_McpToolDescription = &file_protogen_mcp_annotations_proto_extTypes[5]
	// Tool annotations
	// Human-readable title for the tool
	//
	// optional string mcp_tool_title = 50001;
	E_McpToolTitle = &file_protogen_mcp_annotations_proto_extTypes[6]
	// If true, the tool does not modify its environment
	//
	// optional bool mcp_tool_read_only = 50002;
	E_McpToolReadOnly = &file_protogen_mcp_annotations_proto_extTypes[7]
	// If true, the tool may perform destructive updates
	//
	// optional bool mcp_tool_destructive = 50003;
	E_McpToolDestructive = &file_protogen_mcp_annotations_proto_extTypes[8]
	// If true, repeated calls with same args have no additional effect
	//
	// optional bool mcp_tool_idempotent = 50004;
	E_McpToolIdempotent = &file_protogen_mcp_annotations_proto_extTypes[9]
	// If true, tool interacts with external entities
	//
	// optional bool mcp_tool_open_world = 50005;
	E_McpToolOpenWorld = &file_protogen_mcp_annotations_proto_extTypes[10]
	// If true, this method should be exposed as an MCP tool
	//
	// optional bool mcp_expose_as_tool = 50006;
	E_McpExposeAsTool = &file_protogen_mcp_annotations_proto_extTypes[11]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Parameter types
	//
	// optional bool mcp_path_param = 50000;
	E_McpPathParam = &file_protogen_mcp_annotations_proto_extTypes[12]
	// optional bool mcp_query_param = 50001;
	E_McpQueryParam = &file_protogen_mcp_annotations_proto_extTypes[13]
	// optional bool mcp_resource_uri = 50002;
	E_McpResourceUri = &file_protogen_mcp_annotations_proto_extTypes[14]
	// Field properties
	//
	// optional string mcp_field_description = 50003;
	E_McpFieldDescription = &file_protogen_mcp_annotations_proto_extTypes[15]
	// optional bool mcp_required = 50004;
	E_McpRequired = &file_protogen_mcp_annotations_proto_extTypes[16]
	// optional string mcp_validation_pattern = 50005;
	E_McpValidationPattern = &file_protogen_mcp_annotations_proto_extTypes[17]
	// optional string mcp_validation_message = 50006;
	E_McpValidationMessage = &file_protogen_mcp_annotations_proto_extTypes[18]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// Resource template configuration
	//
	// optional bool mcp_resource_template_message = 50000;
	E_McpResourceTemplateMessage = &file_protogen_mcp_annotations_proto_extTypes[19]
	// optional string mcp_uri_template = 50001;
	E_McpUriTemplate = &file_protogen_mcp_annotations_proto_extTypes[20]
	// optional string mcp_template_description = 50002;
	E_McpTemplateDescription = &file_protogen_mcp_annotations_proto_extTypes[21]
	// Message description
	//
	// optional string mcp_message_description = 50003;
	E_McpMessageDescription = &file_protogen_mcp_annotations_proto_extTypes[22]
	// If true, this message should be exposed as an MCP resource
	//
	// optional bool mcp_expose_as_resource = 50004;
	E_McpExposeAsResource = &file_protogen_mcp_annotations_proto_extTypes[23]
)

var File_protogen_mcp_annotations_proto protoreflect.FileDescriptor

const file_protogen_mcp_annotations_proto_rawDesc = "" +
	"\n" +
	"\x1eprotogen/mcp/annotations.proto\x12\x03mcp\x1a google/protobuf/descriptor.proto:B\n" +
	"\vmcp_version\x12\x1f.google.protobuf.ServiceOptions\x18І\x03 \x01(\tR\n" +
	"mcpVersion:<\n" +
	"\bmcp_type\x12\x1f.google.protobuf.ServiceOptions\x18ц\x03 \x01(\tR\amcpType:S\n" +
	"\x14mcp_resource_service\x12\x1f.google.protobuf.ServiceOptions\x18҆\x03 \x01(\bR\x12mcpResourceService:T\n" +
	"\x15mcp_resource_base_uri\x12\x1f.google.protobuf.ServiceOptions\x18ӆ\x03 \x01(\tR\x12mcpResourceBaseUri:W\n" +
	"\x16mcp_server_description\x12\x1f.google.protobuf.ServiceOptions\x18Ԇ\x03 \x01(\tR\x14mcpServerDescription:R\n" +
	"\x14mcp_tool_description\x12\x1e.google.protobuf.MethodOptions\x18І\x03 \x01(\tR\x12mcpToolDescription:F\n" +
	"\x0emcp_tool_title\x12\x1e.google.protobuf.MethodOptions\x18ц\x03 \x01(\tR\fmcpToolTitle:M\n" +
	"\x12mcp_tool_read_only\x12\x1e.google.protobuf.MethodOptions\x18҆\x03 \x01(\bR\x0fmcpToolReadOnly:R\n" +
	"\x14mcp_tool_destructive\x12\x1e.google.protobuf.MethodOptions\x18ӆ\x03 \x01(\bR\x12mcpToolDestructive:P\n" +
	"\x13mcp_tool_idempotent\x12\x1e.google.protobuf.MethodOptions\x18Ԇ\x03 \x01(\bR\x11mcpToolIdempotent:O\n" +
	"\x13mcp_tool_open_world\x12\x1e.google.protobuf.MethodOptions\x18Ն\x03 \x01(\bR\x10mcpToolOpenWorld:M\n" +
	"\x12mcp_expose_as_tool\x12\x1e.google.protobuf.MethodOptions\x18ֆ\x03 \x01(\bR\x0fmcpExposeAsTool:E\n" +
	"\x0emcp_path_param\x12\x1d.google.protobuf.FieldOptions\x18І\x03 \x01(\bR\fmcpPathParam:G\n" +
	"\x0fmcp_query_param\x12\x1d.google.protobuf.FieldOptions\x18ц\x03 \x01(\bR\rmcpQueryParam:I\n" +
	"\x10mcp_resource_uri\x12\x1d.google.protobuf.FieldOptions\x18҆\x03 \x01(\bR\x0emcpResourceUri:S\n" +
	"\x15mcp_field_description\x12\x1d.google.protobuf.FieldOptions\x18ӆ\x03 \x01(\tR\x13mcpFieldDescription:B\n" +
	"\fmcp_required\x12\x1d.google.protobuf.FieldOptions\x18Ԇ\x03 \x01(\bR\vmcpRequired:U\n" +
	"\x16mcp_validation_pattern\x12\x1d.google.protobuf.FieldOptions\x18Ն\x03 \x01(\tR\x14mcpValidationPattern:U\n" +
	"\x16mcp_validation_message\x12\x1d.google.protobuf.FieldOptions\x18ֆ\x03 \x01(\tR\x14mcpValidationMessage:d\n" +
	"\x1dmcp_resource_template_message\x12\x1f.google.protobuf.MessageOptions\x18І\x03 \x01(\bR\x1amcpResourceTemplateMessage:K\n" +
	"\x10mcp_uri_template\x12\x1f.google.protobuf.MessageOptions\x18ц\x03 \x01(\tR\x0emcpUriTemplate:[\n" +
	"\x18mcp_template_description\x12\x1f.google.protobuf.MessageOptions\x18҆\x03 \x01(\tR\x16mcpTemplateDescription:Y\n" +
	"\x17mcp_message_description\x12\x1f.google.protobuf.MessageOptions\x18ӆ\x03 \x01(\tR\x15mcpMessageDescription:V\n" +
	"\x16mcp_expose_as_resource\x12\x1f.google.protobuf.MessageOptions\x18Ԇ\x03 \x01(\bR\x13mcpExposeAsResourceB1Z/github.com/omrikiei/protoc-gen-mcp/protogen/mcpb\x06proto3"

var file_protogen_mcp_annotations_proto_goTypes = []any{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
	(*descriptorpb.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
}
var file_protogen_mcp_annotations_proto_depIdxs = []int32{
	0,  // 0: mcp.mcp_version:extendee -> google.protobuf.ServiceOptions
	0,  // 1: mcp.mcp_type:extendee -> google.protobuf.ServiceOptions
	0,  // 2: mcp.mcp_resource_service:extendee -> google.protobuf.ServiceOptions
	0,  // 3: mcp.mcp_resource_base_uri:extendee -> google.protobuf.ServiceOptions
	0,  // 4: mcp.mcp_server_description:extendee -> google.protobuf.ServiceOptions
	1,  // 5: mcp.mcp_tool_description:extendee -> google.protobuf.MethodOptions
	1,  // 6: mcp.mcp_tool_title:extendee -> google.protobuf.MethodOptions
	1,  // 7: mcp.mcp_tool_read_only:extendee -> google.protobuf.MethodOptions
	1,  // 8: mcp.mcp_tool_destructive:extendee -> google.protobuf.MethodOptions
	1,  // 9: mcp.mcp_tool_idempotent:extendee -> google.protobuf.MethodOptions
	1,  // 10: mcp.mcp_tool_open_world:extendee -> google.protobuf.MethodOptions
	1,  // 11: mcp.mcp_expose_as_tool:extendee -> google.protobuf.MethodOptions
	2,  // 12: mcp.mcp_path_param:extendee -> google.protobuf.FieldOptions
	2,  // 13: mcp.mcp_query_param:extendee -> google.protobuf.FieldOptions
	2,  // 14: mcp.mcp_resource_uri:extendee -> google.protobuf.FieldOptions
	2,  // 15: mcp.mcp_field_description:extendee -> google.protobuf.FieldOptions
	2,  // 16: mcp.mcp_required:extendee -> google.protobuf.FieldOptions
	2,  // 17: mcp.mcp_validation_pattern:extendee -> google.protobuf.FieldOptions
	2,  // 18: mcp.mcp_validation_message:extendee -> google.protobuf.FieldOptions
	3,  // 19: mcp.mcp_resource_template_message:extendee -> google.protobuf.MessageOptions
	3,  // 20: mcp.mcp_uri_template:extendee -> google.protobuf.MessageOptions
	3,  // 21: mcp.mcp_template_description:extendee -> google.protobuf.MessageOptions
	3,  // 22: mcp.mcp_message_description:extendee -> google.protobuf.MessageOptions
	3,  // 23: mcp.mcp_expose_as_resource:extendee -> google.protobuf.MessageOptions
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	0,  // [0:24] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protogen_mcp_annotations_proto_init() }
func file_protogen_mcp_annotations_proto_init() {
	if File_protogen_mcp_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protogen_mcp_annotations_proto_rawDesc), len(file_protogen_mcp_annotations_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 24,
			NumServices:   0,
		},
		GoTypes:           file_protogen_mcp_annotations_proto_goTypes,
		DependencyIndexes: file_protogen_mcp_annotations_proto_depIdxs,
		ExtensionInfos:    file_protogen_mcp_annotations_proto_extTypes,
	}.Build()
	File_protogen_mcp_annotations_proto = out.File
	file_protogen_mcp_annotations_proto_goTypes = nil
	file_protogen_mcp_annotations_proto_depIdxs = nil
}
