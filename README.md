# Simple Static File Server

This is a simple, configurable static file server written in Go. It serves static files from a specified directory with configurable port and base path settings via command-line flags.

## Features

- **Customizable port:** Choose which port your server listens on.
- **Customizable directory:** Serve static files from any directory.
- **Customizable base path:** Set a base URL path for your server.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need Go installed on your machine. To install Go, you can download it from the [official Go website](https://golang.org/dl/).

### Installing

First, clone the repository to your local machine:

```bash
git clone https://github.com/odinnordico/serve.git
cd serve
```

Then, build the program:

```bash
go build
```

### Usage

Run the server using the following command:

```bash
./serve -p="8100" -d="./path/to/your/static/files" -b="/basepath"
```

Here are the flags you can use to configure the server:

- `-p`: Port to listen on. Default is `"8100"`.
- `-d`: Directory of static files to serve. Default is the current directory `"."`.
- `-b`: Base URL path. Default is `"/"`.

### Example

To serve files from the `./static` directory on `localhost:8100` with a base path of `/files`, run:

```bash
./serve -d="./static" -p="8100" -b="/files"
```

Then, navigate to `http://localhost:8100/files` to access the served files.

## Development

This project uses Go Modules for managing dependencies. To add or update dependencies:

```bash
go mod tidy
```

## Contributing

We welcome contributions! Please open an issue if you have suggestions or open a Pull Request with improvements.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details.

