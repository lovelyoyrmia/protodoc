package protodoc

import (
	"os"

	"gopkg.in/yaml.v3"
)

type yamlDoc struct {
	p *IProtodoc
}

func NewYamlDoc(p *IProtodoc) Protodoc {
	return &yamlDoc{p}
}

// Generate generates YAML documentation from the FileDescriptorProto.
//
// This method iterates through the file descriptors associated with the
// Protobuf definitions, extracting message and service information to
// construct a structured YAML document. The resulting YAML document
// includes the API name, detailed message types with their fields,
// and service methods with their respective input and output types.
//
// Returns:
//   - A byte slice containing the YAML representation of the API documentation.
//
// Example usage:
//
//	 // Execute the protodoc to generate API Documentation
//	 if err := yamlDoc.Execute(); err != nil {
//		   return err
//	 }
func (j *yamlDoc) Generate() ([]byte, error) {
	doc := j.p.GenerateAPIDoc()

	return yaml.Marshal(doc)
}

func (j *yamlDoc) Execute() error {
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
