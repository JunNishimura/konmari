package file

import (
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
