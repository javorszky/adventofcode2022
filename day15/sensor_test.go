package day15

import (
	"reflect"
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

func Test_sensor_rowBoundCoordinates(t *testing.T) {
	s := newSensor(coordinate{10, 10}, coordinate{12, 12})

	type args struct {
		row int
	}
	tests := []struct {
		name    string
		row     int
		want    coordinate
		want1   coordinate
		wantErr bool
	}{
		{
			name:    "row goes through sensor",
			row:     10,
			want:    coordinate{6, 10},
			want1:   coordinate{14, 10},
			wantErr: false,
		},
		{
			name:    "row is top tip",
			row:     6,
			want:    coordinate{10, 6},
			want1:   coordinate{10, 6},
			wantErr: false,
		},
		{
			name:    "row is somewhere in the middle",
			row:     8,
			want:    coordinate{8, 8},
			want1:   coordinate{12, 8},
			wantErr: false,
		},
		{
			name:    "row does not intersect",
			row:     0,
			want:    coordinate{},
			want1:   coordinate{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := s.rowBoundCoordinates(tt.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("rowBoundCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rowBoundCoordinates() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("rowBoundCoordinates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
