package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/lovelyoyrmia/protobuf-documentation/protodoc"
)

func main() {

	var name, protodir, destFile, typeName, filename string

	flag.StringVar(&name, "name", "API Documentation", "name is the name of the API Documentation.")
	flag.StringVar(&protodir, "proto_dir", "", "proto_dir is the directory of the all protobuf files.")
	flag.StringVar(&destFile, "dest_file", "", "dest_file is the path name of the API Documentation will be created.")
	flag.StringVar(&typeName, "type", protodoc.ProtodocTypeMD.String(), "type is the API Documentation type.")
	flag.StringVar(&filename, "file_desc", "", "file_desc is the path name of the generated descriptor file.")

	// Help tag
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Println("Usage:")
		flag.PrintDefaults() // Print all flag descriptions
		return
	}

	// Parse the flags
	flag.Parse()

	// Check all required fields
	if err := checkRequiredFields(map[string]string{
		"proto_dir": protodir,
		"dest_file": destFile,
		"file_desc": filename,
	}); err != nil {
		fmt.Println(err)
		fmt.Println("Use --help or -h for usage information.")
		return
	}

	// Valid type documentation
	validTypes := []string{
		protodoc.ProtodocTypeJson.String(),
		protodoc.ProtodocTypeMD.String(),
	}

	// Check if the given command is in valid types
	if typeName == "" && !slices.Contains(validTypes, typeName) {
		fmt.Printf("Error: Invalid command type '%s'.\n", typeName)
		fmt.Println("Valid command types are:", strings.Join(validTypes, ", "))
		fmt.Println("Use --help or -h for usage information.")
		return
	}

	// Run protoc command
	runCommand(protodir, filename)

	// Initialize protodoc
	pbDoc, err := protodoc.New(
		filename,
		destFile,
		protodoc.WithName(name),
		protodoc.WithType(protodoc.ProtodocType(typeName)),
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
	if err := os.Remove(filename); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func runCommand(protoDir, descOut string) {
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
}

// Function to check required fields
func checkRequiredFields(fields map[string]string) error {
	var missingFields []string
	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			missingFields = append(missingFields, fieldName)
		}
	}
	if len(missingFields) > 0 {
		return fmt.Errorf("error: The following fields must not be empty: %s", strings.Join(missingFields, ", "))
	}
	return nil
}
