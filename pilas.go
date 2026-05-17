package main

import "fmt"

// PILA
type Stack struct {
	items []string
}

// Agregar
func (s *Stack) Push(url string) {
	s.items = append(s.items, url)
}

// Sacar
func (s *Stack) Pop() string {

	if len(s.items) == 0 {
		return ""
	}

	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return last
}

// Navegador
type Browser struct {
	current string
	back    Stack
	forward Stack
}

// Visitar página
func (b *Browser) Visit(url string) {

	if b.current != "" {
		b.back.Push(b.current)

		// Limitar historial a 10
		if len(b.back.items) > 10 {
			b.back.items = b.back.items[1:]
		}
	}

	b.current = url

	// Limpiar forward al visitar nueva página
	b.forward.items = []string{}
}

// Retroceder
func (b *Browser) Back() {

	prev := b.back.Pop()

	if prev != "" {
		b.forward.Push(b.current)
		b.current = prev
	}
}

// Avanzar
func (b *Browser) Forward() {

	next := b.forward.Pop()

	if next != "" {
		b.back.Push(b.current)
		b.current = next
	}
}

func main() {

	b := Browser{}

	b.Visit("google.com")
	b.Visit("github.com")
	b.Visit("youtube.com")
	b.Visit("stackoverflow.com")

	fmt.Println("Actual:", b.current)

	b.Back()
	fmt.Println("Back:", b.current)

	b.Back()
	fmt.Println("Back:", b.current)

	b.Forward()
	fmt.Println("Forward:", b.current)
}
