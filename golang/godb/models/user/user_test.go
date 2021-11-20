package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRight(t *testing.T) {
	type args struct {
		atleastHas Rights
		has        Rights
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has right",
			args: args{
				atleastHas: Read,
				has:        AdminRead,
			},
			want: true,
		},
		{
			name: "has not right",
			args: args{
				atleastHas: Write,
				has:        Read,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckRight(tt.args.atleastHas, tt.args.has)
			assert.Equal(t, tt.want, got)
		})
	}
}
