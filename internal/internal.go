package internal

import (
	"fmt"
	"strings"

	"github.com/lovelyoyrmia/protodoc/options"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Function signature for getting extensions
type GetExtensionFunc func(protoreflect.ProtoMessage, protoreflect.ExtensionType) interface{}

// RemoveTypePrefix removes the "TYPE_" prefix from a given field type
// and returns a string representation of the type. It also handles
// optional and repeated labels for both primitive and message types.
func RemoveTypePrefix(field *descriptorpb.FieldDescriptorProto, packageName string) string {
	typeField := field.GetType()
	typeStr := strings.TrimPrefix(typeField.String(), "TYPE_")

	// Prepare the result string based on the field type
	var result string

	// Check if the field is optional
	if field.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL {
		result += "*"
	}

	// Check if the field is repeated
	if field.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		result += "[]"
	}

	// Handle message types
	if typeField == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		messageType := RemovePackagePrefix(field.GetTypeName(), packageName)
		return messageType + result // Append optional/repeated notation
	}

	// For other types, just return the base type
	return strings.ToLower(typeStr) + result
}

// RemovePackagePrefix function to remove the package prefix of the message.
// Example: from ".example.Message" to "@Message"
func RemovePackagePrefix(typeName string, packageName string) string {
	if strings.HasPrefix(typeName, fmt.Sprintf(".%s", packageName)) {
		typeStr := strings.TrimPrefix(typeName, fmt.Sprintf(".%s.", packageName))
		return fmt.Sprintf("#%s", typeStr)
	}

	return typeName
}

// ExtractMethod checks the annotations on the method
func ExtractMethod(method *descriptorpb.MethodDescriptorProto, packageName string, getExtension GetExtensionFunc) *methodOptions {
	option := new(methodOptions)

	option.Name = method.GetName()
	option.InputType = RemovePackagePrefix(method.GetInputType(), packageName)
	option.OutputType = RemovePackagePrefix(method.GetOutputType(), packageName)

	// Retrieve custom options
	if method.GetOptions() != nil {
		if apiOptions, ok := getExtension(method.GetOptions(), options.E_ApiOptions).(*options.APIOptions); ok {
			option.Path = apiOptions.GetPath()
			option.Method = apiOptions.GetMethod()
			option.Summary = apiOptions.GetSummary()
			option.Description = apiOptions.GetDescription()

			queryParams := make([]queryParameters, 0)

			for _, param := range apiOptions.GetQueryParams() {
				queryParam := queryParameters{
					Name:        param.GetName(),
					Type:        param.GetType(),
					Description: param.GetDescription(),
					Required:    param.GetRequired(),
				}
				queryParams = append(queryParams, queryParam)
			}

			option.QueryParameters = queryParams
		}
	}
	return option
}
