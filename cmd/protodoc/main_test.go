package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {

	testCases := []struct {
		name     string
		typeName string
		filePath string
	}{
		{
			name:     "JSON_TYPE",
			typeName: "json",
			filePath: "./api-documentation.json",
		},
		{
			name:     "MARKDOWN_TYPE",
			typeName: "markdown",
			filePath: "./api-documentation.md",
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			cmd := exec.Command("go", "run", ".", "--proto_dir=../../examples", "--doc_opt=source_relative", "--type="+v.typeName)

			err := cmd.Run()
			require.NoError(tt, err)
			require.FileExists(tt, v.filePath)

			err = os.RemoveAll(v.filePath)
			require.NoError(tt, err)
		})
	}

}
