package protodoc

import (
	"fmt"
	"path/filepath"

	"google.golang.org/protobuf/types/descriptorpb"
)

type Option func(*IProtodoc)

// WithFileDescriptor implements option file descriptor proto
func WithFileDescriptor(fileDesc []*descriptorpb.FileDescriptorProto) Option {
	return func(p *IProtodoc) {
		p.FileDescriptors = fileDesc
	}
}

// WithType implements option ProtodocType
func WithType(typeName ProtodocType) Option {
	return func(p *IProtodoc) {
		p.TypeName = typeName
	}
}

// WithName implements option name of API Documentation
func WithName(name string) Option {
	return func(p *IProtodoc) {
		p.Name = name
	}
}

// WithDocOut implements option out directory want to be generated
func WithDocOut(docOut string) Option {
	return func(p *IProtodoc) {
		switch p.TypeName {
		case ProtodocTypeJson:
			p.DestFile = filepath.Join(docOut, fmt.Sprintf("%s.%s", DefaultApiFileName, ProtodocTypeJson.String()))
		case ProtodocTypeMD:
			p.DestFile = filepath.Join(docOut, fmt.Sprintf("%s.%s", DefaultApiFileName, ProtodocTypeMD.String()))
		}
	}
}

// WithCustomTemplate implements option the custom template want to be used
func WithCustomTemplate(customTemplate string) Option {
	return func(p *IProtodoc) {
		p.CustomTemplate = customTemplate
	}
}
