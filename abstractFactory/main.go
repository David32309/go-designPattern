package main

import (
	"fmt"
	"runtime"
)

type Button interface {
	Type() string
}

type WinButton struct{}

func (*WinButton) Type() string { return "WinButton" }

type LinuxButton struct{}

func (*LinuxButton) Type() string { return "LinuxButton" }

type Editor interface {
	Type() string
}

type WinEditor struct{}

func (*WinEditor) Type() string { return "WinEditor" }

type LinuxEditor struct{}

func (*LinuxEditor) Type() string { return "LinuxEditor" }

type Factory interface {
	CreateButton() Button
	CreateEditor() Editor
}

type WinFactory struct{}

func (*WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (*WinFactory) CreateEditor() Editor {
	return &WinEditor{}
}

type LinuxFactory struct{}

func (*LinuxFactory) CreateButton() Button {
	return &LinuxButton{}
}

func (*LinuxFactory) CreateEditor() Editor {
	return &LinuxEditor{}
}

func CreateFactory(os string) Factory {
	switch os {
	case "windows":
		return &WinFactory{}
	case "linux":
		return &LinuxFactory{}
	default:
		return nil
	}
}

func main() {
	factory := CreateFactory(runtime.GOOS)
	if factory == nil {
		return
	}
	fmt.Println("type of button: ", factory.CreateButton().Type())
	fmt.Println("type of editor: ", factory.CreateEditor().Type())
}
