package protodoc

import (
	"os"
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
func (m *mdDoc) Generate() ([]byte, error) {
	doc := m.p.generateAPIDoc()

	return RenderTemplate(ProtodocTypeMD, &doc, "")
}

func (m *mdDoc) Execute() error {
	doc, err := m.Generate()
	if err != nil {
		return err
	}

	file, err := os.Create(m.p.DestFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(doc)
	if err != nil {
		return err
	}

	return nil
}
