package day5

import (
	"reflect"
	"testing"
)

func Test_stack_PopMany(t *testing.T) {
	type fields struct {
		elements []string
	}
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       []string
		wantErr    bool
		wantString string
	}{
		{
			name:       "moves 1 element correctly",
			fields:     fields{elements: []string{"A", "B", "C"}},
			args:       args{n: 1},
			want:       []string{"C"},
			wantErr:    false,
			wantString: "[A B]",
		},
		{
			name:       "moves 4 element correctly",
			fields:     fields{elements: []string{"A", "B", "C", "D", "E", "F"}},
			args:       args{n: 4},
			want:       []string{"C", "D", "E", "F"},
			wantErr:    false,
			wantString: "[A B]",
		},
		{
			name:       "returns error for too many elements",
			fields:     fields{elements: []string{"A", "B", "C", "D", "E", "F"}},
			args:       args{n: 9},
			want:       nil,
			wantErr:    true,
			wantString: "[A B C D E F]",
		},
		{
			name:       "moves all element correctly",
			fields:     fields{elements: []string{"A", "B", "C", "D", "E", "F"}},
			args:       args{n: 6},
			want:       []string{"A", "B", "C", "D", "E", "F"},
			wantErr:    false,
			wantString: "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				elements: tt.fields.elements,
			}
			got, err := s.PopMany(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopMany() got = %v, want %v", got, tt.want)
			}

			gotS := s.String()
			if gotS != tt.wantString {
				t.Errorf("resulting string is not the same. Got %s, want %s", gotS, tt.wantString)
			}
		})
	}
}

func Test_stack_PushMany(t *testing.T) {
	type fields struct {
		elements []string
	}
	type args struct {
		in []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "pushes one to the end",
			fields: fields{elements: []string{"A", "B", "C", "D"}},
			args:   args{in: []string{"E"}},
			want:   "[A B C D E]",
		},
		{
			name:   "pushes several to the end",
			fields: fields{elements: []string{"A", "B", "C", "D"}},
			args:   args{in: []string{"E", "F", "G", "H"}},
			want:   "[A B C D E F G H]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				elements: tt.fields.elements,
			}
			s.PushMany(tt.args.in)
		})
	}
}
