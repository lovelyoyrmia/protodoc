package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lovelyoyrmia/protodoc"
	"github.com/lovelyoyrmia/protodoc/internal"
)

func main() {

	flags := ParseFlags(os.Stdout, os.Args)

	if flags.handleFlags() {
		os.Exit(flags.Code())
	}

	flags.parseOptions()

	// Run protoc command
	flags.runCommand()

	fileDesc, err := internal.ReadFile(protodoc.DefaultDescriptorFile)

	if err != nil {
		fmt.Printf("failed to initialize protoduc, err=%v\n", err)
		return
	}

	// Initialize protodoc
	pbDoc := protodoc.New(
		protodoc.WithDocOut(flags.docOut),
		protodoc.WithName(flags.name),
		protodoc.WithCustomTemplate(flags.customTemplate),
		protodoc.WithType(protodoc.ProtodocType(flags.typeName)),
		protodoc.WithFileDescriptor(fileDesc),
	)

	// Execute the protodoc to generate API Documentation
	if err := pbDoc.Execute(); err != nil {
		fmt.Printf("failed to execute protoduc, err=%v\n", err)
		return
	}

	// Clean Up
	if err := os.Remove(protodoc.DefaultDescriptorFile); err != nil {
		fmt.Printf("failed to remove desc file: err=%v\n", err)
		return
	}

	fmt.Printf("✅ Generated the %s documentation !\n", flags.typeName)
}

func (f *Flags) runCommand() {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()

	// Gather all .proto files
	protoFiles, err := f.getAllProtoFiles()
	if err != nil {
		fmt.Println("error reading proto files!")
		return
	}

	// Prepare the protoc command with all proto files
	cmdArgs := append([]string{"--descriptor_set_out=" + protodoc.DefaultDescriptorFile, "--proto_path=" + f.protoDir}, protoFiles...)

	// Exec command protoc to generate descriptor file
	cmd := exec.Command("protoc", cmdArgs...)

	// Capture output and error
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running protoc:", err)
		fmt.Printf("Output: %s\n", out.String())
		fmt.Printf("Error Output: %s\n", stderr.String())
		return
	}
}

// parseOptions parses the documentation options from the Flags struct.
func (f *Flags) parseOptions() {
	if f.docOpt != "" && f.docOpt == "source_relative" {
		f.sourceRelative = true
		return
	}
}

// handleFlags checks if there's a match and returns true if it was "handled"
func (f *Flags) handleFlags() bool {
	if f.ShowHelp() {
		f.PrintHelp()
		return true
	}

	if f.ShowVersion() {
		f.PrintVersion()
		return true
	}

	// Check all required fields
	if !f.CheckRequiredArgs(map[string]string{
		"proto_dir": f.protoDir,
	}) {
		f.PrintError()
		return true
	}

	if f.ShowValidTypes() {
		f.PrintValidTypes()
		return true
	}

	if !f.CheckCustomTemplate() {
		f.PrintError()
		return true
	}

	return false
}

// loadTemplate tries to load a template from a custom path
func (f *Flags) loadTemplate() (string, error) {
	// Check if the file exists at the given path
	if _, err := os.Stat(f.customTemplatePath); os.IsNotExist(err) {
		return "", err // File does not exist
	}

	// Read the file contents
	content, err := os.ReadFile(f.customTemplatePath)
	if err != nil {
		return "", err // Error reading file
	}

	return string(content), nil
}

func (f *Flags) getAllProtoFiles() ([]string, error) {
	var protoFiles []string

	if !f.sourceRelative {
		err := filepath.Walk(f.protoDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {
				protoFiles = append(protoFiles, path)
			}
			return nil
		})

		if err != nil {
			return []string{}, err
		}

		return protoFiles, nil
	}

	files, err := os.ReadDir(f.protoDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".proto" {
			protoFiles = append(protoFiles, filepath.Join(f.protoDir, file.Name()))
		}
	}

	return protoFiles, nil
}
