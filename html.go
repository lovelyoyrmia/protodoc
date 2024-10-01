package protodoc

import (
	"os"
)

type htmlDoc struct {
	p *IProtodoc
}

func NewHTMLDoc(p *IProtodoc) Protodoc {
	return &htmlDoc{p}
}

// Generate generates HTML documentation from the FileDescriptorProto.
//
// This method processes the file descriptors associated with the
// Protobuf definitions to extract information about message types,
// fields, and service methods. It constructs a structured HTML
// representation of the API documentation.
//
// The generated HTML follows a format that is easy to read
// and can be used for documentation purposes, either in a
// README file or in a dedicated documentation site.
//
// Returns:
//   - A byte slice containing the HTML representation of the API documentation.
//
// Example usage:
//
//	 // Execute the protodoc to generate API Documentation
//	 if err := htmlDoc.Execute(); err != nil {
//		   return err
//	 }
func (m *htmlDoc) Generate() ([]byte, error) {
	doc := m.p.generateAPIDoc()

	return RenderTemplate(ProtodocTypeHTML, &doc, "")
}

func (m *htmlDoc) Execute() error {
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
