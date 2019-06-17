package action

import (
	"fmt"
	"os"
	"bufio"
	// "io/ioutil"
	// "math/rand"
	// "time"
)

var nameHero string
func Greeting(){
	
	
	// fmt.Fscan(os.Stdin,&nameHero)
	
	// fmt.Scanf("%[^\n]s",nameHero)
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	nameHero= myscanner.Text()
	if nameHero!="" {
		fmt.Println("Your name is: ",nameHero)
			ActionGame()
		}else{
				// fmt.Println("Пустое значение! ")
				Greeting()
			}
	// switch nameHero{
	// case " ":
	// 	fmt.Println("Пустое значение! ")
	// 	Greeting()
	// default:
	// 	ActionGame()
	
	// }
}