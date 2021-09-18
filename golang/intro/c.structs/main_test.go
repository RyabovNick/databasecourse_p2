package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rec)
	want := 40.0

	assert.Equal(t, want, got)

	// if got != want {
	// 	t.Errorf("got %.2f want %.2f", got, want)
	// }
}

func TestAreav1(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{12.0, 6.0}
		got := rec.Area()
		want := 72.0

		assert.Equal(t, want, got)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		assert.Equal(t, want, got)
	})
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name string
		s    Shape
		want float64
	}{
		{
			name: "reactangle: first",
			s: Rectangle{
				Width:  5,
				Height: 10,
			},
			want: 50,
		},
		{
			name: "reactangle: second",
			s: Rectangle{
				Width:  15,
				Height: 1,
			},
			want: 15,
		},
		{
			name: "circle: first",
			s: Circle{
				Radius: 10,
			},
			want: 314.1592653589793,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Area()
			assert.Equal(t, tt.want, got)
		})
	}
}
