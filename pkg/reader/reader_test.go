package reader

import (
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

func BenchmarkCheckSumReaderWithBufferSize_1KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_2KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_4KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_8KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_16KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 16*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_32KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_64KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_128KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_256KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_512KB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 512*util.SizeKB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_1MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_2MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_4MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_8MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_32MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_64MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_128MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_256MB_500MB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeMB, 500*util.SizeMB)
}

func BenchmarkCheckSumReaderWithBufferSize_1KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_2KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_4KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_8KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_16KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 16*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_32KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_64KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_128KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_256KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_512KB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 512*util.SizeKB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_1MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_2MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_4MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_8MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_32MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_64MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_128MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_256MB_1GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeMB, util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_1KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_2KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_4KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_8KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_16KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 16*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_32KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_64KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_128KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_256KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_512KB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 512*util.SizeKB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_1MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_2MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 2*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_4MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 4*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_8MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 8*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_32MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 32*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_64MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 64*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_128MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 128*util.SizeMB, 2*util.SizeGB)
}

func BenchmarkCheckSumReaderWithBufferSize_256MB_2GB(b *testing.B) {
	benchmarkCheckSumReaderWithBufferSize(b, 256*util.SizeMB, 2*util.SizeGB)
}
