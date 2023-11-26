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

	cash       float64
	chipValue  float64
	numChips   int
	prizes     []string
	totalgames int
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

func yourWallet(player *Player) {

	// starter money
	player.cash = 100.00
	player.chipValue = 10.00
	player.numChips = 5

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
Blackjack Gambling mechanics Tristan
*/
func bet(x *Player) int {
	var wager int
	fmt.Println("     * Your Chips: ", x.numChips)
	fmt.Println("     * How much would you like to bet?")
	fmt.Scanln(&wager)
	if wager <= x.numChips && wager >= 0 {
		fmt.Println("     * You bet ", wager, "Chips.")
		return wager
	} else {
		fmt.Println("Not enough chips in your wallet")
		bet(x)
	}
	return 0
}
func betResult(wager int, win int, DorN bool) int {
	if DorN == true {
		wager *= 2
	}
	switch win {

	case -1:
		wager *= -1
	case 0:
		wager = 0
	}
	return wager

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
	var wager int

	var DorN = false
	fmt.Println("\n******************** START BLACKJACK GAME ********************")

	wager = bet(player)
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
		fmt.Printf("Would you like to hit or stand? - ( [h]it / [s]tand/ [d]ouble down )\n")
		fmt.Scanln(&decision)
		if decision == "d" {
			DorN = true
			decision = "h"
		}
		if decision == "h" {
			drawCard(player, deck)
			calcScore(player)
			fmt.Printf("\n******** You Hit. ********\n")
			fmt.Printf("Updated hand: %s ", printHand(player.Hand))
			fmt.Printf("\nCurrent Total: %d\n", player.Score)

			if player.Score >= 21 {
				player.IsBusted = true

				wager = betResult(wager, -1, DorN)
				fmt.Printf("Oops, you busted! Dealer wins. Better luck next time %s.\n", player.Name)
				player.numChips += wager
				fmt.Println("You Lost: ", wager, "chips.")
				fmt.Println("\n******************** END OF GAME ********************")
				backmenu(player, 1)
				break

			}
			break
		} else if decision == "s" {
			fmt.Printf("%s stands with a score of %d\n", player.Name, player.Score)
			break
		} else {
			fmt.Println("Invalid entry please enter h for Hit or s for Stand")
		}
	}

	// Dealer reveals second card
	fmt.Printf("\nDealer reveals next card, total hand: %s ", printHand(dealer.Hand))
	fmt.Printf("\nDealer's total: %d ", dealer.Score)

	// Dealer hits if dealer's hand is below 17 and player didn't bust
	for dealer.Score < 17 && player.IsBusted == false {
		drawCard(dealer, deck)
		calcScore(dealer)
		fmt.Println("\n****** Dealer hits. ******\n")
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
		fmt.Printf("\n ***** Dealer stands. *****\n")
	}

	// Score results, if player and dealer doesn't bust:
	fmt.Printf("\nDealer Total: %d", dealer.Score)
	fmt.Printf("\nYour Total: %d\n", player.Score)

	if (player.Score > dealer.Score && player.IsBusted == false) || dealer.IsBusted == true {

		wager = betResult(wager, 1, DorN)
		fmt.Printf("%s, you win!! Good job.\n", player.Name)

		player.numChips += player.Score + wager

		fmt.Println("You Won: ", player.Score+wager, "chips.")

	} else if (player.Score < dealer.Score) || player.IsBusted == true {

		wager = betResult(wager, -1, DorN)
		fmt.Printf("Dealer wins. Better luck next time %s.\n", player.Name)
		player.numChips += wager
		fmt.Println("You Lost: ", wager, "chips.")
	} else {

		fmt.Printf("It's a tie.\n")
	}

	fmt.Println("\n******************** END OF GAME ********************")
	backmenu(player, 1)
}

// Blackjack Game functions (sam)
func playblackjack(player *Player) {
	// Adds to the game count
	player.totalgames += 1

	// Makes the order of cards random each time program starts
	rand.Seed(time.Now().UnixNano())

	// Clears former hands
	player.Hand = nil

	// Create deck
	deck := newDeck()

	//Shuffle deck
	shuffle(&deck)

	// Create Dealer Player
	dealer := Player{Name: "Dealer"}

	// Begin Blackjack
	blackJack(&dealer, &*player, &deck)
}

/*






************** (Brian) Prize Shop Functions **************

 */

func prizeList() (shop Shop) {
	prizes := [10]string{"Starbucks GiftCard", "Gas Coupon", "Scratch-off lottery tickets", "Bottles of liquor", "Movie Tickets", "Spay day tickets", "Football Tickets",
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

// Function for shopping for prizes (brian)
func shopping(player *Player) {
	var choice string
	var itemChoice int
	var numChips = player.numChips

	shop := prizeList()
	fmt.Println("******************* PRIZE SHOP *******************")

	fmt.Println("\nThis is our list of prizes: ")
	for i := 0; i < len(shop); i++ {
		fmt.Println("[", i+1, "]:", shop[i].cost, "chips > ", shop[i].item)
	}

	fmt.Println("\n     *TOTAL CHIPS: ", numChips)
	fmt.Println("\nEnter a prize number for selection, or type [0] to cancel. ")
	fmt.Scanln(&itemChoice)

	for num := range shop {
		if itemChoice-1 == shop[num].itemNum && numChips > shop[num].cost {
			player.numChips = numChips - shop[num].cost
			fmt.Println("\n******** Your Purchase ********")
			fmt.Println("\n     *You bought", shop[num].item)
			player.prizes = append(player.prizes, shop[num].item)
			fmt.Println("     *You have", player.numChips, "chips left.")
			break
		}
		if itemChoice == 0 {
			fmt.Println("Purchase Cancelled.")
			backmenu(player, 2)
			break
		}

	}

	fmt.Println("\nWould you like to purchase another prize? (y/n)")
	fmt.Scanln(&choice)
	if choice == "y" {
		shopping(player)
	} else if choice == "n" {
		fmt.Println("\n***** Thank you for Shopping! *****")
		backmenu(player, 2)
	}
}

/*





************** (Jasmine) User Main Menu Functions **************

 */

func directory(player *Player) {

	fmt.Println("******************* BLACKJACK MENU *******************")

	fmt.Printf("\n  *Welcome to blackjack , %s!\n", player.Name)
	fmt.Println("  *Select a number to begin an activity. \n\n          [1]: Play Blackjack\n          [2]: Go Shopping\n          [3]: View Wallet\n          [4]: Logout")
	fmt.Println("\n******************************************************")
	var menunum int
	fmt.Scanln(&menunum)

	if menunum >= 1 || menunum <= 4 {
		switch {
		case menunum == 1:
			println("You Selected Blackjack.")
			playblackjack(player)
		case menunum == 2:
			println("You Selected Shopping.")
			shopping(player)
		case menunum == 3:
			println("You Selected Wallet.")
			viewwallet(player)
		case menunum == 4:
			logout(player)
			println("Logging You Out...")
		}
	}
	if reflect.TypeOf(menunum).Kind() != reflect.Int {
		println("Input not accepted. Please try again.")
		fmt.Scanln(&menunum)
	}

}

func viewwallet(player *Player) {
	var chipValue = player.chipValue
	var wallet = player.cash
	var numChips = player.numChips

	// for i := 0; chipValue-1 < wallet; i++ {
	// 	wallet = wallet - chipValue
	// 	numChips++
	// }

	fmt.Println("*********************** WALLET ***********************")
	fmt.Println("\n     *Chips: ", numChips)
	fmt.Println("     *Cash: $", wallet)
	fmt.Println("     *Cash Out Rate: $", chipValue)
	fmt.Println("\n******************************************************")
	fmt.Println("\nWhich would you like to convert?")
	fmt.Println("[1]: Convert Chips to Cash")
	fmt.Println("[2]: Convert Cash to Chips")
	fmt.Println("[x]: Cancel")
	var convert string
	fmt.Scanln(&convert)
	if convert == "1" {
		fmt.Println("******* Converting Chips... *******\n")

		player.cash += float64(numChips) * chipValue
		player.numChips = 0

		fmt.Println("     *You Cashed Out for $", player.cash, ".\n")
		fmt.Println("***********************************\n")
	}
	if convert == "2" {
		fmt.Println("******* Converting Cash... *******\n")

		if player.cash >= player.chipValue {
			var buychips int = int(player.cash / chipValue)
			player.numChips += buychips
			player.cash = 0

			fmt.Println("     *You Bought", buychips, "chips .\n")
		} else {
			fmt.Println("Not enough cash. Please convert\n  cash to chips or refresh the game.\n")
		}

		fmt.Println("***********************************\n")
	}
	if convert == "x" {
		backmenu(player, 3)
	}

	backmenu(player, 3)

}

func backmenu(player *Player, taskint int) {

	fmt.Println("Return to menu? (y/n)")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" {
		println("Returning to Menu.\n\n\n\n")
		directory(player)
	}
	if choice == "n" {
		if taskint >= 1 || taskint <= 4 {
			switch {
			case taskint == 1:
				playblackjack(player)
			case taskint == 2:
				shopping(player)
			case taskint == 3:
				viewwallet(player)
			}
		}
	}
}

// logout menu
func logout(player *Player) {
	fmt.Println("********************* GAME RECORD ********************")
	fmt.Println("               ", player.Name, "'s Game Record.")
	fmt.Println("GAMES PLAYED:")
	fmt.Println("      [*] TOTAL GAMES:", player.totalgames)
	fmt.Println("\n\nPRIZES BOUGHT:")
	for i := 0; i < len(player.prizes); i++ {
		fmt.Printf("      [*] %s\n", player.prizes[i])
	}
	fmt.Println("\n\nWALLET: ")
	fmt.Println("      [*] TOTAL CHIPS:", player.numChips)
	fmt.Println("      [*] TOTAL CASH: $", player.cash)
	fmt.Println("\nThank you for playing, please come again!\n")
	fmt.Println("*****************************************************")
	return
}

/*
************** Main Method **************
 */
func main() {

	var player Player

	fmt.Println("What's your name?")
	fmt.Scanln(&player.Name)
	yourWallet(&player)
	directory(&player)

}
