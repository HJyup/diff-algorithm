package hash

import (
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateBlob(filePath, gitDir string) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	header := fmt.Sprintf("blob %d", len(content)) + "\x00"
	blob := append([]byte(header), content...)

	hash := sha1.Sum(blob)
	hashHex := fmt.Sprintf("%x", hash)

	objectDir := filepath.Join(gitDir, "objects", hashHex[:2])
	objectFile := filepath.Join(objectDir, hashHex[2:])

	if _, err := os.Stat(objectFile); err == nil {
		return hashHex, nil
	}

	if err := os.MkdirAll(objectDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create object directory: %v", err)
	}

	outFile, err := os.Create(objectFile)
	if err != nil {
		return "", fmt.Errorf("failed to create object file: %v", err)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)

	zlibWriter := zlib.NewWriter(outFile)
	defer func(zlibWriter *zlib.Writer) {
		_ = zlibWriter.Close()
	}(zlibWriter)

	if _, err := zlibWriter.Write(blob); err != nil {
		return "", fmt.Errorf("failed to write compressed blob: %v", err)
	}

	return hashHex, nil
}

func ReadAndPrintBlob(gitDir, hash string) error {
	objectDir := filepath.Join(gitDir, "objects", hash[:2])
	objectFile := filepath.Join(objectDir, hash[2:])

	file, err := os.Open(objectFile)
	if err != nil {
		return fmt.Errorf("failed to open blob file: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	zlibReader, err := zlib.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create zlib reader: %v", err)
	}
	defer func(zlibReader io.ReadCloser) {
		_ = zlibReader.Close()
	}(zlibReader)

	decompressedData, err := io.ReadAll(zlibReader)
	if err != nil {
		return fmt.Errorf("failed to read decompressed blob: %v", err)
	}

	fmt.Printf("Decompressed Blob Content:\n%s\n", decompressedData)
	return nil
}

func TestBlob(filePath string, rootDir string) {
	hash, err := CreateBlob(filePath, rootDir)
	if err != nil {
		fmt.Printf("Error creating blob: %v\n", err)
		return
	}

	fmt.Printf("Blob created with hash: %s\n", hash)

	if err := ReadAndPrintBlob(rootDir, hash); err != nil {
		fmt.Printf("Error reading blob: %v\n", err)
	}
}
