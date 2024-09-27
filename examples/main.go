package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lovelyoyrmia/protodoc"
)

// This is the main entry point for the documentation generator
func main() {
	protoDir := "."
	descOut := "api_descriptor.desc"

	// Gather all .proto files
	var protoFiles []string

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
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

	// Prepare the protoc command with all proto files
	cmdArgs := append([]string{"--descriptor_set_out=" + descOut, "--proto_path=" + protoDir}, protoFiles...)

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

	// Initialize protodoc
	mdDoc, err := protodoc.New(descOut, "API_DOCUMENTATION.md")

	if err != nil {
		fmt.Printf("failed to initialize, err=%v\n", err)
		return
	}

	// Execute the protodoc to generate API Documentation
	if err := mdDoc.Execute(); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}

	// Initialize protodoc
	jsonDoc, err := protodoc.New(descOut, "API_DOCUMENTATION.json", protodoc.WithType(protodoc.ProtodocTypeJson))

	if err != nil {
		fmt.Printf("failed to initialize, err=%v\n", err)
		return
	}

	// Execute the protodoc to generate API Documentation
	if err := jsonDoc.Execute(); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}

	if err := os.Remove(descOut); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}
}
