package script

import (
	"fmt"
	"math/rand"
	"time"
)
var result2Drag int
func MyRandTwo(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	// randomHero:=rand.Intn(max-min)+min	
	randomDrag:=rand.Intn(max-min)+min
	// result2Hero:=resultHero-randomHero
	result2Drag:=resultDrag-randomDrag
	// fmt.Println("Ваше здоровье", "( - ", randomHero, ")", result2Hero)
	fmt.Println("Здоровье Дракона ", "( - ", randomDrag, ")", result2Drag)
	return rand.Intn(max-min)+min
}

