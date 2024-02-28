package validatorhelper

import "testing"

func TestIsValidUrl(t *testing.T) {
	type args struct {
		urlInput string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid url",
			args: args{
				urlInput: "https://google.com",
			},
			want: true,
		},
		{
			name: "Url with space",
			args: args{
				urlInput: "https://google .com",
			},
			want: false,
		},
		{
			name: "Url with invalid character",
			args: args{
				urlInput: "https://google|.com",
			},
			want: false,
		},
		{
			name: "Url without protocol",
			args: args{
				urlInput: "google.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUrl(tt.args.urlInput); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
