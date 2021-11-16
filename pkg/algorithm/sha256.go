package algorithm

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CooperLi/hash/pkg/reader"
	"github.com/pkg/errors"
)

func SHA256Dir(path string) (string, error) {
	var checkSum string
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		cleanPath := filepath.Clean(path)
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if !info.IsDir() {
			f, err := ioutil.ReadFile(cleanPath) // nolint:gosec
			if err != nil {
				return err
			}
			fileHash := fmt.Sprintf("%x", sha256.Sum256(f))
			checkSum = fmt.Sprintf("%s%s", checkSum, fileHash)
		}
		return nil
	}); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(checkSum))), nil
}

func SHA256FileWithReadFile(path string) (string, error) {
	cleanPath := filepath.Clean(path)
	if f, err := os.Stat(cleanPath); err != nil {
		return "", errors.Errorf("failed to stat file %q", cleanPath)
	} else if f.IsDir() {
		return "", errors.New("failed to hash directory")
	}

	file, err := ioutil.ReadFile(cleanPath) // nolint:gosec
	if err != nil {
		return "", errors.Wrapf(err, "hash file failed")
	}

	return fmt.Sprintf("%x", sha256.Sum256(file)), nil
}

func SHA256FileWithBufioReader(path string) ([]byte, error) {
	cleanPath := filepath.Clean(path)
	if f, err := os.Stat(cleanPath); err != nil {
		return nil, errors.Errorf("failed to stat file %q", cleanPath)
	} else if f.IsDir() {
		return nil, errors.New("failed to hash directory")
	}

	return reader.CheckSumReader(path, sha256.New())
}

func SHA256FileWithBufioReader2(path string) ([]byte, error) {
	return reader.CheckSumReader2(path, sha256.New())
}
