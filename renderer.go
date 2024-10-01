package protodoc

import (
	"bytes"
	"fmt"
	html_template "html/template"
	text_template "text/template"
)

type Processor interface {
	Apply(apiDoc *APIDoc) ([]byte, error)
}

func RenderTemplate(kind ProtodocType, apiDoc *APIDoc, inputTemplate string) ([]byte, error) {
	if inputTemplate != "" {
		processor := &textRenderer{inputTemplate}
		return processor.Apply(apiDoc)
	}

	processor, err := kind.Render()
	if err != nil {
		return nil, err
	}

	return processor.Apply(apiDoc)
}

type textRenderer struct {
	inputTemplate string
}

func (mr *textRenderer) Apply(apiDoc *APIDoc) ([]byte, error) {
	tmpl, err := text_template.New("Text Template").Funcs(funcMap).Parse(mr.inputTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, apiDoc); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf.Bytes(), nil
}

type htmlRenderer struct {
	inputTemplate string
}

func (mr *htmlRenderer) Apply(apiDoc *APIDoc) ([]byte, error) {
	tmpl, err := html_template.New("HTML Template").Funcs(funcMap).Parse(mr.inputTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, apiDoc); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buf.Bytes(), nil
}
