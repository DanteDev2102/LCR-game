package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("welcome to LCR dice game :)")

	g := new()

	fmt.Println("please enter how many players will play the game ?")
	fmt.Println("NOTE : enter number more than 2.")

	// need playes count and will take it as input
	var playersCount int
	for {
		fmt.Scanln(&playersCount)
		if playersCount < 3 {
			fmt.Println("note :enter number More than 2.")
			continue
		}
		break
	}

	for i := 0; i < playersCount; i++ {
		playerName := fmt.Sprintf("Player:%v", i)
		P := g.Join(playerName)
		fmt.Println(fmt.Sprintf("player:%v Joined.", P.name))
	}

	// turn
	turn := g.players[0]

	for {
		// check if player have tokens
		// if not skip him
		if turn.tokens == 0 {
			fmt.Println(fmt.Sprintf("player %v , you have 0 tokens", turn.name))
			turn = turn.right
			continue
		}

		// if yes ask him to hit enter to rolldice
		fmt.Println(fmt.Sprintf("player %v, you have %v tokens , hit any key to roll dice", turn.name, turn.tokens))

		var playerInput string
		fmt.Scanln(&playerInput)
		if playerInput == "EXIT" {
			fmt.Println("you killed Dice game :(")
			return
		}

		// roll dice & apply changes
		diceResult := turn.rollDice()
		fmt.Println("you got :", diceResult)

		for _, p := range g.players {
			fmt.Println(fmt.Sprintf("player %v, have  %v tokens", p.name, p.tokens))
		}

		// check if any one won the game
		// exit
		winner := g.finished()
		if winner != nil {
			fmt.Println("winner:", winner.name)
			return
		}
		// update turn
		turn = turn.right

	}

}

type game struct {
	players []*player
}

// Join add new player to the game
func (g *game) Join(playerName string) *player {
	// crete new player
	p := player{
		name:   playerName,
		tokens: 3,
	}

	playerCount := len(g.players)
	if playerCount > 0 {
		lastplayer := g.players[playerCount-1]
		p.left = lastplayer
		p.right = g.players[0]
		lastplayer.right = &p
		g.players[0].left = &p
	}
	g.players = append(g.players, &p)
	return &p
}

// finished check do we have only one player with Tokens
// this should be done after every turn
func (g *game) finished() (p *player) {
	playersWithtokens := 0
	for _, value := range g.players {
		if value.tokens > 0 {
			playersWithtokens++
			if playersWithtokens > 1 {
				return nil
			}
			p = value
		}
	}
	return
}

// new int game

func new() *game {
	return &game{}
}

// player
type player struct {
	name   string
	tokens int
	right  *player
	left   *player
}

// roll uses rand to return diceface
func (p *player) rollDice() (result []string) {
	// find out how many dices we should roll
	// if user have more than 3 tokens he can only roll 3 dices
	// or he roll exact number of tokens as dices

	dices := p.tokens
	if p.tokens > 2 {
		dices = 3
	}

	// roll the dices and update the value on the players
	for index := 0; index < dices; index++ {
		d := dice{}
		diceResult := d.roll()
		result = append(result, diceResult)

		switch diceResult {
		case "right":
			p.tokens--
			p.right.tokens++
		case "left":
			p.tokens--
			p.left.tokens++
		case "center":
			p.tokens--
		}
	}
	return
}

// dice
type dice struct{}

func (d *dice) roll() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(6)

	switch r {
	case 0:
		return "right"
	case 1:
		return "left"
	case 2:
		return "center"
	default:
		return "DoNothing"
	}
}
