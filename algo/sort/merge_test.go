package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		d Data
	}
	tests := []struct {
		name string
		args args
		want Data
	}{
		{
			name: "Base Example",
			args: args{d: Data{5, 4, 3, 2, 1}},
			want: Data{1, 2, 3, 4, 5},
		},
		{
			name: "Complex Example",
			args: args{Data{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}},
			want: Data{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50},
		},
		{
			name: "Uneven heavy",
			args: args{Data{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1, 2, 3}},
			want: Data{1, 2, 3, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		},
		{
			name: "Uneven light",
			args: args{Data{7, 8, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			want: Data{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 7, 8, 9}},
		{
			name: "Empty Data",
			args: args{Data{}},
			want: Data{}},
		{
			name: "Single item in Data",
			args: args{Data{5}},
			want: Data{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		l Data
		r Data
	}
	tests := []struct {
		name string
		args args
		want Data
	}{
		{name: "Basic1", args: args{l: Data{1}, r: Data{5}}, want: Data{1, 5}},
		{name: "Basic2", args: args{l: Data{5}, r: Data{1}}, want: Data{1, 5}},
		{name: "Basic3", args: args{l: Data{1, 2, 3, 4}, r: Data{5, 6, 7, 8, 9}}, want: Data{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Basic4", args: args{l: Data{1, 4, 7, 9}, r: Data{2, 3, 5, 6, 8}}, want: Data{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Empty right", args: args{l: Data{1, 2, 3, 4, 5}, r: Data{}}, want: Data{1, 2, 3, 4, 5}},
		{name: "Empty left", args: args{l: Data{}, r: Data{1, 2, 3, 4, 5}}, want: Data{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
