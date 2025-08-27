package engine

import (
	"errors"
	"github.com/andre-marcos-perez/um-server/player"
	"testing"
)

func TestGameNew(t *testing.T) {
	t.Parallel()
	t.Run("should create a new game", func(t *testing.T) {
		players := []player.Player{
			*player.NewPlayer(),
			*player.NewPlayer(),
		}
		game, err := NewGame(players)
		if err != nil {
			t.Errorf("error creating game: %v", err)
		}
		expected := 0
		if got := game.DiscardDeck.Len(); got != expected {
			t.Errorf("Discard deck len: got %d, want %d", got, expected)
		}
		expected = GameInitDrawDeckSize - len(players)*GameInitPlayerDeckSize
		if got := game.DrawDeck.Len(); got != expected {
			t.Errorf("Draw deck len: got %d, want %d", got, expected)
		}
		expected = GameInitPlayerDeckSize
		for _, player := range players {
			if got := player.Deck.Len(); got != expected {
				t.Errorf("Player deck len: got %d, want %d", got, expected)
			}
		}
	})
	t.Run("should fail to create a game with 1 player", func(t *testing.T) {
		players := []player.Player{
			*player.NewPlayer(),
		}
		_, err := NewGame(players)
		if err != nil {
			expected := ErrGamePlayerInvalidAmount
			if !errors.Is(err, expected) {
				t.Errorf("expecting error %v, got %v", expected, err)
			}
		}
	})
	t.Run("should fail to create a game with 11 player", func(t *testing.T) {
		players := []player.Player{
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
			*player.NewPlayer(),
		}
		_, err := NewGame(players)
		if err != nil {
			expected := ErrGamePlayerInvalidAmount
			if !errors.Is(err, expected) {
				t.Errorf("expecting error %v, got %v", expected, err)
			}
		}
	})
}
