package day15

import (
	"testing"
)

func Test_sensor_rowInExclusion(t *testing.T) {
	s := newSensor(coordinate{10, 10}, coordinate{12, 12})

	type args struct {
		row int
	}
	tests := []struct {
		name string
		row  int
		want bool
	}{
		{
			name: "row goes through sensor, in exclusion",
			row:  10,
			want: true,
		},
		{
			name: "row is at top tip, in exclusion",
			row:  6,
			want: true,
		},
		{
			name: "row is just above top tip, not in exclusion",
			row:  5,
			want: false,
		},
		{
			name: "row is bottom tip, in exclusion",
			row:  14,
			want: true,
		},
		{
			name: "row is just below bottom tip, not in exclusion",
			row:  15,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.rowInExclusion(tt.row); got != tt.want {
				t.Errorf("rowInExclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
