package main

import (
	"fmt"
)

func newDeck() (deck Deck) {
	//assigned values for the cards, jack = 10 queen=10 king= 10 ace=11
	numValues := [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

	//face value of cards
	faceValues := [13]string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King", "Ace"}

	//suits include Heart, Diamond, Club, Spade
	suits := [4]string{"Heart", "Diamond", "Club", "Spade"}

	var cardColor rune
	var faceCard bool

	for i := 0; i < len(faceValues); i++ {
		for j := 0; j < len(suits); j++ {
			if faceValues[i] == "Jack" || faceValues[i] == "Queen" || faceValues[i] == "King" || faceValues[i] == "Ace" {
				faceCard = true
			} else {
				faceCard = false
			}
			//Switches card color based on suite. 82 = 'R'(Red)   66= 'B' (Black)
			if suits[j] == "Heart" || suits[j] == "Diamond" {
				cardColor = 'R'
			} else {
				cardColor = 'B'
			}
			//Brian-Converted rune cardColor to string for output
			card := cardMaker{
				faceValue:  faceValues[i],
				numValue:   numValues[i],
				suit:       suits[j],
				isFaceCard: faceCard,
				color:      rune(cardColor),
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func prizeList() (shop Shop) {
	prizes := [10]string{"Starbucks GiftCard", "Gas card", "Scratch-off lottery tickets", "Bottles of liquor", "Movie Tickets", "Spay day tickets", "Football Tickets",
		"Dinner for 2", "Smart speakers", "iPod touch"}

	//Based on num of chips
	prizeCost := [10]int{1, 2, 4, 5, 10, 11, 22, 23, 46, 47}

	for i := 0; i < len(prizes); i++ {
		prize := prizeMaker{
			item:    prizes[i],
			cost:    prizeCost[i],
			itemNum: i,
		}
		shop = append(shop, prize)
	}
	return shop
}

type cardMaker struct {
	faceValue  string
	numValue   int
	suit       string
	isFaceCard bool
	color      rune
}

type prizeMaker struct {
	item    string
	cost    int
	itemNum int
}


var wallet float64 = 127.00
var chipValue float64 = 10.00
var numChips int

func login() {
	fmt.Println("\nWelcome to the Blackjack table")
	fmt.Println()
}

func logout() {
	//prototype for converting player chips to cash
	fmt.Println("\nRemaining", numChips, "chips have been converted to cash")
	for i := 0; numChips != 0; i++ {
		wallet = wallet + chipValue
		numChips--
	}

	//fmt.Println("\nRemaining", numChips, "chips have been converted to cash")
	fmt.Println("You have", "$", wallet, "in your wallet")

	fmt.Println("\nThank you for playing, please come again")
	fmt.Println()
	return
}

// Function for shopping for prizes
func shopping() {
	var choice string
	var itemChoice int
	shop := prizeList()
	fmt.Println("\nThis is our list of prizes: ")
	for i := 1; i < len(shop); i++ {
		fmt.Println("\nPrize #", i, ":", shop[i].item, " which cost ", shop[i].cost, "chips")
	}

	fmt.Println("\nWhich prize would you like to purchase: ")
	fmt.Println("\nEnter prize number for selection: ")
	fmt.Scan(&itemChoice)

	for num := range shop {
		if itemChoice == shop[num].itemNum && numChips > shop[num].cost {
			numChips = numChips - shop[num].cost
			fmt.Println("\nThank you for your purchase")
			fmt.Println("\nYou bought", shop[num].item)
			fmt.Println("\nYou have", numChips, "chips left")
			break
		}
	}
	fmt.Println("\nWould you like to purchase another prize?")
	fmt.Println("\nEnter y for Yes and n for No: ")
	fmt.Scan(&choice)
	if choice == "y" {
		shopping()
	} else if choice == "n" {
		return
	}
}

type Shop []prizeMaker
type Deck []cardMaker

func main() {

	login()

	//Uncomment out if you want to see the deck print
	//deck := newDeck()
	//fmt.Println(deck)

	//prototype for converting player cash to chips
	//var wallet float64 = 127.00
	//var chipValue float64 = 10.00
	//var numChips int
	for i := 0; chipValue-1 < wallet; i++ {
		wallet = wallet - chipValue
		numChips++
	}

	fmt.Println("You have", numChips, "chips worth", "$", chipValue, "each")
	fmt.Println("You have", "$", wallet, "in your wallet")

	//Prototype for buying prizes
	var choice string
	//var itemChoice int
	fmt.Println("\nWould you like to purchase a prize?")
	fmt.Println("\nEnter y for Yes and n for No: ")
	fmt.Scan(&choice)
	if choice == "y" {
		shopping()
	}

	logout()
}
