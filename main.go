package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Define image file extensions
    imageExtensions := map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
        ".png":  true,
        ".tiff": true,
        ".bmp":  true,
        ".gif":  true,
    }

    // Get the current working directory
    startDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }

    // Create a new directory for storing the CSV file
    outputDir := filepath.Join(startDir, "_filelisting")
    err = os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        fmt.Println("Error creating output directory:", err)
        return
    }

    // Create the CSV file
    csvFilePath := filepath.Join(outputDir, "_filelisting.csv")
    csvFile, err := os.Create(csvFilePath)
    if err != nil {
        fmt.Println("Error creating CSV file:", err)
        return
    }
    defer csvFile.Close()

    // Create a CSV writer
    writer := csv.NewWriter(csvFile)
    defer writer.Flush()

    // Write the header row
    writer.Write([]string{"filename", "relative path", "type"})

    // Walk the directory tree
    err = filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Get the relative path
        relativePath, err := filepath.Rel(startDir, path)
        if err != nil {
            return err
        }

        // Check the file extension
        ext := strings.ToLower(filepath.Ext(info.Name()))
        fileType := ""
        if imageExtensions[ext] {
            fileType = "image"
        }

        // Write to CSV if it's an image or any file
        writer.Write([]string{info.Name(), relativePath, fileType})
        return nil
    })

    if err != nil {
        fmt.Println("Error walking the path:", err)
    } else {
        fmt.Printf("File listing CSV generated at: %s\n", csvFilePath)
    }
}
