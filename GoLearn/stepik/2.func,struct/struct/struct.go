package main

import "fmt"

type gun struct {
	On    bool
	Ammo  int
	Power int
}

func (g *gun) Shoot() bool {
	if g.On == false {
		return false
	} else if g.Ammo > 0 {
		g.Ammo--
		return true
	}
	return false
}

func (g *gun) RideBike() bool {
	if g.On == false {
		return false
	} else if g.Power > 0 {
		g.Power--
		return true
	}
	return false
}

func main() {

	testStruct := new(gun)
	fmt.Println(testStruct)
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */

}
