package main

import "fmt"

//透過Command Patten實作redo/undo功能

type Command interface {
	Execute()
	UnExecute()
}

type Insert struct {
	doc *string
	pos int
	str string
}

func (a *Insert) Execute() {
	*a.doc = (*a.doc)[:a.pos]+ a.str + (*a.doc)[a.pos:]
}

func (a *Insert) UnExecute() {
	*a.doc = (*a.doc)[:a.pos] + (*a.doc)[a.pos+len(a.str):]
}

//Invoker
type Control struct {
	redo []Command
	undo []Command
}

func (a *Control) Do(c Command) {
	a.undo = append([]Command{c}, a.undo...)
	a.redo = nil
	c.Execute()
}

func (a *Control) Redo() {
	if len(a.redo) == 0 {
		return
	}
	c := a.redo[0]
	a.redo = a.redo[1:]
	a.undo = append([]Command{c}, a.undo...)
	c.Execute()
}

func (a *Control) Undo() {
	if len(a.undo) == 0 {
		return
	}
	c := a.undo[0]
	a.undo = a.undo[1:]
	a.redo = append([]Command{c}, a.redo...)
	c.UnExecute()
}

func main() {
	//Receiver
	doc := "test-"
	var control Control

	control.Do(&Insert{&doc, 5, "1111"})
	fmt.Println("after inserting 1111: ", doc)
	control.Undo()
	fmt.Println("after undoing: ", doc)
	control.Do(&Insert{&doc, 5, "2222"})
	fmt.Println("after inserting 2222: ", doc)
	control.Undo()
	fmt.Println("after undoing: ", doc)
	control.Redo()
	fmt.Println("after redoing: ", doc)
}