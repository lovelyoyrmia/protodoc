# Protobuf API Documentation Generator CLI

This command-line interface (CLI) tool provides a simple way to generate API documentation in various formats (JSON, Markdown, YAML, HTML) from Protocol Buffer (Protobuf) files. It streamlines the documentation process for APIs defined using Protobuf, making it easier to understand and utilize your services.

## Features

- Generate API documentation in JSON format.
- Generate API documentation in Markdown format.
- Generate API documentation in YAML format.
- Generate API documentation in HTML format.
- Support for Protobuf message and service definitions.
- Custom annotations for additional metadata (like paths and methods).

## Installation

To install the Protobuf API Documentation Generator CLI, you can use the following command:

```bash
go install github.com/lovelyoyrmia/protodoc/cmd/protodoc@latest
```

## Usage

### Command-Line Interface

Once installed, you can use the `protodoc` command to generate documentation from your Protobuf files.

#### Basic Command

To generate documentation, run the following command:

```bash
protodoc --proto_dir=./proto --doc_opt=source_relative --type=json --doc_out=.
```

#### Options

- `--proto_dir`: Path to the Protobuf directory (can accept multiple files).
- `--doc_opt`: Documentation options for the generated documentation (`source_relative`).
- `--doc_out`: Output file path for the generated documentation.
- `--type`: Desired output format (`json`, `markdown`, `yaml`, `html`). Default type is `markdown`.
- `--template_path`: Path to the custom template file (need to has `.tmpl` extension).

### Example Commands

1. **Generate JSON Documentation**

   ```bash
   protodoc --proto_dir=./proto --doc_opt=source_relative --type=json --doc_out=.
   ```

2. **Generate Markdown Documentation**

   ```bash
   protodoc --proto_dir=./proto --doc_opt=source_relative --doc_out=.
   ```

3. **Generate YAML Documentation**

   ```bash
   protodoc --proto_dir=./proto --doc_opt=source_relative --type=yaml --doc_out=.
   ```
4. **Generate HTML Documentation**

   ```bash
   protodoc --proto_dir=./proto --doc_opt=source_relative --type=html --doc_out=.
   ```
5. **Generate Documentation with Custom Template**

   ```bash
   protodoc --proto_dir=./proto --doc_opt=source_relative --type=html --doc_out=. --template_path=../custom_template.tmpl
   ```

## API Options

The CLI supports custom API Options in your Protobuf files for additional metadata. Here’s how you can use them:

```protobuf
syntax = "proto3";

import "github.com/lovelyoyrmia/protodoc/options/options.proto";

package yourpackage;

// Service definition with annotations
service YourService {
    // Annotated RPC method
    rpc YourMethod (YourRequest) returns (YourResponse) {
        option (api_options) = {
            summary: ""
            path: "/myapi/mymethod"
            method: "POST"
            query_params: {
                name: "myid",
                type: "int",
                description: "mydescription",
                required: true,
            }
        };
    }
}

// Message definitions
message YourRequest {
    string id = 1; // Unique identifier
}

message YourResponse {
    string result = 1; // Result of the operation
}
```

## Types

There are many data types used in the generated documentation. The table below summarizes the different types and their meanings.

| Type                          | Description                                | Example                          |
|-------------------------------|--------------------------------------------|----------------------------------|
| **Common Types**               | Basic types like string, integer, etc.     | `string`, `int`, etc.            |
| **Optional Types**             | Common types that are optional             | `string*`, `int*`, etc.          |
| **Message Types**              | Custom Protobuf message types              | `#User`, `#Customer`             |
| **Message Types (Optional)**   | Message types that are optional            | `#User*`                         |
| **List Types**                 | Lists of common or message types           | `string[]`, `int[]`, `#User[]`   |

## Examples

For more examples you can see the [examples](./examples/) folders.

## Contribution

Contributions are welcome! If you have suggestions for improvements or new features, feel free to submit a pull request or open an issue. Please ensure to follow the [Code of Conduct](CODE_OF_CONDUCT.md) and the [Contributing Guidelines](CONTRIBUTING.md).

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Conclusion

The Protobuf API Documentation Generator CLI simplifies the process of generating API documentation from Protobuf files. With support for multiple formats and customizable annotations, you can ensure your API is well-documented and easily accessible.

For more information, check out the [GitHub repository](https://github.com/lovelyoyrmia/protodoc) or reach out if you have any questions!