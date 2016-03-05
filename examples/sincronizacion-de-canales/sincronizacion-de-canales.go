// Podemos utilizar canales para sincronizar
// la ejecución a traves de diferentes goroutines.
// Aquí un ejemplo donde utilizamos un bloqueo
// para esperar a que termine de ejectuarse la goroutine.

package main

import "fmt"
import "time"

// Esta es la función que ejecutaremos en una goroutine. Vamos a usar
// el canal de `terminado` para notificar a otra goroutine que el
// trabajo de la función ha terminado.
func trabajo(terminado chan bool) {
    fmt.Print("trabajando...")
    time.Sleep(time.Second)
    fmt.Println("terminado")

    // Enviamos un valor para notificar que terminamos.
    terminado <- true
}

func main() {

    // Comenzamos una goroutine con `trabajo`, y le pasamos el canal al
    // que tiene que notificar
    terminado := make(chan bool, 1)
    go trabajo(terminado)

    // Hay un bloqueo hasta que recibimos la notificación
    // de trabajo en el canal.
    <-terminado
}
