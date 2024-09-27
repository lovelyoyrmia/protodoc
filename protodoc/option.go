package protodoc

type Option func(*protodoc)

func WithType(typeName ProtodocType) Option {
	return func(p *protodoc) {
		p.typeName = typeName
	}
}

func WithName(name string) Option {
	return func(p *protodoc) {
		p.name = name
	}
}
