package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func init() {
	fmt.Println("Initializing Process..")
}

func main() {

	output, err := exec.Command("./youtubedr", "download", "Sv6dMFF_yts").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Move file to videos folder

	// Printing Process
	fmt.Println(string(output))

	defer fmt.Println("Process Finalized!!!")
}

func MoveFile(sourcePath, destinationPath string) error {

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	outputFile, err := os.Create(destinationPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open destination file: %s", err)
	}

	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}
