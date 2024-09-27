package protodoc

import (
	"encoding/json"
	"os"
)

// JsonDoc represents the overall API documentation structure.
type JsonDoc struct {
	// Name of the API Documentation.
	// This should be a descriptive title of the API, providing context to the users.
	Name string `json:"name"`

	// Author indicates the author of the API method documentation.
	Author string `json:"author"`

	// BaseUrl is the root URL for the API, used to construct full endpoint URLs.
	BaseUrl string `json:"base_url"`

	// Messages represent the request and response structures defined in Protobuf.
	// Each MessageDoc corresponds to a specific message type used in the API.
	Messages []MessageDoc `json:"messages"`

	// Services contains the API methods that can be invoked.
	// Each ServiceDoc corresponds to a specific service, detailing its methods.
	Services []ServiceDoc `json:"services"`
}

// MessageDoc describes a Protobuf message used in API requests or responses.
type MessageDoc struct {
	// Name of the message type.
	// This is the identifier for the message, often matching the Protobuf definition.
	Name string `json:"name"`

	// Fields represent the individual attributes within the message.
	// Each FieldDoc provides details about a specific field in the message.
	Fields []FieldDoc `json:"fields"`
}

// FieldDoc describes an individual field within a Protobuf message.
type FieldDoc struct {
	// Name of the field.
	// This is the identifier for the field, as defined in the Protobuf message.
	Name string `json:"name"`

	// Type of the field.
	// This specifies the data type of the field (e.g., string, int32, custom message).
	Type string `json:"type"`
}

// ServiceDoc represents a service that groups related API methods.
type ServiceDoc struct {
	// Name of the service.
	// This is the identifier for the service, often matching the Protobuf service name.
	Name string `json:"name"`

	// Methods represent the individual API methods exposed by the service.
	// Each MethodDoc provides details about a specific method, including input and output types.
	Methods []MethodDoc `json:"methods"`
}

// MethodDoc describes an individual API method within a service.
type MethodDoc struct {
	// Name of the method.
	// This is the identifier for the method, as defined in the Protobuf service.
	Name string `json:"name"`

	// Summary provides a brief overview of what the method does.
	Summary string `json:"summary"`

	// Description offers a detailed explanation of the method's functionality and usage.
	Description string `json:"description"`

	// Path is the API endpoint for the method, specifying how to access it.
	Path string `json:"path"`

	// Method is the HTTP method (GET, POST, etc.) for the API call.
	Method string `json:"method"`

	// InputType represents the expected input message type for the method.
	// This defines the structure of the data that the method will receive.
	// If the input type has an '@' prefix, it means it refers to an existing message type in the Protobuf schema.
	// The default message types are common types (string, int, etc).
	InputType string `json:"input_type"`

	// OutputType represents the expected output message type for the method.
	// This defines the structure of the data that the method will return.
	// If the output type has an '@' prefix, it means it refers to an existing message type in the Protobuf schema.
	// The default message types are common types (string, int, etc).
	OutputType string `json:"output_type"`

	// QueryParams represents a list of query parameters for the method.
	// Each parameter includes details such as name, type, description, and whether it is required.
	QueryParams []*QueryParameterDoc `json:"query_params"`
}

// QueryParameterDoc describes an individual query parameter for an API method.
type QueryParameterDoc struct {
	// Name of the query parameter.
	Name string `json:"name"`

	// Type of the query parameter (e.g., string, int).
	Type string `json:"type"`

	// Description of the query parameter.
	Description string `json:"description"`

	// Indicates whether the query parameter is required.
	Required bool `json:"required"`
}

type jsonDoc struct {
	p *protodoc
}

func newJsonDoc(p *protodoc) Protodoc {
	return &jsonDoc{p}
}

// Generate generates JSON documentation from the FileDescriptorProto.
//
// This method iterates through the file descriptors associated with the
// Protobuf definitions, extracting message and service information to
// construct a structured JSON document. The resulting JSON document
// includes the API name, detailed message types with their fields,
// and service methods with their respective input and output types.
//
// Returns:
//   - A byte slice containing the JSON representation of the API documentation.
//
// Example usage:
//
//	 // Execute the protodoc to generate API Documentation
//	 if err := jsonDoc.Execute(); err != nil {
//		   fmt.Printf("failed to execute, err=%v\n", err)
//		   return
//	 }
func (j *jsonDoc) Generate() []byte {
	doc := JsonDoc{Name: j.p.name}

	for _, fileDescriptor := range j.p.fileDescriptors {
		for _, msg := range fileDescriptor.MessageType {
			message := MessageDoc{Name: msg.GetName()}
			for _, field := range msg.Field {
				typeName := removeTypePrefix(field, fileDescriptor.GetPackage())
				message.Fields = append(message.Fields, FieldDoc{
					Name: field.GetName(),
					Type: typeName,
				})
			}
			doc.Messages = append(doc.Messages, message)
		}

		for _, service := range fileDescriptor.Service {
			serviceDoc := ServiceDoc{Name: service.GetName()}

			for _, method := range service.Method {
				option := extractMethod(method, fileDescriptor.GetPackage())

				// Convert the query parameters
				queryParams := make([]*QueryParameterDoc, 0)

				for _, query := range option.QueryParameters {
					queryParams = append(queryParams, &QueryParameterDoc{
						Name:        query.Name,
						Type:        query.Type,
						Description: query.Description,
						Required:    query.Required,
					})
				}

				serviceDoc.Methods = append(serviceDoc.Methods, MethodDoc{
					Name:        option.Name,
					Summary:     option.Summary,
					Description: option.Description,
					Path:        option.Path,
					Method:      option.Method,
					InputType:   option.InputType,
					OutputType:  option.OutputType,
					QueryParams: queryParams,
				})
			}
			doc.Services = append(doc.Services, serviceDoc)
		}
	}

	res, _ := json.MarshalIndent(doc, "", "  ")
	return res
}

func (j *jsonDoc) Execute() error {
	doc := j.Generate()

	err := os.WriteFile(j.p.destFile, doc, 0644)
	if err != nil {
		return err
	}

	return nil
}
