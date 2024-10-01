package protodoc

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"anchor": func(s string) string {
		return strings.ReplaceAll(s, " ", "-")
	},
}
