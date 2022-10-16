package main

import "fmt"

// Task 1
type People struct {
	id             int
	name           string
	old            int
	follower_count int
}

// old + 1
func (p *People) oldDeath() int {
	return p.old + 1
}

func Task1() {
	p := People{1, "Alex", 30, 10}
	fmt.Println(p.oldDeath())
}

// Task 2
type Car struct {
	id      int
	title   string
	цвет    string
	цена    int
	user_id int
}

func car() {
	Car := Car{1, "BMW", "Black", 1000000, 1}
	fmt.Println(Car)
}

func main() {
	Task1()
	car()
}
