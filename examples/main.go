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

// This is the main examples for the documentation generator
func main() {
	protoDir := "."
	descOut := "api.desc"

	// If the path is source relative.
	// Examples:
	// --proto_dir=. --doc_opt=source_relative
	// this command will scan and files in `./*.proto`
	sourceRelative := true

	// Gather all .proto files
	runCommand(protoDir, sourceRelative)

	// Read file descriptor
	fileDesc, err := internal.ReadFile(protodoc.DefaultDescriptorFile)

	if err != nil {
		fmt.Printf("failed to execute internal.ReadFile, err=%v\n", err)
		return
	}

	// Initialize protodoc type Markdown
	mdDoc := protodoc.New(protodoc.WithFileDescriptor(fileDesc))

	// Execute the protodoc to generate API Documentation
	if err := mdDoc.Execute(); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}

	// Initialize protodoc type JSON
	jsonDoc := protodoc.New(
		protodoc.WithType(protodoc.ProtodocTypeJson),
		protodoc.WithFileDescriptor(fileDesc),
	)

	// Execute the protodoc to generate API Documentation
	if err := jsonDoc.Execute(); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}

	// Initialize protodoc type YAML
	yamlDoc := protodoc.New(
		protodoc.WithType(protodoc.ProtodocTypeYaml),
		protodoc.WithFileDescriptor(fileDesc),
	)

	// Execute the protodoc to generate API Documentation
	if err := yamlDoc.Execute(); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}

	if err := os.Remove(descOut); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}
}

func runCommand(protoDir string, sourceRelative bool) {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()

	// Gather all .proto files
	protoFiles, err := getAllProtoFiles(protoDir, sourceRelative)
	if err != nil {
		fmt.Println("error reading proto files!")
		return
	}

	// Prepare the protoc command with all proto files
	cmdArgs := append([]string{"--descriptor_set_out=" + protodoc.DefaultDescriptorFile, "--proto_path=" + protoDir}, protoFiles...)

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

func getAllProtoFiles(protoDir string, sourceRelative bool) ([]string, error) {
	var protoFiles []string

	if !sourceRelative {
		err := filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
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

	files, err := os.ReadDir(protoDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".proto" {
			protoFiles = append(protoFiles, filepath.Join(protoDir, file.Name()))
		}
	}

	return protoFiles, nil
}
