package repository

import (
	"fmt"
	"io"
	"os"
)

// var filename = "test.txt"

// CreateFile ...
func CreateFile(filename string) string {
	// check if file exists
	filename = filename + ".txt"
	var _, err = os.Stat(filename)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if isError(err) {
			return "Error while creating file: " + err.Error()
		}
		defer file.Close()
	}

	return "File " + filename + " created successfully!!"
}

// WriteFile ...
func WriteFile(filename, text string) string {
	filename = filename + ".txt"
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if isError(err) {
		return err.Error()
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(text)
	if isError(err) {
		return err.Error()
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return err.Error()
	}

	return "File " + filename + " updated successfully."
}

// ReadFile ...
func ReadFile(filename string) {
	// Open file for reading.
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

// DeleteFile ...
func DeleteFile(filename string) {
	// delete file
	var err = os.Remove(filename)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

// isError ...
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
