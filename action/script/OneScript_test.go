package script

import (
	"testing"
	"reflect"
)


func TestOneScript(t *testing.T){
	hpHeroExpected:=300
	
if hpHero!=hpHeroExpected{
		t.Errorf("Ошибка в коде и в здоровье героя")
}
 

	if reflect.TypeOf(hpHero)!=reflect.TypeOf(hpDragon){
		t.Errorf("Типы данных переменныъ отвечающих за здоровье персонажей равны")
		
	}
}
