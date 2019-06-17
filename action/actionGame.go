package action

import (
	"fmt"
	"os"
	"./script"
	"io/ioutil"

)


var req string
var request string

func ActionGame(){
	fmt.Println("Приветствую тебя наш герой - ",nameHero)
	b, err := ioutil.ReadFile("history.txt")
	if err != nil {
        fmt.Print(err)
    }

    // fmt.Println(b) 
	str := string(b) 
	fmt.Println(str) 
	fmt.Println("--------------------")
	fmt.Println("Выбери оружие: ")
	fmt.Println("--------------------")
	fmt.Println("Меч и щит (print 1),\nАвтомат Калашникова (print 2),\nОльга Бузова (print 3)")
	fmt.Fscan(os.Stdin, &req)
		if req=="1"{
	fmt.Println("Вы ударили мечом дракона и защитились щитом")
			script.MyRand(30,60)
		}else if req=="2"{
	fmt.Println("Вы стреляете по дракону калибром 7,62 но защититься нечем")
			script.MyRand(70,120)
		}else if req=="3"{
	fmt.Println("Пока дракон офигивает от Мало Половин вы придумываете план спасения принцессы")
			script.MyRand(120,220)
		}else{
			fmt.Println("Uncorrect!")
			ActionGame()
		}

		fmt.Println("Раунд 2 ")
		fmt.Println("Ты не справишься с этим монстром без особого бонуса.... ")
		fmt.Println("Что б получить его отгадай загадку: Есть всегда у людей Есть всегда у кораблей.")
		fmt.Fscan(os.Stdin, &request)
		if request=="нос" || request=="Нос"{
			fmt.Println("--------------------")
			fmt.Println("Что бы спасти принцессу тебе нужна поддержка эльфиек из улицы Кулиева:  ")
			fmt.Println("--------------------")
			fmt.Println("Деньги (print 1) \n Провиант (print 2) \n Девственность (print 3)")
			fmt.Fscan(os.Stdin, &req)
				if req=="1"{
					fmt.Println("Вы заплатили деньги и они помогли вам побороться с драконом")
					script.MyRandTwo(70,121)
				}else if req=="2"{
			fmt.Println("Вы накормили их и они помогают бороться с драконом")
					script.MyRandTwo(50,81)
				}else if req=="3"{
			fmt.Println("Вы отдаете им самое дорогое что у вас есть и они крошат его)")
					script.MyRandTwo(120,150)
				}else{
					fmt.Println("Uncorrect!")
					ActionGame()
				}
		}else{
			fmt.Println("Uncorrect!")
			ActionGame()
		
		}
		


	
			fmt.Println("----------Раунд 3-----------")
			fmt.Println("Решающая битва цель которой узнать индекс среднего числа массива состоящего из 3 чисел:  ")
			fmt.Println("--------------------")
			fmt.Println("1 (print 1) \n 2 (print 2) \n 3 (print 3)")
			fmt.Fscan(os.Stdin, &req)
				if req=="1"{
					fmt.Println("В решающей схватке вы победили дракона!!!!!")
					script.MyRandFour(200,321)
				}else if req=="2"{
			fmt.Println("Охх как вы промахнулись.....")
					script.MyRandThree(150,281)
				}else if req=="0"{
					fmt.Println("Охх как вы промахнулись.....")
					script.MyRandThree(150,281)
				}else{
					fmt.Println("Вас испепелили!!")
					fmt.Println("Начать заново?")
					fmt.Println("Да (print 1) \n Нет (print 2) \n Не знаю (print 3)")
					fmt.Fscan(os.Stdin, &req)
					if req=="1"{
					ActionGame()
					}else if req=="2"{
				fmt.Println("LOL")
					}else if req=="0"{
						fmt.Println("Охх как вы промахнулись.....")
						script.MyRandThree(150,281)
					}else{
						fmt.Println("Вас испепелили!!")
						fmt.Println("Начать заново?")
						fmt.Println("Да (print 1) \n Нет (print 2) \n Не знаю (print 3)")
						fmt.Fscan(os.Stdin, &req)
						ActionGame()
					}
			
				}
		

	}

