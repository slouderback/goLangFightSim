package main

import "fmt"

type player struct{
	class string
	name string
	health int
	damage int
	inventory[] string
}

func printStats(a player){
	fmt.Println("Class:", a.class, "Name:", a.name, "Health:", a.health, "Damage:", a.damage)
}