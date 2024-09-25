package main

import "fmt"

type Player struct {
	name        string
	HealthPoint int
	Magic       int
}

func NewPlayer(name string, healthPoint int, magic int) *Player {
	return &Player{name, healthPoint, magic}

}

func (p *Player) attack() {
	fmt.Println(p.name, "攻击了")
	p.Magic -= 1
}

func main() {

	stu := NewPlayer("yang", 1, 10)

	stu.attack()
	fmt.Println(stu.Magic)

}
