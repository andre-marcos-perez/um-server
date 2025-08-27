package core

import (
	"encoding/json"
	"testing"
)

func TestCardMarshalJSON(t *testing.T) {
	card := NewCard(Wild, PlusFour)
	t.Run("should marshal card", func(t *testing.T) {
		expected := "{\"suit\":\"wild\",\"rank\":\"+4\"}"
		if got, _ := json.Marshal(card); string(got) != expected {
			t.Errorf("Expected card to be %v, got %v", expected, got)
		}
	})
}
