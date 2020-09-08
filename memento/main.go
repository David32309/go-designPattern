package main

import "fmt"

type Memento struct {
	assets int
}

type Player struct {
	name   string
	assets int
}

func (p *Player) Name() string { return p.name }

func (p *Player) Pay(amount int) {
	p.assets -= amount
}

func (p *Player) Info() string {
	return fmt.Sprintf("name: %s, assets:%d", p.name, p.assets)
}

func (p *Player) SetMemento(m *Memento) {
	if m == nil {
		return
	}
	p.assets = m.assets
}

func (p *Player) Memento() *Memento {
	return &Memento{
		assets: p.assets,
	}
}

type Caretaker struct {
	m map[string]*Memento
}

func (c *Caretaker) Set(key string, memento *Memento) {
	if c.m == nil {
		c.m = make(map[string]*Memento)
	}
	c.m[key] = memento
}

func (c *Caretaker) Get(key string) (*Memento, bool) {
	memento, ok := c.m[key]
	return memento, ok
}

func main() {
	var caretaker Caretaker
	player := Player{
		name:   "Tom",
		assets: 1000,
	}

	caretaker.Set(player.Name(), player.Memento())
	fmt.Printf("initial:\n\t%s\n", player.Info())

	player.Pay(2000)
	fmt.Printf("after paying 2000:\n\t%s\n", player.Info())

	memento, _ := caretaker.Get(player.Name())
	player.SetMemento(memento)
	fmt.Printf("after recovery:\n\t%s\n", player.Info())
}
