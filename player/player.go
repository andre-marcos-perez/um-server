package player

import "github.com/andre-marcos-perez/um-server/core"

type Player struct {
	Deck *core.Deck `json:"deck"`
}

func NewPlayer() *Player {
	return &Player{
		Deck: core.NewDeck(),
	}
}

func (player *Player) Draw(deck *core.Deck) error {
	card, err := deck.Draw()
	if err != nil {
		return err
	}
	player.Deck.Place(*card)
	return nil
}

func (player *Player) Discard() (*core.Card, error) {
	card, err := player.Deck.Draw()
	if err != nil {
		return nil, err
	}
	return card, nil
}
