package toofastwg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []int
	}{
		{
			name: "test",
			args: "127.0.0.1",
			want: []int{135, 445, 623, 2179, 5800, 5900, 5354, 5357, 5040, 7680, 8884},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Scan(tt.args)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
