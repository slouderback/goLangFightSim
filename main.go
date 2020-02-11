package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)


	fmt.Print("Welcome please enter your name: ")
	username, _ := consoleReader.ReadString('\n')
	username = strings.TrimRight(username, "\n")
	fmt.Println("Hello there", username)


	fmt.Println("Choose your class! \n(1) Warrior \n(2) Mage \n(3) Assassin")
	class, _ := consoleReader.ReadString('\n')
	class = strings.TrimRight(class, "\n")

	player := player{}

	if class == "1" {
		fmt.Println("Ah I see you are a Warrior of great strength!")
		player.class = "Warrior"
		player.name = username
		player.health = 15
		player.damage = 6
		player.inventory = append(player.inventory, "Sword")
	} else if class == "2" {
		fmt.Println("Ah I see you are a Mage of great intelligence!")
		player.class = "Mage"
		player.name = username
		player.health = 8
		player.damage = 8
		player.inventory = append(player.inventory, "Ignite")
		player.inventory = append(player.inventory, "Ice Blast")
	} else if class == "3" {
		fmt.Println("Ah I see you are an assassin!")
		player.class = "Assassin"
		player.name = username
		player.health = 6
		player.damage = 10
		player.inventory = append(player.inventory, "Dagger")
	}
	//fmt.Println("Class:", player.class, "Name:", player.name, "Health:", player.health, "Damage:", player.damage)
	fmt.Println("Lets do some training! Hit this practice dummy here. ")
	for i := 0; i < len(player.inventory); i++{
		fmt.Printf("(%d) %s\n", i+1, player.inventory[i])
	}
	attackInput, _ := consoleReader.ReadString('\n')
	attackInput = strings.TrimRight(class, "\n")
	attackOption, _ := strconv.Atoi(attackInput)
	fmt.Printf("%s hit!\n", player.inventory[attackOption-1])

	fmt.Println("Alright time to fight your first real enemy!")
	minion01 := enemy{25, 1, make([]string, 0)}
	minion01.inventory = append(minion01.inventory, "Staff")
	minion01.inventory = append(minion01.inventory, "Big Staff")

	currentTurn := 0
	for player.health > 0 && minion01.health > 0{
		if currentTurn == 0{
			fmt.Println("--------------------")
			fmt.Printf("Your HP: %d\n", player.health)
			fmt.Printf("Monster HP: %d\n", minion01.health)
			fmt.Println("--------------------")
			fmt.Println("Choose your method of attack: ")
			for i := 0; i < len(player.inventory); i++{
				fmt.Printf("(%d) %s\n", i+1, player.inventory[i])
			}
			attackInput, _ := consoleReader.ReadString('\n')
			attackInput = strings.TrimRight(class, "\n")
			attackOption, _ := strconv.Atoi(attackInput)
			fmt.Printf("%s hit for %d damage!\n", player.inventory[attackOption-1], player.damage)
			minion01.health = minion01.health - player.damage
			fmt.Print("\n\n")
			if minion01.health < 1{
				fmt.Println("The monster died!")
			}
			currentTurn = 1
		} else if currentTurn == 1{
			time.Sleep(1 * time.Second)
			fmt.Println("--------------------")
			fmt.Printf("Your HP: %d\n", player.health)
			fmt.Printf("Monster HP: %d\n", minion01.health)
			fmt.Println("--------------------")
			fmt.Printf("The monster attacks and you take %d damage\n", minion01.damage)
			player.health = player.health - minion01.damage
			if player.health > 0{
				fmt.Print("\n\n")
			} else{
				fmt.Println("You died.")
			}
			time.Sleep(2 * time.Second)
			currentTurn = 0
		}
	}
}

