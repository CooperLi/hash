package util

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"runtime"
)

const (
	SizeKB = 1024
	SizeMB = 1024 * SizeKB
	SizeGB = 1024 * SizeMB
)

func TraceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func CreateFileWithSize(path string, size int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.CopyN(file, rand.Reader, int64(size))
	return err
}
