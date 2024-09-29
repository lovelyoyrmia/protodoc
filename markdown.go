package protodoc

import (
	"fmt"
	"os"
	"strings"

	"github.com/lovelyoyrmia/protodoc/internal"
)

type mdDoc struct {
	p *IProtodoc
}

func NewMarkdownDoc(p *IProtodoc) Protodoc {
	return &mdDoc{p}
}

// Generate generates Markdown documentation from the FileDescriptorProto.
//
// This method processes the file descriptors associated with the
// Protobuf definitions to extract information about message types,
// fields, and service methods. It constructs a structured Markdown
// representation of the API documentation.
//
// The generated Markdown follows a format that is easy to read
// and can be used for documentation purposes, either in a
// README file or in a dedicated documentation site.
//
// Returns:
//   - A byte slice containing the Markdown representation of the API documentation.
//
// Example usage:
//
//	 // Execute the protodoc to generate API Documentation
//	 if err := mdDoc.Execute(); err != nil {
//		   return err
//	 }
func (m *mdDoc) Generate() []byte {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", m.p.Name))

	for _, fileDescriptor := range m.p.FileDescriptors {
		packageName := fileDescriptor.GetPackage()

		for _, msg := range fileDescriptor.MessageType {
			sb.WriteString("### Message: " + msg.GetName() + "\n")
			sb.WriteString("| Field Name | Type |\n")
			sb.WriteString("|------------|------|\n")

			for _, field := range msg.Field {
				typeName := internal.RemoveTypePrefix(field, packageName)

				sb.WriteString(
					fmt.Sprintf("| %s | %s |\n",
						field.GetName(),
						typeName,
					),
				)
			}
			sb.WriteString("\n")
		}

		for _, service := range fileDescriptor.Service {
			sb.WriteString("### Service: " + service.GetName() + "\n")

			for _, method := range service.Method {
				inputType := internal.RemovePackagePrefix(method.GetInputType(), packageName)
				outputType := internal.RemovePackagePrefix(method.GetOutputType(), packageName)
				// Add path information (assuming you have a way to derive it)
				path := "/" + service.GetName() + "/" + method.GetName() // Example path

				sb.WriteString(fmt.Sprintf("### Method: %s\n", method.GetName()))
				sb.WriteString(fmt.Sprintf("Endpoint: %s\n", path))
				sb.WriteString(fmt.Sprintf("- **Input Type:** %s\n", inputType))
				sb.WriteString(fmt.Sprintf("- **Output Type:** %s\n\n", outputType))
			}
		}
	}

	return []byte(sb.String())
}

func (m *mdDoc) Execute() error {
	doc := m.Generate()

	file, err := os.Create(m.p.DestFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(doc)
	if err != nil {
		return err
	}

	_, err = file.WriteString(GeneratedComments)
	if err != nil {
		return err
	}

	return nil
}
