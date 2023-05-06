package circular_test

import (
	"testing"

	"garciaericn.com/go-eng/algo/circular"
)

func TestHello(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{name: "Temp Test"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			circular.Hello()
		})
	}
}
