package main

import (
	"fmt"
)

// Старый плеер с разъемом minijack

type MiniJackPlayer interface {
	PlayWithMiniJack()
}

type OldPlayer struct {}

func (o *OldPlayer) PlayWithMiniJack() {
	fmt.Println("... PlayWithMiniJack!")
}

// Новые наушники с разъемом usbc

type USBCHeadPhones interface {
	PlayWithUSBC()
}

type NewHeadPhones struct {}

func (n *NewHeadPhones) PlayWithUSBC() {
	fmt.Println("... PlayWithUSBC!")
}

// Переходник "Адаптер"

type MiniJackToUSBCAdapter struct {
	Player MiniJackPlayer
}

func (a *MiniJackToUSBCAdapter) PlayWithUSBC() {
	fmt.Println("... Using adapter!")
	a.Player.PlayWithMiniJack()
}

func main() {
	oldPlayer := &OldPlayer{}

	var headphones USBCHeadPhones = &MiniJackToUSBCAdapter{Player: oldPlayer}
	headphones.PlayWithUSBC()
}
