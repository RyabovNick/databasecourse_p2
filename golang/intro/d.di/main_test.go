package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Ivan")

	got := buf.String()
	want := "Hello, Ivan"

	assert.Equal(t, want, got)

	// tests := []struct {
	// 	name string
	// 	args string
	// }{
	// 	{
	// 		name: "n",
	// 		args: args{"Ivan"},
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		Greet(tt.args)
	// 	})
	// }
}
