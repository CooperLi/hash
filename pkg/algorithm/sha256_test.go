package algorithm

import (
	"testing"

	"github.com/CooperLi/hash/pkg/util"
	"gotest.tools/v3/fs"
)

func benchmarkSHA256FileWithReadFile(b *testing.B, fileSize int) {
	b.ReportAllocs()
	filepath := fs.NewFile(b, b.Name())
	defer filepath.Remove()
	_ = util.CreateFileWithSize(filepath.Path(), fileSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = SHA256FileWithReadFile(filepath.Path())
	}
}

func benchmarkSHA256FileWithBufioReader(b *testing.B, fileSize int) {
	b.ReportAllocs()
	filepath := fs.NewFile(b, b.Name())
	defer filepath.Remove()
	_ = util.CreateFileWithSize(filepath.Path(), fileSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = SHA256FileWithBufioReader(filepath.Path())
	}
}

func benchmarkSHA256FileWithBufioReader2(b *testing.B, fileSize int) {
	b.ReportAllocs()
	filepath := fs.NewFile(b, b.Name())
	defer filepath.Remove()
	_ = util.CreateFileWithSize(filepath.Path(), fileSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = SHA256FileWithBufioReader2(filepath.Path())
	}
}

func BenchmarkSHA256FileWithReadFile_100MB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 100*util.SizeMB)
}

func BenchmarkSHA256FileWithReadFile_200MB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 200*util.SizeMB)
}

func BenchmarkSHA256FileWithReadFile_500MB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 500*util.SizeMB)
}

func BenchmarkSHA256FileWithReadFile_1GB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, util.SizeGB)
}

func BenchmarkSHA256FileWithReadFile_2GB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 2*util.SizeGB)
}

func BenchmarkSHA256FileWithReadFile_4GB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 4*util.SizeGB)
}

func BenchmarkSHA256FileWithReadFile_8GB(b *testing.B) {
	benchmarkSHA256FileWithReadFile(b, 8*util.SizeGB)
}

func BenchmarkSHA256FileWithBufioReader_500MB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader(b, 500*util.SizeMB)
}

func BenchmarkSHA256FileWithBufioReader_1GB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader(b, util.SizeGB)
}

func BenchmarkSHA256FileWithBufioReader_2GB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader(b, 2*util.SizeGB)
}

func BenchmarkSHA256FileWithBufioReader2_500MB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader2(b, 500*util.SizeMB)
}

func BenchmarkSHA256FileWithBufioReader2_1GB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader2(b, util.SizeGB)
}

func BenchmarkSHA256FileWithBufioReader2_2GB(b *testing.B) {
	benchmarkSHA256FileWithBufioReader2(b, 2*util.SizeGB)
}
