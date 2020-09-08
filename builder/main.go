package main

import "fmt"

type Drink struct {
    size string
    drink string
}

func (d *Drink) SetSize(size string) {
    d.size = size
}

func (d *Drink) Size() string { return d.size}

func (d *Drink) SetDrink(drink string) {
    d.drink = drink
}

func (d *Drink) Drink() string { return d.drink}

type Builder interface {
    CreateDrink()
    Drink() *Drink
    SetSize()
    SetDrink()
}

type TeaBuilder struct {
    drink *Drink
}

func (t *TeaBuilder) CreateDrink() {
    t.drink = &Drink{}
}

func (t *TeaBuilder) Drink() *Drink {return t.drink}

func (t *TeaBuilder) SetSize() {
    t.drink.SetSize("small")
}

func (t *TeaBuilder) SetDrink() {
    t.drink.SetDrink("tea")
}

type CoffeeBuilder struct {
    drink *Drink
}

func (t *CoffeeBuilder) CreateDrink() {
    t.drink = &Drink{}
}

func (t *CoffeeBuilder) Drink() *Drink {return t.drink}

func (t *CoffeeBuilder) SetSize() {
    t.drink.SetSize("large")
}

func (t *CoffeeBuilder) SetDrink() {
    t.drink.SetDrink("coffee")
}

type Waiter struct {
    builder Builder
}

func (w *Waiter) SetBuilder(b Builder) {
    w.builder = b
}

func (w *Waiter) GetDrink() *Drink {
    return w.builder.Drink()
}

func (w *Waiter) ConstructDrink() {
    w.builder.CreateDrink()
    w.builder.SetSize()
    w.builder.SetDrink()
}

func main() {
    var waiter Waiter

    waiter.SetBuilder(&TeaBuilder{})
    waiter.ConstructDrink()
    drink := waiter.GetDrink()
    fmt.Printf("drink: %s, size: %s\n", drink.Drink(), drink.Size())

    waiter.SetBuilder(&CoffeeBuilder{})
    waiter.ConstructDrink()
    drink = waiter.GetDrink()
    fmt.Printf("drink: %s, size: %s\n", drink.Drink(), drink.Size())
}

