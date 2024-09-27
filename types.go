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
