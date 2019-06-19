package action

import (
    "fmt"
	"os"
	// "./raundes"
	// "io/ioutil"
	// "math/rand"
	// "time"
)
var a string
var req5 int
func StartGame(){
	tabulation:="\t\t\t"

	fmt.Println(tabulation, " | Добро пожаловать в игру 'Герой против Дракона | ' ")
	fmt.Println(" Press A ")
	fmt.Fscan(os.Stdin, &a) 
		if a=="A" || a== "a"{
				
			fmt.Println(tabulation,"Сразишься с драконом:(print: 1)")
			fmt.Println(tabulation, "Не я лучше домой ДОМ 2 смотреть (print : 2)")
	
			fmt.Fscan(os.Stdin, &req5) 
				switch(req5){
					case 1:
						fmt.Println("--------------------")
	fmt.Println("Введите свое имя (Например Vasya_Pupkin): ")
						Greeting()
					case 2:
						fmt.Println("Suck me!")
							StartGame()
					default:
							fmt.Println("Uncorrect")
							StartGame()	
						}
						
			}else{
			fmt.Println("Uncorrect! Начать заново.... ")
			StartGame()
		}
	

}