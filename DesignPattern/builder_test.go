package DesignPattern

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	var builder Builder = &CharacterBuilder{}
	var director = &Director{builder: builder}
	var character = director.Create("loader", "AK47")
	fmt.Println(character.GetName() + "," + character.GetArms())
}
