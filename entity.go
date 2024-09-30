package protodoc

// APIDoc represents the overall API documentation structure.
type APIDoc struct {
	// Name of the API Documentation.
	// This should be a descriptive title of the API, providing context to the users.
	Name string `json:"name" yaml:"name"`

	// Author indicates the author of the API method documentation.
	Author string `json:"author" yaml:"author"`

	// BaseUrl is the root URL for the API, used to construct full endpoint URLs.
	BaseUrl string `json:"base_url" yaml:"base_url"`

	// Messages represent the request and response structures defined in Protobuf.
	// Each MessageDoc corresponds to a specific message type used in the API.
	Messages []MessageDoc `json:"messages" yaml:"messages"`

	// Services contains the API methods that can be invoked.
	// Each ServiceDoc corresponds to a specific service, detailing its methods.
	Services []ServiceDoc `json:"services" yaml:"services"`
}

// MessageDoc describes a Protobuf message used in API requests or responses.
type MessageDoc struct {
	// Name of the message type.
	// This is the identifier for the message, often matching the Protobuf definition.
	Name string `json:"name" name:"name"`

	// Fields represent the individual attributes within the message.
	// Each FieldDoc provides details about a specific field in the message.
	Fields []FieldDoc `json:"fields" yaml:"fields"`
}

// FieldDoc describes an individual field within a Protobuf message.
type FieldDoc struct {
	// Name of the field.
	// This is the identifier for the field, as defined in the Protobuf message.
	Name string `json:"name" yaml:"name"`

	// Type of the field.
	// This specifies the data type of the field (e.g., string, int32, custom message).
	Type string `json:"type" yaml:"type"`
}

// ServiceDoc represents a service that groups related API methods.
type ServiceDoc struct {
	// Name of the service.
	// This is the identifier for the service, often matching the Protobuf service name.
	Name string `json:"name" yaml:"name"`

	// Methods represent the individual API methods exposed by the service.
	// Each MethodDoc provides details about a specific method, including input and output types.
	Methods []MethodDoc `json:"methods" yaml:"methods"`
}

// MethodDoc describes an individual API method within a service.
type MethodDoc struct {
	// Name of the method.
	// This is the identifier for the method, as defined in the Protobuf service.
	Name string `json:"name" yaml:"name"`

	// Summary provides a brief overview of what the method does.
	Summary string `json:"summary" yaml:"summary"`

	// Description offers a detailed explanation of the method's functionality and usage.
	Description string `json:"description" yaml:"description"`

	// Path is the API endpoint for the method, specifying how to access it.
	Path string `json:"path" yaml:"path"`

	// Method is the HTTP method (GET, POST, etc.) for the API call.
	Method string `json:"method" yaml:"method"`

	// InputType represents the expected input message type for the method.
	// This defines the structure of the data that the method will receive.
	// If the input type has an '@' prefix, it means it refers to an existing message type in the Protobuf schema.
	// The default message types are common types (string, int, etc).
	InputType string `json:"input_type" yaml:"input_type"`

	// OutputType represents the expected output message type for the method.
	// This defines the structure of the data that the method will return.
	// If the output type has an '@' prefix, it means it refers to an existing message type in the Protobuf schema.
	// The default message types are common types (string, int, etc).
	OutputType string `json:"output_type" yaml:"output_type"`

	// QueryParams represents a list of query parameters for the method.
	// Each parameter includes details such as name, type, description, and whether it is required.
	QueryParams []*QueryParameterDoc `json:"query_params" yaml:"query_params"`
}

// QueryParameterDoc describes an individual query parameter for an API method.
type QueryParameterDoc struct {
	// Name of the query parameter.
	Name string `json:"name" yaml:"name"`

	// Type of the query parameter (e.g., string, int).
	Type string `json:"type" yaml:"type"`

	// Description of the query parameter.
	Description string `json:"description" yaml:"description"`

	// Indicates whether the query parameter is required.
	Required bool `json:"required" yaml:"required"`
}
