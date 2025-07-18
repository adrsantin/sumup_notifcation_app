package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidNotificationType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid Email Notification", "email", true},
		{"Valid SMS Notification", "sms", true},
		{"Invalid Notification Type", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidNotificationType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
