package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct {
	Message string
}

// SepcRequest метод который мы адаптируем
func (a *Adaptee) SepcRequest() string {
	return a.Message
}

type Adapter struct {
	Adaptee *Adaptee
}

// Request Реализация метода
func (adapter *Adapter) Request() string {
	return adapter.Adaptee.SepcRequest()
}

func main() {
	adaptee := &Adaptee{Message: "Это сообщение от адаптируемого класса"}
	adapter := &Adapter{Adaptee: adaptee}

	fmt.Println("Вызов метода Request через адаптер:")
	fmt.Println(adapter.Request())
}
