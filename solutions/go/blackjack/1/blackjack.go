package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0
	switch card {
    case "ace":
        value = 11
        break
    case "ten":
        value = 10
        break
    case "jack":
        value = 10
        break
    case "queen":
        value = 10
        break
    case "king":
        value = 10
        break
    case "nine":
        value = 9
        break
    case "eight":
        value = 8
        break
    case "seven":
        value = 7
        break
    case "six":
        value = 6
        break
    case "five":
        value = 5
        break
    case "four":
        value = 4
        break
    case "three":
        value = 3
        break
    case "two":
        value = 2
        break
    default:
        value = 0
        break
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    sum := ParseCard(card1) + ParseCard(card2)
	if ParseCard(card1) == 11 && ParseCard(card2) == 11 {
        return "P"
    } else if  sum == 21 && ParseCard(dealerCard) < 10{
        return "W"
    } else if  sum == 21 && ParseCard(dealerCard) >= 10{
        return "S"
    } else if sum >= 17 && sum <= 20 {
        return "S"
    } else if sum >= 12 && sum <= 16 && ParseCard(dealerCard) < 7 {
        return "S"
    } else if sum >= 12 && sum <= 16 && ParseCard(dealerCard) >= 7 {
        return "H"
    } 
    return "H"
 }
