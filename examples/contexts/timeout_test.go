package contexts_test

import (
	"testing"
	"time"

	"github.com/eng618/go-eng/examples/contexts"
)

func TestTimeoutContextBeforeTaskIsFinished(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{"One nanosecond", args{time.Nanosecond}},
		{"Two nanoseconds", args{2 * time.Nanosecond}},
		{"Three nanoseconds", args{3 * time.Nanosecond}},
		{"Fifty nanoseconds", args{50 * time.Nanosecond}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contexts.TimeoutContextBeforeTaskIsFinished(tt.args.duration)
		})
	}
}

func TestFinishTaskBeforeContextTimeout(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{"One nanosecond", args{time.Nanosecond}},
		{"Two nanoseconds", args{2 * time.Nanosecond}},
		{"Three nanoseconds", args{3 * time.Nanosecond}},
		{"Fifty nanoseconds", args{50 * time.Nanosecond}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contexts.FinishTaskBeforeContextTimeout(tt.args.duration)
		})
	}
}
