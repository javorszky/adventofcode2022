package day5

import (
	"testing"
)

func Test_stack_Push(t *testing.T) {
	type fields struct {
		elements []string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantPush    string
		wantString  string
		wantPop     string
		wantString2 string
	}{
		{
			name:        "does push pop string correctly",
			fields:      fields{elements: []string{"A", "B", "C"}},
			args:        args{in: "D"},
			wantPush:    "D",
			wantString:  "[A B C D]",
			wantPop:     "D",
			wantString2: "[A B C]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				elements: tt.fields.elements,
			}
			if got := s.Push(tt.args.in); got != tt.wantPush {
				t.Errorf("Push() = %v, want %v", got, tt.wantPush)
			}

			if got := s.String(); got != tt.wantString {
				t.Errorf("String() after push = %v, want %v", got, tt.wantString)
			}

			if got := s.Pop(); got != tt.wantPop {
				t.Errorf("Pop() = %v, want %v", got, tt.wantPop)
			}

			if got := s.String(); got != tt.wantString2 {
				t.Errorf("String() after pop = %v, want %v", got, tt.wantString2)
			}
		})
	}
}
