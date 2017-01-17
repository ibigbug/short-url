package main

import (
	"math"
	"reflect"
	"testing"
)

func TestSimpleShorter_Short(t *testing.T) {
	type fields struct {
		idGen   IdGenerator
		storage Storage
	}
	type args struct {
		original string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantShort string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleShorter{
				idGen:   tt.fields.idGen,
				storage: tt.fields.storage,
			}
			if gotShort := s.Short(tt.args.original); gotShort != tt.wantShort {
				t.Errorf("SimpleShorter.Short() = %v, want %v", gotShort, tt.wantShort)
			}
		})
	}
}

func TestSimpleShorter_Original(t *testing.T) {
	type fields struct {
		idGen   IdGenerator
		storage Storage
	}
	type args struct {
		short string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOriginal string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleShorter{
				idGen:   tt.fields.idGen,
				storage: tt.fields.storage,
			}
			if gotOriginal := s.Original(tt.args.short); gotOriginal != tt.wantOriginal {
				t.Errorf("SimpleShorter.Original() = %v, want %v", gotOriginal, tt.wantOriginal)
			}
		})
	}
}

func TestSimpleIdGen_Gen(t *testing.T) {
	type fields struct {
		start int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"0", fields{0}, "0"},
		{"1", fields{1}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleIdGen{
				start: tt.fields.start,
			}
			if got := s.Gen(); got != tt.want {
				t.Errorf("SimpleIdGen.Gen() = %v, want %v", got, tt.want)
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

func TestNewShortService(t *testing.T) {
	type args struct {
		idGen   IdGenerator
		storage Storage
	}
	tests := []struct {
		name string
		args args
		want *SimpleShorter
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShortService(tt.args.idGen, tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortService() = %v, want %v", got, tt.want)
			}
		})
	}
}
