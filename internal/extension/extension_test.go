package extension

import "testing"

func TestNew(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want Mark
	}{
		{
			name: "true: go",
			args: args{
				filename: "test.go",
			},
			want: Go,
		},
		{
			name: "true: python",
			args: args{
				filename: "test.py",
			},
			want: Python,
		},
		{
			name: "true: javascript",
			args: args{
				filename: "test.js",
			},
			want: JavaScript,
		},
		{
			name: "true: typescript",
			args: args{
				filename: "test.ts",
			},
			want: TypeScript,
		},
		{
			name: "false: no extension",
			args: args{
				filename: "test",
			},
			want: Undefined,
		},
		{
			name: "false: invalid extension",
			args: args{
				filename: "test.xxx",
			},
			want: Undefined,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.filename)
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
