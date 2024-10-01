package protodoc

import (
	"github.com/lovelyoyrmia/protodoc/internal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Protodoc interface {
	Generate() ([]byte, error)
	Execute() error
}

type IProtodoc struct {
	// Name of the API Documentation.
	// Default value is "API Documentation"
	Name string
	// Filename is the name of the generated protobuf file.
	Filename string
	// Destination File is the path name of the API Documentation will be created.
	DestFile string
	// Type Name is the type documentation wants to be generated
	TypeName ProtodocType
	// Custom Template is the custom template for documentation
	CustomTemplate string

	FileDescriptors []*descriptorpb.FileDescriptorProto
}

func New(opts ...Option) Protodoc {
	p := &IProtodoc{
		Name:     DefaultApiDocName,
		Filename: DefaultDescriptorFile,
		DestFile: DefaultApiFileOut,
		TypeName: ProtodocTypeMD,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.DestFile = DefaultApiFileName + p.TypeName.ExtractExtension()

	switch p.TypeName {
	case ProtodocTypeMD:
		return NewMarkdownDoc(p)
	case ProtodocTypeJson:
		return NewJsonDoc(p)
	case ProtodocTypeYaml:
		return NewYamlDoc(p)
	case ProtodocTypeHTML:
		return NewHTMLDoc(p)
	}

	return NewMarkdownDoc(p)
}

func (i *IProtodoc) generateAPIDoc() APIDoc {
	doc := APIDoc{Name: i.Name}

	for _, fileDescriptor := range i.FileDescriptors {
		for _, msg := range fileDescriptor.MessageType {
			message := MessageDoc{Name: msg.GetName()}
			for _, field := range msg.Field {
				typeName := internal.RemoveTypePrefix(field, fileDescriptor.GetPackage())
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
				option := internal.ExtractMethod(method, fileDescriptor.GetPackage(), proto.GetExtension)

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

	return doc
}
