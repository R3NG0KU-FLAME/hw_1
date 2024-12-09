package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Структура для персов//
type Character struct {
	person    string
	hitpoints int
	armor     int
	damage    int
}

// структура для локаций//
type Location struct {
	Name        string
	Description string
	Characters  []Character
	Items       []Item
	Neighbours  []string //связь локаций//
}

// Структура предметов на локациях//
type Item struct {
	Name        string
	Description string
}

// Проверка локаций а так же присутствия на ней предметов и персов//
func printLocations(locations []Location) {
	for _, loc := range locations {
		fmt.Printf("Локация: %s\nОписание: %s\n\n", loc.Name, loc.Description)
		fmt.Println("Персонажи:")
		for _, char := range loc.Characters {
			fmt.Printf(" - %s(Здоровье: %d,Броня: %d, Урон: %d)\n", char.person, char.hitpoints, char.armor, char.damage)
		}
		fmt.Println("Предметы: ")
		for _, item := range loc.Items {
			fmt.Printf("- %s: %s\n", item.Name, item.Description)
		}
		fmt.Println()
	}
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

	//создание предметов на локациях//
	sword := Item{
		Name:        "Меч",
		Description: "Старый ржавый меч лежавший здесь сотни лет",
	}
	shield := Item{
		Name:        "Щит",
		Description: "Щит который изрешетен ударами топора и меча, чем то похож на дуршлаг",
	}
	treasure := Item{
		Name:        "Сокровища",
		Description: "Сокровища которые оставлены здесь много веков назад известным контрабандистом",
	}

	//Создание локаций//

	hall := Location{"Холл замка", "Старый заброшеный замок в рыжем лесу", []Character{warrior}, []Item{shield}, []string{"Коридор замка"}}
	corridor := Location{
		Name:        "Коридор замка",
		Description: "Здесь всегда проходят битвы, здесь очень опасно",
		Characters:  []Character{skeleton},
		Items:       []Item{shield},
		Neighbours:  []string{"Холл замка", "Большой зал"},
	}

	bighall := Location{
		Name:        "Большой зал",
		Description: "Здесь по приданиям должны находиться сокровища",
		Characters:  []Character{warrior},
		Items:       []Item{sword, shield, treasure},
		Neighbours:  []string{"Коридор замка"},
	}
	locations := []Location{hall, corridor, bighall}

	printLocations(locations)

	// Основной цикл игры //
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
