package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := shuffleCards(generateCards())
	dealerHand, playerHand := generateHands(deck)
	outcome := playGame(dealerHand, playerHand, deck)
	fmt.Printf("Game: Winner is %v \n", outcome.winner)
}

func shuffleCards(cards []Card) *[]Card {
	dest := make([]Card, 52)
	rand.Seed(time.Now().Unix())
	perm := rand.Perm(52)
	fmt.Println("permm", perm)
	for i, v := range perm {
		dest[v] = cards[i]
	}
	return &dest
}

func playGame(dealerHand, playerHand Hand, deck *[]Card) gameState {
	playerHand = autoplay(playerHand, deck)
	dealerHand = autoplay(dealerHand, deck)
	fmt.Println("finish player hand", playerHand)
	fmt.Println("finish dealer hand", dealerHand)
	var r gameState
	if playerHand.busted == true && dealerHand.busted == true {
		r = gameState{true, "Tie"}
	} else if playerHand.busted == true {
		r = gameState{true, "dealer"}
	} else if dealerHand.busted == true {
		r = gameState{true, "player"}
	} else if dealerHand.totalValue > playerHand.totalValue {
		r = gameState{true, "dealer"}
	} else if playerHand.totalValue > dealerHand.totalValue {
		r = gameState{true, "player"}
	} else if playerHand.totalValue == dealerHand.totalValue {
		r = gameState{true, "tie"}
	} else {
		r = gameState{false, "error"}
	}
	return r
}

func autoplay(hand Hand, deck *[]Card) Hand {
	for hand.totalValue < 17 {
		var drawnCard Card
		drawnCard, deck = drawCard(*deck)
		hand.cards = append(hand.cards, drawnCard)
		hand.totalValue += drawnCard.value
		if hand.totalValue > 21 {
			hand.busted = true
		}
	}
	return hand
}

type gameState struct {
	isOver bool
	winner string
}

// Card is
type Card struct {
	text  string
	value int
	suit  string
}

// Hand is
type Hand struct {
	cards      []Card
	totalValue int
	busted     bool
}

func generateHands(deck *[]Card) (dealer Hand, player Hand) {
	dealer, deck = generateHand(deck)
	player, deck = generateHand(deck)
	return
}

func generateHand(deck *[]Card) (Hand, *[]Card) {
	pickedCards := make([]Card, 0, 2)
	var drawnCard Card
	drawnCard, deck = drawCard(*deck)
	pickedCards = append(pickedCards, drawnCard)
	drawnCard, deck = drawCard(*deck)
	pickedCards = append(pickedCards, drawnCard)

	totalValue := 0
	for _, v := range pickedCards {
		totalValue += v.value
	}
	return Hand{pickedCards, totalValue, false}, deck
}

func drawCard(deck []Card) (Card, *[]Card) {
	// Shift, draw one card
	cardDrawn, deck := deck[0], deck[1:]
	// Put it to the back
	deck = append(deck, cardDrawn)
	return cardDrawn, &deck
}

func generateCards() []Card {
	cards := make([]Card, 0, 52)
	suits := []string{"Clubs", "Hearts", "Diamonds", "Spades"}
	numbers := generateValues()
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(suits); j++ {
			cardtext := fmt.Sprint(numbers[i], "of", suits[j])
			card := Card{
				suits[j],
				numbers[i],
				cardtext,
			}
			cards = append(cards, card)
		}
	}

	fmt.Println("Card length", len(cards))
	return cards
}

func generateValues() []int {
	numbers := make([]int, 13)

	for i := 0; i < 13; i++ {
		if i+1 > 10 {
			numbers[i] = 10
		} else {
			numbers[i] = i + 1
		}
	}

	return numbers
}
