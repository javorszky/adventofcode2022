package day6

import (
	"testing"
)

func Test_extractMarker(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example 1",
			args: args{
				in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "example 2",
			args: args{
				in: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "example 3",
			args: args{
				in: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "example 4",
			args: args{
				in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "example 5",
			args: args{
				in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want:    11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractMarker(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractMarker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractMarker() got = %v, want %v", got, tt.want)
			}
		})
	}
}
