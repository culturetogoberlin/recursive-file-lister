# File Listing Script

## Overview
This Go project provides a script that recursively scans all directories and subdirectories starting from its execution location and generates a CSV file containing information about all files found. It identifies images based on common extensions and saves the output in a structured format.

### Key Features
- Recursively analyzes directories and subdirectories.
- Generates a CSV file listing file names, relative paths, and image type identifiers.
- Provides a summary report after the scan, showing:
  - Total number of directories analyzed.
  - Total number of files found.
  - Total number of image files identified.
- Waits for user confirmation before execution.
- Keeps the terminal open after completion with an option to press Enter to exit or press `ESC`/type `exit` to close.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (Ensure Go is installed and `GOPATH` is set up).

See [Golang Documentation](https://go.dev/doc/install).

### Steps

1. Clone or download this repository.
2. Navigate to the project folder:
3. Build executable

```bash
go build -o filelisting.exe main.go
```

## Usage

* Copy filelisting.exe to the directory where you want to run the analysis. 
* Doubleclick on filelisting.exe.
* Confirm the execution when prompted.
* View the report upon completion.

## Output

The output CSV file _filelisting.csv is saved in the _filelisting folder within the directory from which the script was run.
The CSV file format:

```bash
filename, relative path, type
example.jpg, subdir/example.jpg, image
```
## Customization

You can modify the imageExtensions map in main.go to include or exclude certain file types based on your requirements:

```bash
imageExtensions := map[string]bool{
    ".jpg":  true,
    ".jpeg": true,
    ".png":  true,
    ".tiff": true,
    ".bmp":  true,
    ".gif":  true,
}
```
Rebuild executable after customization:

```bash
go build -o filelisting.exe main.go
```

## License

This project is open-source and available under the MIT License.

## Author

Developed by Michael Müller, Culture to go GbR.