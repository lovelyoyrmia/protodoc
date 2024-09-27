package protodoc

type methodOptions struct {
	Name            string
	Summary         string
	Description     string
	Path            string
	Method          string
	InputType       string
	OutputType      string
	QueryParameters []queryParameters
}

type queryParameters struct {
	Name        string
	Type        string
	Description string
	Required    bool
}
