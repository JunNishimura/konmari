package extension

import "testing"

func TestIsAcceptible(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				filename: "test.go",
			},
			want: true,
		},
		{
			name: "false: no extension",
			args: args{
				filename: "test",
			},
			want: false,
		},
		{
			name: "false: invalid extension",
			args: args{
				filename: "test.xxx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAcceptible(tt.args.filename)
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
