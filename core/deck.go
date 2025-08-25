package core

import (
	"errors"
	"github.com/andre-marcos-perez/um-server/sequence"
)

var (
	ErrDeckEmpty = errors.New("empty deck")
)

type Deck struct {
	cards sequence.LinkedList[Card]
}

func NewDeck(options ...func(*Deck)) *Deck {
	deck := &Deck{
		cards: *sequence.NewLinkedList[Card](),
	}
	for _, option := range options {
		option(deck)
	}
	return deck
}

func WithCards(cards []Card) func(*Deck) {
	return func(deck *Deck) {
		for _, card := range cards {
			deck.cards.Insert(card)
		}
	}
}

func WithShuffle() func(*Deck) {
	return func(deck *Deck) {
		deck.cards.Sort()
	}
}

func (deck *Deck) Len() int {
	return deck.cards.Len()
}

func (deck *Deck) Draw() (*Card, error) {
	card, err := deck.cards.Delete()
	if errors.Is(err, sequence.ErrSequenceEmpty) {
		return nil, ErrDeckEmpty
	}
	return card, err
}
