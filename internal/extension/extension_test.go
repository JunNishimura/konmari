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
			name: "true: ruby",
			args: args{
				filename: "test.rb",
			},
			want: Ruby,
		},
		{
			name: "true: php",
			args: args{
				filename: "test.php",
			},
			want: PHP,
		},
		{
			name: "true: java",
			args: args{
				filename: "test.java",
			},
			want: Java,
		},
		{
			name: "true: c",
			args: args{
				filename: "test.c",
			},
			want: C,
		},
		{
			name: "true: cpp",
			args: args{
				filename: "test.cpp",
			},
			want: Cpp,
		},
		{
			name: "true: rs",
			args: args{
				filename: "test.rs",
			},
			want: Rust,
		},
		{
			name: "true: swift",
			args: args{
				filename: "test.swift",
			},
			want: Swift,
		},
		{
			name: "true: kotlin",
			args: args{
				filename: "test.kt",
			},
			want: Kotlin,
		},
		{
			name: "true: dart",
			args: args{
				filename: "test.dart",
			},
			want: Dart,
		},
		{
			name: "true: html",
			args: args{
				filename: "test.html",
			},
			want: HTML,
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
