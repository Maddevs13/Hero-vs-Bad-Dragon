package main
import (
	"fmt"
	"os"

    "./action"
)

var menu string
func main(){
	fmt.Println("Герой против Дракона")
	fmt.Println("Начать игру (print 1),\nВыход (print 2)")
	fmt.Fscan(os.Stdin, &menu)
	// fmt.Scanf("%[^\n]s",&menu)
		if menu=="1"{
	action.StartGame()
		}else if menu=="2"{
			
		}else{
			fmt.Println("Uncorrect!")
			// main()	
		}

	// actionGame()
	// endGame()
}