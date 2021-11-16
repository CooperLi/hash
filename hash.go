package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	algo "github.com/CooperLi/hash/pkg/algorithm"
	"github.com/kalafut/imohash"
)

func Hash(path string, algorithm string) (string, error) {
	switch algorithm {
	case SHA256Old:
		hash, err := algo.SHA256FileWithReadFile(path)
		if err != nil {
			return "", err
		}
		return hash, nil
	case SHA256New:
		hash, err := algo.SHA256FileWithBufioReader(path)
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(hash), nil
	case MD5:
		hash, err := algo.MD5File(path)
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(hash), nil
	case SHA256:
		hash, err := algo.SHA256FileWithBufioReader2(path)
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(hash), nil
	case IMOHash:
		hash, err := imohash.SumFile(path)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%x", hash), nil
	default:
		return "", errors.New("invalid algorithm")
	}
}
