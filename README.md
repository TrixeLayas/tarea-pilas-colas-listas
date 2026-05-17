Explicación del Gestor de Tareas en Go (Listas Enlazadas)
Este documento explica paso a paso el código del archivo listas.go, el cual implementa un gestor de tareas utilizando una estructura de datos llamada Lista Enlazada Simplemente (Singly Linked List).

📌 Estructuras Principales
Task: Representa una tarea individual. Tiene un Title (texto de la tarea) y un estado Done (booleano que indica si está completada).
Node (Nodo): Es cada elemento de la lista. Contiene una tarea (value) y un puntero (next) que apunta al siguiente nodo de la lista.
TaskList: Representa la lista completa. Solo necesita guardar una referencia al primer nodo, llamado head (cabeza).
🛠️ Tareas (Métodos de la Lista)
1. Agregar tarea (Add)
Añade una nueva tarea al final de la lista.

Crea un nuevo nodo con la tarea.
Si la lista está vacía (l.head == nil), el nuevo nodo se convierte en el head.
Si no, recorre la lista saltando de nodo en nodo hasta encontrar el último (el que tiene next == nil) y lo conecta con el nuevo nodo.
2. Imprimir tareas (Print)
Muestra todas las tareas en consola.

Empieza desde el head.
Mientras el nodo actual no sea nulo, imprime el título de la tarea y su estado, luego avanza al siguiente nodo.
3. Marcar completada (Complete)
Busca una tarea por su título y la marca como terminada.

Recorre la lista comparando el título de cada nodo.
Cuando lo encuentra, cambia Done a true y termina la búsqueda.
4. Contar tareas (Count)
Devuelve el número total de tareas.

Usa un contador que inicia en 0.
Recorre la lista sumando 1 por cada nodo hasta llegar al final.
🗑️ Eliminar tarea (Delete) y la línea if l.head.value.Title == title
El método de eliminar es uno de los más importantes y tiene una validación especial al inicio.

¿Qué hace la línea if l.head.value.Title == title?
Cuando quieres eliminar una tarea, existen dos escenarios principales en una lista enlazada:

La tarea a eliminar es el primer elemento de la lista (el head).
La tarea a eliminar está en el medio o al final de la lista.
Ese if evalúa exactamente el primer escenario. Desglosemos la línea:

if l.head.value.Title == title {
    l.head = l.head.next
    fmt.Println("Tarea eliminada:", title)
    return
}
Explicación paso a paso:

l: Es la lista de tareas.
l.head: Es el primer nodo de la lista.
l.head.value: Accede a la estructura Task (la tarea) que está guardada dentro del primer nodo.
l.head.value.Title: Accede al texto (título) de esa primera tarea.
== title: Compara si el título de esa primera tarea es exactamente igual al título que queremos eliminar.
¿Qué pasa si la condición se cumple (si la tarea a eliminar es la primera)? Se ejecuta la línea l.head = l.head.next. Esto significa que ahora la "cabeza" de la lista dejará de apuntar al primer elemento y apuntará directamente al segundo elemento (l.head.next). Al hacer esto, el primer elemento queda desconectado de la lista y es eliminado automáticamente de la memoria por Go (Garbage Collector).

Ejemplo visual: Tienes la lista: [Aprender Go] -> [Hacer tarea] -> [Estudiar] Si quieres eliminar "Aprender Go", el código se da cuenta de que es el primer elemento (head). Al hacer l.head = l.head.next, la nueva lista comienza desde el segundo elemento: Nueva lista: [Hacer tarea] -> [Estudiar]

Si la tarea a eliminar no es la primera, el código ignora este if y pasa a un ciclo for para buscar la tarea en el resto de la lista.



Explicación del Sistema de Banco en Go (Colas - Queues)
El archivo colas.go implementa una Cola (Queue) simulando la fila de atención de un banco. Las colas siguen el principio FIFO (First In, First Out - El primero en entrar es el primero en salir).

📌 Estructuras Principales
Client: Representa a un cliente del banco. Tiene un Name y un indicador VIP (booleano).
Queue: Es la estructura de datos Cola. Aquí se usa un arreglo/slice de clientes ([]Client).
Bank: Es el sistema general del banco, que contiene una línea de espera (line Queue).
🛠️ Métodos de la Cola
1. Encolar / Agregar (Enqueue)
Agrega un cliente a la fila. Aquí hay una regla especial de prioridad:

Si es VIP (c.VIP == true): Se inserta al inicio de la lista usando append([]Client{c}, q.items...). ¡Los VIP se saltan la fila!
Si NO es VIP: Se inserta al final de la lista de forma normal usando append(q.items, c).
2. Desencolar / Atender (Dequeue)
Saca al primer cliente de la fila para atenderlo.

Si la cola está vacía, devuelve un cliente vacío.
Si hay clientes, toma el primero: first := q.items[0].
Luego, actualiza la cola para eliminar a ese primero: q.items = q.items[1:] (conserva desde el elemento 1 en adelante).
Retorna al cliente que fue sacado.
3. Ver siguiente (Peek)
Solo mira quién es el primero en la fila sin sacarlo. Útil para saber quién sigue sin alterar la cola.



Explicación del Historial de Navegador en Go (Pilas - Stacks)
El archivo pilas.go implementa una Pila (Stack) simulando el historial de navegación de un navegador web (botones Atrás y Adelante). Las pilas siguen el principio LIFO (Last In, First Out - El último en entrar es el primero en salir).

📌 Estructuras Principales
Stack: La pila en sí, basada en un slice de textos ([]string) donde cada texto es una URL.
Browser: Representa el navegador. Tiene:
current: La página actual donde estás.
back: Una Pila con el historial de páginas anteriores (botón "Atrás").
forward: Una Pila con las páginas a las que puedes avanzar (botón "Adelante").
🛠️ Métodos de la Pila
1. Apilar / Agregar (Push)
Agrega un elemento al tope de la pila.

Usa append(s.items, url) para agregarlo al final del slice (que representa la parte superior de la pila).
2. Desapilar / Sacar (Pop)
Saca el elemento del tope de la pila (el último que fue agregado).

Si está vacía, retorna vacío.
Toma el último elemento: last := s.items[len(s.items)-1].
Actualiza la pila borrando ese último elemento: s.items = s.items[:len(s.items)-1].
Retorna el elemento que sacó.
🌐 Lógica del Navegador
Visitar (Visit): Cuando entras a una nueva página, la página en la que estabas se guarda en la pila back (Atrás). Si el historial supera los 10 elementos, se borra el más antiguo. Además, visitar una nueva página vacía la pila forward (Adelante), igual que en Chrome o Firefox.
Retroceder (Back): Saca la última página de la pila back (Pop). Guarda tu página actual en la pila forward (Push) por si quieres volver, y te lleva a esa página anterior.
Avanzar (Forward): Saca la última página de la pila forward (Pop). Guarda tu página actual en la pila back (Push) y te lleva a esa página hacia adelante.
