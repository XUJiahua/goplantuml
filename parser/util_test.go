package parser

import "testing"

func Test_getCommonPrefix(t *testing.T) {
	type args struct {
		dirs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				dirs: []string{
					"",
				},
			},
			want: "",
		},
		{
			args: args{
				dirs: []string{
					"",
					"/",
				},
			},
			want: "",
		},
		{
			args: args{
				dirs: []string{
					"/a/b/c",
					"/a/b",
				},
			},
			want: "/a/b",
		},
		{
			args: args{
				dirs: []string{
					"/a/b/c",
					"/a/b",
					"/a/b/c/d",
				},
			},
			want: "/a/b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonPrefix(tt.args.dirs); got != tt.want {
				t.Errorf("getCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastSecondDirectory(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				dir: "",
			},
			want: "/",
		},
		{
			args: args{
				dir: "/a",
			},
			want: "/",
		},
		{
			args: args{
				dir: "/a/b",
			},
			want: "/a/",
		},
		{
			args: args{
				dir: "/a/b/",
			},
			want: "/a/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastSecondDirectory(tt.args.dir); got != tt.want {
				t.Errorf("getLastSecondDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
