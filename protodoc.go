package protodoc

import (
	"github.com/lovelyoyrmia/protodoc/internal"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Protodoc interface {
	Generate() []byte
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

	FileDescriptors []*descriptorpb.FileDescriptorProto
}

func New(opts ...Option) (Protodoc, error) {
	fileDescriptor, err := internal.ReadFile(DefaultDescriptorFile)

	if err != nil {
		return nil, err
	}

	p := &IProtodoc{
		Name:            DefaultApiDocName,
		Filename:        DefaultDescriptorFile,
		DestFile:        DefaultApiFileOut,
		FileDescriptors: fileDescriptor,
		TypeName:        ProtodocTypeMD,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.DestFile = DefaultApiFileName + p.TypeName.ExtractExtension()

	switch p.TypeName {
	case ProtodocTypeMD:
		return NewMarkdownDoc(p), nil
	case ProtodocTypeJson:
		return NewJsonDoc(p), nil
	}

	return NewMarkdownDoc(p), nil
}
