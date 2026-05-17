package main

import "fmt"

// Cliente
type Client struct {
	Name string
	VIP  bool
}

// Cola
type Queue struct {
	items []Client
}

// Agregar cliente
func (q *Queue) Enqueue(c Client) {

	// Si es VIP entra al inicio
	if c.VIP {
		q.items = append([]Client{c}, q.items...)
	} else {
		q.items = append(q.items, c)
	}
}

// Atender cliente
func (q *Queue) Dequeue() Client {
	if len(q.items) == 0 {
		return Client{}
	}

	first := q.items[0]
	q.items = q.items[1:]
	return first
}

// Ver siguiente sin eliminar
func (q *Queue) Peek() Client {
	if len(q.items) == 0 {
		return Client{}
	}

	return q.items[0]
}

// Sistema Banco
type Bank struct {
	line Queue
}

// Llegada cliente
func (b *Bank) Arrive(name string, vip bool) {
	b.line.Enqueue(Client{Name: name, VIP: vip})
}

// Atender
func (b *Bank) Attend() {
	client := b.line.Dequeue()

	if client.Name != "" {
		fmt.Println("Atendiendo a:", client.Name)
	} else {
		fmt.Println("No hay clientes")
	}
}

// Ver siguiente
func (b *Bank) Next() {
	client := b.line.Peek()

	if client.Name != "" {
		fmt.Println("Siguiente cliente:", client.Name)
	} else {
		fmt.Println("No hay clientes")
	}
}

func main() {

	bank := Bank{}

	// Llegan clientes
	bank.Arrive("Ana", false)
	bank.Arrive("Luis", false)
	bank.Arrive("Pedro", true) // VIP
	bank.Arrive("Maria", false)

	// Ver siguiente
	bank.Next()

	// Atender
	bank.Attend()
	bank.Attend()
	bank.Attend()
	bank.Attend()
	bank.Attend()
}
