package main

import (
	"math"
	"testing"
)

func TestSimpleIdGen_Gen(t *testing.T) {
	type fields struct {
		start int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"0", fields{-1}, []string{"0", "1", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleIdGen{
				start: tt.fields.start,
			}
			for _, want := range tt.want {
				if got := s.Gen(); got != want {
					t.Errorf("SimpleIdGen.Gen() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestSimpleIdGen_base62(t *testing.T) {
	type fields struct {
		start int64
	}
	type args struct {
		i int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"zero", fields{0}, args{0}, "0"},
		{"one", fields{0}, args{1}, "1"},
		{"max", fields{0}, args{int64(math.Pow(62.0, 6.0)) - 1}, "ZZZZZZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleIdGen{
				start: tt.fields.start,
			}
			if got := s.base62(tt.args.i); got != tt.want {
				t.Errorf("SimpleIdGen.base62() = %v, want %v", got, tt.want)
			}
		})
	}
}
