package file

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/JunNishimura/konmari/internal/extension"
)

func TestNewCleaner(t *testing.T) {
	type args struct {
		filePath string
	}
	type test struct {
		name string
		args args
		want *Cleaner
	}
	tests := []*test{
		func() *test {
			filePath := "main.go"
			cleaner := &Cleaner{
				filePath: filePath,
			}
			ext := extension.New(filePath)
			cleaner.extension = ext

			return &test{
				name: "success",
				args: args{
					filePath: filePath,
				},
				want: cleaner,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCleaner(tt.args.filePath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		content  string
		want     string
		wantErr  error
	}{
		{
			name:     "success: go",
			fileName: "test.go",
			content: `
package main
// import package
import "fmt"
/*
comment
*/
func main() {
	fmt.Println("Hello, Wolrd")// print out hello world
}`,
			want: `
package main

import "fmt"

func main() {
	fmt.Println("Hello, Wolrd")
}`,
			wantErr: nil,
		},
		{
			name:     "success: python",
			fileName: "test.py",
			content: `
# comment
def main():
	'''
	main function
	this is test
	'''
	print("hello") # print hello`,
			want: `

def main():
	
	print("hello") `,
			wantErr: nil,
		},
		{
			name:     "success: javascript",
			fileName: "test.js",
			content: `
// one-line comment
console.log("Hello, World")
/*
multiple 
line
comment
*/
console.log("hogehoge") // test`,
			want: `

console.log("Hello, World")

console.log("hogehoge") `,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			testPath := filepath.Join(tmpDir, tt.fileName)
			f, err := os.Create(testPath)
			if err != nil {
				t.Log(err)
			}

			_, err = f.WriteString(tt.content)
			if err != nil {
				t.Log(err)
			}
			f.Close()

			cleaner := NewCleaner(testPath)
			err = cleaner.Execute(false, defautlPostfix)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got = %v, want = %v", err, tt.wantErr)
			}

			cleanedFileName := addPostfixToFileName(tt.fileName, defautlPostfix)
			testCleanedPath := filepath.Join(tmpDir, cleanedFileName)
			b, err := os.ReadFile(testCleanedPath)
			if err != nil {
				t.Log(err)
			}
			if string(b) != tt.want {
				t.Errorf("got = %s, want = %s", string(b), tt.want)
			}
		})
	}
}
