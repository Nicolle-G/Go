package main

import "fmt"

const size = 5

var table [size][]Peers

type Peers struct {
	clue  string
	value int
}

func IndexHash(clue string) int {

	hash := 0
	runes := []rune(clue)

	for i := 0; i < len(runes); i++ {
		hash = hash*31 + int(runes[i])
	}
	return hash % size
}

func InsertValue(clue string, value int) {

	index := IndexHash(clue)
	peers := Peers{clue: clue, value: value}
	table[index] = append(table[index], peers)
}

func ShowTable() {
	for i := 0; i < len(table); i++ {
		fmt.Printf("Index %d:\n", i)
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("  Clue: %s, Value: %d\n", table[i][j].clue, table[i][j].value)
		}
	}
}

func SearchValue(clue string) int {
	index := IndexHash(clue)
	list := table[index]

	for i := 0; i < len(list); i++ {
		if list[i].clue == clue {
			return list[i].value
		}
	}

	return -1
}

func main() {

	InsertValue("lupe", 5)
	InsertValue("mateo", 10)
	InsertValue("oño", 20)
	InsertValue("pelu", 11)

	ShowTable()

	fmt.Println("Search for values")
	fmt.Println(" clue lupe", SearchValue("lupe"))
	fmt.Println(" clue oto", SearchValue("oto"))

}
