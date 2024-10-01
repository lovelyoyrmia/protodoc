package protodoc

import (
	"encoding/json"
	"os"
)

type jsonDoc struct {
	p *IProtodoc
}

func NewJsonDoc(p *IProtodoc) Protodoc {
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
//		   return err
//	 }
func (j *jsonDoc) Generate() ([]byte, error) {
	doc := j.p.generateAPIDoc()

	return json.MarshalIndent(doc, "", "  ")
}

func (j *jsonDoc) Execute() error {
	doc, err := j.Generate()
	if err != nil {
		return err
	}

	err = os.WriteFile(j.p.DestFile, doc, 0644)
	if err != nil {
		return err
	}

	return nil
}
