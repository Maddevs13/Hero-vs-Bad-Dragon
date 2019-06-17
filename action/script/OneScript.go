package script

import (
	"fmt"
	"math/rand"
	"time"
)
var hpHero=300
var hpDragon=350
var resultHero int
var resultDrag int

func MyRand(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	randomHero:=rand.Intn(max-min)+min	
	randomDrag:=rand.Intn(max-min)+min
	resultHero=hpHero-randomHero
	resultDrag=hpDragon-randomDrag
	fmt.Println("Ваше здоровье", "( - ", randomHero, ")", resultHero)
	fmt.Println("Здоровье Дракона ", "( - ", randomDrag, ")", resultDrag)
	return rand.Intn(max-min)+min
}

