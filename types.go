package protodoc

import (
	_ "embed"
	"errors"
)

type ProtodocType string

const (
	ProtodocTypeJson ProtodocType = "json"
	ProtodocTypeMD   ProtodocType = "markdown"
	ProtodocTypeYaml ProtodocType = "yaml"
	ProtodocTypeHTML ProtodocType = "html"
)

var (
	//go:embed resources/markdown.tmpl
	markdownTmpl []byte
	//go:embed resources/html.tmpl
	htmlTmpl []byte
)

func (p ProtodocType) String() string {
	return string(p)
}

func (p ProtodocType) ExtractExtension() string {
	switch p {
	case ProtodocTypeJson:
		return ".json"
	case ProtodocTypeMD:
		return ".md"
	case ProtodocTypeYaml:
		return ".yaml"
	case ProtodocTypeHTML:
		return ".html"
	}

	return ".md"
}

func (p ProtodocType) Render() (Processor, error) {
	switch p {
	case ProtodocTypeMD:
		return &textRenderer{string(markdownTmpl)}, nil
	case ProtodocTypeHTML:
		return &htmlRenderer{string(htmlTmpl)}, nil
	}

	return nil, errors.New("failed to load renderer")
}
