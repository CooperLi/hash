package reader

import (
	"fmt"
	"testing"

	"github.com/CooperLi/hash/pkg/util"
	"gotest.tools/v3/fs"
)

func benchmarkCheckSumReaderWithBufferSize(b *testing.B, chunkSize, fileSize int) {
	b.ReportAllocs()
	filepath := fs.NewFile(b, b.Name())
	defer filepath.Remove()
	_ = util.CreateFileWithSize(filepath.Path(), fileSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = CheckSumReaderWithSize(filepath.Path(), chunkSize)
	}
}

func BenchmarkCheckSumReaderWithBufferSize(b *testing.B) {
	tests := []struct {
		chunkSize   int
		fileSize    int
		chunkSuffix string
		fileSuffix  string
	}{
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "1KB", chunkSize: 1 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "2KB", chunkSize: 2 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "4KB", chunkSize: 4 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "8KB", chunkSize: 8 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "16KB", chunkSize: 16 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "32KB", chunkSize: 32 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "64KB", chunkSize: 64 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "128KB", chunkSize: 128 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "256KB", chunkSize: 256 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "512KB", chunkSize: 512 * util.SizeKB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "1MB", chunkSize: 1 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "2MB", chunkSize: 2 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "4MB", chunkSize: 4 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "8MB", chunkSize: 8 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "16MB", chunkSize: 16 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "32MB", chunkSize: 32 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "64MB", chunkSize: 64 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "128MB", chunkSize: 128 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "256MB", chunkSize: 256 * util.SizeMB},
		{fileSuffix: "500MB", fileSize: 500 * util.SizeMB, chunkSuffix: "512MB", chunkSize: 512 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "1KB", chunkSize: 1 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "2KB", chunkSize: 2 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "4KB", chunkSize: 4 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "8KB", chunkSize: 8 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "16KB", chunkSize: 16 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "32KB", chunkSize: 32 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "64KB", chunkSize: 64 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "128KB", chunkSize: 128 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "256KB", chunkSize: 256 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "512KB", chunkSize: 512 * util.SizeKB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "1MB", chunkSize: 1 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "2MB", chunkSize: 2 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "4MB", chunkSize: 4 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "8MB", chunkSize: 8 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "16MB", chunkSize: 16 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "32MB", chunkSize: 32 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "64MB", chunkSize: 64 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "128MB", chunkSize: 128 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "256MB", chunkSize: 256 * util.SizeMB},
		{fileSuffix: "1GB", fileSize: 1 * util.SizeGB, chunkSuffix: "512MB", chunkSize: 512 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "1KB", chunkSize: 1 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "2KB", chunkSize: 2 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "4KB", chunkSize: 4 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "8KB", chunkSize: 8 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "16KB", chunkSize: 16 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "32KB", chunkSize: 32 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "64KB", chunkSize: 64 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "128KB", chunkSize: 128 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "256KB", chunkSize: 256 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "512KB", chunkSize: 512 * util.SizeKB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "1MB", chunkSize: 1 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "2MB", chunkSize: 2 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "4MB", chunkSize: 4 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "8MB", chunkSize: 8 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "16MB", chunkSize: 16 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "32MB", chunkSize: 32 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "64MB", chunkSize: 64 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "128MB", chunkSize: 128 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "256MB", chunkSize: 256 * util.SizeMB},
		{fileSuffix: "2GB", fileSize: 2 * util.SizeGB, chunkSuffix: "512MB", chunkSize: 512 * util.SizeMB},
	}

	for _, t := range tests {
		name := fmt.Sprintf("BenchmarkCheckSumReaderWithBufferSize_%s_%s", t.chunkSuffix, t.fileSuffix)
		b.Run(name, func(b *testing.B) {
			benchmarkCheckSumReaderWithBufferSize(b, t.chunkSize, t.fileSize)
		})
	}
}
