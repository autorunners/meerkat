package utils

import (
	"testing"
)

func TestGetTimestamp(t *testing.T) {
	tests := []struct {
		name string
		min  int64
		max  int64
	}{
		{"demo", 0, 1735689600}, // 1970-01-01 00:00:00 UTC 到2025-01-01 00:00:00 UTC
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimestamp()
			if got < tt.min || got > tt.max {
				t.Error("wrong timestamp")
			}
		})
	}
}

func TestGetTimestampMs(t *testing.T) {
	tests := []struct {
		name string
		min  int64
		max  int64
	}{
		{"demo", 0, 1735689600 * 1000}, // 1970-01-01 00:00:00 UTC 到2025-01-01 00:00:00 UTC
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimestampMs()
			if got < tt.min || got > tt.max {
				t.Error("error timestamp")
			}
		})
	}
}
