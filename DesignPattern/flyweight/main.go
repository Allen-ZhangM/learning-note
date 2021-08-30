package main

import "fmt"

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		if dressType == CounterTerroristDressType {
			dress.setColor("white")
		}
		if dressType == TerroristDressType {
			dress.setColor("black")
		}
		fmt.Printf("dress:%+v,dress_point:%p, DressColorType: %s\nDressColor: %s\n", dress, dress, dressType, dress.getColor())
	}

	fmt.Println()

	for _, player := range game.terrorists {
		fmt.Printf("game player terrorists player:%+v,dress:%s\n", player, player.dress)
	}
	for _, player := range game.counterTerrorists {
		fmt.Printf("game player counterTerrorists player:%+v,dress:%s\n", player, player.dress)
	}

}
