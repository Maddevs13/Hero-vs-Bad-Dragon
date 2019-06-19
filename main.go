package main
import (
	"fmt"
	"os"

	"./action"
	
)
var tabulation string

var menu int
func main(){
	tabulation="\n\t\t\t"
	fmt.Println(tabulation," | Герой против Дракона | ")
	fmt.Println("Начать игру (print 1),\nВыход (print 2)")
	fmt.Fscan(os.Stdin, &menu)
	// fmt.Scanf("%[^\n]s",&menu)
		if menu==1{
	action.StartGame()
		}else if menu==2{
			action.EndGame()
		}else{
			fmt.Println(tabulation, "Uncorrect!")
			// main()	
		}

	// actionGame()
	// endGame()
}