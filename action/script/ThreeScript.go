package script

import (
	"fmt"

	"math/rand"
	"time"
)
func MyRandThree(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	randomHero:=rand.Intn(max-min)+min	
	// randomDrag:=rand.Intn(max-min)+min
	result2Hero:=resultHero-randomHero
	// result3Drag:=result2Drag-randomDrag
	fmt.Println("Ваше здоровье", "( - ", randomHero, ")", result2Hero)
	// fmt.Println("Здоровье Дракона ", "( - ", randomDrag, ")", result3Drag)
	return rand.Intn(max-min)+min
}


func MyRandFour(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	// randomHero:=rand.Intn(max-min)+min	
	randomDrag:=rand.Intn(max-min)+min
	// result2Hero:=resultHero-randomHero
	result3Drag:=result2Drag-randomDrag
	// fmt.Println("Ваше здоровье", "( - ", randomHero, ")", result2Hero)
	fmt.Println("Здоровье Дракона ", "( - ", randomDrag, ")", result3Drag)
	return rand.Intn(max-min)+min
}



