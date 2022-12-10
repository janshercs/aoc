package utils

import (
	"bufio"
	"io"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	return os.Open(filename)
}

func ReadFile(f *os.File) ([]byte, error) {
	return io.ReadAll(f)
}

func ReadLine(f *os.File) *bufio.Scanner {
	return bufio.NewScanner(f)
}
