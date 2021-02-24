package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	fmt.Println("Initializing Process..")
}

func main() {

	fmt.Print("Digite o codigo do video: ")
	var input string
	fmt.Scanln(&input)

	output, err := exec.Command("./youtubedr", "download", input).Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Printing Process
	fmt.Println(string(output))

	defer fmt.Println("Process Finalized!!!")
}

// TODO: implementation move file to specific folder
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

// TODO: implementation recursively find files in folder
func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
