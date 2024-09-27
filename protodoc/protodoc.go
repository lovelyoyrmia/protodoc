package protodoc

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

type Protodoc interface {
	Generate() []byte
	Execute() error
}

type protodoc struct {
	// Name of the API Documentation.
	// Default value is "API Documentation"
	name string
	// Filename is the name of the generated protobuf file.
	filename string
	// Destination File is the path name of the API Documentation will be created.
	destFile string
	// Type Name is the type documentation wants to be generated
	typeName ProtodocType

	fileDescriptors []*descriptorpb.FileDescriptorProto
}

func New(filename string, destFile string, opts ...Option) (Protodoc, error) {
	fileDescriptor, err := readFile(filename)

	if err != nil {
		return nil, err
	}

	p := &protodoc{
		name:            defaultApiDocName,
		filename:        filename,
		destFile:        destFile,
		fileDescriptors: fileDescriptor,
		typeName:        ProtodocTypeMD,
	}

	for _, opt := range opts {
		opt(p)
	}

	switch p.typeName {
	case ProtodocTypeMD:
		return newMarkdownDoc(p), nil
	case ProtodocTypeJson:
		return newJsonDoc(p), nil
	}

	return newMarkdownDoc(p), nil
}
