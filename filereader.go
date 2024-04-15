package gosqltemplate

import (
	"bufio"
	"embed"
	"os"
)

type FileReader interface {
	ReadFile(path string) (content []string, err error)
}

type EmbeddedFileReader struct {
	FS embed.FS
}

func (e *EmbeddedFileReader) ReadFile(path string) ([]string, error) {
	file, err := e.FS.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type OSFileReader struct{}

func (o *OSFileReader) ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
