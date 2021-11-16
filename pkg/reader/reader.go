package reader

import (
	"bufio"
	"crypto/sha256"
	"hash"
	"io"
	"os"

	"github.com/pkg/errors"
)

const bufferSize = 64 * 1024

const bufferSize2 = 128 * 1024

func CheckSumReader(path string, hasher hash.Hash) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "hash file failed")
	}
	defer func() {
		if cErr := file.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}()

	reader := bufio.NewReader(file)
	buf := make([]byte, bufferSize)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hasher.Write(buf[:n])
		case io.EOF:
			return hasher.Sum(nil), nil
		default:
			return nil, err
		}
	}
}

func CheckSumReader2(path string, hasher hash.Hash) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cErr := file.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}()

	for buf, reader := make([]byte, bufferSize2), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		hasher.Write(buf[:n])
	}

	return hasher.Sum(nil), nil
}

func CheckSumReaderWithSize(path string, size int) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "hash file failed")
	}
	defer func() {
		if cErr := file.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}()

	hasher := sha256.New()
	reader := bufio.NewReader(file)
	buf := make([]byte, size)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hasher.Write(buf[:n])
		case io.EOF:
			return hasher.Sum(nil), nil
		default:
			return nil, err
		}
	}
}
