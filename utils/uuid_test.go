package utils

import (
	"log"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"demo1", ""},
		{"demo2", ""},
		{"demo3", ""},
		{"demo4", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUUID()
			log.Println(got)

		})
	}
}
