package circular_test

import (
	"testing"

	"github.com/eng618/go-eng/algo/circular"
)

func TestHello(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{name: "Temp Test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			circular.Hello()
		})
	}
}
