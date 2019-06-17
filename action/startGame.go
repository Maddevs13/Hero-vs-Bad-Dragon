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
// var req string
func StartGame(){

	fmt.Println("Добро пожаловать в игру 'Герой против Дракона' ")
	fmt.Println(" Press A ")
	fmt.Fscan(os.Stdin, &a) 
		if a=="A" || a== "a"{
				
			fmt.Println("Сразишься с драконом:(print: yes)")
			fmt.Println("Не я лучше домой ДОМ 2 смотреть (print : no)")
	
			fmt.Fscan(os.Stdin, &req) 
				switch(req){
					case "yes":
						fmt.Println("--------------------")
	fmt.Println("Введите свое имя (Например Vasya_Pupkin): ")
						Greeting()
						
		// test()
					case "no":
						fmt.Println("Suck me!")
							StartGame()
					default:
							fmt.Println("Uncorrect")
							StartGame()	
						}
						
			// DragonBattle()
			// DragonBattle2()
			// raundes.RaundThree()
			}else{
			fmt.Println("Uncorrect! Начать заново.... ")
			StartGame()
		}
	

}