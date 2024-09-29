package main

import (
	"flag"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/lovelyoyrmia/protodoc"
)

const helpMessage = `
This library provides a simple way to generate API documentation in different types format (JSON, Markdown, Yaml) from Protocol Buffer (Protobuf) files.

EXAMPLE: Generate default docs (Markdown)
protodoc --proto_dir=protos/

EXAMPLE: Use a custom type
protodoc --doc_out=. --type=json --proto_dir=protos/

See https://github.com/lovelyoyrmia/protodoc for more details.
`

// Valid type documentation
var validTypes = []string{
	protodoc.ProtodocTypeJson.String(),
	protodoc.ProtodocTypeMD.String(),
}

// Version returns the currently running version of protodoc
func Version() string {
	return protodoc.VERSION
}

// Flags contains details about the CLI invocation of protodoc
type Flags struct {
	appName        string
	flagSet        *flag.FlagSet
	err            error
	showHelp       bool
	showVersion    bool
	name           string
	protoDir       string
	typeName       string
	docOut         string
	docOpt         string
	sourceRelative bool
	writer         io.Writer
}

// Code returns the status code to exit with after handling the supplied flags
func (f *Flags) Code() int {
	if f.err != nil {
		return 1
	}

	return 0
}

// ShowHelp determines whether or not to show the help message
func (f *Flags) ShowHelp() bool {
	return f.err != nil || f.showHelp
}

// ShowVersion determines whether or not to show the version message
func (f *Flags) ShowVersion() bool {
	return f.showVersion
}

// ShowValidTypes determines whether or not the type name is a valid type
func (f *Flags) ShowValidTypes() bool {
	// Check if the given command is in valid types
	if f.typeName == "" && !slices.Contains(validTypes, f.typeName) {
		return true
	}

	return false
}

// CheckRequiredArgs function to check args are required
func (f *Flags) CheckRequiredArgs(fields map[string]string) bool {
	var missingFields []string
	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			missingFields = append(missingFields, fieldName)
		}
	}
	if len(missingFields) > 0 {
		f.err = fmt.Errorf("error: The following fields must not be empty: %s", strings.Join(missingFields, ", "))
		return false
	}
	return true
}

// PrintHelp prints the usage string including all flags to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintHelp() {
	fmt.Fprintf(f.writer, "Usage of %s:\n", f.appName)
	fmt.Fprintf(f.writer, "%s\n", helpMessage)
	fmt.Fprintf(f.writer, "FLAGS\n")
	f.flagSet.PrintDefaults()
}

// PrintValidTypes prints the valid types string to the `io.Writer` that was supplied to the `Flags object`.
func (f *Flags) PrintValidTypes() {
	fmt.Println("Use --help or -h for usage information.")
	fmt.Printf("Error: Invalid command type '%s'.\n", f.typeName)
	fmt.Println("Valid command types are:", strings.Join(validTypes, ", "))
}

// PrintVersion prints the version string to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintVersion() {
	fmt.Fprintf(f.writer, "%s version %s\n", f.appName, Version())
}

// PrintError prints the error string
func (f *Flags) PrintError() {
	if f.err == nil {
		return
	}
	fmt.Println(f.err)
	fmt.Println("Use --help or -h for usage information.")
}

// ParseFlags parses the supplied options are returns a `Flags` object to the caller.
func ParseFlags(w io.Writer, args []string) *Flags {
	f := Flags{appName: args[0], writer: w}

	f.flagSet = flag.NewFlagSet(args[0], flag.ContinueOnError)

	f.flagSet.StringVar(&f.name, "name", protodoc.DefaultApiDocName, "name is the name of the API Documentation.")
	f.flagSet.StringVar(&f.protoDir, "proto_dir", "", "proto_dir is the directory of the all protobuf files.")
	f.flagSet.StringVar(&f.docOut, "doc_out", protodoc.DefaultApiDocsOut, "doc_out is the custom path directory of the API Documentation will be created.")
	f.flagSet.StringVar(&f.typeName, "type", protodoc.ProtodocTypeMD.String(), "type is the API Documentation type.")
	f.flagSet.StringVar(&f.docOpt, "doc_opt", "", "optional documentation options (source_relative)")

	f.flagSet.BoolVar(&f.showHelp, "help", false, "Show this help message")
	f.flagSet.BoolVar(&f.showVersion, "version", false, fmt.Sprintf("Print the current version (%v)", Version()))
	f.flagSet.SetOutput(w)

	// prevent showing help on parse error
	f.flagSet.Usage = func() {}

	f.err = f.flagSet.Parse(args[1:])
	return &f
}
