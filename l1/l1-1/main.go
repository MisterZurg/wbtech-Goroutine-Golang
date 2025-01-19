package main

import (
	"fmt"
	"time"
)

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string `pii:sensetive`

	Action
}

func NewHuman(name string) Human {
	return Human{name: name}
}

type Action struct{}

func (act *Action) Work() {
	time.Sleep(time.Second * 3)
	fmt.Println(Cyan + "Я каменщик работаю 3 дня; без зарплаты!" + Reset)
}

func main() {
	mike := NewHuman("mike")
	mike.Work()
}

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[97m"
)
