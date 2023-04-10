package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

type RatReader interface {
	ReadRat() (*big.Rat, error)
}

type FileRatReader struct {
	filename string
	file     *os.File
	scanner  *bufio.Scanner
}

// Ця функція призначена для читання раціональних чисел з файлу і повернення їх у вигляді масиву []*big.Rat (раціональних чисел).
func ReadRats(filename string) ([]*big.Rat, error) {
	reader, err := NewFileRatReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	rats := []*big.Rat{}
	for {
		rat, err := reader.ReadRat()
		if err == io.EOF {
			break
		}
		if err == nil {
			rats = append(rats, rat)
		} else {
			return nil, err
		}
	}
	return rats, nil
}

func NewFileRatReader(filename string) (*FileRatReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return &FileRatReader{filename, file, scanner}, nil
}

func (r *FileRatReader) ReadRat() (*big.Rat, error) {
	if !r.scanner.Scan() {
		if err := r.scanner.Err(); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}
	line := strings.TrimSpace(r.scanner.Text())
	parts := strings.Split(line, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format: %s", line)
	}
	num := new(big.Int)
	den := new(big.Int)
	if _, ok := num.SetString(parts[0], 10); !ok {
		return nil, fmt.Errorf("invalid numerator: %s", parts[0])
	}
	if _, ok := den.SetString(parts[1], 10); !ok {
		return nil, fmt.Errorf("invalid denominator: %s", parts[1])
	}
	rat := new(big.Rat)
	rat.SetFrac(num, den)
	return rat, nil
}

func (r *FileRatReader) Close() error {
	if err := r.file.Close(); err != nil {
		return err
	}
	return nil
}
