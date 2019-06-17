package action

import (
	"fmt"
	"os"
	
	// "io/ioutil"
	// "math/rand"
	// "time"
)

var nameHero string
func Greeting(){
	fmt.Println("Введите свое имя (Например Vasya_Pupkin): ")
	fmt.Fscan(os.Stdin,&nameHero)
	fmt.Println("Your name is: ",nameHero)
	if nameHero==" " {
		fmt.Println("Пустое значение! ")
		Greeting()
	}else{
		ActionGame()
	}
}