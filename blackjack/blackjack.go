package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0

    switch card {
    case "ace":
        value = 11
    case "two":
        value = 2
    case "three":
        value = 3
    case "four":
        value = 4
    case "five":
        value = 5
    case "six":
        value = 6
    case "seven":
        value = 7
    case "eight":
        value = 8
    case "nine":
        value = 9
    case "ten":
        value = 10
    case "jack":
        value = 10
    case "queen":
        value = 10
    case "king":
        value = 10
    default:
        value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Value := ParseCard(card1)
    card2Value := ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)

    totalValue := card1Value + card2Value

    var action string
    switch {
    case totalValue == 22:
        action = "P"
    case totalValue == 21 && dealerCardValue >= 10:
        action = "S"
    case totalValue == 21:
        action = "W"
    case totalValue >= 17 && totalValue <= 20:
        action = "S"
    case totalValue >= 12 && totalValue <= 16 && dealerCardValue < 7:
        action = "S"
    case totalValue >= 12 && totalValue <= 16 && dealerCardValue >= 7:
        action = "H"
    default:
        action = "H"
    }
    return action
}
