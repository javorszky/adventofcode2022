package day6

import (
	"testing"
)

func Test_extractMarker(t *testing.T) {
	type args struct {
		in string
		n  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example 1, start of packet",
			args: args{
				in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				n:  startOfPacketLength,
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "example 2, start of packet",
			args: args{
				in: "bvwbjplbgvbhsrlpgdmjqwftvncz",
				n:  startOfPacketLength,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "example 3, start of packet",
			args: args{
				in: "nppdvjthqldpwncqszvftbrmjlhg",
				n:  startOfPacketLength,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "example 4, start of packet",
			args: args{
				in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
				n:  startOfPacketLength,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "example 5, start of packet",
			args: args{
				in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
				n:  startOfPacketLength,
			},
			want:    11,
			wantErr: false,
		},

		{
			name: "example 1, start of message",
			args: args{
				in: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				n:  messageMarkerLength,
			},
			want:    19,
			wantErr: false,
		},
		{
			name: "example 2, start of message",
			args: args{
				in: "bvwbjplbgvbhsrlpgdmjqwftvncz",
				n:  messageMarkerLength,
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "example 3, start of message",
			args: args{
				in: "nppdvjthqldpwncqszvftbrmjlhg",
				n:  messageMarkerLength,
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "example 4, start of message",
			args: args{
				in: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
				n:  messageMarkerLength,
			},
			want:    29,
			wantErr: false,
		},
		{
			name: "example 5, start of message",
			args: args{
				in: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
				n:  messageMarkerLength,
			},
			want:    26,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractMarker(tt.args.in, tt.args.n)
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
