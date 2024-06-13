package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("reading the file content failed")
	}

	// file.Close()
	return lines, nil

}

func (fm FileManager) WriteResult(data any) error {

	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) // For example of use of Goroutine

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("failed to encode data to JSON")
	}

	// file.Close()
	return nil

}

func New(input, output string) FileManager {

	return FileManager{
		InputFilePath:  input,
		OutputFilePath: output,
	}

}
