package player

import "testing"

func TestPlayerNew(t *testing.T) {
	player := NewPlayer()
	expected := 0
	if got := player.Deck.Len(); got != expected {
		t.Errorf("player.deck.Len() = %v, want %v", got, expected)
	}
}
