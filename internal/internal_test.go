package internal

import (
	"testing"

	"github.com/lovelyoyrmia/protodoc/options"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Mock implementation of GetExtensionFunc
func mockGetExtension(_ protoreflect.ProtoMessage, ext protoreflect.ExtensionType) interface{} {
	if ext == options.E_ApiOptions {
		return &options.APIOptions{
			Path:        "/test-path",
			Method:      "GET",
			Summary:     "Test summary",
			Description: "Test description",
			QueryParams: []*options.QueryParameter{
				{
					Name:        "param1",
					Type:        "string",
					Description: "First param",
					Required:    true,
				},
			},
		}
	}
	return nil
}
func TestRemoveTypePrefix(t *testing.T) {

	testCases := []struct {
		name          string
		field         *descriptorpb.FieldDescriptorProto
		checkResponse func(t *testing.T, res string)
	}{
		{
			name: "LABEL_OPTIONAL",
			field: &descriptorpb.FieldDescriptorProto{
				Name:     nil,
				Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				TypeName: proto.String("TYPE_STRING"),
			},
			checkResponse: func(t *testing.T, res string) {
				require.Equal(t, "string*", res)
			},
		},
		{
			name: "LABEL_REPEATED",
			field: &descriptorpb.FieldDescriptorProto{
				Name:     nil,
				Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				TypeName: proto.String("TYPE_STRING"),
			},
			checkResponse: func(t *testing.T, res string) {
				require.Equal(t, "string[]", res)
			},
		},
		{
			name: "LABEL_MESSAGE",
			field: &descriptorpb.FieldDescriptorProto{
				Name:     nil,
				Label:    descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum(),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
				TypeName: proto.String(".example.User"),
			},
			checkResponse: func(t *testing.T, res string) {
				require.Equal(t, "#User", res)
			},
		},
		{
			name: "OTHER_LABEL",
			field: &descriptorpb.FieldDescriptorProto{
				Name:     nil,
				Label:    descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum(),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				TypeName: proto.String("TYPE_STRING"),
			},
			checkResponse: func(t *testing.T, res string) {
				require.Equal(t, "string", res)
			},
		},
	}

	packageName := "example"

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			res := RemoveTypePrefix(v.field, packageName)
			v.checkResponse(tt, res)
		})
	}
}

func TestExtractMethod(t *testing.T) {

	testCases := []struct {
		name          string
		method        *descriptorpb.MethodDescriptorProto
		checkResponse func(t *testing.T, res *methodOptions)
	}{
		{
			name: "WITHOUT_API_OPTIONS",
			method: &descriptorpb.MethodDescriptorProto{
				Name:       nil,
				InputType:  proto.String(".example.User"),
				OutputType: proto.String(".example.Customer"),
			},
			checkResponse: func(t *testing.T, res *methodOptions) {
				require.Equal(t, "#User", res.InputType)
				require.Equal(t, "#Customer", res.OutputType)
				require.Empty(t, res.Name)
				require.Empty(t, res.Path)
				require.Empty(t, res.Summary)
				require.Empty(t, res.Description)
				require.Empty(t, res.QueryParameters)
			},
		},
		{
			name: "WITH_API_OPTIONS",
			method: &descriptorpb.MethodDescriptorProto{
				Name:       nil,
				InputType:  proto.String(".example.User"),
				OutputType: proto.String(".example.Customer"),
				Options:    &descriptorpb.MethodOptions{},
			},
			checkResponse: func(t *testing.T, res *methodOptions) {
				require.Equal(t, "#User", res.InputType)
				require.Equal(t, "#Customer", res.OutputType)
				require.Equal(t, "/test-path", res.Path)
				require.Equal(t, "GET", res.Method)
				require.Equal(t, "Test summary", res.Summary)
				require.Equal(t, "Test description", res.Description)
				require.Len(t, res.QueryParameters, 1)
				require.Equal(t, "param1", res.QueryParameters[0].Name)
				require.Equal(t, "string", res.QueryParameters[0].Type)
				require.Equal(t, "First param", res.QueryParameters[0].Description)
				require.True(t, res.QueryParameters[0].Required)
			},
		},
	}

	packageName := "example"

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			res := ExtractMethod(v.method, packageName, mockGetExtension)
			v.checkResponse(tt, res)
		})
	}
}
