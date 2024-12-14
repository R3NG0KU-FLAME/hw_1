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

// структура для инвентаря //
type Inventory struct {
	items []Item
}

func (inv *Inventory) PrintInventory() {
	if len(inv.items) == 0 {
		fmt.Println("Ваш инвентарь пуст.")
		return
	}
	fmt.Println("В вашем инвентаре:")
	for _, item := range inv.items {
		fmt.Println(" -", item.Name)
	}
}
func (inv *Inventory) Additem(item Item) {
	inv.items = append(inv.items, item)
	fmt.Println("Вы подобрали:", item.Name)
}
func (inv *Inventory) Hasitem(itemName string) bool {
	for _, item := range inv.items {
		if strings.ToLower(item.Name) == strings.ToLower(itemName) {
			return true
		}
	}
	return false
}
func printAvailable(location Location) {
	fmt.Println("Вы можете пойти в:")
	for _, v := range location.Neighbours {
		fmt.Println(" -", v)
	}
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
	fmt.Printf("%s нанес %d урона %s\n", opponent.person, opponent.damage, e.person)
}

func battle(warrior *Character, skeleton *Character, reader *bufio.Reader) {
	fmt.Println("Неожиданно на вас нападает скелет! Нужно победить его!")
	for warrior.hitpoints > 0 && skeleton.hitpoints > 0 {
		fmt.Println("Введите команду (атака, инфо герой, инфо скелет):")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {
		case "атака":
			warrior.AttackOpponent(skeleton)
			if skeleton.hitpoints <= 0 {
				fmt.Println("Вы победили скелета!")
				return
			}
			skeleton.AttackOpponent(warrior)
			if warrior.hitpoints <= 0 {
				fmt.Println("Вы были побеждены скелетом!")
				return
			}
		case "инфо герой":
			fmt.Printf("У %s осталось %d здоровья.\n", warrior.person, warrior.hitpoints)
		case "инфо скелет":
			fmt.Printf("У %s осталось %d здоровья.\n", skeleton.person, skeleton.hitpoints)
		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}

// Основная функция(ход боя)
func main() {

	reader := bufio.NewReader(os.Stdin)

	warrior := Character{"Воин", 100, 10, 50}
	skeleton := Character{"Скелет", 100, 5, 30}

	//создание предметов на локациях//
	sword := Item{
		Name:        "Меч",
		Description: "Старый меч,покрытый ржавчиной, но еще внушающий страх",
	}
	shield := Item{
		Name:        "Щит",
		Description: "Изрешеченный щит, который некогда защищал от вражеских ударов",
	}
	treasure := Item{
		Name:        "Сокровища",
		Description: "Легендарные сокровища, оставленные здесь великим и ужасным контрабандистом",
	}
	inventory := Inventory{}

	//Создание локаций//

	hall := Location{"Холл замка", "Вы находитесь в холле старого заброшенного замка, погруженном в полумрак. Пыльные стены и паутина создают атмосферу забвения", []Character{}, []Item{shield}, []string{"Коридор замка"}}

	corridor := Location{
		Name:        "Коридор замка",
		Description: "Темный коридор, в котором, как говорят, обитают призраки прошлого. Здесь кровь стынет в жилах...",
		Characters:  []Character{},
		Items:       []Item{shield},
		Neighbours:  []string{"Холл замка", "Большой зал"},
	}

	bighall := Location{
		Name:        "Большой зал",
		Description: "Великолепный зал, некогда блиставший золотом и драгоценностями, теперь скрывает в себе только тени прошлого",
		Characters:  []Character{},
		Items:       []Item{sword, shield, treasure},
		Neighbours:  []string{"Коридор замка"},
	}
	locations := map[string]Location{
		hall.Name:     hall,
		corridor.Name: corridor,
		bighall.Name:  bighall,
	}
	currentLocation := hall

	fmt.Println("Добро пожаловать в игру!\nВы-отважный воин ,который отправился в путешествие через леса и горы.\n" +
		"В поисках приключений вы наткнулись на старый замок, окутанный тайнами и легендами.\n" +
		"Как только вы вошли, массивные двери за вашей спиной захлопнулись.\n" +
		"Вы обнаружили, что перед дверьми выхода находится большая кнопка, должно быть нужно на нее встать с чем то тяжелым и дверь разблокируется.\n" +
		"Ваша задача - раскрыть секреты замка и найти путь наружу, но будьте осторожны - опасности подстерегают на каждом шагу.\n")

	fmt.Println("Доступные команды:")
	fmt.Println("-осмотреться: осмотреть текущую локацию и увидеть что в ней есть")
	fmt.Println("-доступные направления: посмотреть, куда можно пойти из текущей локации")
	fmt.Println("-взять [предмет]: взять указанный предмет который есть на этой локации")
	fmt.Println("-идти [локация]: переместиться в указанную локацию")
	fmt.Println("-инвентарь: показать содержимое инвентаря")
	fmt.Println("-выход: попытаться покинуть замок(возможно при выполнении некоторых условий)")

	for {
		fmt.Println("\nВы находитесь в:", currentLocation.Name)
		fmt.Println(currentLocation.Description)
		if currentLocation.Name == "коридор замка" {
			battle(&warrior, &skeleton, reader)
			if warrior.hitpoints <= 0 {
				fmt.Println("Игра окончена, вас убил скелет")
			}
		}

		fmt.Print("Введите команду: ")
		comm, _ := reader.ReadString('\n')
		comm = strings.TrimSpace(strings.ToLower(comm))

		parts := strings.SplitN(comm, " ", 2)

		switch comm {
		case "осмотреться":
			fmt.Println("Вы видите:")
			for _, item := range currentLocation.Items {
				fmt.Println(" -", item.Name)
			}
			for _, character := range currentLocation.Characters {
				fmt.Println(" -", character.person)
			}
		case "доступные направления":
			printAvailable(currentLocation)
		case "инвентарь":
			inventory.PrintInventory()
		case "выход":
			if currentLocation.Name == "Холл замка" && inventory.Hasitem("Сокровища") && skeleton.hitpoints <= 0 {
				fmt.Println("Кнопка прожалась благодаря весу сокровища! Вы успешно выбрались из замка!")
				return
			} else {
				fmt.Println("Вы не можете покинуть замок, нужно найти что то увесистое для нажатия кнопки перед выходом!(выход находится в холле замка)")
			}
		default:
			// Обработка команд с аргументами
			if len(parts) > 1 && parts[0] == "взять" {
				itemName := parts[1]
				itemIndex := -1
				for i, item := range currentLocation.Items {
					if strings.ToLower(item.Name) == itemName {
						itemIndex = i
						break
					}
				}
				if itemIndex >= 0 {
					inventory.Additem(currentLocation.Items[itemIndex])
					currentLocation.Items = append(currentLocation.Items[:itemIndex], currentLocation.Items[itemIndex+1:]...)
				} else {
					fmt.Println("Такого предмета нет в этой локации.")
				}
			} else if len(parts) > 1 && parts[0] == "идти" {
				direction := parts[1]
				canMove := false
				for _, neighbour := range currentLocation.Neighbours {
					if strings.ToLower(neighbour) == direction {
						if direction == "холл замка" && (!inventory.Hasitem("Сокровища") || skeleton.hitpoints > 0) {
							fmt.Println("Вы не можете покинуть замок, пока не победите скелета и не найдете сокровища!")
							break
						}
						canMove = true
						currentLocation = locations[neighbour]
						fmt.Println("Вы переместились в:", currentLocation.Name)
						break
					}
				}
				if canMove && currentLocation.Name == "Коридор замка" && skeleton.hitpoints > 0 {
					battle(&warrior, &skeleton, reader)
					if warrior.hitpoints <= 0 {
						fmt.Println("Игра окончена.")
						return
					}
				} else if !canMove {
					fmt.Println("Вы не можете туда пойти")
				}
			} else {
				fmt.Println("Неизвестная команда")
			}
		}
	}
}
