package rules_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/TVolly/goapi-addresses/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCommunityUpdateRule_Load_InvalidBody(t *testing.T) {
	rule := &rules.CommunityUpdateRule{}

	assert.Equal(t, rules.ErrNoBody, rule.Load(nil))

	badReader := strings.NewReader("test")
	assert.Error(t, rule.Load(badReader))
}

func TestCommunityUpdateRule_Load_InvalidPayload(t *testing.T) {
	rule := &rules.CommunityUpdateRule{}
	testCases := []struct {
		name    string
		payload interface{}
	}{
		{
			name:    "invalid payload",
			payload: nil,
		},
		{
			name: "invalid name",
			payload: map[string]string{
				"name": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)

			assert.Error(t, rule.Load(b))
		})
	}
}

func TestCommunityUpdateRule_Load_ValidPayload(t *testing.T) {
	rule := &rules.CommunityUpdateRule{}

	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(map[string]string{
		"name": "Test name",
	})

	assert.NoError(t, rule.Load(b))
}
