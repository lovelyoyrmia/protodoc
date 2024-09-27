package protodoc

type Option func(*IProtodoc)

// WithType implements option ProtodocType
func WithType(typeName ProtodocType) Option {
	return func(p *IProtodoc) {
		p.TypeName = typeName
	}
}

// WithType implements option name of API Documentation
func WithName(name string) Option {
	return func(p *IProtodoc) {
		p.Name = name
	}
}
