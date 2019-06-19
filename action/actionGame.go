package action

import (
	"fmt"
	"os"
	"./script"
	"io/ioutil"

)


var req1 int
var req2 int
var req3 int
var req4 int

var request int

func ActionGame(){
	tabulation:="\t\t\t"
	fmt.Println(tabulation,"---------------------------------------------------------")
	fmt.Println(tabulation,"Приветствую тебя наш герой - ",nameHero)
	fmt.Println(tabulation,"---------------------------------------------------------")
	
	b, err := ioutil.ReadFile("history.txt")
	fmt.Println("Раунд 1")
	if err != nil {
        fmt.Print(err)
    }

    // fmt.Println(b) 
	str := string(b) 
	fmt.Println(str) 
	fmt.Println("--------------------")
	fmt.Println("Выбери оружие: ")
	fmt.Println("--------------------")
	weapon:=make(map[string]int)
	weapon["Меч и щит(print 1)"]=1
	weapon["Автомат Калашникова(print 2)"]=2
	weapon["Ольга Бузова(print 3)"]=3
	// fmt.Println(weapon)
	for key:=range weapon{
		fmt.Println(key)
	}
	// // fmt.Println(" | Меч и щит (print 1) |,\n| Автомат Калашникова (print 2) |,\n| Ольга Бузова (print 3) | ")
	fmt.Fscan(os.Stdin, &req1)
		if req1==1{
	fmt.Println("Вы ударили мечом дракона и защитились щитом")
			script.MyRand(30,60)
		}else if req1==2{
	fmt.Println("Вы стреляете по дракону калибром 7,62 но защититься нечем")
			script.MyRand(70,120)
		}else if req1==3{
	fmt.Println("Пока дракон офигивает от Мало Половин вы придумываете план спасения принцессы")
			script.MyRand(120,220)
		}else{
			fmt.Println("Uncorrect!")
			ActionGame()
		}
		fmt.Println("------------------------------------")
		fmt.Println("Раунд 2 ")
		fmt.Println("Ты не справишься с этим монстром без особого бонуса.... ")
		fmt.Println("------------------------------------")
		fmt.Println("Что б получить его отгадай загадку: Есть всегда у людей Есть всегда у кораблей.")
		fmt.Println("------------------------------------")		
		bonus:=map[string]int{
			"нос - 1":1,
			"лицо - 2":2,
			"трусы - 3":3,
		}

	for key:=range bonus{
		fmt.Println(key)
	}
	fmt.Fscan(os.Stdin, &request)
		if request==1{
			fmt.Println("--------------------")
			fmt.Println("Что бы спасти принцессу тебе нужна поддержка эльфиек из улицы Кулиева:  ")
			fmt.Println("--------------------")
			help:=make(map[string]int)
			help["Деньги - 1"]=1
			help["Провиант - 2"]=2
			help["Девственность - 3"]=3
			for key:=range help{
				fmt.Println(key)
			}
	
			// fmt.Println("Деньги (print 1) \n Провиант (print 2) \n Девственность (print 3)")
			fmt.Fscan(os.Stdin, &req2)
				if req2==1{
					fmt.Println("Вы заплатили деньги и они помогли вам побороться с драконом")
					script.MyRandTwo(70,121)
				}else if req2==2{
			fmt.Println("Вы накормили их и они помогают бороться с драконом")
					script.MyRandTwo(50,81)
				}else if req2==3{
			fmt.Println("Вы отдаете им самое дорогое что у вас есть и они крошат его)")
					script.MyRandTwo(120,150)
				}else{
					fmt.Println("Uncorrect!")
					ActionGame()
				}
		}else{
			fmt.Println("Фильм тупой и еще тупее не с тобой случайно снимали?)")
			EndGame()
			
		}
		


	
			fmt.Println("----------Раунд 3-----------")
			fmt.Println("Решающая битва цель которой узнать индекс среднего числа массива состоящего из 3 чисел:  ")
			fmt.Println("--------------------")
			fmt.Println("1 (print 1) \n 2 (print 2) \n 3 (print 3)")
			fmt.Fscan(os.Stdin, &req3)
				if req3==1{
					fmt.Println("В решающей схватке вы победили дракона!!!!!")
					script.MyRandFour(200,321)
				}else if req3==2{
			fmt.Println("Охх как вы промахнулись.....")
					script.MyRandThree(150,281)
				}else if req3==0{
					fmt.Println("Охх как вы промахнулись.....")
					script.MyRandThree(150,281)
				}else{
					fmt.Println("Вас испепелили!!")
					fmt.Println("Начать заново?")
					fmt.Println("Да (print 1) \n Нет (print 2) \n Не знаю (print 3)")
					fmt.Fscan(os.Stdin, &req4)
					if req4==1{
					ActionGame()
					}else if req4==2{
				fmt.Println("LOL")
					}else if req4==3{
						fmt.Println("Думай быстрее.....")
						}else{
						fmt.Println("Начать заново?")
						fmt.Println("Да (print 1) \n Нет (print 2) \n Не знаю (print 3)")
					}
			
				}
		

	}

