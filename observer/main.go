package main

import "fmt"

type Observer interface {
	OnEvent(interface{})
}

type Subject struct {
	observers []Observer
}

func (m *Subject) Attach(o Observer) {
	m.observers = append(m.observers, o)
}

func (m *Subject) SentEvent(event interface{}) {
	for _, o := range m.observers {
		o.OnEvent(event)
	}
}

type Mouse struct {
	Subject
	x,y int
}

type MouseMove struct {
	X,Y int
}

func (m *Mouse) Move(x,y int) {
	m.x, m.y = x, y
	m.SentEvent(&MouseMove{m.x, m.y})
}

type MouseDown struct {
	X,Y int
}

func (m *Mouse) Down() {
	m.SentEvent(&MouseDown{m.x, m.y})
}

type Keyboard struct {
	Subject
}

type KeyboardDown struct {
	Key string
}

func (m *Keyboard) Down(key string) {
	m.SentEvent(&KeyboardDown{key})
}

type Node struct {
	Name string
}

func (n *Node) OnEvent(e interface{}) {
	switch v := e.(type) {
	case *MouseMove:
		fmt.Printf("[%s] get MouseMove{X:%d, Y:%d}\n", n.Name, v.X, v.Y)
	case *MouseDown:
		fmt.Printf("[%s] get MouseDown{X:%d, Y:%d}\n", n.Name, v.X, v.Y)
	default:
		fmt.Printf("[%s] ignore\n", n.Name)
	}
}

type Editor struct {
	Name string
}

func (e *Editor) OnEvent(event interface{}) {
	switch v := event.(type) {
	case *KeyboardDown:
		fmt.Printf("[%s] get KeyboardDown{Key: %s}\n", e.Name, v.Key)
	default:
		fmt.Printf("[%s] ignore\n", e.Name)
	}
}

func main() {
	mouse := &Mouse{}
	keyboard := &Keyboard{}
	allSubject := []interface{Attach(Observer)}{mouse, keyboard}

	node := &Node{"Node"}
	editor := &Editor{"Editor"}


	for _, subject := range allSubject {
		subject.Attach(node)
		subject.Attach(editor)
	}

	fmt.Println("sent MouseMove{X:100,Y:100}")
	mouse.Move(100,100)

	fmt.Println("\nsent MouseDown{X:100,Y:100}")
	mouse.Down()

	fmt.Println("\nsent KeyboardDown{Key:a}")
	keyboard.Down("a")
}
