package main

import "fmt"

type employee struct {
	person    string
	hitpoints int
	armor     int
	damage    int
}

func newEmployee(person string, hitpoints, armor, damage int) employee {
	return employee{
		person:    person,
		hitpoints: hitpoints,
		armor:     armor,
		damage:    damage,
	}
}

func (e employee) getDamage() string {
	return fmt.Sprintf("Персонаж: %s\nЗдоровье: %d\nБроня: %d\nНаносимый урон: %d\n", e.person, e.hitpoints, e.armor, e.damage)

}

func (e employee) setName(person string) {
	e.person = person

}

func main() {

	employee1 := newEmployee("Воин", 100, 100, 15)
	employee2 := newEmployee("Скелет", 100, 100, 15)

	fmt.Println("Выбери героя: \n")
	fmt.Printf("%s\n", employee1.getDamage())
	fmt.Printf("%s\n", employee2.getDamage())

	a := new

	switch a(employee1.getDamage()) {

	case a(employee1.getDamage()):
		fmt.Scan(&employee1.person)
		fmt.Println("Вы выбрали Воина!\nУдачи в схватке!")

	}

	return
}
