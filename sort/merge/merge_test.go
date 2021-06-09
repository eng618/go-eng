package merge

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		d Data
	}
	tests := []struct {
		name string
		args args
		want Data
	}{
		{"Base Example", args{d: Data{5, 4, 3, 2, 1}}, Data{1, 2, 3, 4, 5}},
		{
			"Complex Example",
			args{Data{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}},
			Data{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50},
		},
		{name: "Empty Data", args: args{Data{}}, want: Data{}},
		{name: "Single item in Data", args: args{Data{5}}, want: Data{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
