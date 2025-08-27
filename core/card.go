package core

import (
	"encoding/json"
	"fmt"
)

type CardSuit uint8

const (
	Red CardSuit = iota
	Blue
	Green
	Yellow
	Wild
)

var cardSuits = [...]string{"red", "blue", "green", "yellow", "wild"}

type CardRank uint8

const (
	Zero CardRank = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Reverse
	Skip
	Any
	PlusTwo
	PlusFour
)

var cardRanks = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "reverse", "skip", "any", "+2", "+4"}

type Card struct {
	suit CardSuit
	rank CardRank
}

func (card Card) String() string {
	return fmt.Sprintf("(%v, %v)", cardSuits[card.suit], cardRanks[card.rank])
}

func NewCard(suit CardSuit, rank CardRank) Card {
	return Card{
		suit: suit,
		rank: rank,
	}
}

func (card Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Suit string `json:"suit"`
		Rank string `json:"rank"`
	}{
		Suit: cardSuits[card.suit],
		Rank: cardRanks[card.rank],
	})
}
