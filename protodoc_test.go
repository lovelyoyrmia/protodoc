package protodoc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"gopkg.in/yaml.v3"
)

// Mock data for testing
var mockFileDescriptor = &descriptorpb.FileDescriptorProto{
	Name:    proto.String("api.proto"),
	Package: proto.String("example"),
	Dependency: []string{
		"google/protobuf/descriptor.proto",
		"options/options.proto",
	},
	MessageType: []*descriptorpb.DescriptorProto{
		{
			Name: proto.String("User"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("name"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					JsonName: proto.String("name"),
				},
				{
					Name:     proto.String("email"),
					Number:   proto.Int32(2),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					JsonName: proto.String("email"),
				},
				{
					Name:     proto.String("age"),
					Number:   proto.Int32(3),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					JsonName: proto.String("age"),
				},
			},
		},
		{
			Name: proto.String("GetUserRequest"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("id"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					JsonName: proto.String("id"),
				},
			},
		},
		{
			Name: proto.String("GetUserResponse"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("user"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
					TypeName: proto.String(".example.User"),
					JsonName: proto.String("user"),
				},
			},
		},
		{
			Name: proto.String("Customer"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("name"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					JsonName: proto.String("name"),
				},
				{
					Name:     proto.String("email"),
					Number:   proto.Int32(2),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					JsonName: proto.String("email"),
				},
				{
					Name:     proto.String("age"),
					Number:   proto.Int32(3),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					JsonName: proto.String("age"),
				},
			},
		},
		{
			Name: proto.String("GetCustomerRequest"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("id"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					JsonName: proto.String("id"),
				},
			},
		},
		{
			Name: proto.String("GetCustomerResponse"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{
					Name:     proto.String("user"),
					Number:   proto.Int32(1),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
					Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
					TypeName: proto.String(".example.Customer"),
					JsonName: proto.String("user"),
				},
			},
		},
	},
}

// Test Generate method of JsonDoc
func TestGenerateJsonDoc(t *testing.T) {
	jsonDoc := APIDoc{
		Name:    "API Documentation",
		Author:  "",
		BaseUrl: "",
		Messages: []MessageDoc{
			{
				Name: "User",
				Fields: []FieldDoc{
					{Name: "name", Type: "string*"},
					{Name: "email", Type: "string*"},
					{Name: "age", Type: "int32*"},
				},
			},
			{
				Name: "GetUserRequest",
				Fields: []FieldDoc{
					{Name: "id", Type: "int32*"},
				},
			},
			{
				Name: "GetUserResponse",
				Fields: []FieldDoc{
					{Name: "user", Type: "#User*"},
				},
			},
			{
				Name: "Customer",
				Fields: []FieldDoc{
					{Name: "name", Type: "string*"},
					{Name: "email", Type: "string*"},
					{Name: "age", Type: "int32*"},
				},
			},
			{
				Name: "GetCustomerRequest",
				Fields: []FieldDoc{
					{Name: "id", Type: "int32*"},
				},
			},
			{
				Name: "GetCustomerResponse",
				Fields: []FieldDoc{
					{Name: "user", Type: "#Customer[]"},
				},
			},
		},
	}

	testCases := []struct {
		name          string
		typeName      ProtodocType
		checkResponse func(t *testing.T, protoDoc Protodoc)
	}{
		{
			name:     "JSON_TYPE",
			typeName: ProtodocTypeJson,
			checkResponse: func(t *testing.T, protoDoc Protodoc) {
				result, err := protoDoc.Generate()
				require.NoError(t, err)
				require.NotEmpty(t, result)

				var actualJson APIDoc
				err = json.Unmarshal(result, &actualJson)
				require.NoError(t, err)

				require.Equal(t, jsonDoc, actualJson)
			},
		},
		{
			name:     "MARKDOWN_TYPE",
			typeName: ProtodocTypeMD,
			checkResponse: func(t *testing.T, protoDoc Protodoc) {
				result, err := protoDoc.Generate()
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
		{
			name:     "YAML_TYPE",
			typeName: ProtodocTypeYaml,
			checkResponse: func(t *testing.T, protoDoc Protodoc) {
				result, err := protoDoc.Generate()
				require.NoError(t, err)
				require.NotEmpty(t, result)

				var actualYaml APIDoc
				err = yaml.Unmarshal(result, &actualYaml)
				require.NoError(t, err)

				require.Empty(t, actualYaml.Services)
			},
		},
		{
			name:     "HTML_TYPE",
			typeName: ProtodocTypeHTML,
			checkResponse: func(t *testing.T, protoDoc Protodoc) {
				result, err := protoDoc.Generate()
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			protoDoc := New(func(p *IProtodoc) {
				p.FileDescriptors = []*descriptorpb.FileDescriptorProto{mockFileDescriptor}
				p.TypeName = v.typeName
			})

			v.checkResponse(tt, protoDoc)
		})
	}
}
