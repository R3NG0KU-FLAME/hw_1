package main

import "fmt"

type employee struct {
	person    string
	hitpoints int
	armor     int
}

func newEmployee(person string, hitpoints, armor int) employee {
	return employee{
		person:    person,
		hitpoints: hitpoints,
		armor:     armor,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Перс: %s\nЗдоровье: %d\nБроня: %d\n", e.person, e.hitpoints, e.armor)

}

func main() {
	employee1 := newEmployee("Воин", 100, 100)
	employee2 := newEmployee("Орк", 100, 100)

	fmt.Printf("%s\n", employee1.getInfo())
	fmt.Printf("%s\n", employee2.getInfo())
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}
