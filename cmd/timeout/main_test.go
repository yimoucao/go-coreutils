package main

import "testing"

func Test_isAllDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name           string
		args           args
		want           uint
		expectAllDigit bool
	}{
		{"empty", args{" "}, 0, false},
		{"invalid", args{"abc"}, 0, false},
		{"digits", args{"123"}, 123, true},
		{"mized", args{"a123"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isAllUnsignedDigit(tt.args.s)
			if got != tt.want {
				t.Errorf("isAllDigit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.expectAllDigit {
				t.Errorf("isAllDigit() got1 = %v, want %v", got1, tt.expectAllDigit)
			}
		})
	}
}
