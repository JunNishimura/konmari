package file

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestExtractFileName(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success: no slash",
			args: args{
				path: "test.txt",
			},
			want: "test.txt",
		},
		{
			name: "success: with slash",
			args: args{
				path: "dir/test.txt",
			},
			want: "test.txt",
		},
		{
			name: "success: back slash",
			args: args{
				path: `\dir\test.txt`,
			},
			want: "test.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractFileName(tt.args.path)
			if got != tt.want {
				t.Errorf("got = %s, want = %s", got, tt.want)
			}
		})
	}
}

func TestExtractFilesFromDirectory(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr error
	}{
		{
			name:    "success",
			want:    []string{"dir/dir2/test.txt", "dir/test.txt"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			dirPath := filepath.Join(tmpDir, "dir")
			if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
				t.Logf("%v: %s", err, dirPath)
			}
			dir2Path := filepath.Join(dirPath, "dir2")
			if err := os.Mkdir(dir2Path, os.ModePerm); err != nil {
				t.Logf("%v: %s", err, dir2Path)
			}
			textPath := filepath.Join(dirPath, "test.txt")
			f, err := os.Create(textPath)
			if err != nil {
				t.Logf("%v: %s", err, textPath)
			}
			f.Close()
			textPath = filepath.Join(dir2Path, "test.txt")
			f, err = os.Create(textPath)
			if err != nil {
				t.Logf("%v: %s", err, textPath)
			}
			f.Close()

			got, err := ExtractFilesFromDirectory(dirPath)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got = %v, want = %v", err, tt.wantErr)
			}

			wantFiles := make([]string, 0)
			for _, path := range tt.want {
				newPath := strings.ReplaceAll(filepath.Join(tmpDir, path), `\`, "/")
				wantFiles = append(wantFiles, newPath)
			}

			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			sort.Slice(wantFiles, func(i, j int) bool { return wantFiles[i] < wantFiles[j] })

			if !reflect.DeepEqual(got, wantFiles) {
				t.Errorf("got = %v, want = %v", got, wantFiles)
			}
		})
	}
}
