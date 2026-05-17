package main

import "fmt"

// Estructura tarea
type Task struct {
	Title string
	Done  bool
}

// Nodo
type Node struct {
	value Task
	next  *Node
}

// Lista enlazada
type TaskList struct {
	head *Node
}

// Agregar tarea
func (l *TaskList) Add(title string) {

	newNode := &Node{
		value: Task{
			Title: title,
			Done:  false,
		},
	}

	// Si está vacía
	if l.head == nil {
		l.head = newNode
		return
	}

	// Buscar final
	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
}

// Imprimir tareas
func (l *TaskList) Print() {

	temp := l.head

	for temp != nil {
		fmt.Println("-", temp.value.Title, "| Done:", temp.value.Done)
		temp = temp.next
	}
}

// Marcar completada
func (l *TaskList) Complete(title string) {

	temp := l.head

	for temp != nil {

		if temp.value.Title == title {
			temp.value.Done = true
			fmt.Println("Tarea completada:", title)
			return
		}

		temp = temp.next
	}

	fmt.Println("Tarea no encontrada")
}

// Eliminar tarea
func (l *TaskList) Delete(title string) {

	// Si lista vacía
	if l.head == nil {
		return
	}

	// Si está al inicio
	if l.head.value.Title == title {
		l.head = l.head.next
		fmt.Println("Tarea eliminada:", title)
		return
	}

	temp := l.head

	for temp.next != nil {

		if temp.next.value.Title == title {
			temp.next = temp.next.next
			fmt.Println("Tarea eliminada:", title)
			return
		}

		temp = temp.next
	}

	fmt.Println("Tarea no encontrada")
}

// Contar tareas
func (l *TaskList) Count() int {

	contador := 0
	temp := l.head

	for temp != nil {
		contador++
		temp = temp.next
	}

	return contador
}

func main() {

	list := TaskList{}

	list.Add("Aprender Go")
	list.Add("Hacer tarea")
	list.Add("Estudiar estructuras")

	fmt.Println("LISTA ORIGINAL")
	list.Print()

	fmt.Println()

	list.Complete("Hacer tarea")

	list.Delete("Aprender Go")

	fmt.Println()

	fmt.Println("LISTA ACTUALIZADA")
	list.Print()

	fmt.Println()

	fmt.Println("Total tareas:", list.Count())
}
