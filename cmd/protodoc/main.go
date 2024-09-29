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
)

func main() {

	flags := ParseFlags(os.Stdout, os.Args)

	if HandleFlags(flags) {
		os.Exit(flags.Code())
	}

	// Run protoc command
	runCommand(flags.protoDir)

	// Initialize protodoc
	pbDoc, err := protodoc.New(
		protodoc.WithDocOut(flags.docOut),
		protodoc.WithName(flags.name),
		protodoc.WithType(protodoc.ProtodocType(flags.typeName)),
	)

	if err != nil {
		fmt.Printf("failed to initialize protoduc, err=%v\n", err)
		return
	}

	// Execute the protodoc to generate API Documentation
	if err := pbDoc.Execute(); err != nil {
		fmt.Printf("failed to execute protoduc, err=%v\n", err)
		return
	}

	// Clean Up
	if err := os.Remove(protodoc.DefaultDescriptorFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Generated the code !")
}

func runCommand(protoDir string) {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()

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

// HandleFlags checks if there's a match and returns true if it was "handled"
func HandleFlags(f *Flags) bool {
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

	return false
}
