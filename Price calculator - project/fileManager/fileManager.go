package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	// Lê uma linha por vez
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in file")
	}

	file.Close()
	return lines, err

}

func (fm FileManager) WriteResult(data interface{}) error {

	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	file.Close()
	return nil
}
