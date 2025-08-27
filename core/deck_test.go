package core

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestDeckNew(t *testing.T) {
	t.Parallel()
	cards := []Card{
		NewCard(Red, Zero),
		NewCard(Red, One),
		NewCard(Red, Two),
	}
	t.Run("should create a new empty deck", func(t *testing.T) {
		deck := NewDeck()
		expected := 0
		if got := deck.Len(); got != expected {
			t.Errorf("Expected deck length to be %d, got %d", expected, got)
		}
	})
	t.Run("should create a new deck with cards", func(t *testing.T) {
		deck := NewDeck(
			WithCards(cards),
		)
		expected := 3
		if got := deck.Len(); got != expected {
			t.Errorf("Expected deck length to be %d, got %d", expected, got)
		}
	})
	t.Run("should create a new deck with cards randomly shuffled", func(t *testing.T) {
		deck := NewDeck(
			WithCards(cards),
			WithShuffle(),
		)
		expected := 3
		if got := deck.Len(); got != expected {
			t.Errorf("Expected deck length to be %d, got %d", expected, got)
		}
	})
}

func TestDeckDraw(t *testing.T) {
	cards := []Card{
		NewCard(Red, Zero),
		NewCard(Red, One),
		NewCard(Red, Two),
	}
	deck := NewDeck(
		WithCards(cards),
	)
	t.Run("should draw cards", func(t *testing.T) {
		expected := cards[2]
		if card, _ := deck.Draw(); *card != expected {
			t.Errorf("Expected deck cards to be %v, got %v", expected, card)
		}
		expected = cards[1]
		if card, _ := deck.Draw(); *card != expected {
			t.Errorf("Expected deck cards to be %v, got %v", expected, card)
		}
		expected = cards[0]
		if card, _ := deck.Draw(); *card != expected {
			t.Errorf("Expected deck cards to be %v, got %v", expected, card)
		}
		if _, err := deck.Draw(); !errors.Is(err, ErrDeckEmpty) {
			t.Errorf("Expected error to be %v, got %v", ErrDeckEmpty, err)
		}
	})
}

func TestDeckPlace(t *testing.T) {
	cards := []Card{
		NewCard(Red, Zero),
		NewCard(Red, One),
		NewCard(Red, Two),
	}
	deck := NewDeck(
		WithCards(cards),
	)
	t.Run("should place cards", func(t *testing.T) {
		expected := NewCard(Green, Nine)
		deck.Place(expected)
		if card, _ := deck.Draw(); *card != expected {
			t.Errorf("Expected deck cards to be %v, got %v", expected, card)
		}
	})
}

func TestDeckMarshalJSON(t *testing.T) {
	cards := []Card{
		NewCard(Red, Zero),
		NewCard(Wild, PlusFour),
	}
	deck := NewDeck(
		WithCards(cards),
	)
	t.Run("should marshal deck", func(t *testing.T) {
		expected := "[{\"suit\":\"wild\",\"rank\":\"+4\"},{\"suit\":\"red\",\"rank\":\"0\"}]"
		if got, _ := json.Marshal(deck); string(got) != expected {
			t.Errorf("Expected deck to be %v, got %v", expected, got)
		}
	})
}
