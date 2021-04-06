// Para esperar a que finalicen varias gorutinas,
// podemos usar un *grupo de espera*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Esta es la función que ejecutaremos en cada gorutina.
// Tenga en cuenta que un grupo de espera debe pasarse
// a las funciones mediante un puntero.
func trabajador(id int, wg *sync.WaitGroup) {
	// A su regreso, notifique al grupo de espera que hemos terminado.
	defer wg.Done()

	fmt.Printf("Trabajador %d comenzando\n", id)

	// Duerme para simular una tarea costosa.
	time.Sleep(time.Second)
	fmt.Printf("Trabajador %d terminado\n", id)
}

func main() {

	// Este grupo de espera se utiliza para esperar
	// a que finalicen todas las gorutinas lanzadas a continuación.
	var wg sync.WaitGroup

	// Inicia varias gorutinas e incremente el contador
	// del grupo de espera para cada trabajador.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go trabajador(i, &wg)
	}

	// Bloquear hasta que el contador del grupo de espera vuelva a 0;
	// todos los trabajadores notificaron que terminaron.
	wg.Wait()
}
