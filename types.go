package protodoc

type ProtodocType string

const (
	ProtodocTypeJson ProtodocType = "json"
	ProtodocTypeMD   ProtodocType = "markdown"
	ProtodocTypeYaml ProtodocType = "yaml"
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
	}

	return ".md"
}
