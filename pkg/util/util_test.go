package util

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateFileWithSize(t *testing.T) {
	type args struct {
		path string
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TC-generate 500MB file",
			args: args{
				path: fmt.Sprintf("/home/%s_500MB", t.Name()),
				size: 500 * SizeMB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFileWithSize(tt.args.path, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("createFileWithSize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				file, _ := os.Stat(tt.args.path)
				if file.Size() != int64(tt.args.size) {
					t.Errorf("createFileWithSize() size = %v, actually %v", tt.args.size, file.Size())
				}
			}
			os.Remove(tt.args.path)
		})
	}
}
