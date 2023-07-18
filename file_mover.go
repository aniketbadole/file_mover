package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Define the source and destination folders
var sourceFolder = "/home/aniket/Downloads"
var destinationFolder = map[string]string{
	"pdf":   "/home/aniket/Downloads/documents",
	"png":   "/home/aniket/images",
	"jpg":   "/home/aniket/images",
	"jpeg":  "/home/aniket/images",
	"tiff":  "/home/aniket/images",
	"gif":   "/home/aniket/images",
	"mp4":   "/home/aniket/videos",
	// Add more file types and their corresponding destination folders if needed
}

// Path to the log file
var logFile = "/home/aniket/Downloads/download_log.txt"

// MoveFile moves the file to the appropriate destination folder
func moveFile(sourceFile, destinationFolder string) {
	// Create the destination folder if it doesn't exist
	if err := os.MkdirAll(destinationFolder, 0755); err != nil {
		log.Printf("Error creating destination folder: %v\n", err)
		return
	}

	// Extract the file name
	fileName := filepath.Base(sourceFile)

	// Generate the destination file path
	destinationFile := filepath.Join(destinationFolder, fileName)

	// Move the file
	if err := os.Rename(sourceFile, destinationFile); err != nil {
		log.Printf("Error moving file: %v\n", err)
		return
	}

	// Get file information
	fileInfo, err := os.Stat(destinationFile)
	if err != nil {
		log.Printf("Error getting file information: %v\n", err)
		return
	}

	// Update the log file
	logEntry := fmt.Sprintf(
		"File: %s\nType: %s\nSource: %s\nDestination: %s\nSize: %d bytes\n=====================\n",
		fileName,
		strings.ToLower(filepath.Ext(fileName)[1:]),
		sourceFile,
		destinationFolder,
		fileInfo.Size(),
	)
	if err := appendToFile(logFile, logEntry); err != nil {
		log.Printf("Error updating log file: %v\n", err)
	}
}

// appendToFile appends content to a file
func appendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}

// Monitor the Downloads folder for changes
func monitorFolder() {
	for {
		err := filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Error accessing path %q: %v\n", path, err)
				return nil
			}
			if !info.IsDir() {
				extension := strings.ToLower(filepath.Ext(path)[1:])
				destination, ok := destinationFolder[extension]
				if ok {
					moveFile(path, destination)
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("Error walking the path: %v\n", err)
		}

		// Delay between scans
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Start monitoring the Downloads folder
	monitorFolder()
}
