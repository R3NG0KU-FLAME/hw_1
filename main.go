package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	person    string
	hitpoints int
	armor     int
	damage    int
}

// вызов инфы//
func (e *Character) Info() {
	fmt.Printf("Персонаж: %s\nЗдоровье: %d\nБроня: %d\nНаносимый урон: %d\n", e.person, e.hitpoints, e.armor, e.damage)

}

// Атака (героя/cкелета) //
func (e *Character) AttackOpponent(opponent *Character) {
	attack := e.damage - opponent.armor
	if attack < 0 {
		attack = 0
	}
	opponent.hitpoints -= attack
	fmt.Printf("%s нанес %d урона %s\n", e.person, opponent.damage, opponent.person)
}

// Основная функция(ход боя)
func main() {

	reader := bufio.NewReader(os.Stdin)

	warrior := Character{"Воин", 100, 10, 15}
	skeleton := Character{"Скелет", 100, 5, 15}

	// Основной цикл игры
	for warrior.hitpoints > 0 && skeleton.hitpoints > 0 {
		fmt.Println("Введите команду (пример: атака воина, атака скелета, инфо герой, инфо скелет):")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {
		case "атака героя":
			warrior.AttackOpponent(&skeleton)
			if skeleton.hitpoints <= 0 {
				fmt.Println("Герой победил!")
				return
			}
		case "атака скелета":
			skeleton.AttackOpponent(&warrior)
			if warrior.hitpoints <= 0 {
				fmt.Println("Герой проиграл!")
				return
			}
		case "инфо герой":
			warrior.Info()
		case "инфо скелет":
			skeleton.Info()
		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}
