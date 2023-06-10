package engfile_test

import (
	"testing"

	engfile "github.com/eng618/go-eng/lib/file"
)

func TestToFile(t *testing.T) {
	type args struct {
		fName string
		input interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "simple", args: args{fName: "simple", input: "This is a simple file"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := engfile.ToFile(tt.args.fName, tt.args.input); got != tt.want {
				t.Errorf("ToFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
