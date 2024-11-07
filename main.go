package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Get the current working directory
    startDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }

    // Display initial prompt
    fmt.Printf("This script will analyze the directories in '%s' and create a list of all files.\n", startDir)
    fmt.Printf("The list will be saved to: '%s/_filelisting/_filelisting.csv'.\n", startDir)
    fmt.Print("Do you want to proceed? (Y/n): ")

    // Read user input for confirmation
    reader := bufio.NewReader(os.Stdin)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)

    // Check user confirmation (case-insensitive)
    if userInput != "" && !(strings.EqualFold(userInput, "y") || strings.EqualFold(userInput, "yes")) {
        fmt.Println("Operation cancelled.")
        waitForExit()
        return
    }

    // Counters for statistics
    dirCount := 0
    fileCount := 0
    imageCount := 0

    // Define image file extensions
    imageExtensions := map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
        ".png":  true,
        ".tiff": true,
        ".bmp":  true,
        ".gif":  true,
    }

    // Create a new directory for storing the CSV file
    outputDir := filepath.Join(startDir, "_filelisting")
    err = os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        fmt.Println("Error creating output directory:", err)
        waitForExit()
        return
    }

    // Create the CSV file
    csvFilePath := filepath.Join(outputDir, "_filelisting.csv")
    csvFile, err := os.Create(csvFilePath)
    if err != nil {
        fmt.Println("Error creating CSV file:", err)
        waitForExit()
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

        // Count directories
        if info.IsDir() {
            dirCount++
            return nil
        }

        // Count files
        fileCount++

        // Get the relative path
        relativePath, err := filepath.Rel(startDir, path)
        if err != nil {
            return err
        }

        // Check the file extension and type
        ext := strings.ToLower(filepath.Ext(info.Name()))
        fileType := ""
        if imageExtensions[ext] {
            fileType = "image"
            imageCount++ // Count image files
        }

        // Write to CSV
        writer.Write([]string{info.Name(), relativePath, fileType})
        return nil
    })

    if err != nil {
        fmt.Println("Error walking the path:", err)
    } else {
        fmt.Println("\nFile listing CSV generated successfully!")
        fmt.Printf("\nSummary report:\n")
        fmt.Printf("Total directories analyzed: %d\n", dirCount)
        fmt.Printf("Total files found: %d\n", fileCount)
        fmt.Printf("Total image files identified: %d\n", imageCount)
        fmt.Printf("CSV file saved to: '%s'\n", csvFilePath)
    }

    // Wait for user input to exit
    waitForExit()
}

// waitForExit keeps the terminal open until the user chooses to exit
func waitForExit() {
    fmt.Println("\nPress Enter to exit or type 'exit' or press 'ESC' to close.")
    reader := bufio.NewReader(os.Stdin)
    for {
        userInput, _ := reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)
        if userInput == "" || strings.EqualFold(userInput, "exit") {
            break
        }
    }
}
