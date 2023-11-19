package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

/*
 ************** Start (Brian) Program Structs **************

 */

type Player struct {
	Name     string
	Hand     []CardMaker
	Score    int
	IsBusted bool
}

type CardMaker struct {
	FaceValue  string
	NumValue   int
	Suit       string
	IsFaceCard bool
	Color      rune
}

type prizeMaker struct {
	item    string
	cost    int
	itemNum int
}

type Shop []prizeMaker
type Deck []CardMaker

var wallet float64 = 127.00
var chipValue float64 = 10.00
var numChips int

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

/*



************** (Brian) Creating Card Deck Functions **************



 */

func newDeck() (deck Deck) {
	//assigned values for the cards, ace = 1, jack = 11 queen=12 king= 13
	numValues := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	//face value of cards
	faceValues := [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	//suits include Heart, Diamond, Club, Spade
	suits := [4]string{" Heart", "Diamond", "Club", " Spade"}

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
			card := CardMaker{
				FaceValue:  faceValues[i],
				NumValue:   numValues[i],
				Suit:       suits[j],
				IsFaceCard: faceCard,
				Color:      rune(cardColor),
			}
			deck = append(deck, card)
		}
	}
	return deck
}

/*





************** (Sam) Blackjack Game Functions **************




 */
// Shuffles deck of cards (Sam)
func shuffle(deck *Deck) {
	rand.Shuffle(len(*deck), func(i, j int) {
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	})
}

// Used to draw a card from the deck for play and remove the card from the deck of cards (Sam)
func drawCard(player *Player, deck *Deck) {
	card := (*deck)[0]
	player.Hand = append(player.Hand, card)
	*deck = (*deck)[1:]
}

// Calculate the player's score (Sam)
func calcScore(player *Player) {
	player.Score = 0
	aceCount := 0

	for _, cardMaker := range player.Hand {
		switch cardMaker.FaceValue {
		case "Jack", "Queen", "King":
			player.Score += 10
		case "Ace":
			player.Score += 11
			aceCount++
		default:
			player.Score += cardMaker.NumValue
		}
	}

	// Accounts for ace being able to be calculated as 1 or 11
	for player.Score > 21 && aceCount > 0 {
		player.Score -= 10
		aceCount--
	}
}

// Print current hand of cards (Sam)
func printHand(hand []CardMaker) string {
	var result string
	for _, cardMaker := range hand {
		result += fmt.Sprintf("\n   *%s of %ss", cardMaker.FaceValue, cardMaker.Suit)
	}
	return result
}

// Blackjack - Game play (Sam)
func blackJack(dealer, player *Player, deck *Deck) {

	fmt.Println("\n******************** START BLACKJACK GAME ********************")

	// Player and Dealer are dealt a card
	drawCard(player, deck)
	drawCard(dealer, deck)

	// Dealer's first card is revealed
	calcScore(dealer)
	fmt.Printf("\nDealer reveals their first card:%s ", printHand(dealer.Hand))
	fmt.Printf("\nDealer's point total: %d\n", dealer.Score)

	// Player and Dealer are dealt two more cards, total is calculated
	drawCard(player, deck)
	drawCard(dealer, deck)
	calcScore(player)
	calcScore(dealer)

	// Show player's hand and total
	fmt.Println("")
	fmt.Printf("Your Hand:%s ", printHand(player.Hand))
	fmt.Printf("\nYour point total: %d\n", player.Score)

	// Let player decide to hit or stand, scores and hands update and print
	var decision string
	for !player.IsBusted {
		fmt.Printf("Would you like to hit or stand? - Enter h for Hit, s for Stand\n")
		fmt.Scanln(&decision)

		if decision == "h" {
			drawCard(player, deck)
			calcScore(player)
			fmt.Printf("****** New Hand ******\n")
			fmt.Printf("Updated hand: %s ", printHand(player.Hand))
			fmt.Printf("\nCurrent Total: %d\n", player.Score)

			if player.Score > 21 {
				player.IsBusted = true
				fmt.Printf("Ooops, you busted! Dealer wins.\n")
				break
			}
		} else if decision == "s" {
			fmt.Printf("%s stands with a score of %d\n", player.Name, player.Score)
			break
		} else {
			fmt.Println("Invalid entry please enter h for Hit or s for Stand")
		}
	}

	// Dealer reveals second card
	fmt.Printf("\nDealer reveals second card, total hand: %s ", printHand(dealer.Hand))
	fmt.Printf("\nDealer's updated total: %d ", dealer.Score)

	// Dealer hits if dealer's hand is below 17 and player didn't bust
	for dealer.Score < 17 && player.IsBusted == false {
		drawCard(dealer, deck)
		calcScore(dealer)
		fmt.Println("\nDealer hits.")
		fmt.Printf("Dealer's updated hand: %s ", printHand(dealer.Hand))
		fmt.Printf("\nDealer's updated point total: %d\n", dealer.Score)

		// Dealer busts if hit goes over 21
		if dealer.Score > 21 {
			fmt.Printf("Dealer busts! You win!")
			dealer.IsBusted = true
			break
		}
	}

	// Dealer stands if score is between 17 and 21
	if dealer.Score >= 17 && dealer.Score <= 21 {
		fmt.Printf("\nDealer stands.\n")
	}

	fmt.Println("\n***** Hand is over *****")

	// Score results, if player and dealer doesn't bust:
	fmt.Printf("\nDealer Total: %d", dealer.Score)
	fmt.Printf("\nYour Total: %d\n", player.Score)

	if (player.Score > dealer.Score && player.IsBusted == false) || dealer.IsBusted == true {
		fmt.Printf("%s, you win!! Good job.\n", player.Name)
	} else if (player.Score < dealer.Score) || player.IsBusted == true {
		fmt.Printf("Dealer wins. Better luck next time %s.\n", player.Name)
	} else {
		fmt.Printf("It's a tie.\n")
	}

	fmt.Println("\n******************** END OF GAME ********************")
	fmt.Println("Return to menu? (Y/N)")
	var back string
	fmt.Scanln(&back)
	if back == "y" {
		directory(*player)
	} else if back == "n" {
		return
	}

}

// Blackjack Game functions (sam)
func playblackjack(x Player) {
	// Makes the order of cards random each time program starts
	rand.Seed(time.Now().UnixNano())

	// Create deck
	deck := newDeck()

	//Shuffle deck
	shuffle(&deck)

	// Create Dealer Player
	dealer := Player{Name: "Dealer"}

	// Begin Blackjack
	blackJack(&dealer, &x, &deck)
}

/*






************** (Brian) Prize Shop Functions **************

 */

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

// prototype for buying prize (brian)
func buyprize(x Player, s int) {

	var choice string
	//var itemChoice int
	fmt.Println("Would you like to purchase a prize? (y/n)")
	fmt.Scan(&choice)
	if choice == "y" {
		fmt.Println("Enter prize number for selection: ")
		fmt.Scan(&s)
	}
	if choice == "n" {
		println("Returning to Menu.")
		directory(x)
	}
}

// Function for shopping for prizes (brian)
func shopping(x Player) {
	var choice string
	var itemChoice int
	shop := prizeList()
	fmt.Println("\nThis is our list of prizes: ")
	for i := 1; i < len(shop); i++ {
		fmt.Println("[", i, "]:", shop[i].cost, "chips > ", shop[i].item)
	}

	fmt.Println("\nEnter prize number for selection: ")
	fmt.Scanln(&itemChoice)

	for num := range shop {
		if itemChoice == shop[num].itemNum && numChips > shop[num].cost {
			numChips = numChips - shop[num].cost
			fmt.Println("\nThank you for your purchase")
			fmt.Println("\nYou bought", shop[num].item)
			fmt.Println("\nYou have", numChips, "chips left")
			break
		}
	}

	fmt.Println("\nWould you like to purchase another prize? (y/n)")
	fmt.Scanln(&choice)
	if choice == "y" {
		shopping(x)
	} else if choice == "n" {

		directory(x)
	}
}

/*





************** (Jasmine) User Main Menu Functions **************

 */

func directory(x Player) {

	fmt.Println("******************* BLACKJACK MENU *******************")

	fmt.Printf("Welcome to blackjack , %s!\n", x.Name)
	fmt.Println("Select a number to begin an activity. \n\n[1]: Play Blackjack\n[2]: Go Shopping\n[3]: View Wallet\n[4]: Logout")

	var menunum int
	fmt.Scanln(&menunum)

	if menunum >= 1 || menunum <= 4 {
		switch {
		case menunum == 1:
			println("You Selected Blackjack.")
			playblackjack(x)
		case menunum == 2:
			println("You Selected Shopping.")
			shopping(x)
		case menunum == 3:
			println("You Selected Wallet.")
			viewwallet(x)
		case menunum == 4:
			println("Logging You Out...")
		}
	}
	if reflect.TypeOf(menunum).Kind() != reflect.Int {
		println("Input not accepted. Please try again.")
		fmt.Scanln(&menunum)
	}

}

func viewwallet(x Player) {

	//prototype for converting player cash to chips
	//var wallet float64 = 127.00
	//var chipValue float64 = 10.00
	//var numChips int
	for i := 0; chipValue-1 < wallet; i++ {
		wallet = wallet - chipValue
		numChips++
	}

	fmt.Println("\nChips Count: ", numChips, "valued at $", chipValue, "each")
	fmt.Println("Cash: $", wallet)
	fmt.Println("\nReturn to menu? (y/n)")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" {
		println("Returning to Menu.")
		directory(x)
	}
	if choice == "n" {
		viewwallet(x)
	}

}

/*
************** Main Method **************
 */
func main() {

	var player Player

	fmt.Println("What's your name?")
	fmt.Scanln(&player.Name)

	directory(player)

	//logout()
}
